package middleware

import (
	"github.com/labstack/echo"
	. "fmt"
)

func ReadCookieMiddleware( next echo.HandlerFunc ) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		if err != nil {
			return err
		}
		Println("cookieName"+cookie.Name)
		Println("cookieName"+cookie.Value)
		middlewareErr := next(c)
		Println("cookieAfter")
		return middlewareErr
	}
}

func CheckCookieForRouter ( c echo.Context ) error {
	cookie, err := c.Cookie("session")
		if err != nil {
			return err
	}
	Println("cookieName "+cookie.Name)
	Println("cookieValue "+cookie.Value)

	if len(cookie.Value) > 0 {
		return c.String(200, "Cookie is already set")
	}
	return c.String(401, "Cookie does not set")
}