package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type GetMetaResponse struct {
	ToolIcons struct {
		Dalle2  string `json:"dalle2"`
		APITool struct {
			Background string `json:"background"`
			Content    string `json:"content"`
		} `json:"api_tool"`
	} `json:"tool_icons"`
}

func (cl *Client) GetMeta(ctx context.Context) (result GetMetaResponse, err error) {
	api := cl.GetAPI(ApiMeta)
	code, body, err := cl.sendGetRequestToAPI(ctx, api)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.Wrap(err, "failed to unmarshal the response")
	}
	return result, nil
}
