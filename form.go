package request

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
)

func newMultipartBody(a *Args) (body io.Reader, contentType string, err error) {
	files := a.Files
	bodyBuffer := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuffer)
	for _, file := range files {
		fileWriter, err := bodyWriter.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, "", err
		}
		_, err = io.Copy(fileWriter, file.File)
		if err != nil {
			return nil, "", err
		}
	}
	if a.Data != nil {
		for k, v := range a.Data {
			bodyWriter.WriteField(k, v)
		}
	}
	contentType = bodyWriter.FormDataContentType()
	defer bodyWriter.Close()
	body = bodyBuffer
	return
}

func newJsonBody(a *Args) (body io.Reader, contentType string, err error) {
	b, err := json.Marshal(a.Json)
	if err != nil {
		return nil, "", err
	}
	return bytes.NewReader(b), defaultJsonType, err
}
