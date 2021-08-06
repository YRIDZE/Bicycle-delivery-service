package repository

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/helpers"
	"github.com/YRIDZE/Bicycle-delivery-service/model"
	"os"
	"strconv"
	"sync"
)

var idSequence int32

type UserRepositoryI interface {
	Create(u *model.User) (*model.User, error)
	Get(email *string) (*model.User, error)
	GetAll(email *string, id *int32) (*[]model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id int) error
}

type UserFileRepository struct {
	idMutex *sync.Mutex
}

func NewUserFileRepository() *UserFileRepository {
	return &UserFileRepository{
		idMutex: &sync.Mutex{},
	}
}

func (ufr *UserFileRepository) Create(user *model.User) (*model.User, error) {
	user.ID = ufr.GetNextID()

	err := helpers.CreateModel("users", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) Get(email *string) (*model.User, error) {
	user, err := helpers.GetModel("users", email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) GetAll(email *string, id *int32) (*[]model.User, error) {
	var usersSearched []model.User
	err := helpers.GetAllModels("users", email, &usersSearched)
	if err != nil {
		return nil, err
	}
	return &usersSearched, nil
}

func (ufr *UserFileRepository) Update(user *model.User) (*model.User, error) {
	err := helpers.UpdateModel("users", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) Delete(id int) error {
	err := helpers.Delete("users", id)
	if err != nil {
		return err
	}
	return nil
}

func (ufr *UserFileRepository) GetNextID() (idSequence int32) {
	ufr.idMutex.Lock()

	bytes, err := os.ReadFile("./datastore/users_id_sequence.txt")
	if err != nil {
		return 0
	}

	idSequence64, _ := strconv.Atoi(string(bytes[:]))
	idSequence = int32(idSequence64) + 1

	newBytes, err := json.Marshal(idSequence)
	if err != nil {
		return 0
	}
	err = os.WriteFile("./datastore/users_id_sequence.txt", newBytes, 0600)
	if err != nil {
		return 0
	}
	ufr.idMutex.Unlock()

	return
}
