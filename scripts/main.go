package main

import (
	"fmt"
	"os"
)

var envVars = []string{
	"POSTGRES_USER",
	"POSTGRES_PASSWORD",
	"POSTGRES_HOST",
	"POSTGRES_PORT",
	"POSTGRES_DB",
	"POSTGRES_URL",
	"JWT_SECRET_KEY",
	"BACKEND_PORT",
	"ALLOW_ORIGIN",
}

// checkEnv checks for empty or unset env variables and prints a list of them.
// It ensures that the necessary env variables required to run the backend project are set.
func CheckEnv() {
	var emptyVars []string = make([]string, 0, len(envVars))

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value == "" {
			emptyVars = append(emptyVars, envVar)
		}
	}

	if len(emptyVars) > 0 {
		fmt.Println("\n➡️ The following env variables are either not set or empty:")
		for _, varName := range emptyVars {
			fmt.Printf("  ❌ %v\n", varName)
		}
		fmt.Println()
		os.Exit(1)
	}

	fmt.Println("✅ All required environment variables are set.")
}

func main() {
	CheckEnv()
}
