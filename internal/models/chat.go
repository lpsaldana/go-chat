package models

type Chat struct{
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  string `json:"userid"`
}