package diskStorage

import "time"

type Image struct {
	Name       string    `db:"name,omitempty"`
	ImgID      int       `db:"id"`
	BinaryData []byte    `db:"binary_data,omitempty"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type UploadResult struct {
	imgID int
	time  time.Time
}

type ImageList struct {
	List []Image
}
