package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type MessagesQuery struct {
	ConversationId string `json:"conversation_id"`
	User           string `json:"user"`
	FirstId        string `json:"first_id"`
	Limit          int    `json:"limit"`
}

type MessagesResponse struct {
	Limit   int  `json:"limit"`
	HasMore bool `json:"has_more"`
	Data    []struct {
		ID             string `json:"id"`
		ConversationID string `json:"conversation_id"`
		Inputs         struct {
			Name string `json:"name"`
		} `json:"inputs"`
		Query              string `json:"query"`
		Answer             string `json:"answer"`
		MessageFiles       []any  `json:"message_files"`
		Feedback           any    `json:"feedback"`
		RetrieverResources []struct {
			Position     int     `json:"position"`
			DatasetID    string  `json:"dataset_id"`
			DatasetName  string  `json:"dataset_name"`
			DocumentID   string  `json:"document_id"`
			DocumentName string  `json:"document_name"`
			SegmentID    string  `json:"segment_id"`
			Score        float64 `json:"score"`
			Content      string  `json:"content"`
		} `json:"retriever_resources"`
		AgentThoughts []any `json:"agent_thoughts"`
		CreatedAt     int   `json:"created_at"`
	} `json:"data"`
}

func (cl *Client) Messages(ctx context.Context, query MessagesQuery) (result MessagesResponse, err error) {
	// TODO
	if query.ConversationId == "" {
		return result, errors.New("conversation_id is required")
	}

	api := cl.GetAPI(ApiMessages)

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
