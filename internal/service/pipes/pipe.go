package pipes

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pipes"
	"github.com/aws/aws-sdk-go-v2/service/pipes/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_pipes_pipe")
func ResourcePipe() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourcePipeCreate,
		ReadWithoutTimeout:   resourcePipeRead,
		UpdateWithoutTimeout: resourcePipeUpdate,
		DeleteWithoutTimeout: resourcePipeDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: verify.SetTagsDiff,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
			"desired_state": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          string(types.RequestedPipeStateRunning),
				ValidateDiagFunc: enum.Validate[types.RequestedPipeState](),
			},
			"enrichment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 1600),
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name_prefix"},
				ValidateFunc:  validation.StringLenBetween(1, 64),
			},
			"name_prefix": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"name"},
				ValidateFunc:  validation.StringLenBetween(1, 64-resource.UniqueIDSuffixLength),
			},
			"role_arn": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidARN,
			},
			"source": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 1600),
			},
			"source_parameters": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_criteria": {
							Type:             schema.TypeList,
							Optional:         true,
							MaxItems:         1,
							DiffSuppressFunc: suppressEmptyConfigurationBlock("source_parameters.0.filter_criteria"),
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 5,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pattern": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringLenBetween(1, 4096),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"tags":     tftags.TagsSchema(),
			"tags_all": tftags.TagsSchemaComputed(),
			"target": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 1600),
			},
		},
	}
}

const (
	ResNamePipe = "Pipe"
)

func resourcePipeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PipesClient()

	name := create.Name(d.Get("name").(string), d.Get("name_prefix").(string))

	input := &pipes.CreatePipeInput{
		DesiredState: types.RequestedPipeState(d.Get("desired_state").(string)),
		Name:         aws.String(name),
		RoleArn:      aws.String(d.Get("role_arn").(string)),
		Source:       aws.String(d.Get("source").(string)),
		Target:       aws.String(d.Get("target").(string)),
	}

	if v, ok := d.Get("description").(string); ok {
		input.Description = aws.String(v)
	}

	if v, ok := d.Get("enrichment").(string); ok && v != "" {
		input.Enrichment = aws.String(v)
	}

	if v, ok := d.Get("source_parameters").([]interface{}); ok && len(v) > 0 && v[0] != nil {
		input.SourceParameters = expandPipeSourceParameters(v[0].(map[string]interface{}))
	}

	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	tags := defaultTagsConfig.MergeTags(tftags.New(ctx, d.Get("tags").(map[string]interface{})))

	if len(tags) > 0 {
		input.Tags = Tags(tags.IgnoreAWS())
	}

	output, err := conn.CreatePipe(ctx, input)
	if err != nil {
		return create.DiagError(names.Pipes, create.ErrActionCreating, ResNamePipe, name, err)
	}

	if output == nil || output.Arn == nil {
		return create.DiagError(names.Pipes, create.ErrActionCreating, ResNamePipe, name, errors.New("empty output"))
	}

	d.SetId(aws.ToString(output.Name))

	if _, err := waitPipeCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return create.DiagError(names.Pipes, create.ErrActionWaitingForCreation, ResNamePipe, d.Id(), err)
	}

	return resourcePipeRead(ctx, d, meta)
}

func resourcePipeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PipesClient()

	output, err := FindPipeByName(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EventBridge Pipes Pipe (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return create.DiagError(names.Pipes, create.ErrActionReading, ResNamePipe, d.Id(), err)
	}

	d.Set("arn", output.Arn)
	d.Set("description", output.Description)
	d.Set("desired_state", output.DesiredState)
	d.Set("enrichment", output.Enrichment)
	d.Set("name", output.Name)
	d.Set("name_prefix", create.NamePrefixFromName(aws.ToString(output.Name)))

	if v := output.SourceParameters; v != nil {
		if err := d.Set("source_parameters", []interface{}{flattenPipeSourceParameters(v)}); err != nil {
			return create.DiagError(names.Pipes, create.ErrActionSetting, ResNamePipe, d.Id(), err)
		}
	}

	d.Set("role_arn", output.RoleArn)
	d.Set("source", output.Source)
	d.Set("target", output.Target)

	tags, err := ListTags(ctx, conn, aws.ToString(output.Arn))
	if err != nil {
		return create.DiagError(names.Pipes, create.ErrActionReading, ResNamePipe, d.Id(), err)
	}

	defaultTagsConfig := meta.(*conns.AWSClient).DefaultTagsConfig
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
	tags = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig)

	if err := d.Set("tags", tags.RemoveDefaultConfig(defaultTagsConfig).Map()); err != nil {
		return create.DiagError(names.Pipes, create.ErrActionSetting, ResNamePipe, d.Id(), err)
	}

	if err := d.Set("tags_all", tags.Map()); err != nil {
		return create.DiagError(names.Pipes, create.ErrActionSetting, ResNamePipe, d.Id(), err)
	}

	return nil
}

func resourcePipeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PipesClient()

	if d.HasChangesExcept("tags", "tags_all") {
		input := &pipes.UpdatePipeInput{
			Description:  aws.String(d.Get("description").(string)),
			DesiredState: types.RequestedPipeState(d.Get("desired_state").(string)),
			Name:         aws.String(d.Id()),
			RoleArn:      aws.String(d.Get("role_arn").(string)),
			Target:       aws.String(d.Get("target").(string)),

			// Omitting the SourceParameters entirely is interpreted as "no change".
			SourceParameters: &types.UpdatePipeSourceParameters{},
		}

		if d.HasChange("enrichment") {
			// Reset state in case it's a deletion.
			input.Enrichment = aws.String("")
		}

		if v, ok := d.Get("enrichment").(string); ok && v != "" {
			input.Enrichment = aws.String(v)
		}

		if d.HasChange("source_parameters.0.filter_criteria") {
			// To unset a parameter, it must be set to an empty object. Nulling a
			// parameter will be interpreted as "no change".
			input.SourceParameters.FilterCriteria = &types.FilterCriteria{}
		}

		if v, ok := d.Get("source_parameters.0.filter_criteria").([]interface{}); ok && len(v) > 0 && v[0] != nil {
			input.SourceParameters.FilterCriteria = expandFilterCriteria(v[0].(map[string]interface{}))
		}

		log.Printf("[DEBUG] Updating EventBridge Pipes Pipe (%s): %#v", d.Id(), input)

		output, err := conn.UpdatePipe(ctx, input)

		if err != nil {
			return create.DiagError(names.Pipes, create.ErrActionUpdating, ResNamePipe, d.Id(), err)
		}

		if _, err := waitPipeUpdated(ctx, conn, aws.ToString(output.Name), d.Timeout(schema.TimeoutUpdate)); err != nil {
			return create.DiagError(names.Pipes, create.ErrActionWaitingForUpdate, ResNamePipe, d.Id(), err)
		}
	}

	if d.HasChange("tags_all") {
		o, n := d.GetChange("tags_all")
		if err := UpdateTags(ctx, conn, d.Get("arn").(string), o, n); err != nil {
			return diag.Errorf("error updating EventBridge Pipes Pipe (%s) tags: %s", d.Id(), err)
		}
	}

	return resourcePipeRead(ctx, d, meta)
}

func resourcePipeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).PipesClient()

	log.Printf("[INFO] Deleting EventBridge Pipes Pipe %s", d.Id())

	_, err := conn.DeletePipe(ctx, &pipes.DeletePipeInput{
		Name: aws.String(d.Id()),
	})

	if err != nil {
		if errs.IsA[*types.NotFoundException](err) {
			return nil
		}

		return create.DiagError(names.Pipes, create.ErrActionDeleting, ResNamePipe, d.Id(), err)
	}

	if _, err := waitPipeDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return create.DiagError(names.Pipes, create.ErrActionWaitingForDeletion, ResNamePipe, d.Id(), err)
	}

	return nil
}

func suppressEmptyConfigurationBlock(key string) schema.SchemaDiffSuppressFunc {
	return func(k, o, n string, d *schema.ResourceData) bool {
		if k != key+".#" {
			return false
		}

		if o == "0" && n == "1" {
			v := d.Get(key).([]interface{})
			return len(v) == 0 || v[0] == nil || len(v[0].(map[string]interface{})) == 0
		}

		return false
	}
}
