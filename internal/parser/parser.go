package parser

import (
	"fmt"
	"strings"
)

const redundantBoundary = 2

// CountParts принимает текст письма в формате SMTP и возвращает количество партов
func CountParts(emailText string) int {
	boundary, countBoundary := getBoundary(emailText)

	if boundary == "" {
		return 0
	}

	return countBoundary
}

// ParseEmail принимает текст письма в формате SMTP и парсит его, возвращает количество партов и их содержимое
func ParseEmail(emailText string) (int, []string) {
	boundary, countBoundary := getBoundary(emailText)

	boundary = "--" + boundary

	result := make([]string, 0, countBoundary)

	for i := 0; i < countBoundary; i++ {
		_, after, _ := strings.Cut(emailText, boundary)

		emailText = after

		before, _, _ := strings.Cut(after, boundary)

		result = append(result, strings.Trim(before, "\n"))
	}

	return countBoundary, result
}

func getBoundary(emailText string) (string, int) {
	_, b, isHereBoundary := strings.Cut(emailText, "boundary=")

	var boundary string

	if isHereBoundary {
		boundary, _, _ = strings.Cut(b, "\r")
	} else {
		return "", 0
	}

	endBoundary := fmt.Sprintf("--%s--", boundary)

	if !strings.Contains(emailText, endBoundary) {
		return "", 0
	}

	countBoundary := strings.Count(emailText, boundary) - redundantBoundary

	return boundary, countBoundary
}
