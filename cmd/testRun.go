/*
Copyright © 2022 Salvatore Fasano fasanosalvatore@hotmail.it

*/
package cmd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/flaky-infra/flaky-cli/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type FlakyResponse struct {
	Message string `json:"message"`
	Data    struct {
		ProjectID string `json:"projectId"`
		TestRunID string `json:"testRunId"`
	} `json:"data"`
}

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if key == "projectFile" {
			file, err := os.Open(val)
			if err != nil {
				return "", nil, err
			}
			defer file.Close()
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			h := sha256.New()
			io.Copy(part, file)
			io.Copy(h, file)
			if _, ok := form["commitId"]; !ok {
				form["commitId"] = hex.EncodeToString(h.Sum(nil))
			}
		} else {
			mp.WriteField(key, val)
		}
	}
	return mp.FormDataContentType(), body, nil
}

var (
	sourceUrl           string
	commitId            string
	moduleName          string
	testMethodName      string
	configurationFolder string
	clusterExecName     string
	watchFlag           bool
	projectVersion      int
	runCmd              = &cobra.Command{
		Use:   "run",
		Short: "The test run command allows you to perform the flakiness analysis of a project",
		Long: `The test run command allows you to carry out the flakiness analysis of a local or remote project,
		correctly configured with a .flaky.yaml file, by sending it to one of the added clusters.`,
		Run: func(cmd *cobra.Command, args []string) {
			var testRuns []map[string]interface{}
			util.GetSliceMapViper("testRuns", &testRuns)

			if testRuns == nil {
				testRuns = make([]map[string]interface{}, 0)
			}

			var target string
			home, _ := os.UserHomeDir()
			flakyPath := filepath.Join(home, ".flaky", "tmp")
			form := map[string]string{"moduleName": moduleName}

			if strings.HasPrefix(sourceUrl, "https://") {
				projectName := strings.Split(sourceUrl, "/")[4]
				form["gitUrl"] = sourceUrl
				sourceUrl = filepath.Join(flakyPath, projectName)
				_, commandErr := exec.Command("bash", "-c", fmt.Sprintf("git clone %s %s && cd %s && git reset --hard %s", form["gitUrl"], sourceUrl, sourceUrl, commitId)).Output()
				if commandErr != nil {
					fmt.Println(commandErr.Error())
				}
				form["commitId"] = commitId
			}

			projectAbsPath, _ := filepath.Abs(sourceUrl)
			form["name"] = filepath.Base(projectAbsPath)
			if configurationFolder != "" {
				configurationAbsPath, _ := filepath.Abs(configurationFolder)
				_, commandErr := exec.Command("cp", "-R", configurationAbsPath, filepath.Join(projectAbsPath)).Output()
				if commandErr != nil {
					fmt.Println(commandErr.Error())
				}
				form["configurationFolder"] = filepath.Base(configurationFolder)
			} else {
				present, _ := filepath.Glob(filepath.Join(projectAbsPath, "*.flaky.yaml"))
				if len(present) == 0 {
					panic("Errore")
				}
			}

			i := strings.LastIndex(testMethodName, ".")
			testMethodName := testMethodName[:i] + strings.Replace(testMethodName[i:], ".", "#", 1)
			form["testMethodName"] = testMethodName
			target, _ = util.CreateTarArchive(projectAbsPath, flakyPath)
			util.Gzip(target, flakyPath)
			form["projectFile"] = target + ".gz"

			ct, body, err1 := createForm(form)
			if err1 != nil {
				panic(err1)
			}

			if clusterExecName == "" {
				clusterExecName = fmt.Sprintf("%s", viper.Get("defaultCluster"))
			}

			var clusters []map[string]interface{}
			util.GetSliceMapViper("clusters", &clusters)

			clusterIndex := isClusterPresent(&clusters, clusterExecName)

			resp, err := http.Post(fmt.Sprintf("%s/api/application", clusters[clusterIndex]["clusterUrl"]), ct, body)

			dir, _ := ioutil.ReadDir(flakyPath)
			for _, d := range dir {
				os.RemoveAll(filepath.Join([]string{flakyPath, d.Name()}...))
			}
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			defer resp.Body.Close()

			bodyResp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			flakyResponse := FlakyResponse{}
			json.Unmarshal(bodyResp, &flakyResponse)

			newTestRun := make(map[string]interface{})
			newTestRun["testRunId"] = flakyResponse.Data.TestRunID
			newTestRun["projectId"] = flakyResponse.Data.ProjectID
			newTestRun["date"] = time.Now().Format("2006-01-02 15:04:05")
			newTestRun["cluster"] = clusterExecName

			testRuns = append(testRuns, newTestRun)

			fmt.Printf("Test run successfully launched, you can check the status using the following id %s\n", flakyResponse.Data.TestRunID)

			viper.Set("testRuns", testRuns)
			viper.WriteConfig()

		},
	}
)

func init() {
	testCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&sourceUrl, "source", "s", ".", "url at which the project can be found, local path or repository url")
	runCmd.Flags().StringVarP(&commitId, "commitId", "i", "", "hash of the commit to refer to if it is a repository (default last commit)")
	runCmd.Flags().StringVarP(&moduleName, "moduleName", "m", ".", "module name")
	runCmd.Flags().StringVarP(&testMethodName, "testName", "t", "", "test name")
	runCmd.Flags().IntVarP(&projectVersion, "projectVersion", "v", 1, "project version")
	runCmd.Flags().StringVarP(&configurationFolder, "configuration", "p", "", "configuration path")
	runCmd.Flags().StringVarP(&clusterExecName, "cluster", "c", "", "cluster name on which to run the project (default cluster)")
	runCmd.Flags().BoolVarP(&watchFlag, "watch", "w", false, "watch test executions with results (default false)")
}
