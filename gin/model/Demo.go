package model

type Demo struct {
	DemoText   string      `json:"demoText" binding:"required"` //json传输过程需要将key值转成小写 并且要求一定不为空
	DemoInt    int         `json:"demoInt,string"`              //将int类型在输出json时转sting
	DemoOption interface{} `json:"demoOption,omitempty"`        //omitempty表示可选，当value为空是不输出
}

func DemoSuccess() Demo {
	DemoSuccess := Demo{
		DemoText: "接受成功",
		DemoInt:  200,
	}
	return DemoSuccess
}
