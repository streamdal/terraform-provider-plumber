package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/terraform-provider-plumber/plumber"
)

func resourceRelay() *schema.Resource {
	return &schema.Resource{
		Description: "Relay messages to your Batch.sh collections",

		CreateContext: resourceRelayCreate,
		ReadContext:   resourceRelayRead,
		UpdateContext: resourceRelayUpdate,
		DeleteContext: resourceRelayDelete,

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Relay ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"collection_token": {
				Description: "Collection Token",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connection_id": {
				Description: "Plumber connection ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"batch_size": {
				Description: "How many messages to send in a single batch",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1000,
			},
			"batch_max_retry": {
				Description: "How many times plumber will try re-sending a batch",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
			},
			"num_workers": {
				Description: "How many workers to launch per relay",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
			},
			"stats_enable": {
				Description: "Display periodic relay stats",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"stats_report_interval_sec": {
				Description: "How often to print stats",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     10,
			},
			"active": {
				Description: "Whether is active or not. Setting to true will start the relay on Create or Update",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"grpc": {
				Description: "Optional GRPC-collector configuration",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description: "Address GRPC Collector Server",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "grpc-collector.batch.sh:9000",
						},
						"connection_timeout": {
							Description: "Connection timeout (in seconds)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
						"disable_tls": {
							Description: "Disable TLS connection",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
			"kafka": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topics": {
							Description: "Topics to read from",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"read_offset": {
							Description: "Specify what offset the consumer should read from (Requires use_consumer_group to be true)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
						},
						"use_consumer_group": {
							Description: "Use a consumer group to read from Kafka",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"consumer_group_name": {
							Description: "Specify a specific group-id to use when reading from kafka",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"max_wait_seconds": {
							Description: "How long to wait for new data when reading batches of messages",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
						},
						"min_bytes": {
							Description: "Minimum number of bytes to fetch in a single kafka request (throughput optimization)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
						},
						"max_bytes": {
							Description: "Maximum number of bytes to fetch in a single kafka request (throughput optimization)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
						},
						"commit_interval_seconds": {
							Description: "How often, in seconds, to commit offsets to broker (0 = synchronous)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
						"rebalance_timeout_seconds": {
							Description: "How long a coordinator will wait for member joins as part of a rebalance",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
						},
						"queue_capacity": {
							Description: "Internal library queue capacity (throughput optimization)",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     1,
						},
					},
				},
			},
			"rabbit": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exchange_name": {
							Description: "Name of the exchange",
							Type:        schema.TypeString,
							Required:    true,
						},
						"queue_name": {
							Description: "Name of the queue where messages will be routed to",
							Type:        schema.TypeString,
							Required:    true,
						},
						"binding_key": {
							Description: "Binding key for topic based exchanges",
							Type:        schema.TypeString,
							Required:    true,
						},
						"consumer_tag": {
							Description: "How to identify the consumer to RabbitMQ",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"queue_exclusive": {
							Description: "Whether plumber should be the only one using the queue",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"queue_declare": {
							Description: "Whether to create/declare the queue (if it does not exist)",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"queue_durable": {
							Description: "Whether the queue should survive after disconnect",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"queue_delete": {
							Description: "Whether to auto-delete the queue after plumber has disconnected",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"auto_ack": {
							Description: "Automatically acknowledge receipt of read/received messages",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					},
				},
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceRelayCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	opts, moreDiags := buildRelayOptions(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	req := &protos.CreateRelayRequest{
		Auth: &common.Auth{Token: pc.Token},
		Opts: opts,
	}

	resp, err := pc.Client.CreateRelay(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.RelayId)

	return diags
}

func buildRelayOptions(d *schema.ResourceData) (*opts.RelayOptions, diag.Diagnostics) {
	var diags diag.Diagnostics

	cfg := &opts.RelayOptions{
		CollectionToken:        d.Get("collection_token").(string),
		BatchSize:              int32(d.Get("batch_size").(int)),
		BatchMaxRetry:          int32(d.Get("batch_max_retry").(int)),
		ConnectionId:           d.Get("connection_id").(string),
		NumWorkers:             int32(d.Get("num_workers").(int)),
		StatsEnable:            false,
		StatsReportIntervalSec: 0,
		XActive:                d.Get("active").(bool),
	}

	switch getBusType(d) {
	case "kafka":
		kafkaData := d.Get("kafka").([]interface{})
		kafkaConfig := kafkaData[0].(map[string]interface{})

		cfg.Kafka = &opts.RelayGroupKafkaOptions{
			Args: &args.KafkaRelayArgs{
				Topics:                  flattenKafkaTopics(kafkaConfig["topics"].([]interface{})),
				ReadOffset:              int64(kafkaConfig["read_offset"].(int)),
				UseConsumerGroup:        kafkaConfig["use_consumer_group"].(bool),
				ConsumerGroupName:       kafkaConfig["consumer_group_name"].(string),
				MaxWaitSeconds:          int32(kafkaConfig["max_wait_seconds"].(int)),
				MinBytes:                int32(kafkaConfig["min_bytes"].(int)),
				MaxBytes:                int32(kafkaConfig["max_bytes"].(int)),
				CommitIntervalSeconds:   int32(kafkaConfig["commit_interval_seconds"].(int)),
				RebalanceTimeoutSeconds: int32(kafkaConfig["rebalance_timeout_seconds"].(int)),
				QueueCapacity:           int32(kafkaConfig["queue_capacity"].(int)),
			},
		}
	case "rabbit":
		rabbitData := d.Get("rabbit").([]interface{})
		config := rabbitData[0].(map[string]interface{})

		cfg.Rabbit = &opts.RelayGroupRabbitOptions{
			Args: &args.RabbitReadArgs{
				ExchangeName:   config["exchange_name"].(string),
				QueueName:      config["queue_name"].(string),
				BindingKey:     config["binding_key"].(string),
				QueueExclusive: config["queue_exclusive"].(bool),
				QueueDeclare:   config["queue_declare"].(bool),
				QueueDurable:   config["queue_durable"].(bool),
				AutoAck:        config["auto_ack"].(bool),
				ConsumerTag:    config["consumer_tag"].(string),
				QueueDelete:    config["queue_delete"].(bool),
				QueueArg:       convertStringMap(config["queue_arg"].(map[string]interface{})),
			},
		}
	default:
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unknown connection type. Must be one of: 'kafka', 'rabbit", // TODO: expand options
		})
	}

	grpcBlock := d.Get("grpc")
	if grpcBlock != nil {
		grpcData, ok := grpcBlock.([]interface{})
		if ok && len(grpcData) > 0 {
			grpcConfig := grpcData[0].(map[string]interface{})
			cfg.XBatchshGrpcAddress = grpcConfig["address"].(string)
			cfg.XBatchshGrpcDisableTls = grpcConfig["disable_tls"].(bool)
			cfg.XBatchshGrpcTimeoutSeconds = int32(grpcConfig["connection_timeout"].(int))
		}
	}

	return cfg, diags
}

func resourceRelayRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	resp, err := pc.Client.GetRelay(ctx, &protos.GetRelayRequest{
		Auth:    &common.Auth{Token: pc.Token},
		RelayId: d.Id(),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("active", resp.Opts.XActive)
	d.Set("collection_token", resp.Opts.CollectionToken)
	d.Set("batch_size", resp.Opts.BatchSize)
	d.Set("batch_max_retry", resp.Opts.BatchMaxRetry)
	d.Set("connection_id", resp.Opts.ConnectionId)
	d.Set("num_workers", resp.Opts.NumWorkers)
	d.Set("stats_enable", resp.Opts.StatsEnable)
	d.Set("stats_report_interval_sec", resp.Opts.StatsReportIntervalSec)
	return diags

}

func resourceRelayUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	opts, moreDiags := buildRelayOptions(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	pc := m.(*plumber.Plumber)

	// Required to be set otherwise it will end up blank in plumber
	opts.XRelayId = d.Id()

	_, err := pc.Client.UpdateRelay(ctx, &protos.UpdateRelayRequest{
		Auth:    &common.Auth{Token: pc.Token},
		RelayId: d.Id(),
		Opts:    opts,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceRelayDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)
	_, err := pc.Client.DeleteRelay(ctx, &protos.DeleteRelayRequest{
		Auth:    &common.Auth{Token: pc.Token},
		RelayId: d.Id(),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
