package main

import (
	"github.com/YRIDZE/Bicycle-delivery-service/handler"
	"github.com/YRIDZE/Bicycle-delivery-service/repository"
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
