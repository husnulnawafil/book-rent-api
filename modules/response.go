package modules

type Response struct {
	Meta       *Meta       `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Page  int32 `json:"page"`
	Total int32 `json:"total"`
	Limit int32 `json:"limit"`
}

func (r *Response) SendResponse(message string, code int, data interface{}, pagination *Pagination) *Response {
	res := &Response{
		Meta: &Meta{
			Message: message,
			Code:    code,
		},
		Data:       data,
		Pagination: pagination,
	}

	return res
}
