package util

import (
	"fmt"
	"os"
)

func getDbCredentials() *DbCredential {

	if !InLambda() {
		if os.Getenv("HOST_NAME") == "" {
			panic("HOST_NAME env variable for db connection is not present")
		}
		if os.Getenv("USER_NAME") == "" {
			panic("USER_NAME env variable for db connection is not present")
		}
		if os.Getenv("PASSWORD") == "" {
			panic("PASSWORD env variable for db connection is not present")
		}
		if os.Getenv("DEFAULT_DB") == "" {
			panic("DEFAULT_DB env variable for db connection is not present")
		}
		return CreateDbCredential(
			os.Getenv("HOST_NAME"),
			os.Getenv("USER_NAME"),
			os.Getenv("PASSWORD"),
			os.Getenv("DEFAULT_DB"),
		)
	}

	params := GetSmmParams()

	return CreateDbCredential(
		params[fmt.Sprintf("/sls-go-pg-rest-demo/%s/db-host", os.Getenv("STAGE"))],
		params[fmt.Sprintf("/sls-go-pg-rest-demo/%s/db-user", os.Getenv("STAGE"))],
		params[fmt.Sprintf("/sls-go-pg-rest-demo/%s/db-password", os.Getenv("STAGE"))],
		params[fmt.Sprintf("/sls-go-pg-rest-demo/%s/db-default-name", os.Getenv("STAGE"))],
	)

}
