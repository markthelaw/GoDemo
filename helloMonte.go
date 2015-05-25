// helloMonte.go
package main

import (
	"fmt"
	"MonteCarloPi/mathutil"
	"net/http"
	"strconv"
	"os"
	"encoding/json"
	//"flag"
	"github.com/gorilla/mux"
)

type MontePI struct {
	Times int64
	Pi float64
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func main() {
	r := mux.NewRouter()
	//r.PathPrefix("/web").Handler(http.FileServer(http.Dir("./web")))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	r.HandleFunc("/monte/{trials}",  helloMonte)
	
	//http.HandleFunc("/monte/", helloMonte)
	fmt.Println("listening...")
	
	serveSingle("/favicon.ico", "./favicon.ico")
	
	http.Handle("/", r)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
	
	//http.ListenAndServe(":8080", nil)
	//fmt.Println("Hello World!")
	//fmt.Println(mathutil.CalculatePi(10000))
}



func helloMonte(w http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	
	//value, _ := strconv.ParseInt(r.URL.Path[1:], 0, 32)
	value, _ := strconv.ParseInt(vars["trials"], 0, 32)
	
	montePi := MontePI{value, mathutil.CalculatePi(int(value))}
	js, err := json.Marshal(montePi)
	
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}

  	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)

	//fmt.Fprintf(w, "Hi there, monte carlo pi of %s try is %f!", r.URL.Path[1:], mathutil.CalculatePi(int(value)))
	//fmt.Fprintf(w, "Hi there, pi is %s!", r.URL.Path[1:])
}