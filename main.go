package main

import (
	"SpiderTool/proxy"
	"fmt"
)

func main() {
	fmt.Println("开始抓取IP......")
	proxy.Get66ip()
	//proxy.GetKuaidaili()
	//proxy.RedisClient()
}

//func hello(done chan bool) {
//	fmt.Println("Hello world goroutine")
//	done <- true
//}
//func main() {
//	done := make(chan bool)
//	go hello(done)
//	//<- done
//	//time.Sleep(1 * time.Second)
//	fmt.Printf("main function, %v", <- done)
//}
