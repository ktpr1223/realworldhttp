package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	fmt.Println(r.Header)
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>2回目以降</body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>初回訪問</body></html>")
	}
}

// 動作確認は、p.40参照
// curl --http1.0 -c cookie.txt http://localhost:18888
// curl --http1.0 -b cookie.txt http://localhost:18888
func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
