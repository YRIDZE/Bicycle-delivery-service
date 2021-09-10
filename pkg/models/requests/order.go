package requests

import "github.com/YRIDZE/Bicycle-delivery-service/pkg/models"

type OrderRequest struct {
	ID       int32                  `json:"id"`
	UserID   int32                  `json:"user_id"`
	Address  string                 `json:"address"`
	Status   string                 `json:"status"`
	Products []models.OrderProducts `json:"products"`
}
