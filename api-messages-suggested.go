package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type MessagesSuggestedQuery struct {
	User string `json:"user"`
}

type MessagesSuggestedResponse struct {
	Result string   `json:"result"`
	Data   []string `json:"data"`
}

func (cl *Client) MessagesSuggested(ctx context.Context, messageId string, query MessagesSuggestedQuery) (result MessagesSuggestedResponse, err error) {
	// TODO
	if messageId == "" {
		return result, errors.Errorf("message_id is required")
	}

	api := cl.GetAPI(ApiMessagesSuggested)
	api = UpdateAPIParam(api, ApiParamMessageId, messageId)

	code, body, err := cl.sendGetRequest(ctx, false, api)

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
