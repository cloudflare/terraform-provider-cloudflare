package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type Config struct {
	Email             string
	APIKey            string
	APIUserServiceKey string
	APIToken          string
	Options           []cloudflare.Option
}

// Client returns a new client for accessing cloudflare.
func (c *Config) Client() (*cloudflare.API, error) {
	var err error
	var client *cloudflare.API
	ctx := context.Background()

	if c.APIToken != "" {
		client, err = cloudflare.NewWithAPIToken(c.APIToken, c.Options...)
	} else {
		client, err = cloudflare.New(c.APIKey, c.Email, c.Options...)
	}
	if err != nil {
		return nil, fmt.Errorf("error creating new Cloudflare client: %w", err)
	}

	if c.APIUserServiceKey != "" {
		client.APIUserServiceKey = c.APIUserServiceKey
	}

	tflog.Info(ctx, fmt.Sprintf("cloudflare Client configured for user: %s", c.Email))
	return client, nil
}
