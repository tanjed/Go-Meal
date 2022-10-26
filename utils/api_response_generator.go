package utils

func GenerateApiResponse(message string, status int) map[string]interface{} {
	response := make(map[string]interface{})
	response["message"] = message
	response["status"] = status
	return response
}