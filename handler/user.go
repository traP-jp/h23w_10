package handler

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h23w_10/pkg/domain"
)

type GetUserByIDResponse struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name,omitempty"`
	IconURL  url.URL         `json:"icon_url,omitempty"`
	UserType domain.UserType `json:"user_type,omitempty"`
}

func (h *Handler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.urepo.FindUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := GetUserByIDResponse{
		ID:       user.ID,
		Name:     user.Name,
		IconURL:  user.IconURL,
		UserType: user.UserType,
	}
	return c.JSON(http.StatusOK, response)
}