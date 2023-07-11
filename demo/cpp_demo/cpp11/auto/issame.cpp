#include <iostream>
int main(){
    auto x = 3;
    auto y = 2.1;
    decltype(x + y) z = x + y;
    std::cout << z  << std::endl;
    if(std::is_same<decltype(z), float>::value){
        std::cout << "type z == float" << std::endl;
    }
    if(std::is_same<decltype(z), double>::value){
        std::cout << "type z == double" << std::endl;
    }
    //output
    //5.1
    //type z == double
}
