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

var (
	inputFileName  string
	outputFileName string
	retainComments bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a sample dotenv file from an existing .env file",
	Long: `Generate a sample dotenv file (.env.example) or a custom output file based on an existing .env file.
The generated file will contain the keys from the .env file with empty values.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open the input .env file
		envFile, err := os.Open(inputFileName)
		if err != nil {
			fmt.Println("Error opening .env file:", err)
			return
		}
		defer envFile.Close()

		// Determine the output file name
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

			// Skip empty lines and comments
			if line == "" || strings.HasPrefix(line, "#") {
				if retainComments {
					_, _ = sampleFile.WriteString(line + "\n")
				}
				continue
			}

			// Split the line into key and value
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]

			// Write the key with an empty value to the sample file
			_, _ = sampleFile.WriteString(key + "=\n")

			// If the line contains a comment, write the comment to the sample file
			if len(parts) > 1 {
				comment := strings.TrimSpace(parts[1])
				if strings.HasPrefix(comment, "#") {
					_, _ = sampleFile.WriteString(comment + "\n")
				}
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

	// Add flags for specifying the input and output file names
	generateCmd.Flags().StringVarP(&inputFileName, "file", "f", ".env", "Input file name (existing .env file)")
	generateCmd.Flags().StringVarP(&outputFileName, "output", "o", "", "Output file name for the sample dotenv file")
	generateCmd.Flags().BoolVarP(&retainComments, "comments", "c", false, "Retain comments from the original .env file")
}
