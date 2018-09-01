package http

import "fmt"

type HttpError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("http error (%d): %s", h.Message, h.Status)
}
