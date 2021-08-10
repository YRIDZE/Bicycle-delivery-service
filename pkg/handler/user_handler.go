package handler

import (
	"fmt"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/model/repository"
)

type UserHandler struct {
	userRepository repository.UserRepositoryI
}

func NewUserHandler(repo repository.UserRepositoryI) *UserHandler {
	return &UserHandler{userRepository: repo}
}

func (uch UserHandler) UserOptions() error {

	user := model.User{
		FirstName: "Ira",
		LastName:  "Laktionova",
		Username:  "Yridze",
		Email:     "Email@com",
		Password:  "123456",
	}

	userUpdate := model.User{
		ID:        34,
		FirstName: "Iya",
		LastName:  "Pupsia",
		Username:  "Yridze",
		Email:     "Email@com",
		Password:  "123456",
	}

	storedUser, err := uch.userRepository.Create(&user)
	if err != nil {
		return err
	}
	fmt.Printf("Added user: \"%v\"\n", storedUser)

	searchVal := "Email2@com"
	searchedUser, err := uch.userRepository.Get(&searchVal)
	if err != nil {
		return err
	}
	fmt.Printf("Searched user: \"%v\"\n", searchedUser)

	searchedUsers, err := uch.userRepository.GetAll()
	if err != nil {
		return err
	}
	fmt.Printf("Searched users: \"%v\"\n", searchedUsers)

	updatedUser, err := uch.userRepository.Update(&userUpdate)
	if err != nil {
		return err
	}
	fmt.Printf("Updated users: \"%v\"\n", updatedUser)

	err = uch.userRepository.Delete(3)
	if err != nil {
		return err
	}

	return nil
}
