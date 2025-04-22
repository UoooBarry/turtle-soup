package client

import "fmt"

type (
	ErrAPIRequest struct {
		Message string
	}

	ErrAPIResponse struct {
		Status string
		Body   string
	}
)

func (e ErrAPIRequest) Error() string {
	return fmt.Sprintf("API request failed: %s", e.Message)
}

func (e ErrAPIResponse) Error() string {
	return fmt.Sprintf("request status error: %s, body: %s", e.Status, e.Body)
}
