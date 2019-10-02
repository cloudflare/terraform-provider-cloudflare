module github.com/terraform-providers/terraform-provider-cloudflare

go 1.12

require (
	github.com/cloudflare/cloudflare-go v0.10.4
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	golang.org/x/net v0.0.0-20191007182048-72f939374954
	google.golang.org/genproto v0.0.0-20191007204434-a023cd5227bd // indirect
)

replace github.com/cloudflare/cloudflare-go => github.com/inge4pres/cloudflare-go v0.10.3
