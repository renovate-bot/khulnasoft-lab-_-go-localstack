// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the URL of an existing Amazon SQS queue. To access a queue that belongs
// to another AWS account, use the QueueOwnerAWSAccountId parameter to specify the
// account ID of the queue's owner. The queue's owner must grant you permission to
// access the queue. For more information about shared queue access, see
<<<<<<< HEAD
// AddPermission or see Allow Developers to Write Messages to a Shared Queue
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
=======
// AddPermission or see Allow Developers to Write Messages to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
// in the Amazon SQS Developer Guide.
func (c *Client) GetQueueUrl(ctx context.Context, params *GetQueueUrlInput, optFns ...func(*Options)) (*GetQueueUrlOutput, error) {
	if params == nil {
		params = &GetQueueUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueueUrl", params, optFns, c.addOperationGetQueueUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

<<<<<<< HEAD
//
type GetQueueUrlInput struct {

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens (-), and underscores (_). Queue URLs
	// and names are case-sensitive.
=======
type GetQueueUrlInput struct {

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens ( - ), and underscores ( _ ). Queue
	// URLs and names are case-sensitive.
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	//
	// This member is required.
	QueueName *string

	// The Amazon Web Services account ID of the account that created the queue.
	QueueOwnerAWSAccountId *string

	noSmithyDocumentSerde
}

<<<<<<< HEAD
// For more information, see Interpreting Responses
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
=======
// For more information, see Interpreting Responses (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
// in the Amazon SQS Developer Guide.
type GetQueueUrlOutput struct {

	// The URL of the queue.
	QueueUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueueUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
<<<<<<< HEAD
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetQueueUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetQueueUrl{}, middleware.After)
	if err != nil {
		return err
	}
=======
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetQueueUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetQueueUrl{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueueUrl"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
<<<<<<< HEAD
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
=======
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
<<<<<<< HEAD
	if err = addClientUserAgent(stack); err != nil {
=======
	if err = addClientUserAgent(stack, options); err != nil {
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	if err = addOpGetQueueUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueUrl(options.Region), middleware.Before); err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
<<<<<<< HEAD
=======
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
	return nil
}

func newServiceMetadataMiddleware_opGetQueueUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
<<<<<<< HEAD
		SigningName:   "sqs",
=======
>>>>>>> 86c663831051e23db463a649fa07cd05ab84e189
		OperationName: "GetQueueUrl",
	}
}
