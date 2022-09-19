package calue

import(
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)
////taskschan中存储要扫描的端口，reschan存储开放的端口号，exitchan存储当前的goroutine是否完成的状态，wgscan 同步goroutine
//Go中提出了一种特殊的类型（channel）,channel类似于一个队列，遵循先进先出
func Scan(ip string, taskschan chan int, reschan chan int, exitchan chan bool, wgscan *sync.WaitGroup){

	defer func() {
		fmt.Println("任务完成")
		exitchan <- true
		wgscan.Done() //表明当前goroutine结束
	}()
	fmt.Println("开始任务")
	for {
		port, ok := <-taskschan
		if !ok {
			break
		}
		_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second)
		if err == nil {
			reschan <- port
			fmt.Println("开放的端口", port)
		}
	}
}