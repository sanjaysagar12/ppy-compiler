package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Define command-line flags
	inputFilePtr := flag.String("i", "", "Input PPY file path (required)")
	outputFilePtr := flag.String("o", "", "Output Python file path (optional, if not provided output is printed to console)")

	// Parse the flags
	flag.Parse()

	// Check if input file is provided
	if *inputFilePtr == "" {
		// Check if a positional argument was provided
		if flag.NArg() > 0 {
			*inputFilePtr = flag.Arg(0)
		} else {
			fmt.Println("Error: Input file is required")
			fmt.Println("Usage: ppy-compiler -i input.ppy [-o output.py]")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	// Read input file
	content, err := os.ReadFile(*inputFilePtr)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	// Convert HTML to Python
	pythonCode := convertHTMLToPython(string(content))

	// Check if output file is specified
	outputFile := *outputFilePtr
	if outputFile != "" {
		// Write to output file
		err = os.WriteFile(outputFile, []byte(pythonCode), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully converted %s to %s\n", *inputFilePtr, outputFile)
	} else {
		// Print to console
		fmt.Println(pythonCode)
	}
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, "'", "\\'")
}

func convertHTMLToPython(html string) string {
	var result strings.Builder

	// First replace PHP-style variable expressions with our markers
	varExprRegex := regexp.MustCompile(`<\?=\s*(.*?)\s*\?>`)
	html = varExprRegex.ReplaceAllString(html, "~~~PPY_VAR~~~$1~~~")

	// Split the content by Python blocks
	pieces := regexp.MustCompile(`(?s)<py\?(.*?)\?>`).Split(html, -1)
	pythonBlocks := regexp.MustCompile(`(?s)<py\?(.*?)\?>`).FindAllStringSubmatch(html, -1)

	// Track indentation level
	indentLevel := 0

	// Process all HTML pieces
	for i, piece := range pieces {
		// Process HTML content
		lines := strings.Split(piece, "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}

			// Check if line contains our variable markers
			if strings.Contains(line, "~~~PPY_VAR~~~") {
				// Split the line by our markers to separate HTML and variables
				parts := regexp.MustCompile(`~~~PPY_VAR~~~(.*?)~~~`).Split(line, -1)
				varMatches := regexp.MustCompile(`~~~PPY_VAR~~~(.*?)~~~`).FindAllStringSubmatch(line, -1)

				// Build a print statement that concatenates variables and strings
				indentation := strings.Repeat("    ", indentLevel)
				printStatement := indentation + "print("

				for j, part := range parts {
					if part != "" {
						printStatement += "'" + escapeQuotes(part) + "'"
					}

					if j < len(varMatches) {
						// Add the variable expression
						if j < len(parts)-1 || parts[j] != "" {
							printStatement += " + str(" + varMatches[j][1] + ")"
						} else {
							printStatement += "str(" + varMatches[j][1] + ")"
						}

						// Add concatenation operator if needed
						if j < len(parts)-1 && parts[j+1] != "" {
							printStatement += " + "
						}
					}
				}

				printStatement += ")"
				result.WriteString(printStatement + "\n")
			} else {
				// Regular HTML line
				indentation := strings.Repeat("    ", indentLevel)
				result.WriteString(fmt.Sprintf("%sprint('%s')\n", indentation, escapeQuotes(line)))
			}
		}

		// Process Python block if there is one
		if i < len(pythonBlocks) {
			pythonCode := strings.TrimSpace(pythonBlocks[i][1])

			// Split the Python code into lines to handle multiple # end statements
			pyLines := strings.Split(pythonCode, "\n")

			for _, pyLine := range pyLines {
				trimmedLine := strings.TrimSpace(pyLine)

				if strings.Contains(trimmedLine, "for ") && strings.Contains(trimmedLine, ":") {
					// This is a loop start - add the line and increase indentation
					indentation := strings.Repeat("    ", indentLevel)
					result.WriteString(indentation + trimmedLine + "\n")
					indentLevel++
				} else if strings.Contains(trimmedLine, "# End") || strings.Contains(trimmedLine, "# end") {
					// This is an end block - decrease indentation before adding the line
					if indentLevel > 0 {
						indentLevel--
					}
					indentation := strings.Repeat("    ", indentLevel)
					result.WriteString(indentation + trimmedLine + "\n")
				} else if trimmedLine != "" {
					// Regular Python code - just add it with the current indentation
					indentation := strings.Repeat("    ", indentLevel)
					result.WriteString(indentation + trimmedLine + "\n")
				}
			}
		}
	}

	return result.String()
}
