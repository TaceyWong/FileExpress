package core

import "sync"

type FileStorage interface {
	SaveFile() error
	DeleteFile() error
	GetFileURL() string
	GetFileResponse() error
}

type SystemFileStorage struct {
	Path string
}

func (s *SystemFileStorage) SaveFile() error {
	return nil
}

func (s *SystemFileStorage) DeleteFile() error {
	return nil
}

func (s *SystemFileStorage) GetFileURL() string {
	return ""
}

func (s *SystemFileStorage) GetFileResponse() error {
	return nil
}

type S3FileStorage struct {
	Path string
}

func (s *S3FileStorage) SaveFile() error {
	return nil
}

func (s *S3FileStorage) DeleteFile() error {
	return nil
}

func (s *S3FileStorage) GetFileURL() string {
	return ""
}

func (s *S3FileStorage) GetFileResponse() error {
	return nil
}

type OnedriveFileStorage struct {
	Path string
}

func (s *OnedriveFileStorage) SaveFile() error {
	return nil
}

func (s *OnedriveFileStorage) DeleteFile() error {
	return nil
}

func (s *OnedriveFileStorage) GetFileURL() string {
	return ""
}

func (s *OnedriveFileStorage) GetFileResponse() error {
	return nil
}

type OpenDALFileStorage struct {
	Path string
}

func (s *OpenDALFileStorage) SaveFile() error {
	return nil
}

func (s *OpenDALFileStorage) DeleteFile() error {
	return nil
}

func (s *OpenDALFileStorage) GetFileURL() string {
	return ""
}

func (s *OpenDALFileStorage) GetFileResponse() error {
	return nil
}

var once sync.Once
var instance FileStorage

func GetFileStorage(storageType string) FileStorage {
	once.Do(func() {
		switch storageType {
		case "system":
			instance = &SystemFileStorage{}
		case "s3":
			instance = &S3FileStorage{}
		case "onedrive":
			instance = &OnedriveFileStorage{}
		case "opendal":
			instance = &OpenDALFileStorage{}
		}
	})
	return instance
}
