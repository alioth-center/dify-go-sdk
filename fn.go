package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

func (cl *Client) sendGetRequest(ctx context.Context, forConsole bool, api string) (httpCode int, bodyText []byte, err error) {
	return cl.doRequest(ctx, http.MethodGet, forConsole, api, nil)
}

func (cl *Client) sendPostRequest(ctx context.Context, forConsole bool, api string, postBody any) (httpCode int, bodyText []byte, err error) {
	var payload *strings.Reader
	if postBody != nil {
		buf, err := json.Marshal(postBody)
		if err != nil {
			return 0, nil, err
		}
		payload = strings.NewReader(string(buf))
	} else {
		payload = nil
	}
	return cl.doRequest(ctx, http.MethodPost, forConsole, api, payload)
}

func (cl *Client) doRequest(ctx context.Context, method string, forConsole bool, api string, body io.Reader) (httpCode int, bodyText []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, method, api, body)
	if err != nil {
		return 0, nil, err
	}

	if forConsole {
		cl.setConsoleAuthorization(req)
	} else {
		cl.setAPIAuthorization(req)
	}

	resp, err := cl.client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()

	bodyText, err = io.ReadAll(resp.Body)
	return resp.StatusCode, bodyText, err
}

func CommonRiskForSendRequest(code int, err error) error {
	return CommonRiskForSendRequestWithCode(code, err, http.StatusOK)
}

func CommonRiskForSendRequestWithCode(code int, err error, targetCode int) error {
	if err != nil {
		return err
	}

	if code != targetCode {
		return errors.Errorf("status code: %d", code)
	}

	return nil
}
