package main

import "fmt"

//func main(){
//	const a int = 5
//	const b int = 6
//	if b <= a {
//		fmt.Print(a > b)
//	} else {
//		fmt.Print("Ты лох")
//	}
//}
//func main(){
//	var a int
//	fmt.Scan(&a)
//	if a > 0{
//		fmt.Print("Число положительное")
//	} else if a < 0 {
//		fmt.Print("Число отрицательное")
//	} else {
//		fmt.Print("Ноль")
//	}
//}

func main() {
	var a int
	fmt.Scan(&a)
	e := a % 10
	s := a / 100
	d := (a%100 - e) / 10
	if (s != d && d != e) && s != e {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
