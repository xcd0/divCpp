package main

import (
	"log"
	"regexp"
	"strings"
)

// 関数の宣言と定義、関数定義の直前にあるコメントを抽出します
func extractFunctionDetailsAndComments(fileInfo *FileInfo) error {
	isInsideComment := false
	isInsideFunction := false
	currentFunction := &FunctionInfo{}

	lines := strings.Split(fileInfo.Content, "\n")
	functionCommentBefore := []string{}

	for i := 0; i < len(lines); {
		line := lines[i]
		trimed := ""                                                   // コメントを削除した文字列
		trimed, isInsideComment = RemoveComment(line, isInsideComment) // isInsideComment コメント内であるかどうかの状態を更新するはす
		log.Printf(">> line[%d] %v : %#v", i, isInsideFunction, lines[i])

		// 現在の行が関数宣言または関数定義の開始行かどうかを判定
		if isFunctionDeclarationOrDefinition(trimed) {
			// 関数宣言または定義の開始行の場合、新しい関数情報を作成
			currentFunction = &FunctionInfo{
				StartLine: i + 1, // 行番号は1から始まるため+1
			}
			isInsideFunction = true
		}

		// 関数直前のコメントを保持する。空行などがあれば破棄する。
		if !isInsideFunction {
			if isInsideComment {
				functionCommentBefore = append(functionCommentBefore, line)
			}
			if strings.Index(line, "//") == 0 {
				functionCommentBefore = append(functionCommentBefore, line)
			}
			if len(re_space.ReplaceAllString(line, "")) == 0 { // 空白かtabしかないような空行
				functionCommentBefore = []string{} // リセットする
			}
		}
		if isInsideFunction {
			currentFunction.FunctionCommentBefore = strings.Join(functionCommentBefore, "\n")
			functionCommentBefore = []string{} // リセットする
			log.Printf(">> currentFunction.FunctionCommentBefore >>: %#v", currentFunction.FunctionCommentBefore)
			// 関数内部の処理
			if strings.HasPrefix(line, "//") {
				// 行がコメント行の場合、関数コメントに追加
				currentFunction.FunctionCommentBefore += line + "\n"
			} else if strings.TrimSpace(line) == "}" {
				// 関数定義の終了行の場合、関数情報を追加して関数外部へ
				currentFunction.EndLine = i + 1 // 行番号は1から始まるため+1
				(*fileInfo).Functions = append((*fileInfo).Functions, *currentFunction)
				currentFunction = nil
				isInsideFunction = false
			} else {
				// その他の行は関数定義の一部として追加
				currentFunction.FunctionDefinitionBody += line + "\n"
			}
			// 現在の行が名前空間を含むかどうかを判定
			namespace := extractNamespace(line)
			if len(namespace) > 0 {
				currentFunction.Namespace = append(currentFunction.Namespace, namespace...)
			}
		} else {
			if isInsideComment {
				// コメント内なので関数外の場合の処理はコメントの蓄積以外不要なはず
			}
		}

		// 関数定義の開始行から終了行までの行数を数える
		if isInsideFunction {
			i++
		}

		// 関数定義の終了行を探す
		if isInsideFunction && strings.TrimSpace(line) == "{" {
			openBraceCount := 1
			closeBraceCount := 0
			i++
			for ; i < len(lines); i++ {
				line = lines[i]
				if strings.Contains(line, "{") {
					openBraceCount++
				}
				if strings.Contains(line, "}") {
					closeBraceCount++
				}
				if openBraceCount == closeBraceCount {
					break
				}
			}
		} else {
			i++
		}
	}

	return nil
}

var re_multi_line_comment *regexp.Regexp = regexp.MustCompile(`/\*.*?\*/`)
var re_space *regexp.Regexp = regexp.MustCompile(`/\s/`)

