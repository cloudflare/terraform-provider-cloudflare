
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "asn"

    item =[ {
      value =[ {
        asn = 345
      }]
      comment = "ASN test"
    },
    {
    value =[ {
        asn = 567
      }]
      comment = "ASN test two"
    }]

  }