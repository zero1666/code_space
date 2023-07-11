#include <iostream>
class Base {
    public:
        int v1;
        int v2;
        Base(){v1 = 1;}
        Base(int v):Base(){ //委托构造，调用Base()构造函数
            v2 = v;
        }
};
class Subclass : public Base{
    public:
        //把基类构造函数继承到派生类中，不再需要书写多个派生类构造函数来完成基类的初始化。
        using Base::Base; 
};

int main(){
    Base b {2};
    std::cout <<"Base: "<<  b.v1 << " " <<  b.v2 << std::endl;
    Subclass s {3};
    std::cout <<"Subclass: "<<  s.v1 << " " <<  s.v2 << std::endl;
    //output
    //Base: 1 2
    //Subclass: 1 3
}

