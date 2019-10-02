module github.com/terraform-providers/terraform-provider-cloudflare

go 1.12

require (
	github.com/cloudflare/cloudflare-go v0.10.3
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/terraform v0.12.8
	github.com/pkg/errors v0.8.1
	golang.org/x/net v0.0.0-20191002035440-2ec189313ef0
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c // indirect
)

replace github.com/cloudflare/cloudflare-go => github.com/inge4pres/cloudflare-go v0.10.3
