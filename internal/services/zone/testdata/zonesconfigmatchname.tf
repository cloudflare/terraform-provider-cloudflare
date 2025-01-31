
data "cloudflare_zones" "%[2]s" {
  filter {
    name   = "baa-com.cfapi.net"
    // This is an ordering fix to ensure that the test suite doesn't assert
    // state before all the resources are available.
    paused = "${cloudflare_zone.foo_net.paused}"
  }
}

%[1]s
