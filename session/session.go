package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	SessionType    string
	CookieSecure   string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool
	// how long should sessions last ?
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 30
	}

	//should the cookie be persisted across requests?

	if strings.ToLower(c.CookiePersist) == "true" {
		persist = true
	}

	// must cookie be secure?
	if strings.ToLower(c.CookieSecure) == "true" {
		secure = true
	}

	// create session manager
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.Secure = secure
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store should be used?

	switch strings.ToLower(c.SessionType) {
	case "redis": // TODO
	case "mysql": // TODO
	case "postgres", "postgresql": // TODO
	default:
		// cookie
	}
	return session
}
