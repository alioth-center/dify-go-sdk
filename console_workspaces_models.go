package dify

// https://dify.lab.io/console/api/workspaces/current/models/model-types/rerank

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type ListWorkspacesRerankModelsResponse struct {
	Data []ListWorkspacesRerankItem `json:"data"`
}

type ListWorkspacesRerankItem struct {
	Provider string `json:"provider"`
	Label    struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"label"`
	IconSmall struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"icon_small"`
	IconLarge struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"icon_large"`
	Status string                      `json:"status"`
	Models []ListWorkspacesRerankModel `json:"models"`
}

type ListWorkspacesRerankModel struct {
	Model string `json:"model"`
	Label struct {
		ZhHans string `json:"zh_Hans"`
		EnUS   string `json:"en_US"`
	} `json:"label"`
	ModelType       string `json:"model_type"`
	Features        any    `json:"features"`
	FetchFrom       string `json:"fetch_from"`
	ModelProperties struct {
		ContextSize int `json:"context_size"`
	} `json:"model_properties"`
	Deprecated bool   `json:"deprecated"`
	Status     string `json:"status"`
}

func (cl *Client) ListWorkspacesRerankModels(ctx context.Context) (result ListWorkspacesRerankModelsResponse, err error) {
	api := cl.GetConsoleAPI(ConsoleApiWorkspacesRerankModel)

	code, body, err := cl.sendGetRequestToConsole(ctx, api)

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

type GetCurrentWorkspaceRerankDefaultModelResponse struct {
	Data any `json:"data"`
}

func (cl *Client) GetCurrentWorkspaceRerankDefaultModel(ctx context.Context) (result GetCurrentWorkspaceRerankDefaultModelResponse, err error) {
	api := cl.GetConsoleAPI(ConsoleApiCurrentWorkspaceRerankModel)

	code, body, err := cl.sendGetRequestToConsole(ctx, api)

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
