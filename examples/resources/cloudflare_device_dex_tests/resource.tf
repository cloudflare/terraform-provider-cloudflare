resource "cloudflare_device_dex_tests" "dex_tests" {
    test_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
    name = "GET dashboard"
    description = "Send a HTTP GET request to the 'home' endpoint of the dash every half hour."
    interval = "0h30m0s"
    enabled = true
    data {
        host = "https://dash.cloudflare.com/home"
        kind = "http"
        method = "GET"
    }
}