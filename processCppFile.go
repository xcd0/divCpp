package main

import "fmt"

// processCppFile は指定された C++ ファイルを処理する関数です。
func processCppFile(filePath string) error {
	// 1. ファイルを開く
	// openAndReadCppFile は指定されたファイルパスから C++ ファイルの内容を読み込み、FileInfo 構造体を作成します。
	fileInfo, err := openAndReadCppFile(filePath)
	if err != nil {
		return err
	}

	fileContent := "" // この変数には readCppFileContent 関数で読み込んだファイルの内容を格納します

	// 2. #include 文とマクロ定義を抽出する
	includeDirectives, err := extractDirectives(fileContent, "#include")
	if err != nil {
		return err
	}
	macroDefinitions, err := extractDirectives(fileContent, "#define")
	if err != nil {
		return err
	}

	// 3. 関数の宣言と定義を抽出する
	functionInfos, err := extractFunctionDeclarationsAndDefinitions(fileContent)
	if err != nil {
		return err
	}

	// 5. 関数の宣言と定義、関数定義の直前にあるコメントを抽出する
	err = extractCommentsBeforeFunctionDefinitions(fileContent, functionInfos)
	if err != nil {
		return err
	}

	// 6. 名前空間を識別する
	err = identifyNamespaces(fileContent, functionInfos)
	if err != nil {
		return err
	}

	// 7. 関数を適切な名前空間に関連付ける
	namespaceInfos := make(map[string]*NamespaceInfo) // この変数には identifyNamespaces 関数で識別した名前空間の情報を格納します

	err = associateFunctionsToAppropriateNamespaces(functionInfos, namespaceInfos)
	if err != nil {
		return err
	}

	// 8. 関数の呼び出し関係を解析する
	err = analyzeFunctionCallRelations(functionInfos)
	if err != nil {
		return err
	}

	// 9. 1つの関数からしか呼ばれていない関数を特定する
	err = identifySingleCallFunctions(functionInfos)
	if err != nil {
		return err
	}

	// 10. 関数のオーバーロード関係を特定する
	err = identifyOverloadedFunctionRelations(functionInfos)
	if err != nil {
		return err
	}

	// 11. 新しい .hpp ファイルを作成する
	err = createNewHeaderFile(fileInfo, functionInfos)
	if err != nil {
		return err
	}

	// 12. 新しい .cpp ファイルを作成する
	err = createNewCppFiles(fileInfo, functionInfos)
	if err != nil {
		return err
	}

	fmt.Printf("Processing completed successfully : %s\n", filePath)
	return nil
}
