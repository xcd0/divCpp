package main

import "strings"

func analyzeFunctionCallRelations(functionInfos map[string]*FunctionInfo) error {
	// ここでは、各関数の呼び出し関係を解析します。
	for callerName, callerInfo := range functionInfos {
		for calleeName, calleeInfo := range functionInfos {
			// 呼び出し関係を解析する簡易的な方法として、関数の本体に他の関数の名前が含まれているかを確認します。
			if callerName != calleeName && strings.Contains(callerInfo.FunctionDefinitionBody, calleeName) {
				calleeInfo.CalledBy = append(calleeInfo.CalledBy, callerName)
			}
		}
	}
	return nil
}

func identifySingleCallFunctions(functionInfos map[string]*FunctionInfo) error {
	// ここでは、1つの関数からしか呼ばれていない関数を特定します。
	for _, funcInfo := range functionInfos {
		if len(funcInfo.CalledBy) == 1 {
			funcInfo.IsSingleCall = true
		}
	}
	return nil
}
