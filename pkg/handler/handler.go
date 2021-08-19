package handler

import (
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/create", h.Create)
	router.HandleFunc("/getByEmail", h.GetByEmail)
	router.HandleFunc("/getAll", h.GetAll)
	router.HandleFunc("/update", h.Update)
	router.HandleFunc("/delete", h.Delete)

	return router
}

//func (h Handler) UserOptions() error {
//
//	user := model.User{
//		FirstName: "Ira",
//		LastName:  "Laktionova",
//		Email:     "Email@com",
//		Password:  "123456",
//	}
//
//	user2 := model.User{
//		ID:        4,
//		FirstName: "Ira2",
//		LastName:  "Laktionova2",
//		Email:     "Email@com",
//		Password:  "123456",
//	}
//
//	_, err := h.services.Create(&user)
//	if err != nil {
//		return err
//	}
//
//	res, err := h.services.GetByEmail(&user.Email)
//	if err != nil {
//		return err
//	}
//	fmt.Println(res)
//
//	res1, err := h.services.GetAll()
//	if err != nil {
//		return err
//	}
//	fmt.Println(res1)
//
//	_ = h.services.Update(&user2)
//
//	_ = h.services.Delete(3)
//
//
//	return nil
//}
