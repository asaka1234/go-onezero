package onezero

import "fmt"

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
}

func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Bearer %s", auth),
	}
}
