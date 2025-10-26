package response

type ServerResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Info    any    `json:"info,omitempty"`
}

func NewServerResponse(message string, data any, info ...any) *ServerResponse {
	resp := &ServerResponse{
		Message: message,
		Data:    data,
	}
	if len(info) > 0 {
		resp.Info = info[0]
	}
	return resp
}
