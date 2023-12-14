// Code generated by smithy-go-codegen DO NOT EDIT.

package b2bi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates some of the parameters for a partnership between a customer and trading
// partner. A partnership represents the connection between you and your trading
// partner. It ties together a profile and one or more trading capabilities.
func (c *Client) UpdatePartnership(ctx context.Context, params *UpdatePartnershipInput, optFns ...func(*Options)) (*UpdatePartnershipOutput, error) {
	if params == nil {
		params = &UpdatePartnershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePartnership", params, optFns, c.addOperationUpdatePartnershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePartnershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePartnershipInput struct {

	// Specifies the unique, system-generated identifier for a partnership.
	//
	// This member is required.
	PartnershipId *string

	// List of the capabilities associated with this partnership.
	Capabilities []string

	// The name of the partnership, used to identify it.
	Name *string

	noSmithyDocumentSerde
}

type UpdatePartnershipOutput struct {

	// Returns a timestamp that identifies the most recent date and time that the
	// partnership was modified.
	//
	// This member is required.
	CreatedAt *time.Time

	// Returns an Amazon Resource Name (ARN) for a specific Amazon Web Services
	// resource, such as a capability, partnership, profile, or transformer.
	//
	// This member is required.
	PartnershipArn *string

	// Returns the unique, system-generated identifier for a partnership.
	//
	// This member is required.
	PartnershipId *string

	// Returns the unique, system-generated identifier for the profile connected to
	// this partnership.
	//
	// This member is required.
	ProfileId *string

	// Returns one or more capabilities associated with this partnership.
	Capabilities []string

	// Returns the email address associated with this trading partner.
	Email *string

	// Returns a timestamp that identifies the most recent date and time that the
	// partnership was modified.
	ModifiedAt *time.Time

	// The name of the partnership, used to identify it.
	Name *string

	// Returns the phone number associated with the partnership.
	Phone *string

	// Returns the unique, system-generated identifier for a trading partner.
	TradingPartnerId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePartnershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdatePartnership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdatePartnership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePartnership"); err != nil {
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
	if err = addOpUpdatePartnershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePartnership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePartnership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePartnership",
	}
}
