
	resource "cloudflare_logpush_job" "%[1]s" {
		enabled          = true
		account_id       = "%[3]s"
		name             = "%[1]s"
		logpull_options  = "fields=Event,EventTimestampMs,Outcome,Exceptions,Logs,ScriptName"
		destination_conf = "r2://%[1]s/date={DATE}?account-id=%[3]s&access-key-id=%[6]s&secret-access-key=%[7]s"
		dataset          = "workers_trace_events"
	}

resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[3]s"
  name = "%[1]s"
  content = "%[2]s"
  module = true
  compatibility_date = "%[4]s"
  compatibility_flags = ["%[5]s"]
  logpush = true
	placement =[ {
		mode = "smart"
	}]

	d1_database_binding =[ {
		name = "MY_DATABASE"
		database_id = "%[8]s"
	}]

	depends_on = [cloudflare_logpush_job.%[1]s]
}