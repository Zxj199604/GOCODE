/**
 * @Author: Zxj
 * @Description:
 * @File: new_make
 * @Version: 1.0.0
 * @Date: 2022/11/11 11:30 下午
 * @Software : GoLand
 */

package main

import (
	"fmt"
)

type SyncedBuffer struct {
	str string
	num int
	sum int
}

func main() {
	// 普通赋值 m 变量输出： { 0 0}
	m := SyncedBuffer{}
	fmt.Println("普通赋值 m 变量输出：", m)

	// new m1 变量输出： &{ 0 0}
	m1 := new(SyncedBuffer)
	fmt.Println("new m1 变量输出：", m1)
	fmt.Println("m1 的地址", &m1)
	fmt.Println("m1的值", *m1)

	// &Type赋值 m2 变量输出： &{ 0 0}
	m2 := &SyncedBuffer{}
	fmt.Println("&Type赋值 m2 变量输出：", m2)

	// &Type赋值 带有初始值的m3 变量输出： &{m3Test 123 44}
	m3 := &SyncedBuffer{"123", 123, 44}
	fmt.Println("&Type赋值 带有初始值的m3 变量输出：", m3)

	//  生成一个int类型切片，初始长度：3，容量：5
	Slice := make([]int, 3, 5)
	fmt.Println(Slice)
	// 生成一个键string值int类型的字典，长度：5
	Map := make(map[string]int,5)
	fmt.Println(Map)
	// 生成一个int类型的通道，长度：5
	Chan := make(chan int, 5)
	fmt.Println(Chan)

}
