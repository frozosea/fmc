package login_provider

import "sync"

type Store struct {
	mu        sync.Mutex
	username  string
	password  string
	basicAuth string
	authToken string
}

func NewStore(username string, password string, basicAuth string, authToken string) *Store {
	return &Store{username: username, password: password, basicAuth: basicAuth, authToken: authToken}
}

func (s *Store) Username() string {
	return s.username
}

func (s *Store) SetUsername(username string) {
	s.username = username
}

func (s *Store) Password() string {
	return s.password
}

func (s *Store) SetPassword(password string) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.password = password
}

func (s *Store) BasicAuth() string {
	return s.basicAuth
}

func (s *Store) SetBasicAuth(basicAuth string) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.basicAuth = basicAuth
}

func (s *Store) AuthToken() string {
	return s.authToken
}

func (s *Store) SetAuthToken(authToken string) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.authToken = authToken
}
