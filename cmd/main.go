package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alioth-center/dify-go-sdk"
)

func main() {
	ctx := context.Background()

	APIKey := os.Getenv("DIFY_API_KEY")
	if APIKey == "" {
		fmt.Println("DIFY_API_KEY is required")
		return
	}

	APIHost := os.Getenv("DIFY_API_HOST")
	if APIHost == "" {
		fmt.Println("DIFY_API_HOST is required")
		return
	}

	ConsoleHost := os.Getenv("DIFY_CONSOLE_HOST")

	client, err := dify.NewClient(dify.ClientConfig{Key: APIKey, Host: APIHost, ConsoleHost: ConsoleHost})
	if err != nil {
		log.Fatalf("failed to create Client: %v\n", err)
		return
	}

	metas, err := client.GetMeta(ctx)
	if err != nil {
		log.Fatalf("failed to get meta: %v\n", err)
		return
	}
	fmt.Println(metas)

	parametersResponse, err := client.GetParameters(ctx)
	if err != nil {
		log.Fatalf("failed to get parameters: %v\n", err)
		return
	}
	fmt.Println(parametersResponse)

	msgID := CompletionMessages(ctx, client)
	//FileUpload(client)
	CompletionMessagesStop(ctx, client)
	MessagesFeedbacks(ctx, client, msgID)
	// TextToAudio(client)

	ConsoleUser := os.Getenv("DIFY_CONSOLE_USER")
	ConsolePass := os.Getenv("DIFY_CONSOLE_PASS")
	if ConsoleUser != "" && ConsolePass != "" {
		log.Println("Get Console Token")
		token := GetUserToken(ctx, client, ConsoleUser, ConsolePass)
		if token == "" {
			log.Fatalf("failed to get console token\n")
		}
		client.ConsoleToken = token

		// Create datasets
		var datasetsID string
		log.Println("Create datasets")
		createResult, err := client.CreateDatasets(ctx, dify.CreateDatasetsPayload{
			Name: "test datasets",
		})
		if err != nil {
			log.Fatalf("failed to create datasets: %v\n", err)
			return
		}
		datasetsID = createResult.ID
		log.Println(createResult)

		// List datasets
		log.Println("List datasets")
		ListResult, err := client.ListDatasets(ctx, dify.ListDatasetsQuery{
			Page:  1,
			Limit: 30,
		})
		if err != nil {
			log.Fatalf("failed to list datasets: %v\n", err)
			return
		}
		if len(ListResult.Data) == 0 {
			log.Fatalf("no datasets found\n")
			return
		}
		for _, dataset := range ListResult.Data {
			if dataset.ID == datasetsID {
				// Delete datasets
				log.Println("Delete datasets")
				result, err := client.DeleteDatasets(datasetsID)
				if err != nil {
					log.Fatalf("failed to delete datasets: %v\n", err)
					return
				}
				log.Println(result)
			}
		}

		// Get the list of rerank models
		log.Println("List rerank models")
		reRankModels, err := client.ListWorkspacesRerankModels(ctx)
		if err != nil {
			log.Println("failed to list rerank models:", err)
		} else {
			log.Println(reRankModels)
		}

		log.Println("Upload file to datasets")
		err = os.WriteFile("testfile-for-dify-database.txt", []byte("test file for dify database"), 0644)
		if err != nil {
			log.Fatalf("failed to create file: %v\n", err)
			return
		}
		//result, err := client.DatasetsFileUpload("testfile-for-dify-database.txt", "testfile-for-dify-database.txt")
		//if err != nil {
		//	log.Fatalf("failed to upload file to datasets: %v\n", err)
		//	return
		//}
		//fileID := result.ID
		//log.Println(result)

		//initResult, err := client.InitDatasetsByUploadFile(ctx, []string{fileID})
		//if err != nil {
		//	log.Fatalf("failed to init datasets by upload file: %v\n", err)
		//	return
		//}
		//log.Println(initResult)

		//initStatus, err := client.InitDatasetsIndexingStatus(initResult.Dataset.ID)
		//if err != nil {
		//	log.Fatalf("failed to get init datasets indexing status: %v\n", err)
		//	return
		//}
		//log.Println(initStatus)
	}
}

func CompletionMessages(ctx context.Context, client *dify.Client) (messageID string) {
	// normal response
	// TODO
	completionMessagesResponse, err := client.CompletionMessages(ctx, dify.CompletionMessagesPayload{
		Inputs: map[string]any{
			"query": "hey",
		},
	})
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesResponse)
	fmt.Println()

	// streaming response
	completionMessagesStreamingResponse, err := client.CompletionMessagesStreaming(ctx, dify.CompletionMessagesPayload{
		Inputs: map[string]any{
			"query": "hey",
		},
	})
	if err != nil {
		log.Fatalf("failed to get completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesStreamingResponse)
	fmt.Println()

	return completionMessagesResponse.MessageID
}

// TODO
//func FileUpload(client *dify.Client) {
//	fileUploadResponse, err := client.FileUpload("./README.md", "readme.md")
//	if err != nil {
//		log.Fatalf("failed to upload file: %v\n", err)
//		return
//	}
//	fmt.Println(fileUploadResponse)
//	fmt.Println()
//}

func CompletionMessagesStop(ctx context.Context, client *dify.Client) {
	completionMessagesStopResponse, err := client.CompletionMessagesStop(ctx, "0d2bd315-d4de-476f-ad5e-faaa00d571ea",
		dify.CompletionMessagesStopPayload{})
	if err != nil {
		log.Fatalf("failed to stop completion messages: %v\n", err)
		return
	}
	fmt.Println(completionMessagesStopResponse)
	fmt.Println()
}

func MessagesFeedbacks(ctx context.Context, client *dify.Client, messageID string) {
	messagesFeedbacksResponse, err := client.MessagesFeedbacks(ctx, messageID, dify.MessagesFeedbacksPayload{
		Rating: "like",
	})
	if err != nil {
		log.Fatalf("failed to get messages feedbacks: %v\n", err)
		return
	}
	fmt.Println(messagesFeedbacksResponse)
	fmt.Println()
}

//func TextToAudio(client *dify.Client) {
//	textToAudioResponse, err := client.TextToAudio("hello world")
//	if err != nil {
//		log.Fatalf("failed to get text to audio: %v\n", err)
//		return
//	}
//	fmt.Println(textToAudioResponse)
//	fmt.Println()
//
//	textToAudioStreamingResponse, err := client.TextToAudioStreaming("hello world")
//	if err != nil {
//		log.Fatalf("failed to get text to audio streaming: %v\n", err)
//		return
//	}
//	fmt.Println(textToAudioStreamingResponse)
//	fmt.Println()
//}

func GetUserToken(ctx context.Context, client *dify.Client, email, password string) string {
	result, err := client.UserLogin(ctx, email, password)
	if err != nil {
		log.Fatalf("failed to login: %v\n", err)
		return ""
	}
	return result.Data
}
