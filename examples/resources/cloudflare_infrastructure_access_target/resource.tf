resource "cloudflare_infrastructure_access_target" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
	hostname   = "example-target"
	ip = {
		ipv4 = {
            ip_addr = "210.26.29.230"
            virtual_network_id = "238dccd1-149b-463d-8228-560ab83a54fd"
        }
        ipv6 = {
            ip_addr = "24c0:64e8:f0b4:8dbf:7104:72b0:ef8f:f5e0"
            virtual_network_id = "238dccd1-149b-463d-8228-560ab83a54fd"
        }
	}
}

resource "cloudflare_infrastructure_access_target" "ipv4_only_example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
	hostname   = "example-ipv4-only"
	ip = {
		ipv4 = {
            ip_addr = "210.26.29.230"
            virtual_network_id = "238dccd1-149b-463d-8228-560ab83a54fd"
        }
	}
}
