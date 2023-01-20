package provider

import (
	"context"
	"errors"
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

func (c *Config) Client(ctx context.Context) (*cloudflare.API, error) {
	var err error
	var client *cloudflare.API

	if c.APIUserServiceKey != "" {
		client, err = cloudflare.NewWithUserServiceKey(c.APIUserServiceKey, c.Options...)
	} else if c.APIToken != "" {
		client, err = cloudflare.NewWithAPIToken(c.APIToken, c.Options...)
	} else if c.APIKey != "" {
		client, err = cloudflare.New(c.APIKey, c.Email, c.Options...)
	} else {
		return nil, errors.New("no credentials detected")
	}

	if err != nil {
		return nil, fmt.Errorf("error creating new Cloudflare client: %w", err)
	}

	tflog.Info(ctx, fmt.Sprintf("cloudflare Client configured for user: %s", c.Email))
	return client, nil
}
