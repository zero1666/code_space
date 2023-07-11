// g++ -std=c++11 main.cpp foo.o -o main
#include "foo.h"
#include <iostream>
#include <functional>

int main(){
    std::cout << "Result from C code: " << add(1, 2) <<std::endl;
    return 0;
}
