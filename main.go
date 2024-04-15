package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open the existing .env file
    envFile, err := os.Open(".env")
    if err != nil {
        fmt.Println("Error opening .env file:", err)
        return
    }
    defer envFile.Close()

    // Create a new file for the sample dotenv
    sampleFile, err := os.Create("sample.env")
    if err != nil {
        fmt.Println("Error creating sample.env file:", err)
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

    fmt.Println("Sample dotenv file generated successfully!")
}
