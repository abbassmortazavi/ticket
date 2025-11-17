package session

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	Store       *sessions.CookieStore
	sessionName string
)

// Init initializes the session store
func Init(authKey, encryptionKey []byte, appName string) {
	log.Println("Initializing session")
	Store = sessions.NewCookieStore(authKey, encryptionKey)
	sessionName = appName
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

// Set sets a session value
func Set(w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	log.Println("inja Set session")
	session, err := Store.Get(r, sessionName)
	log.Println("err=======>", err)
	if err != nil {
		return err
	}
	log.Println("Set", key, value)
	session.Values[key] = value
	return session.Save(r, w)
}

// Get gets a session value
func Get(r *http.Request, key string) (interface{}, error) {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		return nil, err
	}
	return session.Values[key], nil
}

// GetString gets a session value as string
func GetString(r *http.Request, key string) string {
	value, err := Get(r, key)
	if err != nil {
		return ""
	}
	return value.(string)
}

// Flash gets and removes a session value (like your Flash function)
func Flash(r *http.Request, w http.ResponseWriter, key string) string {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		return ""
	}
	value := session.Values[key]
	delete(session.Values, key)
	err = session.Save(r, w)
	if err != nil {
		return ""
	}
	return value.(string)
}
func Delete(w http.ResponseWriter, r *http.Request, key string) error {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		return err
	}
	delete(session.Values, key)
	return session.Save(r, w)
}

// Destroy destroys the entire session
func Destroy(w http.ResponseWriter, r *http.Request, key string) error {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	session.Values = map[interface{}]interface{}{}

	return session.Save(r, w)
}

// Exists checks if a session value exists
func Exists(r *http.Request, key string) bool {
	session, err := Store.Get(r, sessionName)
	if err != nil {
		return false
	}

	_, exists := session.Values[key]
	return exists
}
