package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	goreloaded "goreloaded/modifications"
)

func main() {
	// Check if proper arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run program.go input_file output_file")
		os.Exit(1)
	}

	// Get input and output file names from command line arguments
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	sample, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer sample.Close()

	result, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer result.Close()

	// Create a scanner to read from the input file
	scanner := bufio.NewScanner(sample)

	// Create a writer to write to the output file
	writer := bufio.NewWriter(result)
	defer writer.Flush() // Ensure the writer is flushed when done

	// Iterate over each line of the input file
	for scanner.Scan() {
		lines := scanner.Text()
		line := strings.Split(lines, " ")

		// Calling on the functions that apply the modifications
		goreloaded.BintoInt(line)
		goreloaded.HextoInt(line)
		goreloaded.ToUpper(line)
		goreloaded.ToLower(line)
		goreloaded.Capitalize(line)
		goreloaded.AtoAn(line)

		modifiedLine := strings.Join(line, " ")

		// Write the modified line to the output file
		_, err := writer.WriteString(modifiedLine + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			os.Exit(1)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
		os.Exit(1)
	}

	fmt.Println("Output written to", outputFile)
}
