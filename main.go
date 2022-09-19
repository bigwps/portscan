package port_scan

import (
	"calue"
	"fmt"
	"net"
	"strconv"
	"time"
)  

func port_scan() {

	// [...]代表不确定数组个数的情况
	// make() 函数来创建一个切片 同时创建好相关数组： var slice1 []type =make([]type, len),
	// 这里  len  是数组的长度并且也是  slice  的初始长度
	 
	defer_port := [...]int{21, 22, 23, 25, 80, 443, 8080,
		110, 135, 139, 445, 389, 489, 587, 1433, 1434,
		1521, 1522, 1723, 2121, 3306, 3389, 4899, 5631,
		5632, 5800, 5900, 7071, 43958, 65500, 4444, 8888,
		6789, 4848, 5985, 5986, 8081, 8089, 8443, 10000,
		6379, 7001, 7002}
	var res []int = make([]int, 0)
	start := time.Now()
	// for - range结构是g特有的迭代结构, for ix, val := range coll { }  ,ix为索引
	for _, port := range defer_port {
		//DialTimeout参数1，扫描使用的协议，参数2，IP+端口号，参数3，设置连接超时的时间, strconv.Itoa()将整形转换为字符串
		_, err := net.DialTimeout("tcp", "127.0.0.1:"+strconv.Itoa(port), 3*time.Second) 
		// 判断是否超时，err是否为空
		if err == nil{
			fmt.Printf("端口开放%d\n",port)
			res = append(res, port)
		}else{
			fmt.Printf("端口关闭%d\n",port)
		}
	}
	end := time.Since(start)
	fmt.Println("花费时间",end)
	fmt.Println("开放端口",res)

}
