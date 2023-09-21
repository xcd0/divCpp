#include <iostream>

// bbbb

// aaaa
// aaaa
// aaaa
int func1();
int func2(){ // bbbb
	//}
	/*
	 * }
	 * */
	puts("");/*ccc*/
	puts("");/*{}*/
	puts("");/*}*/
	}
template<class T>
T func3(){
	return 0;
}
int func1();

// main comment
int main(){
	func1();
}
namespace aaaa{
	int func1();
	namespace aaaa2{
		int func1();
	}
	namespace {
		int func1();
	}
}

namespace {
	auto lambda1 = [](){
		puts("noname::lambda");
	};
	int func1(){
		puts("noname::func1");
		lambda1();
	}
}
namespace aaaa{
	// aaa::func1 comment
	int func1(){
		puts("aaa::func1");
	}
	namespace aaaa2{
		int func1(){ puts("aaa2::func1"); }
		int
		func2
		// }}}}}}
		(
		)
		{
		puts
		(
		"aaa2::func1"
		)
		;
		}// }}}}}}
	}
	namespace {
		auto lambda1 = [](){
			puts("aaaa::noname::lambda");
		};
		int func1(){
			puts("aaaa::noname::func1");
			lambda1();
		}
	}
}
namespace bbbb{
	// なにもない
}
