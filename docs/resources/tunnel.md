---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "plumber_tunnel Resource - terraform-provider-plumber"
subcategory: ""
description: |-
  Relay messages to your Batch.sh collections
---

# Resource: plumber_tunnel

A tunnel replay destination allows plumber to act as a relay for messages inside your infrastructure, avoiding the need to
punch holes in firewalls in order to replay messages to your message bus from Batch.sh.

This resource allows you to create/update/delete tunnel replay destination configurations.

## Example Usage

---

```hcl
resource "plumber_connection" "infra_kafka" {
  name = "local kafka on my infrastructure"
  kafka {
    address = ["10.1.80.253:9092"]
    connection_timeout = 5
  }
}

resource "plumber_tunnel" "my_replay_tunnel" {
  name = "Tunnel to infra kafka"
  
  # Obtained from your https://console.batch.sh account
  batchsh_api_token = "batchsh_....."

  connection_id = plumber_connection.infra_kafka.id

  kafka {
    topics = ["new_orders"]
    headers = {
      replayed = "true"
      additional = "some additional val"
    }
  }
}
```

## Argument Reference

---

- `name` - Name of the tunnel.
- `active` - Whether the tunnel is active.
- `connection_id` - (String) Plumber connection ID
- `batchsh_api_token` - (String) API Token for https://console.batch.sh - (NOT Plumber Token)

At least one message bus relay configuration must be specified:

- `kafka` - (Block List, Max: 1) Replay messages to a Kafka connection - see [below for nested schema](#nestedblock--kafka)

<a id="nestedblock--kafka"></a>
### Nested Schema for `kafka`

Optional:

- `headers` - (Map of String) Headers to write messages with
- `key` - (String) Key to write messages with
- `topics` - (List of String) Topics to replay

## Attributes Reference

---

In addition to all arguments above, the following attributes are returned:

- `id` - (String) The ID of the configured tunnel destination


## Import

---

You can import an existing tunnel configuration from plumber by specifying its ID:

```bash
$ terraform import plumber_tunnel.my_kafka_dest d743952d-6dc4-472b-849f-44015c8af3fb
```