#include <iostream>

// 值捕获
void learn_lambda_func_1() {
    int value_1 = 1;
    auto copy_value_1 = [value_1] {
        return value_1;
    };
    value_1 = 100;
    auto stored_value_1 = copy_value_1();
    std::cout<< "value_1: " << value_1 << " stored_value_1: " << stored_value_1 << std::endl;
    // 这时, stored_value_1 == 1, 而 value_1 == 100.
    // 因为 copy_value_1 在创建时就保存了一份 value_1 的拷贝
}

// 引用捕获
void learn_lambda_func_2() {
    int value_2 = 1;
    auto copy_value_2 = [&value_2] {
        return value_2;
    };
    value_2 = 100;
    auto stored_value_2 = copy_value_2();
    std::cout<< "value_2: " << value_2 << " stored_value_2: " << stored_value_2 << std::endl;
    // 这时, stored_value_2 == 100, value_1 == 100.
    // 因为 copy_value_2 保存的是引用
}

using foo = void(int);
void functional(foo f){
    f(1);
}


int main(){
    learn_lambda_func_1();
    learn_lambda_func_2();

    auto f = [](int value){
        std::cout << value << std::endl;
    };
    functional(f); // 传递闭包对象，隐式转换为 foo* 类型的函数指针
    f(2); //lambda 表达式调用
    return 0;
}

