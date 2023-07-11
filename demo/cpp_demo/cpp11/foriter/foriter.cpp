#include <iostream>
#include <vector>
#include <algorithm>

int main(){
    std::vector<int> vec  {1,2,3,4};
    auto itr = std::find(vec.begin(), vec.end(), 3);
    if ( itr != vec.end()){
        *itr = 8;
    }

    for (auto element: vec){
        std::cout << element << " " ;// 只读，不能修改数组值
    } 
    std::cout << std::endl;

    for (auto& element: vec){
        element += 1; // 可读写，能修改数组值
    } 
    for (auto element: vec){
        std::cout << element <<" ";// 只读，不能修改数组值
    } 
    std::cout << std::endl;
    //output:
    //1 2 8 4
    //2 3 9 5
}
