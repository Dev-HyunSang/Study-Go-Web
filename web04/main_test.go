package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	// 개인 프로필 사진 ^__^
	path := "/Users/hyun.sang/Dev/images/photo_2021-03-01_18-45-56.jpg"
	file, _ := os.Open(path)

	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

}
