package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

/*
// handler default
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		os.Exit(1)
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}*/

// handler cookie
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		os.Exit(1)
	}
	fmt.Println(string(dump))

	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>welcome back</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body>hello</body></html>\n")
	}
}

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %s\n", r.URL.String())
	fmt.Printf("Query: %v\n", r.URL.Query())
	fmt.Printf("Proto: %s\n", r.Proto)
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Header: %v\n", r.Header)

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510a27804 73e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/digest", handlerDigest)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
