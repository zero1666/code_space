#include <iostream>

int foo(int a, int b, int c) {
    std::cout << a + b + c << std::endl;
    return 0;
}
int main() {
    // 将参数b=1,c=2绑定到函数 foo 上
    // 但是使用 std::placeholders::_1 来对第一个参数进行占位
    auto bindFoo = std::bind(foo, std::placeholders::_1, 1,2);
    // 这时调用 bindFoo 时，只需要提供第一个参数即可
    bindFoo(1);
    //output : 4
}
