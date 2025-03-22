package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type MessagesFeedbacksPayload struct {
	Rating  string `json:"rating"`
	User    string `json:"user"`
	Content string `json:"content"`
}

type MessagesFeedbacksResponse struct {
	Result string `json:"result"`
}

func (cl *Client) MessagesFeedbacks(ctx context.Context, messageId string, payload MessagesFeedbacksPayload) (result MessagesFeedbacksResponse, err error) {
	if messageId == "" {
		return result, errors.Errorf("message_id is required")
	}

	api := cl.GetAPI(ApiMessagesFeedbacks)
	api = UpdateAPIParam(api, ApiParamMessageId, messageId)

	code, body, err := cl.sendPostRequestToAPI(ctx, api, payload)

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
