package main

//参考学习：狼组安全：https://cloud.tencent.com/developer/article/1785578

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"xc10/task"
)

func main() {

	var scanports []string
	var portinput []string
	// [...]代表不确定数组个数的情况
	// make() 函数来创建一个切片 同时创建好相关数组： var slice1 []type =make([]type, len),
	//使用  make()  函数来给它分配内存。这里先声明了一个int通道 chan，然后创建了它
	// 这里  len  是数组的长度并且也是  slice  的初始长度

	defaultports := [...]string{"21", "22", "23", "25", "80", "443", "8080",
		"110", "135", "139", "445", "389", "489", "587", "1433", "1434",
		"1521", "1522", "1723", "2121", "3306", "3389", "4899", "5631",
		"5632", "5800", "5900", "7071", "43958", "65500", "4444", "8888",
		"6789", "4848", "5985", "5986", "8081", "8089", "8443", "10000",
		"6379", "7001", "7002"}

	var ip string
	var ports string

	flag.StringVar(&ip, "i", "127.0.0.1", "扫描的IP地址,默认本机地址")
	flag.StringVar(&ports, "p", "", "扫描的端口地址，默认使用默认端口")
	// flag.IntVar(&gonum, "g", 4, "开启的goroutine的数量,默认为4")
	flag.Parse()

	//扫描哪些端口
	if len(ports) != 0 {
		portinput = strings.Split(ports, "-")
		a, _ := strconv.Atoi(portinput[0])
		b, _ := strconv.Atoi(portinput[1])
		for i := a; i < b+1; i++ {
			scanports = append(scanports, strconv.Itoa(i))
		}
	// 不指定则默认
	} else {
		scanports = defaultports[:]
	}

	allport := make(chan string, len(scanports))
	openport := make(chan string, len(scanports))
	exitchan := make(chan bool, 4)

	// 声明一个等待组，对一组等待任务只需要一个等待组，而不需要每一个任务都使用一个等待组。
	var wgp sync.WaitGroup

	for _, value := range scanports {
		allport <- value
	}
	// 关闭 allport 通道
	close(allport)

	start := time.Now()

	// 等待组4,四个goroutine执行扫描任务
	for i := 0; i < 4; i++ {
		wgp.Add(1)
		go task.Scan(ip, allport, openport, exitchan, &wgp)
	}
	// 等待所有任务完成
	wgp.Wait()

	// 判断4个协程是否都执行完了，当执行完了之后才能关闭openport才能关闭
	for i := 0; i < 4; i++ {
		<-exitchan
	}

	close(exitchan)
	close(openport)

	end := time.Since(start)

	for {
		open, ok := <-openport
		if !ok {
			break
		}
		fmt.Println("开放的端口：", open)
	}
	fmt.Println("花费的时间：", end)
}
