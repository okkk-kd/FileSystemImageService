package models

type UploadFileRequest struct {
	ClientID int
	FileName string
	File     []byte
}

type GetFileList struct {
	ClientID int
}

type GetFileByName struct {
	ClientID int
	Name     string
}
