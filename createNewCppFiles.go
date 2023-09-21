package main

import (
	"fmt"
	"os"
)

func createNewCppFiles(fileInfo *FileInfo, functionInfos map[string]*FunctionInfo) error {
	// 元のファイル名を取得します (拡張子を除く)
	baseFileName := fileInfo.FileName[:len(fileInfo.FileName)-len(".cpp")]

	// 各関数について新しい .cpp ファイルを作成します
	for funcName, funcInfo := range functionInfos {
		// 新しい .cpp ファイルの名前を決定します
		newCppFileName := fmt.Sprintf("%s_%s.cpp", baseFileName, funcName)

		// 新しい .cpp ファイルを作成します
		newCppFile, err := os.Create(newCppFileName)
		if err != nil {
			return fmt.Errorf("could not create file %s: %v", newCppFileName, err)
		}
		defer newCppFile.Close()

		// 新しい .hpp ファイルをインクルードするコードを記述します
		hppFileName := fmt.Sprintf("%s_functions.hpp", baseFileName)
		includeDirective := fmt.Sprintf("#include \"%s\"\n", hppFileName)
		_, err = newCppFile.WriteString(includeDirective)
		if err != nil {
			return fmt.Errorf("could not write to file %s: %v", newCppFileName, err)
		}

		// 関数定義を記述します
		_, err = newCppFile.WriteString(funcInfo.Definition)
		if err != nil {
			return fmt.Errorf("could not write to file %s: %v", newCppFileName, err)
		}
	}

	return nil
}
