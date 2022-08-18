terraform {
  required_providers {
    plumber = {
      version = "0.1.1"
      source  = "batch.sh/tf/plumber"
    }
  }
}

provider "plumber" {
  plumber_token      = "batchcorp"
  address            = "localhost:9090"
  connection_timeout = 10
}


resource "plumber_connection" "test_rabbit" {
  name = "test rabbit"
  rabbit {
    address = "amqp://localhost:5672"
  }
}

#resource "plumber_relay" "testrabbitrelay" {
#  connection_id = plumber_connection.test_rabbit.id
#  collection_token = "4b9784ae-f97b-4606-b172-a78444291db3"
#  active = "true"
#
#  # optional, for local testing
#  grpc {
#    address = "localhost:9000"
#    connection_timeout = 5
#    disable_tls = true
#  }
#
#  # Need at least one connection
#  rabbit {
#    exchange_name = "events"
#    queue_name = "testin"
#    queue_declare = true
#    binding_key = "#"
#  }
#}

resource "plumber_tunnel" "dynamictest" {
  # optional, for local testing
  dproxy {
    address = "localhost:9001"
    connection_timeout = 5
    disable_tls = true
  }

  name = "terraform test3"

  active = false

  connection_id = plumber_connection.test_rabbit.id

  batchsh_api_token = "batchsh_319041f4b82fb7c0fe04b2598449a3e07effe66c2af5d54d13d6f6b1d2bb"

  # Dev token: "batchsh_3b17c235a49a871d2c9715c40acdef33c9bfe6e1bc881e61f5659022eac9"

  rabbit {
    exchange_name = "events"
    routing_key = "testing"
  }
}


#
#data "plumber_relay" "test" {
#  filter {
#    name   = "connection_id"
#    values = ["75b40ebe-e757-4376-b100-dd77df1e1d2c"]
#  }
#}
#
#output "testin" {
#  value = data.plumber_relay.test.id
#}

#data "plumber_connection" "testkafkadata" {
#  filter {
#    name   = "name"
#    values = ["test kafka 3"]
#  }
#}
#
#output "testkafka" {
#  value = data.plumber_connection.testkafkadata.id
#}

#
#resource "plumber_relay" "testkafkarelay" {
#  connection_id = plumber_connection.testkafka.id
#  collection_token = "48b30466-e3cb-4a58-9905-45b74284709f"
#
#  # optional, for local testing
#  grpc {
#    address = "localhost:9000"
#    connection_timeout = 5
#    disable_tls = true
#  }
#
#  # Need at least one connection
#  kafka {
#    topics = ["tfrelaytest"]
#    consumer_group_name = "plumber"
#  }
#}

#resource "plumber_connection" "testkafka2" {
#  name = "test kafka 3"
#  kafka {
#    address = ["localhost:9092"]
#    connection_timeout = 8
#  }
#}
#resource "plumber_tunnel" "dynamictest" {
#  # optional, for local testing
#  dproxy {
#    address = "localhost:9001"
#    connection_timeout = 5
#    disable_tls = true
#  }
#
#  name = "terraform test3"
#
#  active = false
#
#  connection_id = plumber_connection.testkafka2.id
#
#  batchsh_api_token = "batchsh_319041f4b82fb7c0fe04b2598449a3e07effe66c2af5d54d13d6f6b1d2bb"
#
#  # Dev token: "batchsh_3b17c235a49a871d2c9715c40acdef33c9bfe6e1bc881e61f5659022eac9"
#
#  kafka {
#    topics = ["tfrelaytest2"]
##    headers = {
##      sample = "sample val"
##      additional = "additional val"
##    }
#  }
#}