// https://go.dev/play/p/TaR0yKG46FS
// 文字列から/**/のコメントを削除する。boolの引数と戻り値は、1文字目をコメント内であると見做して処理するか、処理後の文字列の末尾がコメント内か、を表す。
func RemoveComment(input string, isFirstInnerComment bool) (string, bool) {
	//log.SetFlags(log.Ltime | log.Lshortfile)
	i := strings.Index(input, "//")
	if i != -1 {
		input = input[:i] // "//"があればそれ以降を削除する。
	}
	if isFirstInnerComment { // 1文字目をコメント内であると見做す場合
		input = "/*" + input // 先頭に/*をつける
	}
	isLastInnerComment := false
	for { // 行頭1文字目がコメントでないと見做して "/*/"を処理する。
		input = strings.Replace(input, "/*/", "/*", 1)             // 文字列内の /*/ を /* に置き換える。
		input = re_multi_line_comment.ReplaceAllString(input, " ") // 文字列内の /* aaaa */ を削除する。
		s := strings.Index(input, "/*")
		if s == -1 {
			break
		}
		e := strings.Index(input[:s], "*/")
		if s > e || e != -1 {
			input = input[:s]
		}

	}
	return input, isLastInnerComment
}

func isFunctionDeclarationOrDefinition(line string) bool {
	// カンマと括弧をカウントする変数を初期化
	commaCount := 0
	openParenCount := 0
	closeParenCount := 0
	insideComment := false

	for _, char := range line {
		switch char {
		case '/':
			if !insideComment && commaCount == 0 && openParenCount == closeParenCount {
				// '/'がコメントの開始である場合、コメントが開始したことを記録
				insideComment = true
			} else if insideComment && len(line) > 1 && line[0] == '*' && line[1] == '/' {
				// コメントが終了した場合、コメントが終了したことを記録
				insideComment = false
			}
		case '*':
			if insideComment && len(line) > 1 && line[0] == '/' && line[1] == '*' {
				// コメントが開始した場合、コメントが開始したことを記録
				insideComment = true
			}
		case '(':
			if !insideComment {
				// '('がコメント内ではない場合、開き括弧のカウントを増やす
				openParenCount++
			}
		case ')':
			if !insideComment {
				// ')'がコメント内ではない場合、閉じ括弧のカウントを増やす
				closeParenCount++
			}
		case ',':
			if !insideComment {
				// ','がコメント内ではない場合、カンマのカウントを増やす
				commaCount++
			}
		}
	}
	// 関数宣言または定義の条件を満たすかどうかを判定
	return commaCount > 0 && openParenCount > 0 && openParenCount == closeParenCount
}

func extractNamespace(line string) []string {
	// 名前空間を抽出するためのカスタム処理を実装する
	// ここでは単純な例として、空白で区切られた単語を名前空間として扱う
	parts := strings.Fields(line)
	var namespace []string
	for _, part := range parts {
		namespace = append(namespace, part)
		// カンマや括弧などで名前空間が終了する場合、ここで終了する
		if strings.Contains(part, ",") || strings.Contains(part, "(") || strings.Contains(part, ")") {
			break
		}
	}
	return namespace
}

/*
// 関数の宣言と定義、関数定義の直前にあるコメントを抽出します
func extractFunctionDetailsAndComments(fileContent string) (map[string]*FunctionInfo, error) {
	// C++ ファイルの内容から関数の宣言、定義、および関数定義の直前にあるコメントを抽出します。
	// この関数では、関数のシグネチャ（関数名と引数の型の組み合わせ）をキーとして、関数の情報を保持する FunctionInfo 構造体のインスタンスを値とするマップを作成します。
	// FunctionInfo 構造体には、関数の名前、引数、戻り値の型、定義が存在する行番号、および関数定義の直前にあるコメントなどの情報を格納します。

	// ファイルの内容から関数の宣言と定義を正規表現などを用いて抽出する

	// 抽出した関数定義の直前にあるコメントを抽出する

	// 抽出した関数の情報を FunctionInfo 構造体に格納する

	// 関数のシグネチャを作成する（関数名と引数の型の組み合わせを用いて）

	// 関数のシグネチャをキーとして、FunctionInfo 構造体のインスタンスをマップに格納する
}

*/
