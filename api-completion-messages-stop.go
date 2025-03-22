package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type CompletionMessagesStopPayload struct {
	User string `json:"user"`
}

type CompletionMessagesStopResponse struct {
	Result string `json:"result"`
}

func (cl *Client) CompletionMessagesStop(ctx context.Context, taskId string, payload CompletionMessagesStopPayload) (result CompletionMessagesStopResponse, err error) {
	if taskId == "" {
		return result, errors.Errorf("task_id is required")
	}

	api := cl.GetAPI(ApiCompletionMessagesStop)
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
