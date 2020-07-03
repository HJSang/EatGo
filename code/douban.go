package main

import (
	"fmt"
	"net/http"
	"strconv"
	"regexp"
	"io"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get error:", err1)
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n ==0 {
			break
		}
		if err2 !=nil && err2!= io.EOF{
			err=err2
			return
		}
		result += string(buf[:n])
	}
	fmt.Println("result:", result)
	return 
}

func SpiderPage(idx int){
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx-1)*25) + "&filter="
	fmt.Println("Url:",url)
	result, err :=  HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	ret := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	fileName := ret.FindAllStringSubmatch(result, -1)
	fmt.Println("fileName:", fileName)
	for _, name := range fileName {
		fmt.Println("Name:", name[0])
	}
}

func toWork(start, end int) {
	fmt.Printf("reading from %d to %d pages\n", start, end)
	for i:=start; i<=end; i++ {
		SpiderPage(i)
	}
}

func main() {
	var start, end int
	fmt.Print("Please input the start page (>=1)")
	fmt.Scan(&start)
	fmt.Print("Please input the end pages (>=start)")
	fmt.Scan(&end)
	toWork(start,end)
}
