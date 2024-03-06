package responses

func NewSuccessResponseBody[T interface{}](payload T) ResponseWrapper[T] {
	return ResponseWrapper[T]{
		Data:    payload,
		Success: true,
		Error:   "",
	}
}

func NewFailureResponseBody(err error) ResponseWrapper[any] {
	return ResponseWrapper[any]{
		Data:    nil,
		Success: false,
		Error:   err.Error(),
	}
}

type ResponseWrapper[T interface{}] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
