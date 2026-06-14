package file

import (
	"io"
	"os"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UploadFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	destPath := filePath + ".uploaded"
	err = os.WriteFile(destPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DownloadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Service) DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}
