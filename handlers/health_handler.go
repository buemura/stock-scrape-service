package handlers

type HealthHandlerResponse struct {
	Message string `json:"message"`
}

func HealthHandler() HealthHandlerResponse {
	return HealthHandlerResponse{
		Message: "API working",
	}
}