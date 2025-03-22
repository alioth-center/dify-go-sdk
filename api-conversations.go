package dify

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type ConversationsQuery struct {
	User   string `json:"user"`
	LastId string `json:"last_id"`
	Limit  int    `json:"limit"`
	SortBy string `json:"sort_by"`
}

type ConversationsResponse struct {
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
	Data    []struct {
		ID     string `json:"id"`
		Name   string `json:"name,omitempty"`
		Inputs struct {
			Book   string `json:"book"`
			MyName string `json:"myName"`
		} `json:"inputs,omitempty"`
		Status    string `json:"status,omitempty"`
		CreatedAt int    `json:"created_at,omitempty"`
	} `json:"data"`
}

func (cl *Client) Conversations(ctx context.Context, query ConversationsQuery) (result ConversationsResponse, err error) {
	api := cl.GetAPI(ApiConversations)

	// TODO
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

type DeleteConversationsPayload struct {
	User string `json:"user"`
}

type DeleteConversationsResponse struct {
	Result string `json:"result"`
}

func (cl *Client) DeleteConversation(ctx context.Context, conversationId string, payload DeleteConversationsPayload) (result DeleteConversationsResponse, err error) {
	if conversationId == "" {
		return result, errors.Errorf("conversation_id is required")
	}

	api := cl.GetAPI(ApiConversationsDelete)
	api = UpdateAPIParam(api, ApiParamConversationId, conversationId)

	buf, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest(http.MethodDelete, api, bytes.NewBuffer(buf))
	if err != nil {
		return result, errors.Errorf("could not create a new request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+cl.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, errors.Errorf("status code: %d, could not read the body", resp.StatusCode)
		}
		return result, errors.Errorf("status code: %d, %s", resp.StatusCode, bodyText)
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bodyText, &result)
	if err != nil {
		return result, errors.Wrap(err, "failed to unmarshal the response")
	}
	return result, nil
}

type RenameConversationsPayload struct {
	User string `json:"user"`
}

type RenameConversationsResponse struct {
	Result string `json:"result"`
}

func (cl *Client) RenameConversation(ctx context.Context, conversationId string, payload RenameConversationsPayload) (result RenameConversationsResponse, err error) {
	if conversationId == "" {
		return result, errors.Errorf("conversation_id is required")
	}

	api := cl.GetAPI(ApiConversationsRename)
	api = UpdateAPIParam(api, ApiParamConversationId, conversationId)

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
