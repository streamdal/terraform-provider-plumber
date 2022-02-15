terraform {
  required_providers {
    plumber = {
      version = "0.1.0"
      source  = "batchcorp/plumber"
    }
  }
}

provider "plumber" {
  plumber_token      = "batchcorp"
  address            = "localhost:9090"
  connection_timeout = 10
}

# First create the connection
resource "plumber_connection" "infra_kafka" {
  name = "local kafka on my infrastructure"
  kafka {
    address = ["10.1.80.253:9092"]
    connection_timeout = 5
  }
}

# Now create/start the tunnel
# It will then appear under https://console.batch.sh/destinations
resource "plumber_tunnel" "my_replay_tunnel" {
  name = "Tunnel to infra kafka"

  # Obtained from your https://console.batch.sh account
  batchsh_api_token = "batchsh_....."

  connection_id = plumber_connection.infra_kafka.id

  active = true

  # Specify the kafka topic that messages will be written to when replaying to this tunnel
  kafka {
    topics = ["new_orders"]
    headers = {
      replayed = "true"
      additional = "some additional val"
    }
  }
}