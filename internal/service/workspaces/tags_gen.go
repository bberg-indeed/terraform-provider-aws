// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package workspaces

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// ListTags lists workspaces service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn *workspaces.WorkSpaces, identifier string) (tftags.KeyValueTags, error) {
	input := &workspaces.DescribeTagsInput{
		ResourceId: aws.String(identifier),
	}

	output, err := conn.DescribeTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.TagList), nil
}

// []*SERVICE.Tag handling

// Tags returns workspaces service tags.
func Tags(tags tftags.KeyValueTags) []*workspaces.Tag {
	result := make([]*workspaces.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &workspaces.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from workspaces service tags.
func KeyValueTags(tags []*workspaces.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}

// UpdateTags updates workspaces service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn *workspaces.WorkSpaces, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &workspaces.DeleteTagsInput{
			ResourceId: aws.String(identifier),
			TagKeys:    aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.DeleteTags(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &workspaces.CreateTagsInput{
			ResourceId: aws.String(identifier),
			Tags:       Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.CreateTags(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
