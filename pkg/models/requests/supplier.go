package requests

type SupplierRequest struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
