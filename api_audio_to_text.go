package dify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type AudioToTextResponse struct {
	Text string `json:"text"`
}

func (cl *Client) AudioToText(filePath string) (result AudioToTextResponse, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	fw, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return result, errors.Errorf("error creating form file: %v", err)
	}

	fd, err := os.Open(filePath)
	if err != nil {
		return result, errors.Errorf("error opening file: %v", err)
	}
	defer fd.Close()

	_, err = io.Copy(fw, fd)
	if err != nil {
		return result, errors.Errorf("error copying file: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return result, errors.Errorf("error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", cl.GetAPI(ApiAudioToText), payload)
	if err != nil {
		return result, errors.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.ApiKey))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := cl.client.Do(req)
	if err != nil {
		return result, errors.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, errors.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, errors.Errorf("could not read the body: %v", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.Wrap(err, "failed to unmarshal the response")
	}

	return result, nil
}
