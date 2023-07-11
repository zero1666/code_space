#include <iostream>

enum class new_enum : unsigned int {
    value1,
    value2,
    value3 = 100,
    value4 = 100
};
// 枚举类声明的枚举获得枚举值时，必须显式进行类型转换，
// 可以通过重载 << 这个算符来进行输出
template<typename T>
std::ostream& operator<<(typename std::enable_if<std::is_enum<T>::value, std::ostream>::type& stream, const T& e)
{
    return stream << static_cast<typename std::underlying_type<T>::type>(e);
}

int main(){
    //相同枚举值之间如果指定的值相同，那么可以进行比较
    if (new_enum::value3 == new_enum::value4) {
        // 会输出
        std::cout << "new_enum::value3 == new_enum::value4" << std::endl;
    }
    std::cout << new_enum::value3 << std::endl;
}
