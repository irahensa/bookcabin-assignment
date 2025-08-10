package server

import (
	"fmt"
	"net/http"

	"github.com/irahensa/bookcabin-assignment/backend/config"
	"github.com/irahensa/bookcabin-assignment/backend/handler"
	"github.com/irahensa/bookcabin-assignment/backend/resource"
	"github.com/irahensa/bookcabin-assignment/backend/usecase"
)

func Serve() {
	res, err := resource.InitResource()
	if err != nil {
		panic("failed resource init")
	}

	uc := usecase.InitUsecase(res)
	app := handler.InitHandler(uc)

	serveMux := http.NewServeMux()
	app.RegisterRoutes(serveMux)

	addr := fmt.Sprintf(":%s", config.GetConfig().Server.Port)
	fmt.Printf("Listening and serving on %s \n", addr)
	http.ListenAndServe(addr, serveMux)
}
