// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_vpcs")
func DataSourceVPCs() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceVPCsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			names.AttrFilter: customFiltersSchema(),
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			names.AttrTags: tftags.TagsSchemaComputed(),
		},
	}
}

func dataSourceVPCsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DescribeVpcsInput{}

	if tags, tagsOk := d.GetOk(names.AttrTags); tagsOk {
		input.Filters = append(input.Filters, newTagFilterList(
			Tags(tftags.New(ctx, tags.(map[string]interface{}))),
		)...)
	}

	if filters, filtersOk := d.GetOk(names.AttrFilter); filtersOk {
		input.Filters = append(input.Filters,
			newCustomFilterList(filters.(*schema.Set))...)
	}

	if len(input.Filters) == 0 {
		input.Filters = nil
	}

	output, err := FindVPCs(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 VPCs: %s", err)
	}

	var vpcIDs []string

	for _, v := range output {
		vpcIDs = append(vpcIDs, aws.StringValue(v.VpcId))
	}

	d.SetId(meta.(*conns.AWSClient).Region)
	d.Set("ids", vpcIDs)

	return diags
}
