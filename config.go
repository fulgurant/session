package session

// Config values for session management
type Config struct {
	KeyLength int    `kong:"prefix='session-',default='32',help='Length of session key',hidden=''"`
	Encoding  string `kong:"prefix='session-',default='base64',enum='none,base64',help='Encoding'"`
}
