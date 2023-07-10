package main

import (
	"fmt"
	"github.com/zhexiao/mtef-go/eqn"
	"net/http"
)

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":33333", nil)
}
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//路径中获取参数
	paraName := r.URL.Path[1:]
	// latex := eqn.Convert("E:/go_env/works/src/mtef-go-master/test/oleObject1.bin")
	latex := eqn.Convert(paraName)
	fmt.Println(paraName)
	fmt.Fprintf(w, latex)
}
