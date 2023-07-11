#include <iostream>
#include <memory>
void foo(std::shared_ptr<int> i)
{
    (*i)++;
}
int main()
{
    // auto pointer = new int(10); // 非法, 不允许直接赋值
    // 构造了一个 std::shared_ptr
    auto pointer = std::make_shared<int>(10);
    foo(pointer);
    std::cout << *pointer << std::endl; // 11
    // 离开作用域前，shared_ptr 会被析构，从而释放内存
    return 0;
}
