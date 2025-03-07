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
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v8/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v8/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newFeedItemTargetClientHook clientHook

// FeedItemTargetCallOptions contains the retry settings for each method of FeedItemTargetClient.
type FeedItemTargetCallOptions struct {
	GetFeedItemTarget     []gax.CallOption
	MutateFeedItemTargets []gax.CallOption
}

func defaultFeedItemTargetGRPCClientOptions() []option.ClientOption {
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

func defaultFeedItemTargetCallOptions() *FeedItemTargetCallOptions {
	return &FeedItemTargetCallOptions{
		GetFeedItemTarget: []gax.CallOption{
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
		MutateFeedItemTargets: []gax.CallOption{
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

// internalFeedItemTargetClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedItemTargetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedItemTarget(context.Context, *servicespb.GetFeedItemTargetRequest, ...gax.CallOption) (*resourcespb.FeedItemTarget, error)
	MutateFeedItemTargets(context.Context, *servicespb.MutateFeedItemTargetsRequest, ...gax.CallOption) (*servicespb.MutateFeedItemTargetsResponse, error)
}

// FeedItemTargetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage feed item targets.
type FeedItemTargetClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedItemTargetClient

	// The call options for this service.
	CallOptions *FeedItemTargetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedItemTargetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedItemTargetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedItemTargetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedItemTarget returns the requested feed item targets in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemTargetClient) GetFeedItemTarget(ctx context.Context, req *servicespb.GetFeedItemTargetRequest, opts ...gax.CallOption) (*resourcespb.FeedItemTarget, error) {
	return c.internalClient.GetFeedItemTarget(ctx, req, opts...)
}

// MutateFeedItemTargets creates or removes feed item targets. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CriterionError (at )
// DatabaseError (at )
// DistinctError (at )
// FeedItemTargetError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NotEmptyError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *FeedItemTargetClient) MutateFeedItemTargets(ctx context.Context, req *servicespb.MutateFeedItemTargetsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemTargetsResponse, error) {
	return c.internalClient.MutateFeedItemTargets(ctx, req, opts...)
}

// feedItemTargetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedItemTargetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedItemTargetClient
	CallOptions **FeedItemTargetCallOptions

	// The gRPC API client.
	feedItemTargetClient servicespb.FeedItemTargetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedItemTargetClient creates a new feed item target service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage feed item targets.
func NewFeedItemTargetClient(ctx context.Context, opts ...option.ClientOption) (*FeedItemTargetClient, error) {
	clientOpts := defaultFeedItemTargetGRPCClientOptions()
	if newFeedItemTargetClientHook != nil {
		hookOpts, err := newFeedItemTargetClientHook(ctx, clientHookParams{})
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
	client := FeedItemTargetClient{CallOptions: defaultFeedItemTargetCallOptions()}

	c := &feedItemTargetGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		feedItemTargetClient: servicespb.NewFeedItemTargetServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedItemTargetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedItemTargetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedItemTargetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedItemTargetGRPCClient) GetFeedItemTarget(ctx context.Context, req *servicespb.GetFeedItemTargetRequest, opts ...gax.CallOption) (*resourcespb.FeedItemTarget, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedItemTarget[0:len((*c.CallOptions).GetFeedItemTarget):len((*c.CallOptions).GetFeedItemTarget)], opts...)
	var resp *resourcespb.FeedItemTarget
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemTargetClient.GetFeedItemTarget(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *feedItemTargetGRPCClient) MutateFeedItemTargets(ctx context.Context, req *servicespb.MutateFeedItemTargetsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemTargetsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateFeedItemTargets[0:len((*c.CallOptions).MutateFeedItemTargets):len((*c.CallOptions).MutateFeedItemTargets)], opts...)
	var resp *servicespb.MutateFeedItemTargetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemTargetClient.MutateFeedItemTargets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
