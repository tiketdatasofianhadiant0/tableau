package models

import (
	"bytes"
	"fmt"
	"github.com/antchfx/xmlquery"
)

type ErrorBody struct {
	Error *Error `json:"error,omitempty"`
}

func NewErrorBody(b []byte) (*ErrorBody, error) {
	body := ErrorBody{}
	if err := json.Unmarshal(b, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func NewErrorBodyXML(b []byte) (*ErrorBody, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	q := xmlquery.FindOne(doc, "//error")
	if q == nil {
		return nil, fmt.Errorf("unknown error response")
	}

	e := Error{
		Code: q.SelectAttr("code"),
	}

	if n := q.SelectElement("summary"); n != nil {
		e.Summary = n.InnerText()
	}

	if n := q.SelectElement("detail"); n != nil {
		e.Detail = n.InnerText()
	}

	return &ErrorBody{Error: &e}, nil
}
