#include <iostream>

using namespace std;
class Person
{
public:
    Person(int v) {
        value = v;
        std::cout << "Cons" <<value<< std::endl;
    }
    ~Person() {
        std::cout << "Des" <<value<< std::endl;
    }
    int value;
};

int main()
{
    std::shared_ptr<Person> p1(new Person(1));// Person(1)的引用计数为1

    std::shared_ptr<Person> p2 = std::make_shared<Person>(2); // 引用计数1

    p1.reset(new Person(3));// 首先生成新对象，然后引用计数减1，引用计数变为0，故析构Person(1)
                            // 最后将新对象的指针交给智能指针

    std::shared_ptr<Person> p3 = p1;//现在p1和p3同时指向Person(3)，Person(3)的引用计数为2

    p1.reset();//Person(3)的引用计数减小1，变为1
    p3.reset();//Person(3)的引用计数减小1, 变为0，析构Person(3)
    return 0;
}
