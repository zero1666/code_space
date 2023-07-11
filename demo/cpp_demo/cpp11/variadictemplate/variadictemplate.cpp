//g++ -std=c++11 variadictemplate.cpp -o out
#include <iostream>
#include <map>
#include <vector>
//变长类模版参数：C++11允许任意个数、类别的模板参数
template <typename... ts> class Magic{};

// 推荐，手动的定义至少一个模板参数,防止产生参数个数为0的情况
template<typename Require, typename... Args> class Magic1{};

// 变长函数模版参数,使用 ... 表示不定长参数
template<typename... Args> 
void magic( Args... args){
    //使用 sizeof... 来计算参数的个数
    std::cout<< sizeof...(args) <<std::endl; 
}

//变长模版参数解包 Example: 递归模版函数形式, 
//前提：变长部分需要保证至少有一个模版参数，另需要定义一个终止递归的函数
template<typename T0>
void printf(T0 value){
    std::cout << value << std::endl;
}
template<typename T, typename... Args>
void printf(T value, Args... args){
    std::cout << value << std::endl;
    printf(args...);
}

int main(){
    class Magic<int,
            std::vector<int>,
            std::map<std::string, std::vector<int>>> darkMagic;//ok
    class Magic<> nothing; //ok, 个数为0的模板参数也是可以的,但不建议这样使用
    class Magic1<int> oneMagic; // ok,这种情况下，至少需要有一个模版参数

    magic();
    magic(1);// 1
    magic(1, " "); //  2

    // C++11  使用递归模版函数对 变长参数进行解包
    std::cout << "[C++11  使用递归模版函数对 变长参数进行解包]" << std::endl;
    printf(1, 2, "testtemplate", 1.1);
    //output
    //0
    //1
    //2
    //[C++11  使用递归模版函数对 变长参数进行解包]
    //1
    //2
    //testtemplate
    //1.1
}
