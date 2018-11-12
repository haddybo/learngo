package main

import (
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil) //第二个参数一般传nil
	if err != nil {
		panic(err)
	}
}
