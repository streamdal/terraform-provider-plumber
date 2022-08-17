package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/terraform-provider-plumber/plumber"
)

func resourceConnection() *schema.Resource {
	return &schema.Resource{
		Description: "Message Bus Connections",

		CreateContext: resourceConnectionCreate,
		ReadContext:   resourceConnectionRead,
		UpdateContext: resourceConnectionUpdate,
		DeleteContext: resourceConnectionDelete,

		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Connection ID",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"notes": {
				Description: "Customer Notes",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"kafka": {
				Description: "Arguments for a Kafka connection",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description: "Address of Kafka brokers",
							Type:        schema.TypeList,
							Required:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"connection_timeout": {
							Description: "Connection timeout (in seconds)",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"use_tls": {
							Description: "Force TLS connection",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tls_skip_verify": {
							Description: "Skip TLS certificate verification",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"sasl_type": {
							Description:  "SASL Authentication Type",
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "NONE",
							ValidateFunc: validation.StringInSlice([]string{"", "none", "scram", "plain"}, true),
						},
						"sasl_username": {
							Description: "SASL Username",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"sasl_password": {
							Description: "SASL Password",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"rabbit": {
				Description: "Arguments for a RabbitMQ connection",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Description: "Full DSN of RabbitMQ cluister (ex: amqp://bus.domain.com:5672)",
							Type:        schema.TypeString,
							Required:    true,
						},
						"use_tls": {
							Description: "Force TLS connection (regardless of DSN)",
							Type:        schema.TypeBool,
							Optional:    true,
						},
						"tls_skip_verify": {
							Description: "Skip TLS certificate verification",
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

func resourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	resp, err := pc.Client.GetConnection(ctx, &protos.GetConnectionRequest{
		Auth:         &common.Auth{Token: pc.Token},
		ConnectionId: d.Id(),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	opts := resp.GetOptions()

	d.SetId(opts.GetXId())
	d.Set("name", opts.GetName())
	d.Set("notes", opts.GetNotes())

	if kafka := opts.GetKafka(); kafka != nil {
		d.Set("kafka", flattenKafkaConnection(kafka))
	}

	if rabbit := opts.GetRabbit(); rabbit != nil {
		d.Set("rabbit", flattenRabbitConnection(rabbit))
	}
	// TODO: expand for other message buses

	return diags
}

func flattenKafkaConnection(k *args.KafkaConn) []map[string]interface{} {
	return []map[string]interface{}{{
		"address":            k.Address,
		"connection_timeout": int(k.TimeoutSeconds),
		"use_tls":            k.UseTls,
		"tls_skip_verify":    k.TlsSkipVerify,
		"sasl_type":          k.SaslType.String(),
		"sasl_username":      k.SaslUsername,
		"sasl_password":      k.SaslPassword,
	}}
}

func flattenRabbitConnection(k *args.RabbitConn) []map[string]interface{} {
	return []map[string]interface{}{{
		"address":         k.Address,
		"use_tls":         k.UseTls,
		"tls_skip_verify": k.TlsSkipVerify,
	}}
}

func resourceConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)

	conn, moreDiags := buildConnection(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	resp, err := pc.Client.CreateConnection(ctx, &protos.CreateConnectionRequest{
		Auth:    &common.Auth{Token: pc.Token},
		Options: conn,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ConnectionId)

	return diags
}

func resourceConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var conn *opts.ConnectionOptions

	conn, moreDiags := buildConnection(d)
	if moreDiags.HasError() {
		return append(diags, moreDiags...)
	}

	pc := m.(*plumber.Plumber)

	// Required to be set otherwise it will end up blank in Plumber
	conn.XId = d.Id()

	_, err := pc.Client.UpdateConnection(ctx, &protos.UpdateConnectionRequest{
		Auth:         &common.Auth{Token: pc.Token},
		ConnectionId: d.Id(),
		Options:      conn,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	pc := m.(*plumber.Plumber)
	_, err := pc.Client.DeleteConnection(ctx, &protos.DeleteConnectionRequest{
		Auth:         &common.Auth{Token: pc.Token},
		ConnectionId: d.Id(),
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func buildConnection(d *schema.ResourceData) (*opts.ConnectionOptions, diag.Diagnostics) {
	var diags diag.Diagnostics
	var conn *opts.ConnectionOptions

	// TODO: expand options
	switch getBusType(d) {
	case "kafka":
		conn, diags = buildConnectionKafka(d)
		if diags.HasError() {
			return nil, diags
		}
		return conn, diags
	case "rabbit":
		conn, diags = buildConnectionRabbit(d)
		if diags.HasError() {
			return nil, diags
		}
		return conn, diags
	default:
		return nil, append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unknown connection type. Must be one of: 'kafka', 'rabbit", // TODO: expand options
		})
	}
}

// buildConnectionKafka builds a plumber-schemas ConnectionOptions from a tf resource
func buildConnectionKafka(d *schema.ResourceData) (*opts.ConnectionOptions, diag.Diagnostics) {
	var diags diag.Diagnostics

	kafkaData := d.Get("kafka").([]interface{})
	data := kafkaData[0].(map[string]interface{})

	return &opts.ConnectionOptions{
		Name:  d.Get("name").(string),
		Notes: d.Get("notes").(string),
		Conn: &opts.ConnectionOptions_Kafka{
			Kafka: &args.KafkaConn{
				Address:        interfaceToStrings(data["address"]),
				TimeoutSeconds: int32(data["connection_timeout"].(int)),
				UseTls:         data["use_tls"].(bool),
				TlsSkipVerify:  data["tls_skip_verify"].(bool),
				SaslType:       buildConnectionKafkaSASLType(data["sasl_type"].(string)),
				SaslUsername:   data["sasl_username"].(string),
				SaslPassword:   data["sasl_password"].(string),
			},
		},
	}, diags
}

// buildConnectionRabbit builds a plumber-schemas ConnectionOptions from a tf resource
func buildConnectionRabbit(d *schema.ResourceData) (*opts.ConnectionOptions, diag.Diagnostics) {
	var diags diag.Diagnostics

	kafkaData := d.Get("rabbit").([]interface{})
	data := kafkaData[0].(map[string]interface{})

	return &opts.ConnectionOptions{
		Name:  d.Get("name").(string),
		Notes: d.Get("notes").(string),
		Conn: &opts.ConnectionOptions_Rabbit{
			Rabbit: &args.RabbitConn{
				Address:       data["address"].(string),
				UseTls:        data["use_tls"].(bool),
				TlsSkipVerify: data["tls_skip_verify"].(bool),
			},
		},
	}, diags
}

// buildConnectionKafkaSASLType converts a string version of SASLType, aka "plain", "scram" to the protobuf
// version for the gRPC call
func buildConnectionKafkaSASLType(saslType string) args.SASLType {
	switch strings.ToLower(saslType) {
	case "plain":
		return args.SASLType_PLAIN
	case "scram":
		return args.SASLType_SCRAM
	default:
		return args.SASLType_NONE
	}
}

// interfaceToStrings converts an interface{} value to []string
// This is needed when a nested resource, say "kafka" has a value that is a schema.TypeList of schema.TypeString
func interfaceToStrings(value interface{}) []string {
	brokers := make([]string, 0)
	for _, v := range value.([]interface{}) {
		brokers = append(brokers, v.(string))
	}

	return brokers
}
