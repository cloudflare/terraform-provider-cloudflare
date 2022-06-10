# Import an account-scoped job.
$ terraform import cloudflare_logpush_job.example account/<account_id>/<job_id>

# Import a zone-scoped job.
$ terraform import cloudflare_logpush_job.example zone/<zone_id>/<job_id>
