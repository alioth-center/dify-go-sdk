package dify

import "strings"

const (
	ApiCompletionMessages     = "/completion-messages"
	ApiCompletionMessagesStop = "/completion-messages/:task_id/stop"

	ApiChatMessages     = "/chat-messages"
	ApiChatMessagesStop = "/chat-messages/:task_id/stop"

	ApiMessages          = "/messages"
	ApiMessagesSuggested = "/messages/:message_id/suggested"
	ApiMessagesFeedbacks = "/messages/:message_id/feedbacks"

	ApiConversations       = "/conversations"
	ApiConversationsDelete = "/conversations/:conversation_id"
	ApiConversationsRename = "/conversations/:conversation_id/name"

	ApiFileUpload = "/files/upload"
	ApiParameters = "/parameters"
	ApiMeta       = "/meta"

	ApiAudioToText = "/audio-to-text"
	ApiTextToAudio = "/text-to-audio"

	ApiParamTaskId         = ":task_id"
	ApiParamMessageId      = ":message_id"
	ApiParamConversationId = ":conversation_id"

	ConsoleApiFileUpload = "/files/upload?source=datasets"
	ConsoleApiLogin      = "/login"

	ConsoleApiParamDatasetsId = ":datasets_id"

	ConsoleApiDatasetsCreate     = "/datasets"
	ConsoleApiDatasetsList       = "/datasets"
	ConsoleApiDatasetsDelete     = "/datasets/:datasets_id"
	ConsoleApiDatasetsInit       = "/datasets/init"
	ConsoleApiDatasetsInitStatus = "/datasets/:datasets_id/indexing-status"

	ConsoleApiWorkspacesRerankModel       = "/workspaces/current/models/model-types/rerank"
	ConsoleApiCurrentWorkspaceRerankModel = "/workspaces/current/default-model?model_type=rerank"
)

func (cl *Client) GetAPI(api string) string {
	return cl.Host + api
}

func (cl *Client) GetConsoleAPI(api string) string {
	return cl.ConsoleHost + api
}

func UpdateAPIParam(api, key, value string) string {
	return strings.ReplaceAll(api, key, value)
}
