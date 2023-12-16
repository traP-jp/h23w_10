package trap

import (
	"context"
	"fmt"

	"github.com/traPtitech/go-traq"
)

func (t *TrapService) PostQuestionInfo(title, content, url string) error {
	_, res, err := t.traqClient.MessageApi.
		PostMessage(getAuth(context.Background(), t.botToken), t.channelID).
		PostMessageRequest(*traq.NewPostMessageRequest(fmt.Sprintf("## [%s](%s)\n%s", title, url, content))).
		Execute()
	defer res.Body.Close()

	return err
}
