// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a specified graph snapshot.
func (c *Client) GetGraphSnapshot(ctx context.Context, params *GetGraphSnapshotInput, optFns ...func(*Options)) (*GetGraphSnapshotOutput, error) {
	if params == nil {
		params = &GetGraphSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGraphSnapshot", params, optFns, c.addOperationGetGraphSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGraphSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGraphSnapshotInput struct {

	// The ID of the snapshot to retrieve.
	//
	// This member is required.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

func (in *GetGraphSnapshotInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type GetGraphSnapshotOutput struct {

	// The ARN of the graph snapshot.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph snapshot.
	//
	// This member is required.
	Id *string

	// The snapshot name. For example: my-snapshot-1 . The name must contain from 1 to
	// 63 letters, numbers, or hyphens, and its first character must be a letter. It
	// cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	Name *string

	// The ID of the KMS key used to encrypt and decrypt the snapshot.
	KmsKeyIdentifier *string

	// The time when the snapshot was created.
	SnapshotCreateTime *time.Time

	// The graph identifier for the graph for which a snapshot is to be created.
	SourceGraphId *string

	// The status of the graph snapshot.
	Status types.SnapshotStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGraphSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGraphSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGraphSnapshot"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetGraphSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGraphSnapshot(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetGraphSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGraphSnapshot",
	}
}
