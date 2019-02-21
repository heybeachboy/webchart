package controller

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter,r *http.Request) {

	fmt.Fprintf(w,"Hello:%s",r.Host)
	//http.ServeFile(w,r,"home.html")

}

func Echo(w http.ResponseWriter,r *http.Request) {

}
