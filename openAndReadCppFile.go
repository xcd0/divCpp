package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// openAndReadCppFile は指定されたファイルパスから C++ ファイルの内容を読み込み、FileInfo 構造体を作成します。
func openAndReadCppFile(filePath string) (*FileInfo, error) {
	// ファイルが存在するかどうかを確認します
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %v", err)
	}

	// ファイルの内容を読み込みます
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	return &FileInfo{
		FilePath: filePath,
		Content:  string(content),
	}, nil
}
