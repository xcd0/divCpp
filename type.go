package main

import "fmt"

type FunctionInfo struct {
	FullName  string // 完全修飾関数名。名前空間、クラス名、関数名、および関数のシグネチャ（引数の型と数）を含む。形式: "Namespace1::Namespace2::ClassName::FunctionName(ArgType1, ArgType2, ...)"
	Signature string // 関数のシグネチャ。引数の型と数を表す。形式: "ArgType1, ArgType2, ..."

	FunctionCommentBefore  string   //関数の直前に記述されたコメント。Doxygenコメントも含む。
	FunctionDeclaration    string   //関数の宣言。
	FunctionDefinitionBody string   // 関数の本体。関数の開始から終了までのコードを含む。
	StartLine              int      // 関数定義が開始する行番号。
	EndLine                int      // 関数定義が終了する行番号。
	Namespace              []string // 関数が属する名前空間のリスト。ネストした名前空間を考慮し、最も外側の名前空間から順にリストに含める。

	Overloads    []string // この関数とオーバーロード関係にある他の関数の完全修飾名のリスト。
	CalledBy     []string // この関数を呼び出している他の関数の完全修飾名のリスト。
	Calls        []string // この関数から呼び出される他の関数の完全修飾名のリスト。
	IsSingleCall bool     // この関数が1つの関数からしか呼ばれていないかどうかを示すフラグ
}

type FileInfo struct {
	FilePath                 string                   // 元のC++ファイルパス
	Content                  string                   // 内容
	FileName                 string                   // 元のC++ファイルの名前。
	FileNameWithoutExtension string                   // 元のC++ファイルの名前の拡張子を除いた文字列。
	IncludeDirectives        []string                 // 元のC++ファイルに含まれる#includeディレクティブのリスト。
	MacroDefinitions         []string                 // 元のC++ファイルに含まれるマクロ定義のリスト。
	FunctionMap              map[string]*FunctionInfo // 完全修飾関数名をキーとして関数情報を保持するマップ。これにより、各関数の詳細情報を迅速に取得できる。
	Functions                []FunctionInfo           // 各関数についての情報を保持
}

type NamespaceInfo struct {
	Name   string         // 名前空間の名前。
	Parent *NamespaceInfo // 親名前空間へのポインタ。ネストした名前空間を表現するために使用する。
}

func (fi *FunctionInfo) Print() string {
	return fmt.Sprintf(`
	: FullName                     : %v
	: Signature                    : %v
	: FunctionCommentBefore        : %v
	: FunctionDeclaration          : %v
	: FunctionDefinitionBody len   : %v
	: StartLine                    : %v
	: EndLine                      : %v
	: Namespace                    : %v
	: Overloads                    : %v
	: CalledBy                     : %v
	: Calls                        : %v
	: IsSingleCall                 : %v`,
		fi.FullName, fi.Signature, fi.FunctionCommentBefore, fi.FunctionDeclaration,
		len(fi.FunctionDefinitionBody),
		fi.StartLine, fi.EndLine, fi.Namespace, fi.Overloads, fi.CalledBy, fi.Calls, fi.IsSingleCall)
}

func (fi *FileInfo) Print() string {
	return fmt.Sprintf(`
	: FilePath                 : %v
	: Content len              : %v
	: FileName                 : %v
	: FileNameWithoutExtension : %v
	: IncludeDirectives        : %v
	: MacroDefinitions         : %v
	: Functions                : %v
`, fi.FilePath, len(fi.Content), fi.FileName, fi.FileNameWithoutExtension, fi.IncludeDirectives, fi.MacroDefinitions, fi.Functions)
}
