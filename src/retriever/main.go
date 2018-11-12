package main

import (
	"retriever/mock"
	"fmt"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string //不用加func关键字
}

func download(r Retriever) string  {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster)  {
	poster.Post(url, map[string]string {
		"name":"ccmouse",
		"course": "golang",
	})
}

const url  = "http://www.imooc.com"

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents":"another fake imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{    //取地址让r成为一个指针接收者
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(r)

	if mockRetriever, ok := r.(*mock.Retriever); ok {    //r肚子里边是指针接收者
		fmt.Println(mockRetriever.Content)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v \n", r,r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}
