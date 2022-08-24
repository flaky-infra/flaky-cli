/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/flaky-infra/flaky-cli/util"
	"github.com/spf13/cobra"
)

type TestExecutionResult struct {
	Stacktrace  string `json:"stacktrace"`
	Classname   string `json:"classname"`
	Displayname string `json:"displayName"`
	Name        string `json:"name"`
	Message     string `json:"message"`
}

type TestFailures struct {
	TestCases []TestExecutionResult `json:"testCases"`
}

type ScenarioExecutionResult struct {
	ScenarioName          string         `json:"scenarioName"`
	ScenarioConfiguration string         `json:"scenarioConfiguration"`
	NumberOfRuns          int            `json:"numberOfRuns"`
	ExecutionWithFailure  int            `json:"executionWithFailure"`
	FailureRate           int            `json:"failureRate"`
	ExecutionLog          string         `json:"executionLog"`
	TestExecutionsResults []TestFailures `json:"testExecutionsResults"`
}

type TestRun struct {
	TestMethodName           string                    `json:"testMethodName"`
	ConfigFolderPath         string                    `json:"configFolderPath"`
	RootCause                ScenarioExecutionResult   `json:"rootCause"`
	ScenarioExecutionsResult []ScenarioExecutionResult `json:"scenarioExecutionsResult"`
}

type TestRunResponse struct {
	Message string `json:"message"`
	Data    struct {
		TestRun TestRun `json:"testRun"`
	} `json:"data"`
}

func isTestRunPresent(testRuns *[]map[string]interface{}, testRunId string) int {
	testRunIndex := -1
	for i, testRun := range *testRuns {
		if testRun["testRunId"] == testRunId {
			testRunIndex = i
			break
		}
	}
	return testRunIndex
}

var inspectCmd = &cobra.Command{
	Use:   "inspect [id of execution]",
	Short: "The test inspect command allows you to view details relating to an execution",
	Long:  `The test inspect command allows you to view details relating to an execution.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		testRunId := args[0]

		var testRuns []map[string]interface{}
		util.GetSliceMapViper("testRuns", &testRuns)

		testRunIndex := isTestRunPresent(&testRuns, testRunId)

		if testRunIndex == -1 {
			fmt.Printf("Error: Test run %s not found.\n", testRunId)
			os.Exit(0)
		}

		localTestRun := testRuns[testRunIndex]

		var clusters []map[string]interface{}
		util.GetSliceMapViper("clusters", &clusters)

		clusterIndex := isClusterPresent(&clusters, fmt.Sprintf("%s", localTestRun["cluster"]))

		resp, err := http.Get(fmt.Sprintf("%s/api/testRun/%s", clusters[clusterIndex]["clusterUrl"], testRunId))
		defer resp.Body.Close()

		bodyResp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		testRun := TestRunResponse{}
		err = json.Unmarshal(bodyResp, &testRun)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(testRun.Data.TestRun.RootCause.ScenarioName)

	},
}

func init() {
	testCmd.AddCommand(inspectCmd)
}
