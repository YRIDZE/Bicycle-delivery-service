package services

import (
	"github.com/YRIDZE/Bicycle-delivery-service/conf"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/requests"
	yolo_log "github.com/YRIDZE/yolo-log"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	cfg      conf.ConfigToken
	logger   *yolo_log.Logger
	userRepo db_repository.UserRepositoryI
}

func NewUserService(cfg *conf.ConfigToken, logger *yolo_log.Logger, userRepo db_repository.UserRepositoryI) *UserService {
	return &UserService{
		cfg:      *cfg,
		logger:   logger,
		userRepo: userRepo,
	}
}

func (u UserService) Create(user *requests.UserRequest) (*models.User, error) {
	return u.userRepo.Create(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  generatePasswordHash(user.Password),
		},
	)
}

func (u UserService) GetByEmail(email string) (*models.User, error) {
	return u.userRepo.GetByEmail(email)
}

func (u UserService) GetByID(id int32) (*models.User, error) {
	return u.userRepo.GetByID(id)
}

func (u UserService) EmailExist(email string) (int, error) {
	return u.userRepo.EmailExist(email)
}

func generatePasswordHash(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}
