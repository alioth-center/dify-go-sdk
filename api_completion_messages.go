package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type CompletionMessagesPayload struct {
	Inputs         map[string]any `json:"inputs"`
	User           string         `json:"user"`
	ConversationID string         `json:"conversation_id,omitempty"`
}

type completionMessagesPayloadWithResponseMode struct {
	CompletionMessagesPayload
	ResponseMode EnumResponseMode `json:"response_mode,omitempty"`
}

type CompletionMessagesResponse struct {
	Event     string `json:"event"`
	TaskID    string `json:"task_id"`
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	Mode      string `json:"mode"`
	Answer    string `json:"answer"`
	Metadata  any    `json:"metadata"`
	CreatedAt int    `json:"created_at"`
}

func (cl *Client) CompletionMessages(ctx context.Context, payload CompletionMessagesPayload) (result CompletionMessagesResponse, err error) {
	if len(payload.Inputs) == 0 {
		payload.Inputs = make(map[string]any)
	}

	request := completionMessagesPayloadWithResponseMode{
		CompletionMessagesPayload: payload,
		ResponseMode:              ResponseModeBlocking,
	}
	api := cl.GetAPI(ApiCompletionMessages)
	code, body, err := cl.sendPostRequestToAPI(ctx, api, request)

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

func (cl *Client) CompletionMessagesStreaming(ctx context.Context, payload CompletionMessagesPayload) (result string, err error) {
	if len(payload.Inputs) == 0 {
		payload.Inputs = make(map[string]any)
	}

	request := completionMessagesPayloadWithResponseMode{
		CompletionMessagesPayload: payload,
		ResponseMode:              ResponseModeStreaming,
	}
	api := cl.GetAPI(ApiCompletionMessages)
	code, body, err := cl.sendPostRequestToAPI(ctx, api, request)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	// TODO
	return string(body), nil
}
