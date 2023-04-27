package project02

import (
	"fmt"
)

func Switchfunc1() {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好!\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("您的等级是 %s\n", grade)
}
func Switchfunc2() {
	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型
	var x interface{}
	x = false
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型:%T", i)
	}
}

func Switchfunc3() {
	var match int = 1
	//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	//switch 从第一个判断表达式为 true 的 case 开始执行，如果 case 带有 fallthrough，程序会继续执行下一条 case，且它不会去判断下一个 case 的表达式是否为 true。
	switch match {
	case 1:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case 2:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	default:
		fmt.Println("3、默认 case")
	}
}

func Switchfunc4() {
	switch 'a' {
	case 'a':
		//fmt.Println("a")
	case 'b':
		fmt.Println("b")
	default:
		fmt.Printf("default")
	}
}
