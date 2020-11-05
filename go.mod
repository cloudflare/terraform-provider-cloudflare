module github.com/cloudflare/terraform-provider-cloudflare

go 1.15

require (
	github.com/cloudflare/cloudflare-go v0.13.4
	github.com/hashicorp/go-cleanhttp v0.5.1
	github.com/hashicorp/terraform-plugin-sdk v1.16.0
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20201026091529-146b70c837a4
)

replace github.com/cloudflare/cloudflare-go v0.13.4 => github.com/UrosSimovic/cloudflare-go v0.11.2-0.20201105063726-a3eaafdd1160
