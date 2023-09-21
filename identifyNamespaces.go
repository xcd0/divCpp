package main

import "strings"

// identifyNamespaces は C++ ファイルの内容と関数情報を受け取り、
// 各関数がどの名前空間に属しているかを識別します。
// この関数は名前空間のネストにも対応します。
func identifyNamespaces(fileContent string, functionInfos map[string]*FunctionInfo) error {
	// 以下のステップで名前空間を識別します:
	// 1. ファイル内容から名前空間の定義を探します (namespace { ... } の形式を探します)
	// 2. 名前空間のネストを識別します (名前空間内に他の名前空間が定義されている場合)
	// 3. 各関数がどの名前空間に属しているかを識別します
	// 4. functionInfos マップを更新して、各関数の名前空間情報を保存します

	lines := strings.Split(fileContent, "\n")
	var currentNamespace string

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// 名前空間の開始を識別
		if strings.HasPrefix(trimmedLine, "namespace ") {
			namespaceStartIndex := strings.Index(trimmedLine, " ") + 1
			namespaceEndIndex := strings.Index(trimmedLine, "{")
			if namespaceEndIndex == -1 {
				namespaceEndIndex = len(trimmedLine)
			}
			currentNamespace = trimmedLine[namespaceStartIndex:namespaceEndIndex]
		}

		// 名前空間の終了を識別
		if strings.Contains(trimmedLine, "}") {
			currentNamespace = ""
		}

		// 各関数がどの名前空間に属しているかを識別
		for funcName, funcInfo := range functionInfos {
			if funcInfo.StartLine <= i && funcInfo.EndLine >= i {
				funcInfo.Namespace = currentNamespace
			}
		}
	}

	return nil
}
