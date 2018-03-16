package main

import(
	"fmt"
	"Calc/Operations"
	"net/http"
	"strings"
	"html"
	"time"
	"log"
	)



func main() {
	log.Println("Here the synatx to query the Request\n")
	log.Print("Adding Numbers -> 1 + 2 ==> 1A2  ----- A is used as replacement for +")
	log.Print("Adding Numbers -> 1 - 2 ==> 1S2  ----- S is used as replacement for -")
	log.Print("Adding Numbers -> 1 * 2 ==> 1M2  ----- M is used as replacement for *")
	log.Print("Adding Numbers -> 1 / 2 ==> 1D2  ----- D is used as replacement for /")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path,"/")

		if urlPathElements[1] == "CalcServer" {
			exp := strings.TrimSpace(urlPathElements[2])
			s := []rune(exp)
			for i:= 0; i <len(s); i++ {
				if s[i] == rune('A') {
					s[i] = rune('+')
				}
				if s[i] == rune('S') {
					s[i] = rune('-')
				}
				if s[i] == rune('M') {
					s[i] = rune('*') 
				}
				if s[i] == rune('D') {
					s[i] = rune('/')
				}
				
			}
			expr := string(s)
			res, err := Operations.Cal(expr)
			if err !=nil{
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(res))

			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})
	log.Println("Listening....")
	s := &http.Server {
		Addr: ":3000",
		ReadTimeout : 10 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,
	}
	
	s.ListenAndServe()	
	
	
}