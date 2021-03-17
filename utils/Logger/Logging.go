package Logger

import (
	"fmt"
	"net/http"
)

func ErorrDebugger(err error,w http.ResponseWriter){
	if err != nil {
		debugView(w,err)
		fmt.Println(err,w)
	}
}

func ErorrLog(err error){
	if err != nil {
		fmt.Println(err)
	}
}
