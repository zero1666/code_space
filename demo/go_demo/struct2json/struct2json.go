package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Address 地址结构体（内层结构体）
type Address struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Street   string `json:"street"`
	ZipCode  string `json:"zip_code"`
}

// Person 人员结构体（外层结构体，包含 Address）
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Email   string  `json:"email"`
	Address Address `json:"address"` // 嵌套的 Address 结构体
}

func main() {
	// 创建一个包含嵌套结构体的实例
	person := Person{
		Name:  "张三",
		Age:   28,
		Email: "zhangsan@example.com",
		Address: Address{
			Province: "广东省",
			City:     "深圳市",
			Street:   "南山区科技园",
			ZipCode:  "518000",
		},
	}

	// 结构体转 JSON（格式化输出）
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalf("JSON 序列化失败: %v", err)
	}

	fmt.Println("结构体转 JSON:")
	fmt.Println(string(jsonData))

	// JSON 转结构体
	jsonStr := `{
  "name": "李四",
  "age": 35,
  "email": "lisi@example.com",
  "address": {
    "province": "北京市",
    "city": "北京市",
    "street": "朝阳区建国路",
    "zip_code": "100000"
  }
}`

	var newPerson Person
	err = json.Unmarshal([]byte(jsonStr), &newPerson)
	if err != nil {
		log.Fatalf("JSON 反序列化失败: %v", err)
	}

	fmt.Println("\nJSON 转结构体:")
	fmt.Printf("姓名: %s\n", newPerson.Name)
	fmt.Printf("年龄: %d\n", newPerson.Age)
	fmt.Printf("邮箱: %s\n", newPerson.Email)
	fmt.Printf("地址: %s %s %s (邮编: %s)\n",
		newPerson.Address.Province,
		newPerson.Address.City,
		newPerson.Address.Street,
		newPerson.Address.ZipCode)
}
