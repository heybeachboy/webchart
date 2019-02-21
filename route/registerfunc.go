package route

import (
	"net/http"
	"webchart/controller"
)

func registerFunc() {
	 http.FileServer(http.Dir("./webchart/html"))
	 http.HandleFunc("/",controller.Home)
     http.HandleFunc("/echo",controller.Echo)
}

/**
 *start http service
 */
func StartHttp(addr string)(error) {
	 registerFunc()
	 return http.ListenAndServe(addr,nil)
}
