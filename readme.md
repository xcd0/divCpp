# 概略
与えられたcppファイルに含まれる関数群を、複数のファイルに分割する。

## 仕様

### 入力
* C++のソースコードが記述された任意個数の .cpp ファイル

### 出力
* 入力ファイル1つにつき、それぞれ下記を生成する。
	* 分割された .cpp ファイル群  
		元の .cpp ファイル内にある各関数を、それぞれ独立した .cpp ファイルに分割して出力します。ただし、1つの関数からしか呼ばれていない関数は個別のファイルを作成せず、呼び出し元の関数と同じ .cpp ファイルに記述します。また, オーバーロードされた関数は同一の.cppファイルにまとめます。各ファイルの名前は "元ファイル名_関数名.cpp" とします。また, これらの .cpp ファイルは新しく生成されたヘッダーファイルをインクルードします。
	* ヘッダーファイル  
		元の .cpp ファイルに存在する #include 文とマクロ定義、および分割した各関数の宣言を含む .hpp ファイルを作成します。このファイルの名前は "元ファイル名_functions.hpp" とします。また, 元ファイル内の関数定義の直前に記述されているコメントを、関数宣言の直前にも記述します。

### 処理フロー
1. 元の .cpp ファイルを読み込みます。
2. #include 文とマクロ定義を抽出します。
3. 各関数の宣言と定義を抽出します。この際、関数定義の直前に記述されているコメントも抽出します。
4. 名前空間を適切に識別し、関数がどの名前空間に属しているかを特定します。この際、名前空間がネストしている場合も適切にハンドリングします。
5. 各関数の呼び出し関係を解析し、1つの関数からしか呼ばれていない関数を特定します。また、関数のオーバーロード関係も特定します。
6. 新しい .hpp ファイルを作成し、抽出した #include 文、マクロ定義、および関数宣言を記述します。関数宣言の直前には、抽出したコメントも記述します。
7. 各関数について、新しい .cpp ファイルを作成し、関数定義と新しい .hpp ファイルを #include するコードを記述します。ただし、1つの関数からしか呼ばれていない関数は、呼び出し元の関数と同じ .cpp ファイルに記述します。また, オーバーロードされた関数は同一の.cppファイルにまとめます。この際、関数が属していた名前空間を保持するようにします。名前空間がネストしている場合も、それを保持し、適切に記述します。

## 必要な機能:
4.1 ファイル読み込み機能
4.1.1 ファイルオープン: 指定された .cpp ファイルを開く機能
4.1.2 ファイル読み込み: ファイルの内容を読み込む機能
4.2 テキスト解析機能
4.2.1 #include 文抽出: #include 文を抽出する機能
4.2.2 マクロ定義抽出: マクロ定義を抽出する機能
4.2.3 関数宣言と定義抽出: 関数の宣言と定義を抽出する機能
4.2.4 コメント抽出: 関数定義の直前にあるコメントを抽出する機能
4.3 名前空間解析機能
4.3.1 名前空間識別: 名前空間を識別する機能
4.3.2 名前空間関連付け: 関数を適切な名前空間に関連付ける機能
4.4 関数呼び出し解析機能
4.4.1 呼び出し関係解析: 各関数の呼び出し関係を解析する機能
4.4.2 単一呼び出し関数特定: 1つの関数からしか呼ばれていない関数を特定する機能
4.4.3 オーバーロード関係特定: 関数のオーバーロード関係を特定する機能
4.5 ファイル書き出し機能
4.5.1 .hpp ファイル作成: 新しい .hpp ファイルを作成する機能
4.5.2 .cpp ファイル作成: 新しい .cpp ファイルを作成する機能
4.5.3 コード記述: 必要なコードを .cpp と .hpp ファイルに記述する機能
