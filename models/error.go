package models

type Error struct {
	Summary string `json:"summary,omitempty"`
	Detail  string `json:"detail,omitempty"`
	Code    string `json:"code,omitempty"`
}

func (e Error) Error() string {
	return e.Detail
}
