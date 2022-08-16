package models

type QueryWorkbookBody struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Workbooks  *struct {
		Workbook []Workbook `json:"workbook,omitempty"`
	} `json:"workbooks,omitempty"`
}
