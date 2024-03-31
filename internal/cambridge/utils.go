package cambridge

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func getQuestionType(xml string) string {
	re := regexp.MustCompile(QUESTION_IDENTIFIER_REGEX)

	// Find the match
	match := re.FindStringSubmatch(xml)

	if len(match) > 1 {
		return match[1]
	}

	return ""
}

func cleanSimpleQuestionValue(matches [][]string) string {
	var cleanedValues []string
	for _, match := range matches {
		cleanedValues = append(cleanedValues, match[1])
	}
	cleanedString := strings.Join(cleanedValues, ", ")
	return cleanedString
}

// Custom comparison function
func compare(a, b string) bool {
	// Check if both strings are equal, return false
	if a == b {
		return false
	}

	// If one string is "LearningObjectInfo.xml", it should be prioritized
	if a == LEARNING_OBJECTINFO_XML {
		return true
	} else if b == LEARNING_OBJECTINFO_XML {
		return false
	}

	// Extract the numeric part from strings
	aNumber := extractNumber(a)
	bNumber := extractNumber(b)

	// Compare the numeric parts
	return aNumber < bNumber
}

// Function to extract the numeric part from a string
func extractNumber(s string) int {
	// Split the string by '.xml' and take the first part
	parts := strings.Split(s, ".xml")
	num := parts[0]

	// Extract the numeric part from the string
	var number int
	fmt.Sscanf(num, "cat%d", &number)
	return number
}

func traverse(n *html.Node, correctAnswersText *string) {
	if n.Type == html.TextNode && strings.Contains(n.Data, CORRECT_ANSWERS_FINDER) {
		for c := n.NextSibling; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				*correctAnswersText += strings.TrimSpace(c.Data)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, correctAnswersText)
	}
}

func cleanString(input string) string {
	endIndex := strings.Index(input, "</p>")
	if endIndex != -1 {
		textPart := input[:endIndex]
		textPart = strings.TrimSpace(textPart)
		return textPart
	}
	return input
}
