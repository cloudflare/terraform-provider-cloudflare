resource "cloudflare_device_managed_networks" "%s" {
  account_id = "%s"
  name       = "tf-test-managed-network-%s"
  type       = "tls"

  config {
    tls_sockaddr = "example.com:443"
    sha256       = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
  }
}
