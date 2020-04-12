package queue

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQS provides operations for SQS. It requests to SQS
// via `sqs.SQS` module.
type SQS struct {
	sqs      *sqs.SQS
	queueURL string
}

// NewSQS returns a new SQS.
func NewSQS(queueURL string, p client.ConfigProvider, cfgs ...*aws.Config) *SQS {
	return &SQS{
		sqs:      sqs.New(p, cfgs...),
		queueURL: queueURL,
	}
}

// Send sends a message to SQS via `sqs.SQS` module. You can use `Opt`
// function if you want to set some options for `SendMessageInput`.
func (s *SQS) Send(body string, opts ...Opt) (SendMessageOutput, error) {
	smi := &sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(s.queueURL),
	}

	for _, f := range opts {
		f(smi)
	}

	return s.sqs.SendMessage(smi)
}

// SendWithContext sends a message to SQS via `sqs.SQS` module. You can use `Opt`
// function if you want to set some options for `SendMessageInput`.
func (s *SQS) SendWithContext(ctx context.Context, body string, opts ...Opt) (SendMessageOutput, error) {
	smi := &sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(s.queueURL),
	}

	for _, f := range opts {
		f(smi)
	}

	return s.sqs.SendMessageWithContext(ctx, smi)
}

// SQSDelaySeconds returns the function that sets delay seconds option
// for `SendMessageInput`.
func SQSDelaySeconds(v int64) func(*sqs.SendMessageInput) {
	return func(s *sqs.SendMessageInput) {
		s.SetDelaySeconds(v)
	}
}

// SQSMessageAttributes returns the function that sets message attibutes
// option for `SendMessageInput`.
func SQSMessageAttributes(v map[string]*sqs.MessageAttributeValue) func(*sqs.SendMessageInput) {
	return func(s *sqs.SendMessageInput) {
		s.SetMessageAttributes(v)
	}
}
