package dify

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func (cl *Client) DeleteDatasets(datasets_id string) (ok bool, err error) {
	if datasets_id == "" {
		return false, errors.Errorf("datasets_id is required")
	}

	api := cl.GetConsoleAPI(ConsoleApiDatasetsDelete)
	api = UpdateAPIParam(api, ConsoleApiParamDatasetsId, datasets_id)

	req, err := http.NewRequest("DELETE", api, nil)
	if err != nil {
		return false, errors.Errorf("could not create a new request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.ConsoleToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, errors.Errorf("status code: %d, could not read the body", resp.StatusCode)
		}
		return false, errors.Errorf("status code: %d, %s", resp.StatusCode, bodyText)
	}

	return true, nil
}

type CreateDatasetsPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// TODO
}

type CreateDatasetsResponse struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Description            any    `json:"description"`
	Provider               string `json:"provider"`
	Permission             string `json:"permission"`
	DataSourceType         any    `json:"data_source_type"`
	IndexingTechnique      any    `json:"indexing_technique"`
	AppCount               int    `json:"app_count"`
	DocumentCount          int    `json:"document_count"`
	WordCount              int    `json:"word_count"`
	CreatedBy              string `json:"created_by"`
	CreatedAt              int    `json:"created_at"`
	UpdatedBy              string `json:"updated_by"`
	UpdatedAt              int    `json:"updated_at"`
	EmbeddingModel         any    `json:"embedding_model"`
	EmbeddingModelProvider any    `json:"embedding_model_provider"`
	EmbeddingAvailable     any    `json:"embedding_available"`
	RetrievalModelDict     struct {
		SearchMethod    string `json:"search_method"`
		RerankingEnable bool   `json:"reranking_enable"`
		RerankingModel  struct {
			RerankingProviderName string `json:"reranking_provider_name"`
			RerankingModelName    string `json:"reranking_model_name"`
		} `json:"reranking_model"`
		TopK                  int  `json:"top_k"`
		ScoreThresholdEnabled bool `json:"score_threshold_enabled"`
		ScoreThreshold        any  `json:"score_threshold"`
	} `json:"retrieval_model_dict"`
	Tags []any `json:"tags"`
}

func (cl *Client) CreateDatasets(ctx context.Context, payload CreateDatasetsPayload) (result CreateDatasetsResponse, err error) {
	api := cl.GetConsoleAPI(ConsoleApiDatasetsCreate)

	code, body, err := cl.sendPostRequestToConsole(ctx, api, payload)

	err = CommonRiskForSendRequestWithCode(code, err, http.StatusCreated)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.Wrap(err, "failed to unmarshal the response")
	}
	return result, nil
}

type ListDatasetsQuery struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListDatasetsResponse struct {
	Page    int                `json:"page"`
	Limit   int                `json:"limit"`
	Total   int                `json:"total"`
	HasMore bool               `json:"has_more"`
	Data    []ListDatasetsItem `json:"data"`
}

type ListDatasetsItem struct {
	ID             string                  `json:"id"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	Mode           string                  `json:"mode"`
	Icon           string                  `json:"icon"`
	IconBackground string                  `json:"icon_background"`
	ModelConfig    ListDatasetsModelConfig `json:"model_config"`
	CreatedAt      int                     `json:"created_at"`
	Tags           []any                   `json:"tags"`
}

type ListDatasetsModelConfig struct {
	Model     ListDatasetsModelConfigDetail `json:"model"`
	PrePrompt string                        `json:"pre_prompt"`
}

type ListDatasetsModelConfigDetail struct {
	Provider         string `json:"provider"`
	Name             string `json:"name"`
	Mode             string `json:"mode"`
	CompletionParams struct {
	} `json:"completion_params"`
}

func (cl *Client) ListDatasets(ctx context.Context, query ListDatasetsQuery) (result ListDatasetsResponse, err error) {
	if query.Page < 1 {
		return result, errors.Errorf("page should be greater than 0")
	}
	if query.Limit < 1 {
		return result, errors.Errorf("limit should be greater than 0")
	}

	api := cl.GetConsoleAPI(ConsoleApiDatasetsList)
	api = fmt.Sprintf("%s?page=%d&limit=%d", api, query.Page, query.Limit)

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
