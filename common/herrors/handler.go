package herrors

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	he := new(echo.HTTPError)
	ok := errors.As(err, &he)

	if !ok {
		ErrorHandler(NewInternalServer(err), c)
		return
	}

	if he.Internal != nil {
		c.Logger().Error("internal error: ", he.Internal)
	}

	if se := c.JSON(he.Code, NewResponse(he.Message)); se != nil {
		c.Logger().Error("send error: ", se)
	}
}
