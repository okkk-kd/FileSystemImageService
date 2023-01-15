package diskStorage

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"sync"
	"time"
)

type diskStorage struct {
	db          *sqlx.DB
	mx          sync.RWMutex
	data        map[int]map[string]Image
	buf         int
	bufMax      int
	storagePath string
}

type DiskStorage interface {
	Upload(clientID int, fileName string, file []byte) error
	Download(clientID int, fileName string) (string, error)
	GetList(clientID int) (ImageList, error)

	write(data Image, clientID int, name string) error
	path(clientID int) string
	file(clientID, imgID int) string
}

func NewDiskStorage(db *sqlx.DB, buf int) (obj DiskStorage, err error) {
	data := make(map[int]map[string]Image)
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	return &diskStorage{
		bufMax:      buf,
		storagePath: dir,
		db:          db,
		data:        data,
	}, err
}

func (d *diskStorage) Upload(clientID int, fileName string, file []byte) error {
	var imgID int
	if _, err := os.Stat(d.path(clientID)); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(d.path(clientID), os.ModePerm); err != nil {
			return err
		}
		fmt.Println("dfgdsg")
	}
	if err := d.db.Get(&imgID, queryUploadFile, clientID, fileName); err != nil {
		return err
	}
	rawFile, err := os.Create(d.file(clientID, imgID))
	if err != nil {
		return err
	}
	if _, err := rawFile.Write(file); err != nil {
		return err
	}
	if err := d.write(
		Image{
			Name:       fileName,
			ImgID:      imgID,
			BinaryData: file,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		clientID,
		fileName,
	); err != nil {
		return err
	}
	return nil
}

func (d *diskStorage) Download(clientID int, fileName string) (string, error) {
	var img Image
	el, ok := d.data[clientID]
	if ok {
		img, ok = el[fileName]
		if ok {
			return string(img.BinaryData), nil
		}
	}
	if err := d.db.Get(&img, queryDownloadFile, clientID, fileName); err != nil {
		return "", err
	}
	file, err := os.Open(d.file(clientID, img.ImgID))

	if err != nil {
		return "", err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return "", statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	img.BinaryData = bytes
	if err := d.write(img, clientID, fileName); err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (d *diskStorage) GetList(clientID int) (ImageList, error) {
	var list ImageList
	if err := d.db.Select(&list.List, queryGetList, clientID); err != nil {
		return ImageList{}, err
	}
	return list, nil
}

func (d *diskStorage) write(data Image, clientID int, name string) error {
	defer d.mx.Unlock()
	d.mx.Lock()
	if d.buf >= d.bufMax {
		for key, _ := range d.data[clientID] {
			delete(d.data[clientID], key)
			break
		}
		d.buf--
	}
	_, ok := d.data[clientID]
	if ok {
		buf := d.data[clientID]
		buf[name] = data
		d.buf++
	} else {
		d.data[clientID] = map[string]Image{
			name: data,
		}
		d.buf++
	}
	return nil
}

func (d *diskStorage) path(clientID int) string {
	return fmt.Sprintf("%s\\%d", d.storagePath, clientID)
}

func (d *diskStorage) file(clientID, imgID int) string {
	return fmt.Sprintf("%s\\%d\\%d.bin", d.storagePath, clientID, imgID)
}
