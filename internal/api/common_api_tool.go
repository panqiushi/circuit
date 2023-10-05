package api

func ResponseBuilder(data interface{}, code int, message string) map[string]interface{} {
	return map[string]interface{}{
		"data":    data,
		"code":    code,
		"message": message,
	}
}
