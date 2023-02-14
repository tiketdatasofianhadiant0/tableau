package models

import "strconv"

type Error struct {
	Summary string `json:"summary,omitempty"`
	Detail  string `json:"detail,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (e Error) Error() string {
	return e.Detail
}

func (e Error) String() string {
	return e.Summary + "; " + e.Detail
}

func (e Error) IsHttpCode(httpCode int) bool {
	if len(e.Code) < 3 {
		return false
	}

	if code, _ := strconv.Atoi(e.Code[:3]); code == httpCode {
		return true
	}

	return false
}
