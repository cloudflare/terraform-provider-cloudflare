resource "cloudflare_device_managed_networks" "managed_networks" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "managed-network-1"
  type       = "tls"
  config {
    tls_sockaddr = "foobar:1234"
    sha256       = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
}
