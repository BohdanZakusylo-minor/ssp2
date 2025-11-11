package domain

import "errors"

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (m *Message) Validate() error {
	if m.Author == "" {
		return errors.New("author field is required")
	}

	if m.Body == "" {
		return errors.New("body field is required")
	}

	return nil
}
