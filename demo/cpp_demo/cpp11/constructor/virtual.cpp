struct Base0 {
    virtual void foo(int);
};
struct SubClass: Base0 {
    virtual void foo(int) override; // 合法
    //virtual void foo(float) override; // 非法, 父类没有此虚函数
};

struct Base {
    virtual void foo(int)final;
};

struct SubClass1 final: Base {
}; // 合法

/*
struct SubClass2 : SubClass1 {
}; // 非法, SubClass1 已 final
*/

/*
struct SubClass3: Base {
    void foo() override; // 非法, foo 已 final
};
*/
struct SubClass4: Base {
    void foo() ; // 合法, 已经非重载版本 
};

int main(){
}

