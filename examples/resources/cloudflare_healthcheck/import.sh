# Use the Zone ID and Healthcheck ID to import.
$ terraform import cloudflare_healthcheck.example <zone_id>/<healthcheck_id>

# 9a7806061c88ada191ed06f989cc3dac - the zone ID
# 699d98642c564d2e855e9661899b7252 - healthcheck ID as returned by [API](https://api.cloudflare.com/#health-checks-list-health-checks)