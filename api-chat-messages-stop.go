package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type ChatMessagesStopPayload struct {
	User string `json:"user"`
}

type ChatMessagesStopResponse struct {
	Result string `json:"result"`
}

func (cl *Client) ChatMessagesStop(ctx context.Context, taskId string, payload ChatMessagesStopPayload) (result ChatMessagesStopResponse, err error) {
	if taskId == "" {
		return result, errors.New("task_id is required")
	}

	api := cl.GetAPI(ApiChatMessagesStop)
	api = UpdateAPIParam(api, ApiParamTaskId, taskId)

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
