# SF service

simple wrapper for managing aws step-functions.

### Usage

* List StateMachines

```go

package main

import (
  "fmt"
  "github.com/scottjustin5000/sf-service/monitor"
)

func main() {
  resp, err := monitor.ListStateMachines()
  if err !=nil {
    fmt.Println("err", err)
  }

  for _, v := range resp {
    fmt.Println("name:", v.Name)
    fmt.Println("arn:", v.Arn)
  }
}

```

* Get Failures

```go
package main

import (
  "fmt"
  "github.com/scottjustin5000/sf-service/monitor"
)

func main() {
  resp, err := monitor.GetFailures("state-machine-arn")
  if err != nil {
    fmt.Println("error:", err)
  }
  for _, v := range resp {
    fmt.Println("ex-arn:", v)
  }
}

```

* Get Inputs

```go
import (
  "github.com/scottjustin5000/sqs-monitor/monitor"
)


func main() {
  resp, err :=monitor.GetInput("execution-arn")
  if err!= nil {
    fmt.Println("error:", err)
  }
  fmt.Println("input:", resp)
}

```

Assumes ~/.aws profile is present
