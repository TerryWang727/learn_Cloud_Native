package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Cloud Native</h1>"))
	// 03.设置version
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
	// 02.将requst中的header 设置到 reponse中
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Add(k, vv)
		}
	}
	// 04.记录日志并输出
	clientIP := getClientIP(r)
	//fmt.Println(r.RemoteAddr)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientIP)
}

func Handler(writer http.ResponseWriter, req *http.Request) {
	version := os.Getenv("VERSION")
	//req.Response.Header = req.Header
	//解析参数，默认是不会解析的
	if err := req.ParseForm(); err != nil {
		panicError("ParseForm: ", err)
	}
	fmt.Println("VERSION: ", version)
	fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	clientIP := getClientIP(req)
	fmt.Println(clientIP)
	writer.Header().Set("VERSION", version)
	fmt.Println("header: ", req.Header)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Header {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
		writer.Header().Add(k, strings.Join(v, ""))
		//for _, vv := range v {
		//	writer.Header().Add(k, vv)
		//}
		fmt.Fprint(writer, k, ": ")
		fmt.Fprintln(writer, strings.Join(v, ""))
	}

	fmt.Fprintln(writer, writer.Header())
	fmt.Fprintln(writer, "加薪!") //这个写入到w的是输出到客户端的
	fmt.Fprintf(writer, version)
	//writer.Write([]byte(req.Header))
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
//解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func panicError(reason string, err error) {
	if err != nil {
		log.Fatal(reason, err)
		panic(err)
	}
}

type WelcomeHandlerStruct struct {}

func (*WelcomeHandlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	//init router
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/wxk", Handler)
	mux.Handle("/welcome", &WelcomeHandlerStruct{})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Starting HTTP server...")
	//设置监听的端口
	if err := http.ListenAndServe(":2022", mux); err != nil {
		panicError("ListenAndServe", err)
	}
}
