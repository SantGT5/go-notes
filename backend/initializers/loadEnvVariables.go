package initializers

import (
	"fmt"
	"os"
)

// envVars is a list of essential env variables required to run the Backend project.
var envVars []string = []string{
	"BACKEND_PORT",
	"POSTGRES_URL",
	"JWT_SECRET_KEY",
	"BASE_URL_V1",
	"ALLOW_ORIGINS",
}

// LoadEnvVariables checks for empty or unset env variables and prints a list of them.
// It ensures that the necessary env variables required to run the backend project are set.
func LoadEnvVariables() {
	var emptyVars []string = make([]string, 0, len(envVars))

	for _, envVar := range envVars {
		value := os.Getenv(envVar)
		if value == "" {
			emptyVars = append(emptyVars, envVar)
		}
	}

	if len(emptyVars) > 0 {
		fmt.Println("\n➡️ The following env variables are either not set or empty:")
		for _, varName := range emptyVars {
			fmt.Printf("	❌ %v", varName)
		}
		fmt.Println("\n ")
	}
}
