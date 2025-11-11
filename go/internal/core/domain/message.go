package domain

import "errors"

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

// I know that technicaly speaking I need to move the validation to business logic (app), but I wanted to keep the
// data and data-oriented functions in one place.
func (m *Message) Validate() error {
	if m.Author == "" {
		return errors.New("author field is required")
	}

	if m.Body == "" {
		return errors.New("body field is required")
	}

	return nil
}
