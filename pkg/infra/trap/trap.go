package trap

import (
	"context"
	"net/http"
	"net/url"

	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traPtitech/go-traq"
)

type TrapService struct {
	traqClient *traq.APIClient
}

func NewTrapService(traqClient *traq.APIClient) *TrapService {
	return &TrapService{
		traqClient: traqClient,
	}
}

func (t *TrapService) GetMe(ctx context.Context, token string) (*domain.User, error) {
	user, res, err := t.traqClient.MeApi.GetMe(getAuth(ctx, token)).Execute()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	baseURL, err := url.Parse("https://q.trap.jp/api/v3")
	if err != nil {
		return nil, err
	}
	url := baseURL.JoinPath("public", "icon", user.Name)
	result := &domain.User{
		ID:          user.Id,
		Name:        user.Name,
		DisplayName: user.DisplayName,
		IconURL:     *url,
		UserType:    "trap",
	}
	return result, nil
}

func getAuth(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, traq.ContextAccessToken, token)
}
