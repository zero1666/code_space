#include <iostream>
#include <vector>

template<typename T, typename U>
class MagicType {
public:
    T dark;
    U magic;
};

typedef int (*process)(void *);
using NewProcess = int(*)(void *); // 跟typedef 一样，生成函数的别名

/*
// 不合法
template<typename T>
typedef MagicType<std::vector<T>, std::string> FakeDarkMagic; //typedef 不支持模版别名 
 * */
template<typename T>
using TrueDarkMagic = MagicType<std::vector<T>, std::string>; //新功能，生成模版类型别名

int main() {
    TrueDarkMagic<bool> you;
}
