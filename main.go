package main

import (
	"fmt"
	"os"
)

func main() {
	// コマンドライン引数から C++ ファイルのパスを取得します
	args := os.Args[1:]

	// 引数が指定されていない場合はエラーメッセージを表示します
	if len(args) == 0 {
		fmt.Println("Error: No file paths provided")
		return
	}

	// 各ファイルパスに対して processCppFile 関数を呼び出します
	for _, filePath := range args {
		fmt.Printf("Processing file: %s\n", filePath)
		err := processCppFile(filePath)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Successfully processed file: %s\n", filePath)
		}
	}
}
