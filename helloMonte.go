// helloMonte.go
package main

import (
	"fmt"
	"MonteCarloPi/mathutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", helloMonte)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
    //http.ListenAndServe(":8080", nil)
	//fmt.Println("Hello World!")
	//fmt.Println(mathutil.CalculatePi(10000))
}

func helloMonte(w http.ResponseWriter, r *http.Request){
	value, _ := strconv.ParseInt(r.URL.Path[1:], 0, 32)
	
	fmt.Fprintf(w, "Hi there, monte carlo pi of %s try is %f!", r.URL.Path[1:], mathutil.CalculatePi(int(value)))
	//fmt.Fprintf(w, "Hi there, pi is %s!", r.URL.Path[1:])
}