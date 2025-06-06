package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInputFilePath() (string, error) {
	if len(os.Args) > 1 {
		return strings.TrimSpace(os.Args[1]), nil
	}
	fmt.Print("Enter file path: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}
	return "", fmt.Errorf("no input provided")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func commonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a[:minLen]
}

func main() {
	inputPath, err := getInputFilePath()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if !fileExists(inputPath) {
		fmt.Println("Error: File does not exist.")
		return
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var outputGroups [][]string
	for i := 0; i < len(lines)-1; i++ {
		base := lines[i]
		prefix := ""
		group := []string{base}
		for j := i + 1; j < len(lines); j++ {
			p := commonPrefix(base, lines[j])
			if p == "" {
				break
			}
			base = p
			prefix = p
			group = append(group, lines[j])
		}
		if len(group) >= 3 {
			outputGroups = append(outputGroups, group)
			fmt.Println("Common prefix:", prefix)
			i += len(group) - 1 // skip next already grouped
		}
	}

	if len(outputGroups) == 0 {
		fmt.Println("No matching prefix groups found.")
		return
	}

	outputPath := inputPath + ".out"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, group := range outputGroups {
		for _, line := range group {
			writer.WriteString(line + "\n")
		}
		writer.WriteString("\n")
	}
	writer.Flush()

	fmt.Printf("Filtered output written to: %s\n", outputPath)
}
