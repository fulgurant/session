package session

import "errors"

// Errors
var (
	ErrNotFound = errors.New("not found")
)

type Session struct {
	id     string
	values map[string]interface{}
}

func newSession(id string) *Session {
	return &Session{
		id:     id,
		values: map[string]interface{}{},
	}
}

func (s *Session) Id() string {
	return s.id
}

func (s *Session) Get(key string) (interface{}, error) {
	v, ok := s.values[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (s *Session) Set(key string, value interface{}) error {
	s.values[key] = value
	return nil
}
