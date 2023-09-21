package main

// identifyOverloadedFunctionRelations は関数のオーバーロード関係を特定する関数です。
// この関数では、関数の名前と引数の型を基に、オーバーロードされた関数を特定し、
// それらの関数をグループ化します。
func identifyOverloadedFunctionRelations(functionInfos map[string]*FunctionInfo) error {
	// 関数名をキーとして、その関数がオーバーロードされている関数のリストを保持するマップを作成します
	overloadedFunctions := make(map[string][]*FunctionInfo)

	// 各関数についてオーバーロード関係を調べます
	for _, funcInfo := range functionInfos {
		// ここでは簡略化のため、関数名だけを使用してオーバーロード関係を特定します
		// (実際には引数の型も考慮する必要があります)
		overloadedFunctions[funcInfo.Name] = append(overloadedFunctions[funcInfo.Name], funcInfo)
	}

	// オーバーロード関係を特定した後、それを FunctionInfo 構造体に格納します
	for _, funcInfos := range overloadedFunctions {
		// オーバーロードされた関数が1つしかない場合は、オーバーロード関係は存在しません
		if len(funcInfos) < 2 {
			continue
		}

		// オーバーロード関係にある関数をグループ化します
		for _, funcInfo := range funcInfos {
			funcInfo.OverloadedFunctions = funcInfos
		}
	}

	return nil
}
