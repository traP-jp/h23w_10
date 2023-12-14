package handler

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

func (h *Handler) GetAuthParams(c echo.Context) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	state := make([]byte, 16)
	_, _ = rand.Read(state)
	hexState := hex.EncodeToString(state)
	sess.Values["state"] = hexState

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	url := h.oauth2Config.AuthCodeURL(hexState)

	return c.JSON(http.StatusOK, map[string]string{"url": url})
}

// TODO: traqのgetmeを叩いて、ユーザー情報をdbに保存する
func (h *Handler) Oauth2Callback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	sess, err := session.Get(sessionName, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sessState, ok := sess.Values["state"].(string)
	if !ok || sessState != state {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid state")
	}

	token, err := h.oauth2Config.Exchange(c.Request().Context(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	sess.Values["token"] = token

	// get user from traq
	user, err := h.trapSvc.GetMe(c.Request().Context(), token.AccessToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	sess.Values["userID"] = user.ID

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// check user exists
	_, err = h.urepo.FindUserByID(user.ID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if err == nil {
		return c.Redirect(http.StatusFound, "/")
	}

	// create user
	_, err = h.urepo.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusFound, "/")
}
