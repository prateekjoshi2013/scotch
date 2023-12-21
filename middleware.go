package scotch

import "net/http"

// The LoadAndSave method uses the s.Session field, which is of type sessions.Session,
// to load and save session data. This is done by calling the Load and Save methods on the session object,
// which retrieve and store data in a backend data store such as a database
func (s *Scotch) SessionLoad(next http.Handler) http.Handler {
	return s.Session.LoadAndSave(next)
}
