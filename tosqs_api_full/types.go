package main

type Query struct {
	ID             uint   `json:"id" `
	Description    string `json:"description" `
	Query          string `json:"query" `
	CountRows      int    `json:"countRows" `
	QueryCountRows string `json:"countRowsQuery" `
	Selected       string
}

type Queue struct {
	ID       uint   `json:"id" `
	Name     string `json:"name" `
	Selected string
}

type FormElements struct {
	QueryOptionList []Query
	QueueOptionList []Queue
}
