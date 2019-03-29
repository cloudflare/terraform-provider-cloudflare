module github.com/terraform-providers/terraform-provider-cloudflare

go 1.12

require (
	github.com/cloudflare/cloudflare-go v0.8.6-0.20190328165822-4034ff974d99
	github.com/hashicorp/go-cleanhttp v0.5.0
	github.com/hashicorp/terraform v0.12.0-beta1
	github.com/pkg/errors v0.8.1
)

// Necessary to build without depending on bazaar
replace (
	labix.org/v2/mgo => gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
)
