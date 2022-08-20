package herrors

type Response struct {
	Error interface{} `json:"error"`
}

func NewResponse(msg interface{}) Response {
	return Response{Error: msg}
}
