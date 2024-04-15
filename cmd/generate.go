/*
Copyright © 2024 Wickes1 <27881570+wickes1@users.noreply.github.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var outputFileName string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a sample dotenv file from an existing .env file",
	Long: `Generate a sample dotenv file (.env.example) or a custom output file based on an existing .env file.
The generated file will contain the keys from the .env file with empty values.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open the existing .env file
		envFile, err := os.Open(".env")
		if err != nil {
			fmt.Println("Error opening .env file:", err)
			return
		}
		defer envFile.Close()

		// Determine the output file name
		outputFileName, _ := cmd.Flags().GetString("output")
		if outputFileName == "" {
			outputFileName = ".env.example"
		}

		// Create a new file for the sample dotenv
		sampleFile, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println("Error creating", outputFileName, "file:", err)
			return
		}
		defer sampleFile.Close()

		// Read each line of the .env file
		scanner := bufio.NewScanner(envFile)
		for scanner.Scan() {
			line := scanner.Text()
			// Split the line into key-value pair
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				// Write key to the sample dotenv file
				fmt.Fprintf(sampleFile, "%s=\n", key)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading .env file:", err)
			return
		}

		fmt.Println("Sample dotenv file has been saved to:", outputFileName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Add a flag for specifying the output file name
	generateCmd.Flags().StringVarP(&outputFileName, "output", "o", "", "Output file name for the sample dotenv file")
}
