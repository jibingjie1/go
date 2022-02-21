package main

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy() *Data {
	// 实例化c为Data类型
	var c Data

	//返回函数局部变量地址
	return &c
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	println(Sunday, Monday, Tuesday)
}
