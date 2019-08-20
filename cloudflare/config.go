package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

type Config struct {
	Email    string
	Token    string
	APIToken string
	Options  []cloudflare.Option
}

// Client returns a new client for accessing cloudflare.
func (c *Config) Client() (*cloudflare.API, error) {
	var err error
	var client *cloudflare.API

	if c.APIToken != "" {
		client, err = cloudflare.NewWithAPIToken(c.APIToken, c.Options...)
	} else {
		client, err = cloudflare.New(c.Token, c.Email, c.Options...)
	}
	if err != nil {
		return nil, fmt.Errorf("Error creating new Cloudflare client: %s", err)
	}

	log.Printf("[INFO] Cloudflare Client configured for user: %s", c.Email)
	return client, nil
}
