/**
 * @Author: Zxj
 * @Description:
 * @File: Select_test.go
 * @Version: 1.0.0
 * @Date: 2022/2/17 10:31 下午
 * @Software : GoLand
 */

package project02

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test_Selectfun1(t *testing.T) {
	size := 10
	ch := make(chan int, size)
	for i := 0; i < size; i++ {
		ch <- 1
	}
	ch2 := make(chan int, size)
	for i := 0; i < size; i++ {
		ch2 <- 2
	}
	ch3 := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case v := <-ch:
			fmt.Printf("%s:%v\n", "ch1", v)
		case b := <-ch2:
			fmt.Printf("%s:%v\n", "ch2", b)
		case ch3 <- 10:
			fmt.Print("write\n")
		default:
			fmt.Println("none")
		}
	}
}
func Test_selectfun2(t *testing.T) {
	//超时用法
	ch := make(chan int)
	go func(c chan int) {
		// 修改时间后,再查看执行结果
		time.Sleep(time.Second * 1)
		c <- 1
	}(ch)
	select {
	case v := <-ch:
		fmt.Print(v)
	case <-time.After(2 * time.Second): // 等待 2s
		fmt.Println("no case ok")
	}
}
func Test_selectfun3(t *testing.T) {
	// 空select{} 会死锁
	//select {}
	fmt.Println(runtime.NumCPU())
	quit := make(chan bool)
	for i := 0; i != runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-quit:
					fmt.Printf("case\n")
					//break
				default:
					fmt.Printf("default\n")
				}
			}
		}()
	}
	time.Sleep(time.Second * 35)
	//for i := 0; i != runtime.NumCPU(); i++ {
	//	quit <- true
	//}
}
