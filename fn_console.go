package dify

import (
	"context"
	"net/http"
)

func (cl *Client) setConsoleAuthorization(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+cl.ConsoleToken)
	req.Header.Set("Content-Type", "application/json")
}

func (cl *Client) sendGetRequestToConsole(ctx context.Context, api string) (httpCode int, bodyText []byte, err error) {
	return cl.sendGetRequest(ctx, true, api)
}

func (cl *Client) sendPostRequestToConsole(ctx context.Context, api string, postBody any) (httpCode int, bodyText []byte, err error) {
	return cl.sendPostRequest(ctx, true, api, postBody)
}
