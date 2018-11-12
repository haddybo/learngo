package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()     //remember close

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	//all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)


}

//this func can determine what encode is in io.Reader
func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)     //use bufio.RewReader for Peek
	if err != nil {
		panic(err)
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}