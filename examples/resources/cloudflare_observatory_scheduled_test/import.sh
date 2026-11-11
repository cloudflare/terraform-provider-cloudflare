# Import using zone_id and url (region defaults to us-central1)
$ terraform import cloudflare_observatory_scheduled_test.example '<zone_id>/<url>'

# Import with an explicit region (required when the schedule is not in us-central1)
$ terraform import cloudflare_observatory_scheduled_test.example '<zone_id>/<url>/<region>'
