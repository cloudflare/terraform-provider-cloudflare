# import hostname associations for the active Cloudflare Managed CA 
$ terraform import cloudflare_authenticated_origin_pulls.example <zone_id>

# import hostname associations for the specified mTLS certificate
$ terraform import cloudflare_authenticated_origin_pulls.example <zone_id>/<mtls_certificate_id>