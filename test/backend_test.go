package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestTerraformS3Backend(t *testing.T) {
	t.Parallel()
	testRandomValue := strings.ToLower(random.UniqueId())
	expectedBucketName := fmt.Sprintf("terratest-backend-%s", testRandomValue)
	expectedDynamoDbTableName := fmt.Sprintf("terratest-backend-%s", testRandomValue)
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"bucket_name":         expectedBucketName,
			"dynamodb_table_name": expectedDynamoDbTableName,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
	actualBucketArn := terraform.Output(t, terraformOptions, "s3_bucket_arn")
	actualDynamoDbTableArn := terraform.Output(t, terraformOptions, "dynamodb_table_arn")
	expectedBucketArn := fmt.Sprintf("arn:aws:s3:::%s", expectedBucketName)
	expectedDynamoDbTableArn := fmt.Sprintf("arn:aws:dynamodb:%s:\\d{12}:table/%s", awsRegion, expectedDynamoDbTableName)
	assert.Equal(t, expectedBucketArn, actualBucketArn)
	assert.Regexp(t, expectedDynamoDbTableArn, actualDynamoDbTableArn)

}
