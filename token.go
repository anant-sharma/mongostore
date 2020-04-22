package mongostore

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// TokenGetSeter Interface
type TokenGetSeter interface {
	GetToken(req *http.Request, name string) (string, error)
	SetToken(rw http.ResponseWriter, name, value string, options *sessions.Options)
}

// CookieToken Struct
type CookieToken struct{}

// GetToken Struct
func (c *CookieToken) GetToken(req *http.Request, name string) (string, error) {
	cook, err := req.Cookie(name)
	if err != nil {
		return "", err
	}

	return cook.Value, nil
}

// SetToken Struct
func (c *CookieToken) SetToken(rw http.ResponseWriter, name, value string,
	options *sessions.Options) {
	http.SetCookie(rw, sessions.NewCookie(name, value, options))
}
