resource "cloudflare_ai_gateway_dynamic_routing" "example_ai_gateway_dynamic_routing" {
  account_id = "0d37909e38d3e99c29fa2cd343ac421a"
  gateway_id = "54442216"
  elements = [{
    id = "id"
    outputs = {
      next = {
        element_id = "elementId"
      }
    }
    type = "start"
  }]
  name = "name"
}
