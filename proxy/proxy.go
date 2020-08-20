package proxy

import (
	"SpiderTool/conf"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gocolly/colly"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Get66ip() {
	collector := colly.NewCollector()
	var ips []string
	var ports []string
	//var proxies []string
	client, config := RedisClient()
	var wg sync.WaitGroup


	collector.OnRequest(func(request *colly.Request) {
		fmt.Println(fmt.Sprintf("请求页面 %s\n", request.URL))
	})

	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println(fmt.Sprintf("出现错误 %s", err))
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println(fmt.Sprintf("已访问页面 %s", response.Request.URL))
	})

	collector.OnHTML("#main tr:first-child~tr td:first-child", func(element *colly.HTMLElement) {
		ips = append(ips, element.Text)

	})
	collector.OnHTML("#main tr:first-child~tr td:nth-child(2)", func(element *colly.HTMLElement) {
		ports = append(ports, element.Text)
	})

	collector.OnScraped(func(response *colly.Response) {
		for i := 0; i < len(ips); i++ {
			proxy := fmt.Sprintf("%s:%s", ips[i], ports[i])
			wg.Add(1)
			//done := make(chan bool)
			fmt.Printf("测试第%d个代理。。。\n", i)
			go TestProxy(proxy, &wg, client, config)
			//Add(client, config, proxy)
			//if <-done == true {
			//	AddProxy(client, config, proxy)
			//}
		}
		wg.Wait()
		ips = ips[:0]
		ports = ports[:0]
	})

	for page := 1; page <= 10; page++ {
		url := fmt.Sprintf("http://www.66ip.cn/%s.html", strconv.Itoa(page))
		collector.Visit(url)
	}
	//fmt.Println(fmt.Sprintf("代理地址：%s", proxies))

}

func GetKuaidaili() {
	collector := colly.NewCollector()
	var ips []string
	var ports []string
	//var proxies []string
	client, config := RedisClient()
	var wg sync.WaitGroup

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println(fmt.Sprintf("请求页面 %s\n", request.URL))
	})

	collector.OnError(func(_ *colly.Response, err error) {
		fmt.Println(fmt.Sprintf("出现错误 %s", err))
	})

	collector.OnResponse(func(response *colly.Response) {
		fmt.Println(fmt.Sprintf("已访问页面 %s", response.Request.URL))

	})

	collector.OnHTML("[data-title=\"IP\"]", func(element *colly.HTMLElement) {
		ips = append(ips, element.Text)
	})
	collector.OnHTML("[data-title=\"PORT\"]", func(element *colly.HTMLElement) {
		ports = append(ports, element.Text)
	})

	collector.OnScraped(func(response *colly.Response) {
		for i := 0; i < len(ips); i++ {
			proxy := fmt.Sprintf("%s:%s", ips[i], ports[i])
			wg.Add(1)
			//done := make(chan bool)
			fmt.Printf("测试第%d个代理。。。\n", i)
			go TestProxy(proxy, &wg, client, config)
			//Add(client, config, proxy)
			//if <-done == true {
			//	AddProxy(client, config, proxy)
			//}
		}
		wg.Wait()
		ips = ips[:0]
		ports = ports[:0]
	})

	for page := 1; page <= 10; page++ {
		url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%s", strconv.Itoa(page))
		collector.Visit(url)
	}
	//fmt.Println(fmt.Sprintf("代理地址：%s", proxies))
	//proxies = All(client, config)
	////fmt.Println(All(client, config))
	//for index, proxy := range proxies {
	//	done := make(chan bool)
	//	go TestProxy(proxy, done, &wg)
	//	fmt.Printf("测试第%d个代理%t", index, <-done)
	//}

}

func TestProxy(proxyAddr string, wg *sync.WaitGroup, clinet *redis.Client, config *conf.Config) (ip string, status int) {
	var testUrl string
	if strings.Contains(proxyAddr, "https") {
		testUrl = "https://icanhazip.com"
	} else {
		testUrl = "http://icanhazip.com"
	}

	defer wg.Done()
	proxy, err := url.Parse(proxyAddr)
	//if err != nil {
	// fmt.Println(fmt.Sprintf("代理：%s", err))
	//
	//} else {
	//	fmt.Println(fmt.Sprintf("代理为：%s", proxy))
	//}
	transport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	httpClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	begin := time.Now()

	res, err := httpClient.Get(testUrl)

	if err != nil {
		fmt.Println(err)
		return
	}
	//
	if res.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("测试错误：%s", err))
		//done <- false
		return
	}else {
		Max(clinet, config, proxyAddr)
		delay := int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000)
		fmt.Println(fmt.Sprintf("测试代理成功：%s已添加，响应代码：%d，延迟：%dms", proxyAddr, res.StatusCode, delay))
		return
	}

	//done <- true

	//
	//defer res.Body.Close()
	//
	//if res.StatusCode != http.StatusOK {
	//	log.Println(err)
	//	return
	//}
	//
	return proxyAddr, res.StatusCode
}
