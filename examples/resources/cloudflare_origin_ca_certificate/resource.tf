# Create a CSR and generate a CA certificate
resource "tls_private_key" "example" {
  algorithm = "RSA"
}

resource "tls_cert_request" "example" {
  key_algorithm   = tls_private_key.example.algorithm
  private_key_pem = tls_private_key.example.private_key_pem

  subject {
    common_name  = ""
    organization = "Terraform Test"
  }
}

resource "cloudflare_origin_ca_certificate" "example" {
  csr                = tls_cert_request.example.cert_request_pem
  hostnames          = ["example.com"]
  request_type       = "origin-rsa"
  requested_validity = 7
}
