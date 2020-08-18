package main

import "fmt"

/**
	case后不需要手动添加break，默认只会执行一个case。

 	switch expr {
	case val1:

	case val2:
	case val3,val4:

	}

 	switch {
	case 条件1:

	case 条件2:
	case 条件3,条件4:	// 满足一个条件即可执行

	}


*/
func main() {
	var name string = "benny"
	switch name {
	case "benny":
		fmt.Printf("hello %s\n ", name)

	}

	switch expr {

	}

	var score = 92
	var grade string
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80 && score < 90:
		grade = "B"
	case score >= 60 && score < 80:
		grade = "C"
	case score < 60:
		grade = "D"
	}

	fmt.Println(grade)

}
