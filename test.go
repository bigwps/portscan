package main

import "fmt"

func main(){
	var res []int = make([]int, 10)

	for i := 0; i < len(res); i++{
		res[i] = 5 * i
	}

	for i := 0;i < len(res); i++{
		fmt.Printf("res at %d is %d",i , res[i])
	}
	fmt.Printf("the len of res is %d",len(res))
	fmt.Printf("the cap of res is %d",cap(res))

}