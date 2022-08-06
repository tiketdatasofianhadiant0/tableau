package models

type ErrorBody struct {
	Error *Error `json:"error,omitempty"`
}

func NewErrorBody(bytes []byte) (*ErrorBody, error) {
	body := ErrorBody{}
	if err := json.Unmarshal(bytes, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
