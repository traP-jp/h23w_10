package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/traPtitech/go-traq"
)

var (
	botToken  string
	client    *traq.APIClient
	channelID string
)

func main() {
	botToken = os.Getenv("BOT_TOKEN")
	client = traq.NewAPIClient(traq.NewConfiguration())
	channelID = os.Getenv("CHANNEL_ID")

	http.HandleFunc("/question", postNewQuestionInfo)

	log.Println("start server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}

type Request struct {
	QuestionTitle   string `json:"question_title,omitempty"`
	QuestionContent string `json:"question_content,omitempty"`
	QuestionURL     string `json:"question_url,omitempty"`
}

func postNewQuestionInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = getAuth(ctx, botToken)

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg, _, err := client.MessageApi.
		PostMessage(ctx, channelID).
		PostMessageRequest(*traq.NewPostMessageRequest(
			fmt.Sprintf("## [%s](%s)\n%s", req.QuestionTitle, req.QuestionURL, req.QuestionContent),
		)).
		Execute()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("posted message: %s\n", msg.Id)

	w.WriteHeader(http.StatusOK)
}

func getAuth(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, traq.ContextAccessToken, token)
}
