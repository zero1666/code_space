#include <iostream>
#include <memory>

 template<typename T, typename ...Args>
 std::unique_ptr<T> make_unique( Args&& ...args ) {
   return std::unique_ptr<T>( new T( std::forward<Args>(args)... ) );
 }


struct Foo {
    Foo() { std::cout << "Con Foo::Foo" << std::endl; }
    ~Foo() { std::cout << "Des Foo::~Foo" << std::endl; }
    void foo() { std::cout << "Foo::foo" << std::endl; }
};

void f(const Foo &) {
    std::cout << "f(const Foo&)" << std::endl;
}

int main() {
    std::unique_ptr<Foo> p1(make_unique<Foo>());// 构造函数输出
    if (p1) p1->foo();// p1 不空, 输出
    {
        std::unique_ptr<Foo> p2(std::move(p1));
        f(*p2);// p2 不空, 输出
        if(p2) p2->foo(); // p2 不空, 输出
        if(p1) p1->foo();// p1 为空, 无输出
        p1 = std::move(p2);
        if(p2) p2->foo();// p2 为空, 无输出
        std::cout << "p2 被销毁" << std::endl;
    }
    if (p1) p1->foo();// p1 不空, 输出
    // Foo 的实例会在离开作用域时被销毁
}
