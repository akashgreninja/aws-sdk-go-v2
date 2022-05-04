// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Lightsail load balancer. To learn more about deciding whether to load
// balance your application, see Configure your Lightsail instances for load
// balancing
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/configure-lightsail-instances-for-load-balancing).
// You can create up to 5 load balancers per AWS Region in your account. When you
// create a load balancer, you can specify a unique name and port settings. To
// change additional load balancer settings, use the UpdateLoadBalancerAttribute
// operation. The create load balancer operation supports tag-based access control
// via request tags. For more information, see the Amazon Lightsail Developer Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancer", params, optFns, c.addOperationCreateLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoadBalancerInput struct {

	// The instance port where you're creating your load balancer.
	//
	// This member is required.
	InstancePort int32

	// The name of your load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The optional alternative domains and subdomains to use with your SSL/TLS
	// certificate (e.g., www.example.com, example.com, m.example.com,
	// blog.example.com).
	CertificateAlternativeNames []string

	// The domain name with which your certificate is associated (e.g., example.com).
	// If you specify certificateDomainName, then certificateName is required (and
	// vice-versa).
	CertificateDomainName *string

	// The name of the SSL/TLS certificate. If you specify certificateName, then
	// certificateDomainName is required (and vice-versa).
	CertificateName *string

	// The path you provided to perform the load balancer health check. If you didn't
	// specify a health check path, Lightsail uses the root path of your website (e.g.,
	// "/"). You may want to specify a custom health check path other than the root of
	// your application if your home page loads slowly or has a lot of media or
	// scripting on it.
	HealthCheckPath *string

	// The IP address type for the load balancer. The possible values are ipv4 for IPv4
	// only, and dualstack for IPv4 and IPv6. The default value is dualstack.
	IpAddressType types.IpAddressType

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []types.Tag

	// The name of the TLS policy to apply to the load balancer. Use the
	// GetLoadBalancerTlsPolicies
	// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetLoadBalancerTlsPolicies.html)
	// action to get a list of TLS policy names that you can specify. For more
	// information about load balancer TLS policies, see Configuring TLS security
	// policies on your Amazon Lightsail load balancers
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configure-load-balancer-tls-security-policy)
	// in the Amazon Lightsail Developer Guide.
	TlsPolicyName *string

	noSmithyDocumentSerde
}

type CreateLoadBalancerOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancer(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opCreateLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateLoadBalancer",
	}
}
