# Import an account-scoped job.
$ terraform import cloudflare_logpush_job.example account/<accountID>/<jobID>

# Import a zone-scoped job.
$ terraform import cloudflare_logpush_job.example zone/<zoneID>/<jobID>
