package helper

import (
	"bufio"
	"os"
	"strings"
)

func GetModuleName() string {
	file, err := os.Open("go.mod")
	if err != nil {
		return "github.com/zaheershaikh936/gogen" // Fallback
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return "github.com/zaheershaikh936/gogen" // Fallback
}
