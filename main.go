package main

import (
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/model"
	"github.com/YRIDZE/Bicycle-delivery-service/repository"
)

func main() {
	user := model.User{
		FirstName: "Ira",
		LastName:  "Laktionova",
		Username:  "Yridze",
		Email:     "Email@com",
		Password:  "123456",
	}

	userUpdate := model.User{
		ID:        22,
		FirstName: "Iya",
		LastName:  "Pupsia",
		Username:  "Yridze",
		Email:     "Email@com",
		Password:  "123456",
	}
	userRepository := repository.NewUserFileRepository()
	storedUser, err := userRepository.Create(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Added user: \"%v\"\n", storedUser)

	searchVal := "Email2@com"
	searchedUser, err := userRepository.Get(&searchVal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Searched user: \"%v\"\n", searchedUser)

	searchVal2 := "Email@com"
	searchedUsers, err := userRepository.GetAll(&searchVal2, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Searched users: \"%v\"\n", searchedUsers)

	updatedUser, err := userRepository.Update(&userUpdate)
	if err != nil {
		fmt.Println("Update error: ", err.Error())
		return
	}
	fmt.Printf("Updated users: \"%v\"\n", updatedUser)

	err = userRepository.Delete(24)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
