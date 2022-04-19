package plugins

type responseParams struct{
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

func MakeResponse (code int, message string) responseParams {
	response := responseParams{
		StatusCode: code,
		Message: message,
	}
	return response
}