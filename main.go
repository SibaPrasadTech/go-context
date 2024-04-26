package main

import (
	//"context"
	"fmt"
	"sync"
	"time"
	//"sync"
)

func main(){
	fmt.Println("Let's learn Golang contexts >>>>>>>")
	// var numberList []int
	// fmt.Println(numberList)
	// fmt.Println(len(numberList))
	// fmt.Println(numberList == nil)
	// //numberList[0] = 5
	// numberList = append(numberList, 6)
	// fmt.Println(numberList)
	// fmt.Println(len(numberList))
	// fmt.Println(numberList == nil)
	list := make([]int, 5)
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(list == nil)
	var chars [3]string = [...]string{"a","b","c"}
	fmt.Println(chars)

	// for i:=0;i<len(chars);i++ {
	// 	fmt.Println(chars[i])
	// }

	// for _,char := range(chars) {
	// 	fmt.Println(char)
	// }
	// var i = 0;
	// var k = len(chars);
	// for {
	// 	if (i>=k) {
	// 		break
	// 	}
	// 	fmt.Println(chars[i]);
	// 	i++;
	// }
	var wg sync.WaitGroup
	numOfItems := 3
	wg.Add(numOfItems)
	ch := make(chan int, numOfItems);
	go func(wg *sync.WaitGroup) {
		time.Sleep(time.Second);
		ch <- 10
		wg.Done()
		ch <- 6
		time.Sleep(time.Second*3);
		wg.Done()
		ch <- 15
		wg.Done()
		//close(ch) // this can be done by the main go routine
	}(&wg)
	wg.Wait()
	close(ch)
	for result := range ch {
		fmt.Println(result)
	}
	// x := <- ch // blocking
	// fmt.Println(x)
	// y := <- ch // blocking
	// fmt.Println(y)
	// z := <- ch // blocking
	// fmt.Println(z)
	//fmt.Println(numberList[0])
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	
	// generator := func(dataItem string, stream chan interface{}){
	// 	for {
	// 		select {
	// 		case <- ctx.DOne():
	// 			return
	// 		case stream <- dataItem:		
	// 		}
	// 	}
	// }

	// infiniteApples := make(chan interface{})
	// go generator( dataItem: "apple", infiniteApples)
}