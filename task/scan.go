package task

import (
	"net"
	"sync"
	"time"
)

////taskschan中存储要扫描的端口，reschan存储开放的端口号，exitchan存储当前的goroutine是否完成的状态，wgscan 同步goroutine
func Scan(ip string, allport chan string, openport chan string, exitchan chan bool, wgscan *sync.WaitGroup) {
	// defer关键字在函数执行完后再执行func，表示函数完成时等待组减一
	defer func() {
		exitchan <- true
		wgscan.Done()
	}()
	// 通信操作符<-,信息按照箭头的方向流，理解为传送，A := <- B 可以理解为B值传送到A
	for {
		port, ok := <-allport
		// 非运算！
		if !ok {
			break
		}
		_, err := net.DialTimeout("tcp", ip+":"+port, time.Second)
		// 指针、切片、映射、通道、函数和接口的零值则是 nil。布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""
		if err == nil {
			openport <- port
		}
	}
}
