
resource "cloudflare_zero_trust_device_managed_networks" "%[1]s" {
  account_id                = "%[2]s"
  name                      = "%[1]s"
  type                      = "tls"
  config = {
  tls_sockaddr = "foobar:1234"
	sha256 = "b5bb9d8014a0f9b1d61e21e796d78dccdf1352f23cd32812f4850b878ae4944c"
}
}
