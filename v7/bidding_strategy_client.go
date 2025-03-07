// Copyright 2021 Google LLC
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
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v7/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v7/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newBiddingStrategyClientHook clientHook

// BiddingStrategyCallOptions contains the retry settings for each method of BiddingStrategyClient.
type BiddingStrategyCallOptions struct {
	GetBiddingStrategy      []gax.CallOption
	MutateBiddingStrategies []gax.CallOption
}

func defaultBiddingStrategyGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBiddingStrategyCallOptions() *BiddingStrategyCallOptions {
	return &BiddingStrategyCallOptions{
		GetBiddingStrategy: []gax.CallOption{
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
		MutateBiddingStrategies: []gax.CallOption{
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

// internalBiddingStrategyClient is an interface that defines the methods availaible from Google Ads API.
type internalBiddingStrategyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetBiddingStrategy(context.Context, *servicespb.GetBiddingStrategyRequest, ...gax.CallOption) (*resourcespb.BiddingStrategy, error)
	MutateBiddingStrategies(context.Context, *servicespb.MutateBiddingStrategiesRequest, ...gax.CallOption) (*servicespb.MutateBiddingStrategiesResponse, error)
}

// BiddingStrategyClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage bidding strategies.
type BiddingStrategyClient struct {
	// The internal transport-dependent client.
	internalClient internalBiddingStrategyClient

	// The call options for this service.
	CallOptions *BiddingStrategyCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BiddingStrategyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BiddingStrategyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *BiddingStrategyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetBiddingStrategy returns the requested bidding strategy in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *BiddingStrategyClient) GetBiddingStrategy(ctx context.Context, req *servicespb.GetBiddingStrategyRequest, opts ...gax.CallOption) (*resourcespb.BiddingStrategy, error) {
	return c.internalClient.GetBiddingStrategy(ctx, req, opts...)
}

// MutateBiddingStrategies creates, updates, or removes bidding strategies. Operation statuses are
// returned.
//
// List of thrown errors:
// AdxError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// BiddingError (at )
// BiddingStrategyError (at )
// ContextError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *BiddingStrategyClient) MutateBiddingStrategies(ctx context.Context, req *servicespb.MutateBiddingStrategiesRequest, opts ...gax.CallOption) (*servicespb.MutateBiddingStrategiesResponse, error) {
	return c.internalClient.MutateBiddingStrategies(ctx, req, opts...)
}

// biddingStrategyGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type biddingStrategyGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing BiddingStrategyClient
	CallOptions **BiddingStrategyCallOptions

	// The gRPC API client.
	biddingStrategyClient servicespb.BiddingStrategyServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBiddingStrategyClient creates a new bidding strategy service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage bidding strategies.
func NewBiddingStrategyClient(ctx context.Context, opts ...option.ClientOption) (*BiddingStrategyClient, error) {
	clientOpts := defaultBiddingStrategyGRPCClientOptions()
	if newBiddingStrategyClientHook != nil {
		hookOpts, err := newBiddingStrategyClientHook(ctx, clientHookParams{})
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
	client := BiddingStrategyClient{CallOptions: defaultBiddingStrategyCallOptions()}

	c := &biddingStrategyGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		biddingStrategyClient: servicespb.NewBiddingStrategyServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *biddingStrategyGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *biddingStrategyGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *biddingStrategyGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *biddingStrategyGRPCClient) GetBiddingStrategy(ctx context.Context, req *servicespb.GetBiddingStrategyRequest, opts ...gax.CallOption) (*resourcespb.BiddingStrategy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetBiddingStrategy[0:len((*c.CallOptions).GetBiddingStrategy):len((*c.CallOptions).GetBiddingStrategy)], opts...)
	var resp *resourcespb.BiddingStrategy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.biddingStrategyClient.GetBiddingStrategy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *biddingStrategyGRPCClient) MutateBiddingStrategies(ctx context.Context, req *servicespb.MutateBiddingStrategiesRequest, opts ...gax.CallOption) (*servicespb.MutateBiddingStrategiesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateBiddingStrategies[0:len((*c.CallOptions).MutateBiddingStrategies):len((*c.CallOptions).MutateBiddingStrategies)], opts...)
	var resp *servicespb.MutateBiddingStrategiesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.biddingStrategyClient.MutateBiddingStrategies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
