package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// Errors
var (
	ErrSessionNotFound = errors.New("session not found")
)

// Handler is used to handle sessions
type Handler struct {
	options *Options

	sessions map[string]*Session
}

// New instantiates a session generator
func New(options *Options) (*Handler, error) {
	if options.config == nil {
		return nil, errors.New("nil config not allowed")
	}

	return &Handler{
		options:  options,
		sessions: map[string]*Session{},
	}, nil
}

// New Creates a new session with a random id
func (h *Handler) New() (*Session, error) {
	id, err := h.generateID()
	if err != nil {
		return nil, err
	}
	s := newSession(id)
	h.sessions[s.id] = s
	return s, nil
}

func (h *Handler) Get(id string) (*Session, error) {
	s, found := h.sessions[id]
	if !found {
		return nil, ErrSessionNotFound
	}
	return s, nil
}

// generateID creates a random session id
func (h *Handler) generateID() (string, error) {
	b := make([]byte, h.options.config.KeyLength)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	switch h.options.config.Encoding {
	case "base64":
		return base64.StdEncoding.EncodeToString(b), nil
	default:
		return string(b), nil
	}
}
