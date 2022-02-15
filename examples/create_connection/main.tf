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


resource "plumber_connection" "test_kafka" {
  name = "test kafka"
  kafka {
    address = ["localhost:9092"]
    connection_timeout = 5
  }
}