package muxclient_test

import (
	"context"
	"testing"

	cfv1 "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/provider"
	"github.com/stretchr/testify/assert"
)

func TestMuxClientInitialised(t *testing.T) {
	c1 := muxclient.Client{}
	assert.Empty(t, c1)

	v1Config := provider.Config{}
	v1Config.Email = "user@example.com"
	cfv1Client, _ := v1Config.Client(context.TODO())

	cfv2Client := cfv2.NewClient(
		option.WithAPIEmail("user@example.com"),
	)
	c2 := &muxclient.Client{
		V1: cfv1Client,
		V2: cfv2Client,
	}

	assert.NotEmpty(t, c2)
	assert.IsType(t, &cfv1.API{}, cfv1Client)
	assert.IsType(t, &cfv2.Client{}, cfv2Client)
	assert.IsType(t, &muxclient.Client{}, c2)
}
