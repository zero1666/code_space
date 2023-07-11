// initlist.cpp
// g++ -std=c++11 initlist.cpp -o out
#include <iostream>
#include <vector>
#include <map>
#include <initializer_list>

using namespace std;

template<typename T>
void printvector(const string& name, std::vector<T>& v){
    std::cout << "print " << name << ": " <<std::endl;
    for(auto it = v.begin(); it != v.end(); ++it){
        std::cout << *it << " ";
    }
    std::cout <<std::endl;
}

class MagicFoo {
    public:
        MagicFoo(){} 
        MagicFoo(std::initializer_list<string> list) {
            for (auto it = list.begin(); it != list.end(); ++it){
                vec.push_back(*it);
            }
        }

    public:
        std::vector<std::string> vec  {"magic", "default"}; // ok, deault value
        std::vector<std::string> vec2 = {"magic2", "default2"}; // ok, deault value
        int int1 = (5); // ok, deault value
        const static  int cs_int1 = 1; //ok, static const 常量初始化依然可以用
        static const int cs_int2 {2}; //ok, static const 常量初始化依然可以用
       // static int s_int2 {3}; // error,类静态变量初始化必须在类外
        static int s_int1; // ok, 类静态变量初始化必须在类外
        static int s_int2; // ok, 类静态变量初始化必须在类外
};

int MagicFoo::s_int1 = 1; // 类静态变量初始化
int MagicFoo::s_int2 {2}; // 类静态变量初始化


int main(){
   // std::vector<int> v1 {1, 3, 5}; // ok， 初始化列表初始化
    std::vector<int> v1 = {1, 3, 5}; // ok, 与上面等价
    std::vector<std::string> v2  {"test1", "test2", "test3"}; // ok, 支持所有基本类型
    int arr[] {1,3}; //ok
    std::map<int ,string> m1 =  {{3,"value3"}, {2,"value2"}}; //ok
    printvector("v1", v1);
    printvector("v2", v2);

    MagicFoo magicFoo; 
    printvector("magicFoo.vec", magicFoo.vec);
    printvector("magicFoo.vec2", magicFoo.vec2);
    std::cout << magicFoo.int1 <<endl;

    MagicFoo magicFoo2 = {"foo2", "foo2"};
    printvector("magicFoo2.vec", magicFoo2.vec); // 注意，vec = magic default foo2 foo2
    std::cout <<"class const static:" << MagicFoo::cs_int1 << " " << MagicFoo::cs_int2 << std::endl;
    std::cout <<"class static: " << MagicFoo::cs_int1 << " " << MagicFoo::cs_int2 << std::endl;
}
