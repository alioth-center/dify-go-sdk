package dify

import (
	"context"
	"net/http"
)

func (cl *Client) sendGetRequestToAPI(ctx context.Context, api string) (httpCode int, bodyText []byte, err error) {
	return cl.sendGetRequest(ctx, false, api)
}

func (cl *Client) sendPostRequestToAPI(ctx context.Context, api string, postBody any) (httpCode int, body []byte, err error) {
	return cl.sendPostRequest(ctx, false, api, postBody)
}

func (cl *Client) setAPIAuthorization(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+cl.ApiKey)
	req.Header.Set("Content-Type", "application/json")
}
