package dify

type TextToAudioPayload struct {
	MessageId string `json:"message_id"`
	Text      string `json:"text"`
	User      string `json:"user"`
}

// TODO
//func (cl *Client) TextToAudio(ctx context.Context, payload TextToAudioPayload) (result any, err error) {
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	_ = writer.WriteField("text", text)
//	_ = writer.WriteField("user", cl.User)
//	_ = writer.WriteField("streaming", "false")
//	err = writer.Close()
//	if err != nil {
//		return result, errors.Errorf("error closing writer: %v", err)
//	}
//
//	api := cl.GetAPI(ApiTextToAudio)
//	code, body, err := cl.sendPostRequestToAPI(ctx, api, payload)
//
//	err = CommonRiskForSendRequest(code, err)
//	if err != nil {
//		return result, err
//	}
//
//	err = json.Unmarshal(body, &result)
//	if err != nil {
//		return result, errors.Wrap(err, "failed to unmarshal the response")
//	}
//	return result, nil
//}
//
//func (cl *Client) TextToAudioStreaming(text string) (result any, err error) {
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	_ = writer.WriteField("text", text)
//	_ = writer.WriteField("user", cl.User)
//	_ = writer.WriteField("streaming", "true")
//	err = writer.Close()
//	if err != nil {
//		return result, errors.Errorf("error closing writer: %v", err)
//	}
//
//	req, err := http.NewRequest("POST", cl.GetAPI(ApiTextToAudio), payload)
//	if err != nil {
//		return result, errors.Errorf("error creating request: %v", err)
//	}
//
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.ApiKey))
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	resp, err := cl.client.Do(req)
//	if err != nil {
//		return result, errors.Errorf("error sending request: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return result, errors.Errorf("status code: %d", resp.StatusCode)
//	}
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return result, errors.Errorf("could not read the body: %v", err)
//	}
//	return body, nil
//}
