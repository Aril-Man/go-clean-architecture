package response

type Response struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Status    bool   `json:"status" example:"false"`
	ErrorCode int    `json:"error_code" example:"500"`
	Message   string `json:"message" example:"Something wrong with this API"`
}

type UserResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
