package main

import (
	"testing"
)

func TestSpectrumApplicationTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "remove optional id attribute",
			Config: `resource "cloudflare_spectrum_application" "example" {
  id           = "some-user-provided-id"
  zone_id      = "example.com"
  protocol     = "tcp/443"
  dns {
    type = "CNAME"
    name = "secure.example.com"
  }
  
  origin_direct = ["203.0.113.1"]
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/443"
  dns {
    type = "CNAME"
    name = "secure.example.com"
  }

  origin_direct = ["203.0.113.1"]
}`},
		},
		{
			Name: "preserve existing origin_port attribute when no id present",
			Config: `resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/3306"
  dns {
    type = "ADDRESS"
    name = "db.example.com"
  }
  
  origin_port   = 3306
  origin_direct = ["203.0.113.1"]
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/3306"
  dns {
    type = "ADDRESS"
    name = "db.example.com"
  }

  origin_port   = 3306
  origin_direct = ["203.0.113.1"]
}`},
		},
		{
			Name: "remove id from complex spectrum application",
			Config: `resource "cloudflare_spectrum_application" "example" {
  id            = "user-specified-id"
  zone_id       = "example.com"
  protocol      = "tcp/443"
  dns {
    type = "CNAME"
    name = "app.example.com"
  }
  
  edge_ips {
    type         = "dynamic"
    connectivity = "all"
  }
  
  origin_dns {
    name = "backend.example.com"
  }
  
  tls              = "flexible"
  argo_smart_routing = true
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id       = "example.com"
  protocol      = "tcp/443"
  dns {
    type = "CNAME"
    name = "app.example.com"
  }

  edge_ips {
    type         = "dynamic"
    connectivity = "all"
  }

  origin_dns {
    name = "backend.example.com"
  }

  tls              = "flexible"
  argo_smart_routing = true
}`},
		},
		{
			Name: "no transformation needed when id not present",
			Config: `resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/22"
  dns {
    type = "CNAME"
    name = "ssh.example.com"
  }
  
  origin_direct = ["203.0.113.1"]
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/22"
  dns {
    type = "CNAME"
    name = "ssh.example.com"
  }

  origin_direct = ["203.0.113.1"]
}`},
		},
		{
			Name: "convert origin_port_range to origin_port string format",
			Config: `resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/3306-3310"
  dns {
    type = "ADDRESS"
    name = "db.example.com"
  }
  
  origin_port_range {
    start = 3306
    end   = 3310
  }
  origin_direct = ["203.0.113.1"]
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/3306-3310"
  dns {
    type = "ADDRESS"
    name = "db.example.com"
  }

  origin_direct = ["203.0.113.1"]
  origin_port   = "3306-3310"
}`},
		},
		{
			Name: "remove id and convert origin_port_range simultaneously",
			Config: `resource "cloudflare_spectrum_application" "example" {
  id           = "user-provided-id"
  zone_id      = "example.com"
  protocol     = "tcp/8080"
  dns {
    type = "CNAME"
    name = "app.example.com"
  }
  
  origin_port_range {
    start = 8080
    end   = 8090
  }
  origin_direct = ["203.0.113.1"]
}`,
			Expected: []string{`resource "cloudflare_spectrum_application" "example" {
  zone_id      = "example.com"
  protocol     = "tcp/8080"
  dns {
    type = "CNAME"
    name = "app.example.com"
  }

  origin_direct = ["203.0.113.1"]
  origin_port   = "8080-8090"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}