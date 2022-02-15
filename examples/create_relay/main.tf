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

# First create a connection
resource "plumber_connection" "test_kafka" {
  name = "test kafka"
  active = true
  kafka {
    address = ["localhost:9092"]
    connection_timeout = 5
  }
}

# Now create a relay for the above connection
resource "plumber_relay" "my_kafka_relay" {
  connection_id = plumber_connection.test_kafka.id
  collection_token = "48b30466-e3cb-4a58-9905-45b74284709f"

  # Relay details for your kafka connection
  kafka {
    topics = ["new_orders"]
    consumer_group_name = "plumber"
  }
}