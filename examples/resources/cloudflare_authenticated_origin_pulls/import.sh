# global
$ terraform import cloudflare_authenticated_origin_pulls.example <zone_id>

# per zone
$ terraform import cloudflare_authenticated_origin_pulls.example <zone_id>/<certificate_id>

# per hostname
$ terraform import cloudflare_authenticated_origin_pulls.example <zone_id>/<certificate_id>/<hostname>
