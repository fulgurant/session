package session

import "encoding/base64"

// Handler is used to handle sessions
type Handler struct {
	options *Options
}

// New instantiates a session generator
func New(options *Options) (*Handler, error) {
	return &Handler{
		options: options,
	}, nil
}

// New generates a new session string
func (h *Handler) New() (string, error) {
	b := make([]byte, h.options.config.KeyLength)

	switch h.options.config.Encoding {
	case "base64":
		return base64.StdEncoding.EncodeToString(b), nil
	default:
		return string(b), nil
	}
}
