package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
var catch=make(map[string]string)
func main() {

//	var url string
	
	fmt.Println("Starting server on port 8000")
	//fmt.Println("Enter the url")
	//fmt.Scan(&url)
	server:=http.Server{Addr: ":8000"}
	http.HandleFunc("/", Handler)
	//MakeGetReq("http://"+url)
	log.Fatal(server.ListenAndServe())	
}


func Caching(url string,w http.ResponseWriter) {
	check:=catch[url]
	if check=="" {
		_,data:=MakeGetReq("http://"+url)	
		catch[url]=data
	}else {
		fmt.Println(catch[url])
		w.Write([]byte(catch[url]))
		fmt.Println("From cache",catch[url])	
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %q","Hitansu")
	theurl:=r.URL.Path[1:]
	Caching(theurl,w)
	//statuscode,data:=MakeGetReq("http://"+theurl)
	//w.WriteHeader(statuscode)
	//w.Header().Get("Content-Type")
}

func MakeGetReq(url string) (int,string) {
	
	res,err:=http.Get(url)

	if err!=nil {
		fmt.Println(err)
		return 0,""
	}
	defer res.Body.Close()
	
	if res.StatusCode!=200 {
		fmt.Println("Status code not 200")
		return 0,""
	}
	data,err:=io.ReadAll(res.Body)

	if err!=nil {
		fmt.Println(err)
		return 0,""
	}
	//fmt.Println(string(data))
	return res.StatusCode,string(data)
}
