package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q","Hitansu")
	theurl:=r.URL.Path[1:]
	MakeGetReq("http://"+theurl)
}

func MakeGetReq(url string) {
	
	res,err:=http.Get(url)

	if err!=nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	
	if res.StatusCode!=200 {
		fmt.Println("Status code not 200")
		return
	}
	data,err:=io.ReadAll(res.Body)

	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
