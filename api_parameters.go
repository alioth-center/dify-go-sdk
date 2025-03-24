package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type GetParametersResponse struct {
	OpeningStatement              string `json:"opening_statement"`
	SuggestedQuestions            []any  `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"`
	} `json:"suggested_questions_after_answer"`
	SpeechToText struct {
		Enabled bool `json:"enabled"`
	} `json:"speech_to_text"`
	TextToSpeech struct {
		Enabled  bool   `json:"enabled"`
		Voice    string `json:"voice"`
		Language string `json:"language"`
	} `json:"text_to_speech"`
	RetrieverResource struct {
		Enabled bool `json:"enabled"`
	} `json:"retriever_resource"`
	AnnotationReply struct {
		Enabled bool `json:"enabled"`
	} `json:"annotation_reply"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	UserInputForm []struct {
		Paragraph struct {
			Label    string `json:"label"`
			Variable string `json:"variable"`
			Required bool   `json:"required"`
			Default  string `json:"default"`
		} `json:"paragraph"`
	} `json:"user_input_form"`
	SensitiveWordAvoidance struct {
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
		Configs []any  `json:"configs"`
	} `json:"sensitive_word_avoidance"`
	FileUpload struct {
		Image struct {
			Enabled         bool     `json:"enabled"`
			NumberLimits    int      `json:"number_limits"`
			Detail          string   `json:"detail"`
			TransferMethods []string `json:"transfer_methods"`
		} `json:"image"`
	} `json:"file_upload"`
	SystemParameters struct {
		ImageFileSizeLimit string `json:"image_file_size_limit"`
	} `json:"system_parameters"`
}

func (cl *Client) GetParameters(ctx context.Context) (result GetParametersResponse, err error) {
	api := cl.GetAPI(ApiParameters)
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
