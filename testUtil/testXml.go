package main;

import (
	"encoding/xml"
	"fmt"
	"time"
	"sync"
	"runtime"
)

//注意，结构体中的字段必须是可导出的
type Books struct {
	XMLName xml.Name `xml:"books"`;
	Nums    int      `xml:"nums,attr"`;
	Book    []Book   `xml:"book"`;
}

type Book struct {
	XMLName xml.Name `xml:"book"`;
	Name    string   `xml:"name,attr"`;
	Money   string   `xml:"Money,attr"`;
	Author  string   `xml:"author"`;
	Time    string   `xml:"time"`;
}

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()

	var wg sync.WaitGroup
	for i := 100000; i > 0; i-- {

		wg.Add(1)
		go func(wg *sync.WaitGroup) () {
			bs := Books{Nums: 666};
			//通过append添加book数据
			bs.Book = append(bs.Book, Book{Name: "小红", Money: "57.6$", Author: "阿三", Time: "2018年6月3日"});
			bs.Book = append(bs.Book, Book{Name: "小绿", Money: "79.9$", Author: "阿四", Time: "2018年6月5日"});
			//通过MarshalIndent，让xml数据输出好看点
			data, _ := xml.MarshalIndent(&bs, "", "  ");
			fmt.Println(string(data));
			wg.Done()
		}(&wg)

	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("%v s done!", time.Now().Sub(startTime).Seconds()))
}
