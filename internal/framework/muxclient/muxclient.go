package muxclient

import (
	cfv1 "github.com/cloudflare/cloudflare-go"
	cfv2 "github.com/cloudflare/cloudflare-go/v2"
	cfv6 "github.com/cloudflare/cloudflare-go/v6"
)

// Client is the intermediatry structure to allow us to run both versions of the
// Go SDK alongside one another without collisions.
//
// The usage is resource dependent and no safe guards are included here.
type Client struct {
	V1 *cfv1.API
	V2 *cfv2.Client
	V6 *cfv6.Client
}
