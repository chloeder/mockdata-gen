package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Declare variables
	var inputPath, outputPath string
	var helpPath bool

	// Get input and output paths
	flag.StringVar(&inputPath, "i", "", "Location of the JSON file to be read")
	flag.StringVar(&inputPath, "input", "", "Location of the JSON file to be read")
	flag.StringVar(&outputPath, "o", "", "Location of the JSON file output")
	flag.StringVar(&outputPath, "output", "", "Location of the JSON file output")

	// Get help path
	flag.BoolVar(&helpPath, "h", false, "Help")
	flag.BoolVar(&helpPath, "help", false, "Help")

	// Parse the flags
	flag.Parse()

	// Check if the help, input, output flag is set
	if helpPath || inputPath == "" || outputPath == "" {
		printUsage()
		os.Exit(0)
	}

	// Validate the input and output path
	if err := validateInputPath(inputPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}

	if err := validateOutputPath(outputPath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("Usage  : [-i || --input] <input file> [-o || --output] <output file>")
	fmt.Println("Help   : [--help] or [-h] for help")
	fmt.Println("Input  : [--input] or [-i] for input file")
	fmt.Println("Output : [--output] or [-o] for output file")
}

func validateInputPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	return nil
}

func validateOutputPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	overwriteFile()
	return nil
}

func overwriteFile() {
	fmt.Println("Output file already exists. Do you want to overwrite it? [y/n]")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" {
		fmt.Println("Aborting...")
		os.Exit(0)
	}
}
