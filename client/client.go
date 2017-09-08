package client

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sfn"
)

func NewSFNClient(key string, secret string, region string) (*sfn.SFN, error) {

  var svc *sfn.SFN
  var regionVal = ""
  if(region !=""){
    regionVal =  "us-west-2"
  }

  if key != "" && secret != "" {
    awsConfig := &aws.Config {
      Credentials: credentials.NewStaticCredentials(key, secret, ""),
      Region: aws.String(regionVal),
    }
    sess, err := session.NewSession()
    if err != nil {
      fmt.Println("failed to create session,", err)
      return nil, err
    }
    svc = sfn.New(sess, awsConfig)
  }else {
    svc = sfn.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})
  }

  return svc, nil
}