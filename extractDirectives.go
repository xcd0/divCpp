package main

import (
	"bufio"
	"strings"
)

// extractDirectives は指定されたプレフィックスで始まるディレクティブを抽出する関数です。
// この関数は #include または #define などのプレフィックスを受け取り、
// それで始まる行を抽出して返します。
func extractDirectives(fileContent string, prefix string) ([]string, error) {
	var directives []string
	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, prefix) {
			directives = append(directives, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return directives, nil
}
