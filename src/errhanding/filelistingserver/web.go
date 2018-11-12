package main

import (
	"net/http"
	"errhanding/filelistingserver/filelisting"
	"os"
	_ "net/http/pprof"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler)  func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//Panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			//TODO 日志颜色设置
			log.Printf("Error occured handling request: %s", err.Error())

			//user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			//system error
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err) :
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main()  {
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		 panic(nil)
	}
}
