package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(writer http.ResponseWriter, req *http.Request) {
	version := os.Getenv("VERSION")
	//req.Response.Header = req.Header
	//解析参数，默认是不会解析的
	if err := req.ParseForm(); err != nil {
		log.Fatal("ParseForm: ", err)
	}
	fmt.Println("VERSION: ", version)
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("host: ", req.Host)
	fmt.Println("header: ", req.Header)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Header {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		writer.Header().Add(k, strings.Join(v, ""))
		fmt.Fprint(writer, k, ": ")
		fmt.Fprintln(writer, strings.Join(v, ""))
	}
	fmt.Fprintln(writer, writer.Header())
	fmt.Fprintln(writer, "加薪!") //这个写入到w的是输出到客户端的
	fmt.Fprintf(writer, version)
	//writer.Write([]byte(req.Header))
}

type WelcomeHandlerStruct struct {

}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func (*WelcomeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main () {
	//init router
	http.HandleFunc("/wxk", Handler) //设置访问的路由
	http.Handle("/welcome", &WelcomeHandlerStruct{})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	log.Println("Starting HTTP server...")
	err := http.ListenAndServe(":2022", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

