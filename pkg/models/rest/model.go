package model

import "mime/multipart"

// Resume ...
type Resume struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Email   string `json:"email"`
	Project string `json:"project"`
	File    File   `json:"file"`
}

// File ...
type File struct {
	Filename string `json:"filename"`
	Content  *multipart.FileHeader
}
