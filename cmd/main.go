package main

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/handler"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model/repository"
)

func main() {
	//userBDRepository := repository.UserDBRepository{}
	userFileRepository := repository.NewUserFileRepository()

	userHandler := handler.NewUserHandler(userFileRepository)

	err := userHandler.UserOptions()
	if err != nil {
		panic(err.Error())
	}

}
