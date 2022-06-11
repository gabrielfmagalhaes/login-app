package out

type ControllerResponse struct {
	Code    interface{} `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
