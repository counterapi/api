package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"time"
)

type SNSClient struct {
	Client *sns.Client
}

func NewSNS(cfg aws.Config) *SNSClient {
	return &SNSClient{Client: sns.NewFromConfig(cfg)}
}

func (c SNSClient) Publish(ctx context.Context, topicArn *string, message *string) error {
	publishInput := sns.PublishInput{
		TopicArn:               topicArn,
		Message:                message,
		MessageGroupId:         aws.String(fmt.Sprintf("counterAPI-%s", time.Now().UnixMicro())),
		MessageDeduplicationId: aws.String(fmt.Sprintf("counterAPI-%s", time.Now().UnixMicro())),
	}

	_, err := c.Client.Publish(ctx, &publishInput)
	if err != nil {
		return err
	}

	return err
}
