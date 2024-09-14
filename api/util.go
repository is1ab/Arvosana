package api

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/is1ab/Arvosana/service/logger"
	"github.com/labstack/echo/v4"
)

func RegisterUtil(e *echo.Group) {
	type GetValidUserRequest struct {
		StudentId string `json:"student_id" query:"student_id"`
		Password  string `json:"password" query:"password"`
		Mode      string `json:"mode" query:"mode"`
	}

	e.GET("/valid_user", func(c echo.Context) error {
		ctx := c.Request().Context()
		l := logger.Ctx(ctx)

		// https://github.com/is1ab/Paper-Schedule-System/blob/master/api/auth/ntut_auth_util.py
		var data GetValidUserRequest
		err := c.Bind(&data)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.Unwrap(err))
		}

		payload := url.Values{}
		payload.Set("muid", data.StudentId)
		payload.Set("mpassword", data.Password)
		payload.Set("forceMobile", "app")
		payload.Set("md5Code", "1111")
		payload.Set("ssoId", "")

		switch data.Mode {
		case "ntut":
			client := &http.Client{}
			req, err := http.NewRequest("POST", "https://app.ntut.edu.tw/login.do", strings.NewReader(payload.Encode()))
			if err != nil {
				l.Errorln(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			resp, err := client.Do(req)
			if err != nil {
				l.Errorln(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
			defer resp.Body.Close()

		case "local":

		default:
			return echo.NewHTTPError(http.StatusBadRequest, "invalid mode '%s', only 'ntut' and 'local' are available", data.Mode)
		}

		return c.JSON(http.StatusOK, echo.Map{})
	})
}
