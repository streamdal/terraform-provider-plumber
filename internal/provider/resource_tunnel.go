package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/terraform-provider-plumber/plumber"
)

func resourceTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "A tunnel allows Plumber to act as a replay destination within your infrastructure",

		CreateContext: resourceTunnelCreate,
		ReadContext:   resourceTunnelRead,
		UpdateContext: resourceTunnelUpdate,
		DeleteContext: resourceTunnelDelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"connection_id": {
				Description: "Plumber connection ID",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of this tunnel, which will appear in https://console.batch.sh Destinations",
				Type:        schema.TypeString,
				Required:    true,
			},
			"batchsh_api_token": {
				Description: "API Token for https://console.batch.sh (NOT Plumber Token)",
				Type:        schema.TypeString,
				Required:    true,
			},
			"active": {
				Description: "Whether the tunnel is active or not. Setting to true will start the tunnel on Create/Update",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"dproxy": {
				Description: "Optional dProxy configuration",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description: "Address dProxy GRPC Server",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "dproxy.batch.sh:443",
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
				Description: "Replay messages to a Kafka connection",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topics": {
							Description: "Topics to replay",
							Type:        schema.TypeList,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"key": {
							Description: "Key to write messages with",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"headers": {
							Description: "Headers to write messages with",
							Type:        schema.TypeMap,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"rabbit": {
				Description: "Replay messages to a RabbitMQ connection",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exchange_name": {
							Description: "Exchange to write message(s)",
							Type:        schema.TypeString,
							Required:    true,
						},
						"routing_key": {
							Description: "Routing key to write message(s) to",
							Type:        schema.TypeString,
							Required:    true,
						},
						"app_id": {
							Description: "Fills message properties $app_id with this value",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"exchange_type": {
							Description: "The type of exchange we are working with: direct,topic,headers,fanout",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"exchange_declare": {
							Description: "Whether to declare an exchange (if it does not exist)",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"exchange_durable": {
							Description: "Whether to make a declared exchange durable",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"exchange_auto_delete": {
							Description: "Whether to auto-delete the exchange (after writes)",
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

func resourceTunnelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	opts, moreDiags := buildTunnel(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	req := &protos.CreateTunnelRequest{
		Auth: &common.Auth{Token: pc.Token},
		Opts: opts,
	}

	resp, err := pc.Client.CreateTunnel(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.TunnelId)

	return diags
}

func buildTunnel(d *schema.ResourceData) (*opts.TunnelOptions, diag.Diagnostics) {
	var diags diag.Diagnostics
	var opts *opts.TunnelOptions

	switch getBusType(d) {
	case "kafka":
		opts, diags = buildTunnelOptionsKafka(d)
		if diags.HasError() {
			return nil, diags
		}
		return opts, diags
	case "rabbit":
		opts, diags = buildTunnelOptionsRabbit(d)
		if diags.HasError() {
			return nil, diags
		}
		return opts, diags
	default:
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unknown tunnel message bus type. Must be one of: 'kafka', 'rabbit'", // TODO: expand options
		})
	}

	return nil, append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unknown tunnel message bus type. Must be one of: 'kafka'", // TODO: expand options
	})
}

func buildTunnelOptionsRabbit(d *schema.ResourceData) (*opts.TunnelOptions, diag.Diagnostics) {
	var diags diag.Diagnostics

	kafkaData := d.Get("rabbit").([]interface{})
	config := kafkaData[0].(map[string]interface{})

	opts := &opts.TunnelOptions{
		ApiToken:     d.Get("batchsh_api_token").(string),
		ConnectionId: d.Get("connection_id").(string),
		Name:         d.Get("name").(string),
		XActive:      d.Get("active").(bool),
		Rabbit: &opts.TunnelGroupRabbitOptions{
			//Args: &args.RabbitWriteArgs{
			//	Key:     kafkaConfig["key"].(string),
			//	Headers: convertStringMap(kafkaConfig["headers"].(map[string]interface{})),
			//	Topics:  flattenKafkaTopics(kafkaConfig["topics"].([]interface{})),
			//},
			Args: &args.RabbitWriteArgs{
				ExchangeName:         config["exchange_name"].(string),
				RoutingKey:           config["routing_key"].(string),
				AppId:                config["app_id"].(string),
				ExchangeType:         config["exchange_type"].(string),
				ExchangeDeclare:      false,
				ExchangeDurable:      false,
				ExchangeAutoDelete:   false,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
		},
	}

	//dproxyBlock := d.Get("dproxy")
	//if dproxyBlock != nil {
	//	dproxyData, ok := dproxyBlock.([]interface{})
	//	if ok && len(dproxyData) > 0 {
	//		dproxyConfig := dproxyData[0].(map[string]interface{})
	//		opts.XGrpcAddress = dproxyConfig["address"].(string)
	//		opts.XGrpcInsecure = dproxyConfig["disable_tls"].(bool)
	//		opts.XGrpcTimeoutSeconds = uint32(dproxyConfig["connection_timeout"].(int))
	//	}
	//}

	return opts, diags
}

func buildTunnelOptionsKafka(d *schema.ResourceData) (*opts.TunnelOptions, diag.Diagnostics) {
	var diags diag.Diagnostics

	kafkaData := d.Get("kafka").([]interface{})
	kafkaConfig := kafkaData[0].(map[string]interface{})

	opts := &opts.TunnelOptions{
		ApiToken:     d.Get("batchsh_api_token").(string),
		ConnectionId: d.Get("connection_id").(string),
		Name:         d.Get("name").(string),
		XActive:      d.Get("active").(bool),
		Kafka: &opts.TunnelGroupKafkaOptions{
			Args: &args.KafkaWriteArgs{
				Key:     kafkaConfig["key"].(string),
				Headers: convertStringMap(kafkaConfig["headers"].(map[string]interface{})),
				Topics:  flattenKafkaTopics(kafkaConfig["topics"].([]interface{})),
			},
		},
	}

	//dproxyBlock := d.Get("dproxy")
	//if dproxyBlock != nil {
	//	dproxyData, ok := dproxyBlock.([]interface{})
	//	if ok && len(dproxyData) > 0 {
	//		dproxyConfig := dproxyData[0].(map[string]interface{})
	//		opts.XGrpcAddress = dproxyConfig["address"].(string)
	//		opts.XGrpcInsecure = dproxyConfig["disable_tls"].(bool)
	//		opts.XGrpcTimeoutSeconds = uint32(dproxyConfig["connection_timeout"].(int))
	//	}
	//}

	return opts, diags
}

func convertStringMap(input map[string]interface{}) map[string]string {
	result := make(map[string]string)

	for k, v := range input {
		result[k] = fmt.Sprintf("%s", v)
	}

	return result
}

// flattenKafkaTopics interface{} of type []interface{} to []string
func flattenKafkaTopics(input []interface{}) []string {
	var result []string

	for _, elem := range input {
		result = append(result, elem.(string))
	}

	return result
}

func flattenKafkaWriteArgs(args *args.KafkaWriteArgs) []map[string]interface{} {
	return []map[string]interface{}{
		{
			"key":     args.Key,
			"headers": args.Headers,
			"topics":  args.Topics,
		},
	}
}

func resourceTunnelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	resp, err := pc.Client.GetTunnel(ctx, &protos.GetTunnelRequest{
		Auth:     &common.Auth{Token: pc.Token},
		TunnelId: d.Id(),
	})
	if err != nil {
		if strings.Contains(err.Error(), "tunnel not found") {
			d.SetId("")
			return diag.FromErr(err)
		}
		return diag.FromErr(err)
	}

	opts := resp.GetOpts()

	d.SetId(opts.XTunnelId)
	d.Set("connection_id", opts.ConnectionId)

	if kafka := opts.GetKafka(); kafka != nil {
		d.Set("kafka", flattenKafkaWriteArgs(kafka.Args))
	}
	// TODO: expand for other message buses

	return diags
}

func resourceTunnelUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	opts, moreDiags := buildTunnel(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	// Required to be set otherwise it will end up blank in Plumber
	opts.XTunnelId = d.Id()

	pc := m.(*plumber.Plumber)

	_, err := pc.Client.UpdateTunnel(ctx, &protos.UpdateTunnelRequest{
		Auth:     &common.Auth{Token: pc.Token},
		TunnelId: d.Id(),
		Opts:     opts,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceTunnelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)
	_, err := pc.Client.DeleteTunnel(ctx, &protos.DeleteTunnelRequest{
		Auth:     &common.Auth{Token: pc.Token},
		TunnelId: d.Id(),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
