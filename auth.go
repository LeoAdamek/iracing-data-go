package iracing

import (
	"net/http"
)

// Credentials needed to authenticate with the API
type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// A CredentialsProvider is a func which attempts
// to return credentials
type CredentialsProvider func() (*Credentials, error)

// Static Credentials (pass in the credentials)
func StaticCredentials(username string, password string) CredentialsProvider {
	return func() (*Credentials, error) {
		return &Credentials{
			Username: username,
			Password: password,
		}, nil
	}
}

func (c *Client) Login() error {
	credentials, err := c.credentials()

	if err != nil {
		return err
	}

	return c.json(http.MethodPost, Host+"/auth", credentials, nil)
}
