package file_repository

import (
	"encoding/json"
	"github.com/YRIDZE/Bicycle-delivery-service/internal/helpers"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"os"
	"strconv"
	"sync"
)

type UserRepositoryI interface {
	Create(user *models.User) (int32, error)
	GetByEmail(email *string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(user *models.User) (*models.User, error)
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

func (ufr *UserFileRepository) Create(user *models.User) (*models.User, error) {
	user.ID = ufr.GetNextID()

	err := helpers.Create("users", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) Get(email *string) (*models.User, error) {
	user, err := helpers.Get("users", email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) GetAll() (*[]models.User, error) {
	usersSearched, err := helpers.GetAll("users")
	if err != nil {
		return nil, err
	}
	return &usersSearched, nil
}

func (ufr *UserFileRepository) Update(user *models.User) (*models.User, error) {
	err := helpers.Update("users", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ufr *UserFileRepository) Delete(id int) error {
	//err := helpers.Delete("users", id)
	//if err != nil {
	//	return err
	//}
	//return nil
	panic("implement me")
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
