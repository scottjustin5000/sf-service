package monitor 

import (
  "fmt"
  "time"
  "strings"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/sfn"

  "github.com/scottjustin5000/sf-service/client"
)

var sfnsvc *sfn.SFN

type StateMachineResult struct {
    Name  string
    Arn string 
}

type ExecutionResult struct {
  Name string
  Status string
  Start time.Time
  ExecutionTime float64
}

func getSfnClient() *sfn.SFN {
  if sfnsvc != nil {
    return sfnsvc
  }

  sfnsvc, err := client.NewSFNClient("","","")
  if err != nil {
    fmt.Println("failed to create session")
    return nil
  }
  return sfnsvc
}

func getByStatus(arn string, status string) ([]string, error) {
  var svc = getSfnClient()
   params := &sfn.ListExecutionsInput {
      MaxResults: aws.Int64(100),
      StateMachineArn: &arn,
      StatusFilter: aws.String(status),
  }
  resp, err := svc.ListExecutions(params)
  if err != nil {
    fmt.Println("error:", err)
    return nil, err
  }
  var result []string
  for _, ex := range resp.Executions {
    result = append(result, *ex.ExecutionArn)
  }

  return result, nil
}

func GetFailures(arn string) ([]string, error) {
  return getByStatus(arn, "FAILED")
}

func GetSucesses(arn string) ([]string, error) {
  return getByStatus(arn, "SUCCEEDED")
}

func GetInput(exArn string) (string, error){
  var svc = getSfnClient()
   params := &sfn.DescribeExecutionInput {
    ExecutionArn: &exArn,
  }
  resp, err := svc.DescribeExecution(params)

  if err != nil {
    fmt.Println("error:", err)
    return "", err
  }

  return *resp.Input, nil

}

func ListStateMachines() ([]StateMachineResult, error) {
  var svc = getSfnClient()
   params := &sfn.ListStateMachinesInput {
      MaxResults: aws.Int64(100),
  }
  resp, err := svc.ListStateMachines(params)

  if err != nil {
    fmt.Println("errr", err)
    return nil, err
  }

  var result []StateMachineResult

  for _, machine := range resp.StateMachines {
    name := *machine.Name
    arn := *machine.StateMachineArn
    mach := StateMachineResult{name, arn}
    result = append(result, mach)
  }

  return result, nil

}

func filterExecutionPage(page *sfn.ListExecutionsOutput, pattern string) ExecutionResult {
  var execution ExecutionResult
  for _, ex := range page.Executions {
    var name = *ex.Name
    if strings.ToLower(name) == strings.ToLower(pattern) {
      start := *ex.StartDate
      durration :=  start.Sub(*ex.StopDate).Seconds() * 1000
      name := name
      status := *ex.Status
      execution = ExecutionResult{ name, status, start, durration }
    }
  }
  return execution
}

func FindExecution(stateMachine string, pattern string) ExecutionResult {
  var svc = getSfnClient()
  var execution ExecutionResult
  params := &sfn.ListExecutionsInput {
      MaxResults: aws.Int64(100),
      StateMachineArn: &stateMachine,
    }
  err := svc.ListExecutionsPages(params,
    func(page *sfn.ListExecutionsOutput, lastPage bool) bool {
        execution = filterExecutionPage(page, pattern)
        if execution.Name != "" || lastPage {
          return false
        }
        return true
    })
  if err != nil {
    fmt.Println(err)
  }
  return execution
}



