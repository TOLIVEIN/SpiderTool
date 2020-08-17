package proxy

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Get66ip() {
	collector := colly.NewCollector()
	var ips []string
	var ports []string
	var proxies []string
	client, config := RedisClient()

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
			AddProxy(client, config, proxy)
			//_, testResult := TestProxy(proxy)
			//if testResult == http.StatusOK {
			//	proxies = append(proxies, proxy)
			//}
		}
		ips = ips[:0]
		ports = ports[:0]

		//fmt.Println(fmt.Sprintf("代理地址：%s", proxies))
	})

	for page := 1; page <= 10; page++ {
		url := fmt.Sprintf("http://www.66ip.cn/%s.html", strconv.Itoa(page))
		collector.Visit(url)
	}
	fmt.Println(fmt.Sprintf("代理地址：%s", proxies))

}

func GetKuaidaili() {
	collector := colly.NewCollector()
	var ips []string
	var ports []string
	var proxies []string
	client, config := RedisClient()

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
			fmt.Println(proxy)
			AddProxy(client, config, proxy)
			//_, testResult := TestProxy(proxy)
			//if testResult == http.StatusOK {
			//	proxies = append(proxies, proxy)
			//}
		}
		ips = ips[:0]
		ports = ports[:0]


		//fmt.Println(fmt.Sprintf("代理地址：%s", proxies))
	})

	for page := 1; page <= 10; page++ {
		url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%s", strconv.Itoa(page))
		collector.Visit(url)
	}
	fmt.Println(fmt.Sprintf("代理地址：%s", proxies))
	fmt.Println(All(client, config))

}

func TestProxy(proxyAddr string) (ip string, status int) {
	var testUrl string
	if strings.Contains(proxyAddr, "https") {
		testUrl = "https://icanhazip.com"
	} else {
		testUrl = "http://icanhazip.com"
	}

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
	//
	//begin := time.Now()
	//
	res, err := httpClient.Get(testUrl)

	if err != nil {
		fmt.Println(err)
		return
	}
	//
	if res.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("测试错误：%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("测试代理成功：%s，%s", proxy, res))


	//
	//defer res.Body.Close()
	//delay := int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000)
	//
	//if res.StatusCode != http.StatusOK {
	//	log.Println(err)
	//	return
	//}
	//
	return proxyAddr, res.StatusCode
}
