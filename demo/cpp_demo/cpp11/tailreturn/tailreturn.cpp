#include <iostream>

//after C++11
template<typename T, typename U>
auto add (T x, U y) -> decltype(x+y) {
    return x + y;
}

int main(){
    auto w = add<int, double>(1, 2.0);
    if(std::is_same<decltype(w), double>::value){
        std::cout << "w is double : ";
    }
    std::cout << w << std::endl;

    auto v = add(1.0, 5);
    if(std::is_same<decltype(v), double>::value){
        std::cout << "v is double : ";
    }
    std::cout << v << std::endl;
    //output
    //w is double : 3
    //v is double : 6
}

