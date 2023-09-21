package main

import (
	"fmt"
	"log"
)

// processCppFile は指定された C++ ファイルを処理する関数です。
func processCppFile(filePath string) error {
	// 1. ファイルを開く
	// openAndReadCppFile は指定されたファイルパスから C++ ファイルの内容を読み込み、FileInfo 構造体を作成します。
	fileInfo, err := openAndReadCppFile(filePath)
	if err != nil {
		return err
	}
	//log.Printf("fileInfo : %v", fileInfo)
	log.Printf("--------------------------------------------------------------------------------")
	log.Printf("fileInfo : %v", fileInfo.Print())
	log.Printf("--------------------------------------------------------------------------------")

	// 2. #include 文とマクロ定義を抽出する
	includeDirectives, err := extractDirectives(fileInfo.Content, "#include")
	if err != nil {
		return err
	}
	fileInfo.IncludeDirectives = includeDirectives

	macroDefinitions, err := extractDirectives(fileInfo.Content, "#define")
	if err != nil {
		return err
	}
	fileInfo.MacroDefinitions = macroDefinitions
	log.Printf("FileInfo : %v", fileInfo.Print())
	log.Printf("--------------------------------------------------------------------------------")

	// 5. 関数の宣言と定義、関数定義の直前にあるコメントを抽出する
	// func extractFunctionDetailsAndComments(fileInfo *FileInfo) error
	if err := extractFunctionDetailsAndComments(fileInfo); err != nil {
		return err
	}
	for i, fi := range fileInfo.Functions {
		log.Printf("FunctionInfo[%d] : %s", i, fi.Print())
	}
	log.Printf("--------------------------------------------------------------------------------")

	// // 6. 名前空間を識別する
	// var namespaceInfos map[string]*NamespaceInfo
	// err = identifyNamespaces(fileInfo.Content, functionInfos)
	// if err != nil {
	// 	return err
	// }

	// 7. 関数を適切な名前空間に関連付ける
	//namespaceInfos := make(map[string]*NamespaceInfo) // この変数には identifyNamespaces 関数で識別した名前空間の情報を格納します

	// // func associateFunctionsToAppropriateNamespaces(functionInfos map[string]*FunctionInfo, namespaceInfos map[string]*NamespaceInfo) error {
	// err = associateFunctionsToAppropriateNamespaces(functionInfos, namespaceInfos)
	// if err != nil {
	// 	return err
	// }

	// // 8. 関数の呼び出し関係を解析する
	// //func analyzeFunctionCallRelations(functionInfos map[string]*FunctionInfo) error {
	// err = analyzeFunctionCallRelations(functionInfos)
	// if err != nil {
	// 	return err
	// }

	// // 9. 1つの関数からしか呼ばれていない関数を特定する
	// err = identifySingleCallFunctions(functionInfos)
	// if err != nil {
	// 	return err
	// }

	// // 10. 関数のオーバーロード関係を特定する
	// err = identifyOverloadedFunctionRelations(functionInfos)
	// if err != nil {
	// 	return err
	// }

	// // 11. 新しい .hpp ファイルを作成する
	// err = createNewHeaderFile(fileInfo, functionInfos)
	// if err != nil {
	// 	return err
	// }

	// // 12. 新しい .cpp ファイルを作成する
	// err = createNewCppFiles(fileInfo, functionInfos)
	// if err != nil {
	// 	return err
	// }

	fmt.Printf("Processing completed successfully : %s\n", filePath)
	return nil
}
