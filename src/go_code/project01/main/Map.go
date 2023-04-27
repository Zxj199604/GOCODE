package main

import "fmt"

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
*/
type funcname struct {
	name string
}

func (func1 funcname) function1() {
	var countryCapitalMap1 map[string]string
	countryCapitalMap1 = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap1["France"] = "巴黎"
	countryCapitalMap1["Italy"] = "罗马"
	countryCapitalMap1["Japan"] = "东京"
	countryCapitalMap1["India"] = "新德里"
	for country := range countryCapitalMap1 {
		fmt.Println(country, "首都是", countryCapitalMap1[country])
	}
	for s, s2 := range countryCapitalMap1 {
		fmt.Println(s)  //key
		fmt.Println(s2) //value
	}
}

func (func2 funcname) function2() {
	//创建map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}
