// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/scotthenley/go-genproto-googleads/pb/v9/resources"
	servicespb "github.com/scotthenley/go-genproto-googleads/pb/v9/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAdParameterClientHook clientHook

// AdParameterCallOptions contains the retry settings for each method of AdParameterClient.
type AdParameterCallOptions struct {
	GetAdParameter     []gax.CallOption
	MutateAdParameters []gax.CallOption
}

func defaultAdParameterGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAdParameterCallOptions() *AdParameterCallOptions {
	return &AdParameterCallOptions{
		GetAdParameter: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MutateAdParameters: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalAdParameterClient is an interface that defines the methods availaible from Google Ads API.
type internalAdParameterClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAdParameter(context.Context, *servicespb.GetAdParameterRequest, ...gax.CallOption) (*resourcespb.AdParameter, error)
	MutateAdParameters(context.Context, *servicespb.MutateAdParametersRequest, ...gax.CallOption) (*servicespb.MutateAdParametersResponse, error)
}

// AdParameterClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad parameters.
type AdParameterClient struct {
	// The internal transport-dependent client.
	internalClient internalAdParameterClient

	// The call options for this service.
	CallOptions *AdParameterCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdParameterClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdParameterClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdParameterClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAdParameter returns the requested ad parameter in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdParameterClient) GetAdParameter(ctx context.Context, req *servicespb.GetAdParameterRequest, opts ...gax.CallOption) (*resourcespb.AdParameter, error) {
	return c.internalClient.GetAdParameter(ctx, req, opts...)
}

// MutateAdParameters creates, updates, or removes ad parameters. Operation statuses are
// returned.
//
// List of thrown errors:
// AdParameterError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// ContextError (at )
// DatabaseError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdParameterClient) MutateAdParameters(ctx context.Context, req *servicespb.MutateAdParametersRequest, opts ...gax.CallOption) (*servicespb.MutateAdParametersResponse, error) {
	return c.internalClient.MutateAdParameters(ctx, req, opts...)
}

// adParameterGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adParameterGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdParameterClient
	CallOptions **AdParameterCallOptions

	// The gRPC API client.
	adParameterClient servicespb.AdParameterServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdParameterClient creates a new ad parameter service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad parameters.
func NewAdParameterClient(ctx context.Context, opts ...option.ClientOption) (*AdParameterClient, error) {
	clientOpts := defaultAdParameterGRPCClientOptions()
	if newAdParameterClientHook != nil {
		hookOpts, err := newAdParameterClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdParameterClient{CallOptions: defaultAdParameterCallOptions()}

	c := &adParameterGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		adParameterClient: servicespb.NewAdParameterServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adParameterGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adParameterGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adParameterGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adParameterGRPCClient) GetAdParameter(ctx context.Context, req *servicespb.GetAdParameterRequest, opts ...gax.CallOption) (*resourcespb.AdParameter, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAdParameter[0:len((*c.CallOptions).GetAdParameter):len((*c.CallOptions).GetAdParameter)], opts...)
	var resp *resourcespb.AdParameter
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adParameterClient.GetAdParameter(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adParameterGRPCClient) MutateAdParameters(ctx context.Context, req *servicespb.MutateAdParametersRequest, opts ...gax.CallOption) (*servicespb.MutateAdParametersResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdParameters[0:len((*c.CallOptions).MutateAdParameters):len((*c.CallOptions).MutateAdParameters)], opts...)
	var resp *servicespb.MutateAdParametersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adParameterClient.MutateAdParameters(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
