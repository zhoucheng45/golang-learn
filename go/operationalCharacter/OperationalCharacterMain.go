package main

import "fmt"

func main() {
	testIf()
}

func testOC() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int
	var p **int
	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)
	/*  & 和 * 运算符实例 */
	ptr = &a
	/* 'ptr' 包含了 'a' 变量的地址 */
	p = &ptr
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
	fmt.Printf("ptr 为 %d\n", ptr)
	fmt.Printf("p 为 %T\n", p)
	fmt.Printf("*p 为 %d\n", *p)
}

func testIf() {

	var age = 12
	if age > 18 {
		fmt.Println("已经成年")
	} else {
		fmt.Println("未成年")
	}
}
