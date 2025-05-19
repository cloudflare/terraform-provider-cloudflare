
# resource "cloudflare_workers_script" "%[1]s" {
# 	account_id  = "%[2]s"
# 	script_name = "%[1]s"
# 	content = "addEventListener('fetch', event => {event.respondWith(new Response('test'))});"
# }

resource "cloudflare_workers_cron_trigger" "%[1]s" {
	account_id  = "%[2]s"
	script_name = "mute-truth-fdb1" # cloudflare_workers_script.%[1]s.name
	schedules   = [
  	{
   	  cron = "*/5 * * * *"     # every 5 minutes
  	},
  	{
  		cron = "10 7 * * mon-fri" # 7:10am every weekday
  	}
	]
}
