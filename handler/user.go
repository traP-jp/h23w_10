package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
	"github.com/traP-jp/h23w_10/pkg/domain/repository"
)

type GetUserByIDResponse struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	IconURL  string          `json:"icon_url,omitempty"`
	UserType domain.UserType `json:"user_type,omitempty"`
}

type GetMeResponse struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	IconURL  string          `json:"icon_url,omitempty"`
	UserType domain.UserType `json:"user_type,omitempty"`
}

func (h *Handler) GetUserMe(c echo.Context) error {
	sess, err := session.Get(sessionName, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userID, ok := sess.Values["userID"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	user, err := h.urepo.FindUserByID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := GetMeResponse{
		ID:       user.ID,
		Name:     user.Name,
		IconURL:  user.IconURL.String(),
		UserType: user.UserType,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.urepo.FindUserByID(id)
	if errors.Is(err, repository.ErrNotFound) {
		return c.JSON(http.StatusNotFound, err.Error())
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := GetUserByIDResponse{
		ID:       user.ID,
		Name:     user.Name,
		IconURL:  user.IconURL.String(),
		UserType: user.UserType,
	}
	return c.JSON(http.StatusOK, response)
}
