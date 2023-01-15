package util

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
	"os"
	"sync"
)

var ssmParams map[string]string
var lock = &sync.Mutex{}

func GetSmmParams() map[string]string {
	if ssmParams == nil {
		lock.Lock()
		defer lock.Unlock()
		if ssmParams == nil {
			ssmParams = fetchSSMParams()
		}
	}
	return ssmParams
}

func fetchSSMParams() map[string]string {
	config := map[string]string{}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	if err != nil {
		log.Fatal(err)
	}
	svc := ssm.New(sess)
	path := fmt.Sprintf("/sls-go-pg-rest-demo/%s/", os.Getenv("STAGE"))
	fmt.Println("path", path)
	ssmQuery := &ssm.GetParametersByPathInput{
		Recursive:      aws.Bool(true),
		Path:           aws.String(path),
		MaxResults:     aws.Int64(10),
		NextToken:      nil,
		WithDecryption: aws.Bool(true),
	}

	for {
		results, err := svc.GetParametersByPath(ssmQuery)
		fmt.Println("results", results)
		if err != nil {
			log.Fatal(err)
		}
		updateResult(&config, results)
		if results.NextToken == nil {
			break
		}
		ssmQuery.NextToken = results.NextToken
	}
	return config
}

func updateResult(m *map[string]string, results *ssm.GetParametersByPathOutput) {
	for _, r := range results.Parameters {
		(*m)[*r.Name] = *r.Value
	}
}
