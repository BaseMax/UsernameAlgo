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
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	outputPrefixes := []string{}
	i := 0
	for i < len(lines)-1 {
		j := i + 1
		currentPrefix := commonPrefix(lines[i], lines[j])
		groupCount := 2

		for j+1 < len(lines) {
			nextPrefix := commonPrefix(currentPrefix, lines[j+1])
			if nextPrefix == "" {
				break
			}
			currentPrefix = nextPrefix
			j++
			groupCount++
		}

		if groupCount >= 3 && currentPrefix != "" {
			outputPrefixes = append(outputPrefixes, currentPrefix)
			i = j + 1
		} else {
			i++
		}
	}

	if len(outputPrefixes) == 0 {
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
	for _, prefix := range outputPrefixes {
		writer.WriteString(prefix + "\n")
	}
	writer.Flush()

	fmt.Printf("✅ Common prefixes written to: %s\n", outputPath)
}
