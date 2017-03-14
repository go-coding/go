package main

import "fmt"

func SimplePanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic info is ", err)
		}
	}()
	panic("SimplePanicRecover function ed")
}

func main() {
	SimplePanicRecover()
	test := fmt.Sprintf("test ddd %d",2)
	fmt.Printf(test)
}
