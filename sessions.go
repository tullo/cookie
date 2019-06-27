package cookie

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	cookies = make(map[string]*sessions.CookieStore)

	ErrCookieNotFound = errors.New("Cookie not found")
)

type CookieOptions struct {
	AuthenticationKey []byte
	EncryptionKey     []byte
	Domain            string
	Path              string
	MaxAge            int
	Secure            bool
}

func NewCookieStore(name string, cookieOptions CookieOptions) {
	if cookies[name] == nil {
		cookies[name] = sessions.NewCookieStore(cookieOptions.AuthenticationKey, cookieOptions.EncryptionKey)
		cookies[name].Options = &sessions.Options{
			Domain: cookieOptions.Domain,
			Path:   cookieOptions.Path,
			MaxAge: cookieOptions.MaxAge,
			Secure: cookieOptions.Secure,
		}
	}
}

func GetSession(r *http.Request, name string) (*sessions.Session, error) {
	if cookies[name] == nil {
		return nil, ErrCookieNotFound
	}
	return cookies[name].Get(r, name)
}

func ClearSession(sess *sessions.Session) {
	for key := range sess.Values {
		delete(sess.Values, key)
	}
}
