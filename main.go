package main

import (
	//"fmt"
	// "net/http"
	"o2o-center-api/routers"
)

func main() {
	//register router
	r :=routers.SetupRouter()
	// if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
	// 	fmt.Printf("server startup failed, err:%v\n", err)
	// }
	r.Run(":9000")
	
}
