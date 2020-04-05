package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	s := []int{1, 1, 1}
	add(s)
	fmt.Println(s)
}

func add(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	      i++
	  }
	*/
	var s2 []int
	copy(s2, s)

	for i := range s2 {
		s2[i] += 1
	}
	fmt.Println(s2)
}

//func main() {
//	defer fmt.Println("defer main")
//	var user = os.Getenv("USER_")
//
//	go func() {
//		defer func() {
//			fmt.Println("defer caller")
//			if err := recover(); err != nil {
//				fmt.Println("recover success. err: ", err)
//			}
//		}()
//
//		func() {
//			defer func() {
//				fmt.Println("defer here")
//			}()
//
//			if user == "" {
//				panic("should set user env.")
//			}
//
//			// 此处不会执行
//			fmt.Println("after panic")
//		}()
//	}()
//
//	time.Sleep(100)
//	fmt.Println("end of main function")
//}
