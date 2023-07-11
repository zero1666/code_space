#include <iostream>
#include<string>

void reference(std::string& str){
    std::cout<< "左值" << std::endl;
}

void reference(std::string&& str){
    std::cout << "右值" <<std::endl;
}

int main(){
    std::string lv1 = "string, "; //lv1 左值
    //std::string && r1 = lv1; // error, 右值不能直接引用左值
    std::string&& rv1 = std::move(lv1); //ok,  std::move() 将左值转换为右值引用   
    std::cout << rv1 << std::endl; //string, 

    const std::string& lv2 = lv1 + lv1; //ok, 常量左值引用可延长临时变量的声明周期
    //    lv2 += "test";// error, 常量引用无法修改
    std::cout << lv2 << std::endl;

    std::string&& rv2 = lv1 + lv2;// ok, 右值引用延长临时对象生命周期
    rv2 += "Test" ; //ok,  右值引用依然是左值，非常量引用可修改临时变量
    std::cout << rv2 << std::endl;

    reference(rv2);// 输出左值
    reference(std::move(rv2));// 输出右值
    reference(std::move("test"));// 输出右值
}
