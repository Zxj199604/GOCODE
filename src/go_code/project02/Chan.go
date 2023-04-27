package project02

import (
	"fmt"
	"time"
)

func function1() {

	s := []int{1, 2, 3, 4, 5, 6, 7}
	c := make(chan int)
	go sum(s[len(s)-1:], c)
	time.Sleep(5000)
	go sum(s[len(s)-2:], c)
	time.Sleep(5000)
	go sum(s[len(s)-3:], c)
	time.Sleep(5000)
	go sum(s[len(s)-4:], c)
	time.Sleep(5000)
	go sum(s, c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func function2() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为3
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func function3() {
	//遍历通道和关闭通道
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}



func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}
