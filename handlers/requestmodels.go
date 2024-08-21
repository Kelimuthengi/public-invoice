package handlers

type ParentInput struct {
	Username        string         `json:"username" binding:"required"`
	Email           string         `json:"email" binding:"required"`
	Address         string         `json:"address" binding:"required"`
	Phonenumber     string         `json:"phonenumber" binding:"required"`
	Students        []StudentInput `json:"students"`
	HouseUnitNameID uint           `json:"houseId"`
	// UserID      string `json:"userid"`
}

type StudentInput struct {
	Username        string `json:"username" binding:"required"`
	AdmissionNumber string `json:"admissionnumber" binding:"required"`
	Stream          string `json:"stream" binding:"required"`
	Boardingstatus  bool   `json:"boardingstatus" binding:"required"`
	Hostelname      string `json:"hostelname"`
	// ParentID        string `json:"parentid"`
}

type RequestParams struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	SearchValue string `json:"searchValue"`
}

type ProductInput struct {
	ProductName string `json:"productName" binding:"required"`
	ProductType string `json:"productType" `
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Discount    int    `json:"discount"`
	UserId      uint   `json:"userid" binding:"required"`
	Id          uint   `json:"id"`
}

type InvoiceInput struct {
	InvoiceNumber string `json:"invoiceNumber" binding:"required"`
	// InvoiceDue    string           `json:"invoiceDue" binding:"required"`
	Products     []ProductInput `json:"products" binding:"required"`
	StudentID    uint           `json:"studentId"`
	ParentID     uint           `json:"parentId"`
	HouseNo      string         `json:"houseno"`
	MeterReading float64        `json:"meterreading"`
	UserId       uint           `json:"userid" binding:"required"`
}

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Phonenumber string `json:"phonenumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type ResponseHandler struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty" `
	Count   int         `json:"count,omitempty"`
	Token   string      `json:"token,omitempty"`
}

type LoginHandler struct {
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
}

type TransactionRequest struct {
	TransactionRef string                `json:"transRef" binding:"required"`
	Products       []InvoiceProductInput `json:"products" binding:"required"`
	InvoiceNumber  string                `json:"invoiceNumber" binding:"required"`
	Amount         float64               `json:"amount" binding:"required"`
}

type InvoiceProductInput struct {
	ProductId uint `json:"productId"`
}

type UnitNameInput struct {
	Unitname string `json:"unitname" binding:"required"`
}

type UnitsInput struct {
	Price    float64 `json:"price" binding:"required"`
	Discount float64 `json:"discount"`
	HouseNo  string  `json:"houseno" binding:"required"`
	ID       uint    `json:"id"`
}

type SendMailInput struct {
	TenantId      uint   `json:"tenantId" binding:"required"`
	InvoiceNumber string `json:"invoicenumber" binding:"required"`
}
