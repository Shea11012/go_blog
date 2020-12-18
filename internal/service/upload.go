package service

import (
	"errors"
	"github.com/shea11012/go_blog/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	AccessUrl string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo,error) {
	fileName := upload.GetFileName(header.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	if !upload.CheckContainExt(fileType,fileName) {
		return nil,errors.New("file suffix is not supported")
	}

	if upload.CheckSavePath(dst) {
		err := upload.CreateSavePath(uploadSavePath,os.ModePerm)
		if err != nil {
			return nil,errors.New("failed to create save directory")
		}
	}

	if upload.CheckMaxSize(fileType,file) {
		return nil,errors.New("exceeded maximum file limit")
	}

	if upload.CheckPermission(dst) {
		return nil,errors.New("insufficient file permissions")
	}

	if err := upload.SaveFile(header,dst); err != nil {
		return nil,err
	}

	accessUrl := upload.GetFileUrl(fileName)

	return &FileInfo{
		Name: fileName,
		AccessUrl: accessUrl,
	},nil
}
