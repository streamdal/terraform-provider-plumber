# Testing locally

1. `make`
2. Ensure the  block has the `address` key set and the provider block has the correct provider config and version:
   ```hcl
   
   terraform {
     required_providers {
       plumber = {
         version = "0.1.0"
         source  = "batch.sh/tf/plumber"
       }
     }
   }

   provider "plumber" {
     plumber_token = "batchcorp"
     address = "localhost:9090"
     connection_timeout = 10
   }
   ```
3. ```tf init && tf apply```

## Testing tunnels locally

Specify the `dproxy` block to override defaults to point to local dProxy instance:

```hcl
resource "plumber_tunnel" "kafka_tunnel" {
  grpc dproxy {
    address            = "localhost:9001"
    connection_timeout = 5
    disable_tls        = true
  }
  
  #...rest of config

}
```

## Testing relays locally

Specify the `grpc` block to override defaults to local for grpc-collector:

```hcl
resource "plumber_relay" "kafkarelay" {
  grpc {
    address            = "localhost:9000"
    connection_timeout = 5
    disable_tls        = true
  }
  
  #...rest of config

}
```