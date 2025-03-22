package dify

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
)

type ChatMessagesPayload struct {
	/**
	允许传入 App 定义的各变量值。
	inputs 参数包含了多组键值对（ApiKey/Value pairs），每组的键对应一个特定变量，每组的值则是该变量的具体值。
	如果变量是文件类型，请指定一个包含以下 files 中所述键的对象。
	默认 {}
	*/
	Inputs map[string]any `json:"inputs"`
	Query  string         `json:"query"` // 用户输入/提问内容。
	// （选填）会话 ID，需要基于之前的聊天记录继续对话，必须传之前消息的 conversation_id。
	ConversationID string `json:"conversation_id,omitempty"`
	// 用户标识，用于定义终端用户的身份，方便检索、统计。 由开发者定义规则，需保证用户标识在应用内唯一。
	User string `json:"user"`
	// 文件列表，适用于传入文件结合文本理解并回答问题，仅当模型支持 Vision 能力时可用。
	Files []ChatMessagesPayloadFile `json:"files,omitempty"`
}

type chatMessagesPayloadWithResponseMode struct {
	ChatMessagesPayload
	ResponseMode EnumResponseMode `json:"response_mode"`
}

type ChatMessagesPayloadFile struct {
	Type           EnumFileType           `json:"type"`
	TransferMethod FileTransferMethodEnum `json:"transfer_method"`          // 传递方式
	URL            string                 `json:"url,omitempty"`            // 图片地址。（仅当传递方式为 remote_url 时）
	UploadFileID   string                 `json:"upload_file_id,omitempty"` // 上传文件 ID。（仅当传递方式为 local_file 时）。
}

type ChatCompletionResponse struct {
	MessageID      string `json:"message_id"`
	ConversationID string `json:"conversation_id"`
	Mode           string `json:"mode"`
	Answer         string `json:"answer"`
	Metadata       any    `json:"metadata"`
	CreatedAt      int    `json:"created_at"`
}

func (cl *Client) ChatMessages(ctx context.Context, payload ChatMessagesPayload) (result ChatCompletionResponse, err error) {
	if len(payload.Inputs) == 0 {
		payload.Inputs = make(map[string]any)
	}

	request := chatMessagesPayloadWithResponseMode{
		ChatMessagesPayload: payload,
		ResponseMode:        ResponseModeBlocking,
	}
	code, body, err := cl.sendPostRequestToAPI(ctx, cl.GetAPI(ApiChatMessages), request)

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

func (cl *Client) ChatMessagesStreaming(ctx context.Context, payload ChatMessagesPayload) (result string, err error) {
	if len(payload.Inputs) == 0 {
		payload.Inputs = make(map[string]any)
	}

	request := chatMessagesPayloadWithResponseMode{
		ChatMessagesPayload: payload,
		ResponseMode:        ResponseModeStreaming,
	}
	code, body, err := cl.sendPostRequestToAPI(ctx, cl.GetAPI(ApiChatMessages), request)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}
	// TODO streaming
	// if !strings.Contains(resp.Header.Get("Content-Type"), "text/event-stream") {
	// 	return "", errors.Errorf("response is not a streaming response")
	// }

	return string(body), nil
}
