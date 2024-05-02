// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindKeyByID(ctx context.Context, conn *kms.KMS, id string) (*kms.KeyMetadata, error) {
	input := &kms.DescribeKeyInput{
		KeyId: aws.String(id),
	}

	output, err := conn.DescribeKeyWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, kms.ErrCodeNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.KeyMetadata == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	keyMetadata := output.KeyMetadata

	// Once the CMK is in the pending (replica) deletion state Terraform considers it logically deleted.
	if state := aws.StringValue(keyMetadata.KeyState); state == kms.KeyStatePendingDeletion || state == kms.KeyStatePendingReplicaDeletion {
		return nil, &retry.NotFoundError{
			Message:     state,
			LastRequest: input,
		}
	}

	return keyMetadata, nil
}

func findDefaultKey(ctx context.Context, client *conns.AWSClient, service, region string) (string, error) {
	conn := client.KMSConnForRegion(ctx, region)

	k, err := FindKeyByID(ctx, conn, fmt.Sprintf("alias/aws/%s", service)) //default key
	if err != nil {
		return "", fmt.Errorf("finding default key: %s", err)
	}

	return aws.StringValue(k.Arn), nil
}

func FindKeyPolicyByKeyIDAndPolicyName(ctx context.Context, conn *kms.KMS, keyID, policyName string) (*string, error) {
	input := &kms.GetKeyPolicyInput{
		KeyId:      aws.String(keyID),
		PolicyName: aws.String(policyName),
	}

	output, err := conn.GetKeyPolicyWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, kms.ErrCodeNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.Policy, nil
}

func FindKeyRotationEnabledByKeyID(ctx context.Context, conn *kms.KMS, keyID string) (*bool, error) {
	input := &kms.GetKeyRotationStatusInput{
		KeyId: aws.String(keyID),
	}

	output, err := conn.GetKeyRotationStatusWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, kms.ErrCodeNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.KeyRotationEnabled, nil
}
