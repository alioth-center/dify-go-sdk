package dify

// TODO
//func (cl *Client) DatasetsFileUpload(filePath string, fileName string) (result FileUploadResponse, err error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return result, errors.Errorf("error opening file: %v", err)
//	}
//	defer file.Close()
//
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//
//	part, err := writer.CreateFormFile("file", fileName)
//	if err != nil {
//		return result, errors.Errorf("error creating form file: %v", err)
//	}
//	_, err = io.Copy(part, file)
//	if err != nil {
//		return result, errors.Errorf("error copying file: %v", err)
//	}
//
//	_ = writer.WriteField("user", cl.User)
//	err = writer.Close()
//	if err != nil {
//		return result, errors.Errorf("error closing writer: %v", err)
//	}
//
//	req, err := http.NewRequest("POST", cl.GetConsoleAPI(ConsoleApiFileUpload), body)
//	if err != nil {
//		return result, errors.Errorf("error creating request: %v", err)
//	}
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.ConsoleToken))
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		return result, errors.Errorf("error sending request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 201 {
//		return result, errors.Errorf("status code: %d, create file failed", resp.StatusCode)
//	}
//
//	bodyText, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return result, errors.Errorf("could not read the body: %v", err)
//	}
//
//	err = json.Unmarshal(bodyText, &result)
//	if err != nil {
//		return result, errors.Wrap(err, "failed to unmarshal the response")
//	}
//	return result, nil
//}
