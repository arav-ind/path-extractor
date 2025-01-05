package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extractPathsFromSVG(svgContent string) map[string]string {
	// Regex to match the `d="..."` attributes in the SVG paths
	re := regexp.MustCompile(`d="([^"]+)"`)
	matches := re.FindAllStringSubmatch(svgContent, -1)

	// Map to store the extracted path data
	pathData := make(map[string]string)
	for i, match := range matches {
		// Generate unique keys like A1, A2, ...
		key := fmt.Sprintf("A%d", i+1)
		pathData[key] = match[1]
	}
	return pathData
}

// Function to generate JS code from the path data
func generateJSCode(pathData map[string]string) (string, string) {
	var sb strings.Builder
	sb.WriteString("export const drawPath = {\n")

	// Create a slice to store district codes (keys)
	var districtCodes []string

	// Iterate over the map and append each path to the JS object
	for key, path := range pathData {
		if !strings.HasPrefix(path, "m ") {
			continue
		}
		sb.WriteString(fmt.Sprintf("  %s: `%s`,\n", key, path))
		// Add the key (district code) to the array
		districtCodes = append(districtCodes, fmt.Sprintf("\"%s\"", key))
	}

	// Remove the last comma and close the object
	sb.WriteString("};\n")

	// Create the districtCode export string
	districtCode := fmt.Sprintf("export const districtCodes = [%s];\n", strings.Join(districtCodes, ", "))

	// Return both JS code and districtCode export
	return sb.String(), districtCode
}

func main() {
	// Define the path to the SVG file (assuming it's in the root directory)
	svgFilePath := "./input.svg"

	// Open the SVG file
	file, err := os.Open(svgFilePath)
	if err != nil {
		fmt.Println("Error opening SVG file:", err)
		return
	}
	defer file.Close()

	// Read the SVG file content
	var svgContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		svgContent += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading SVG file:", err)
		return
	}

	// Extract paths from the SVG content
	pathData := extractPathsFromSVG(svgContent)

	// Generate JavaScript code and the districtCode export
	jsCode, districtCode := generateJSCode(pathData)

	// Write the JS code to a file
	jsFilePath := "drawpath.js"
	err = os.WriteFile(jsFilePath, []byte(jsCode+districtCode), 0644)
	if err != nil {
		fmt.Println("Error writing JS file:", err)
		return
	}

	fmt.Printf("JS file generated: %s\n", jsFilePath)
}
