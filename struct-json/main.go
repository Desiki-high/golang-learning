package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	StudentId      string `json:"id"`
	StudentName    string `json:"name"`
	StudentClass   string `json:"class"`
	StudentTeacher string `json:"teacher"`
}

type StudentNoJson struct {
	StudentId      string
	StudentName    string
	StudentClass   string
	StudentTeacher string
}

type StudentWithOption struct {
	StudentId      string //默认使用原定义中的值
	StudentName    string `json:"studentName"`     // 解析（encode/decode） 的时候，使用 `name`，而不是 `Field`
	StudentClass   string `json:"class,omitempty"` // 解析的时候使用 `class`，如果值为空，就忽略它
	StudentTeacher string `json:"-"`               // 解析的时候忽略该字段。默认情况下会解析这个字段，因为它是大写字母开头
}

func main() {
	//NO.1 with json struct tag
	s := &Student{StudentId: "1", StudentName: "ding", StudentClass: "0903", StudentTeacher: "dean"}
	jsonString, _ := json.Marshal(s)
	fmt.Println(string(jsonString))
	//{"id":"1","name":"ding","class":"0903","teacher":"dean"}

	newStudent := new(Student)
	_ = json.Unmarshal(jsonString, newStudent)
	fmt.Printf("%+v\n", newStudent)
	//&{StudentId:1 StudentName:ding StudentClass:0903 StudentTeacher:dean}
	// Unmarshal 是怎么找到结构体中对应的值呢？比如给定一个 JSON key Filed，它是这样查找的：
	// 首先查找 tag 名字（关于 JSON tag 的解释参看下一节）为 Field 的字段
	// 然后查找名字为 Field 的字段
	// 最后再找名字为 FiElD 等大小写不敏感的匹配字段。
	// 如果都没有找到，就直接忽略这个 key，也不会报错。这对于要从众多数据中只选择部分来使用非常方便。

	//NO.2 without json struct tag
	so := &StudentNoJson{StudentId: "1", StudentName: "ding", StudentClass: "0903", StudentTeacher: "dean"}
	jsonStringO, _ := json.Marshal(so)

	fmt.Println(string(jsonStringO))
	//{"StudentId":"1","StudentName":"ding","StudentClass":"0903","StudentTeacher":"dean"}

	//NO.3 StudentWithOption
	studentWO := new(StudentWithOption)
	js, _ := json.Marshal(studentWO)

	fmt.Println(string(js))
	//{"StudentId":"","studentName":""}

	studentWO2 := &StudentWithOption{StudentId: "1", StudentName: "ding", StudentClass: "0903", StudentTeacher: "dean"}
	js2, _ := json.Marshal(studentWO2)

	fmt.Println(string(js2))
	//{"StudentId":"1","studentName":"ding","class":"0903"}
}
