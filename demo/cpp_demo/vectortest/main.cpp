#include <iostream>
#include <vector>

using namespace std;

int printVector(vector<int> & a1){
    for(auto iter = a1.begin(); iter != a1.end(); iter++){
        cout << *iter << endl;
    }
    return 0;
}

int main(){
    std::vector<int> arr1;
    arr1.push_back(1);
    arr1.push_back(2);
    //arr1.push_back(3);
    //arr1.push_back(4);
   // arr1.push_back(5);
   // arr1.push_back(6);
    //arr1.push_back(7);
   // arr1.push_back(8);

    printVector(arr1);

    auto first = std::begin(arr1);
    std::cout << *first << endl;
    cout<< "after add ........" << endl;
    for (int i = 0; i <4 && first != arr1.end(); i++, ++first){
        cout << *first <<endl;
    }
    cout<< "result ........" << endl;
    arr1.insert(first, 100);
    printVector(arr1);

    return 0;
}


