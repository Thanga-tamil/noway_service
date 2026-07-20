package response

func Success(message string, statusCode int, data any) any {

	return map[string]any{
		"data": data,
		"message": message,
		"status": statusCode,
	}

}

func Error(message string, statusCode int) any {

	return map[string]any{
		"message": message,
		"status": statusCode,
	}

}
