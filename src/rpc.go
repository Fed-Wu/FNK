package main
import (
	"fmt"
	"time"
	"io/ioutil"
    "net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err == nil{}
	fmt.Println("[", req.RemoteAddr,"]: ", req.Method, req.URL)
	fmt.Println("body: ", string(body))
    str := "{\"jsonrpc\":\"2.0\",\"id\":0,\"result\":[\"0x0000000000000000000000000000000000000000\",\"0x0000000000000000000000000000000000000001\"]}"
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(str))
}
func main() {
    http.HandleFunc("/", index)
    s := &http.Server{
        Addr:           ":8545",
        ReadTimeout:    30 * time.Second,
        WriteTimeout:   30 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
