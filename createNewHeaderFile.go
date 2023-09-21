package main

import (
	"fmt"
	"os"
)

func createNewHeaderFile(fileInfo *FileInfo, functionInfos map[string]*FunctionInfo) error {
	// 新しい .hpp ファイルの名前を生成します
	headerFileName := fmt.Sprintf("%s_functions.hpp", fileInfo.FileNameWithoutExtension)

	// 新しい .hpp ファイルを作成します
	file, err := os.Create(headerFileName)
	if err != nil {
		return fmt.Errorf("could not create header file: %w", err)
	}
	defer file.Close()

	// #include 文とマクロ定義を .hpp ファイルに記述します
	for _, includeDirective := range fileInfo.IncludeDirectives {
		_, err := file.WriteString(includeDirective + "\n")
		if err != nil {
			return fmt.Errorf("could not write to header file: %w", err)
		}
	}

	for _, macroDefinition := range fileInfo.MacroDefinitions {
		_, err := file.WriteString(macroDefinition + "\n")
		if err != nil {
			return fmt.Errorf("could not write to header file: %w", err)
		}
	}

	// 関数宣言と関連するコメントを .hpp ファイルに記述します
	for _, functionInfo := range functionInfos {
		// 関数定義の直前にあるコメントを記述します
		if functionInfo.CommentBefore != "" {
			_, err := file.WriteString(functionInfo.CommentBefore + "\n")
			if err != nil {
				return fmt.Errorf("could not write to header file: %w", err)
			}
		}

		// 関数宣言を記述します
		_, err := file.WriteString(functionInfo.Declaration + ";\n")
		if err != nil {
			return fmt.Errorf("could not write to header file: %w", err)
		}
	}

	fmt.Printf("Successfully created header file: %s\n", headerFileName)
	return nil
}
