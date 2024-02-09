terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

provider "cloudflare" {
  api_token = "joOPhd4H2g3J3CI2_hOJwwgatJWXWixDo2Iww8Au"
}

variable "zone_id" {
  default = "cea4663809ffcd5afa285c3fca3d5eef"
}

variable "account_id" {
  default = "d4a55165ba5ec426089bd63e723ec25e"
}

variable "domain" {
  default = "somesite.xyz"
}

resource "cloudflare_zaraz_config" "examplecom" {
  zone_id = var.zone_id
  config {
    debug_key = "sdsdds"
    triggers = {
     trigger1 = {
        exclude_rules = []
        load_rules = [
          {
            id = "f754b30b-a0c5-4027-b72f-513a7ace637c"
            match = "{{ client.__zarazTrack }}"
            op = "EQUALS",
            value = "CustomEvent"
          },
          {
            action = "scrollDepth"
            id = "iQss"
            settings = {
               positions = "50%" 
            }
          },
          {
            id = "45bcc075-7c5d-4eac-bc95-ec038f949d80"
            action = "clickListener"
            settings = { type = "css", selector = "a", wait_for_tags = 0 }
          }
        ]
        description = "my triger"
        name = "my trigger"
      }
    }
    tools = {
      tool1 = {
        name = "Tool1"
        component = "http"
        enabled = true
        default_fields = {}
        actions = {"1234" = {
          blocking_triggers = []
          firing_triggers = []
          data = {}
          action_type = "pageview"
        }}
        permissions = ["execute_unsafe_scripts"]
        settings = {}
        type = "component"
        worker = {
          escaped_worker_name = "custom-mc-cf-zaraz-latest"
          worker_tag = "0c3d528e0b1a4219be0e14ea6ecdc785"
        }
      }
      tool2 = {
        name = "Tool2"
        component = "html"
        enabled = false
        # other fields
        default_fields = {}
        actions = {}
        permissions = []
        settings = {}
        type = "component"
      }
    }
  }
}

