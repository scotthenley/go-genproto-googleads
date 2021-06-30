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
	resourcespb "google.golang.org/genproto/googleapis/ads/googleads/v7/resources"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v7/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newTopicConstantClientHook clientHook

// TopicConstantCallOptions contains the retry settings for each method of TopicConstantClient.
type TopicConstantCallOptions struct {
	GetTopicConstant []gax.CallOption
}

func defaultTopicConstantGRPCClientOptions() []option.ClientOption {
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

func defaultTopicConstantCallOptions() *TopicConstantCallOptions {
	return &TopicConstantCallOptions{
		GetTopicConstant: []gax.CallOption{
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

// internalTopicConstantClient is an interface that defines the methods availaible from Google Ads API.
type internalTopicConstantClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetTopicConstant(context.Context, *servicespb.GetTopicConstantRequest, ...gax.CallOption) (*resourcespb.TopicConstant, error)
}

// TopicConstantClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch topic constants.
type TopicConstantClient struct {
	// The internal transport-dependent client.
	internalClient internalTopicConstantClient

	// The call options for this service.
	CallOptions *TopicConstantCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TopicConstantClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TopicConstantClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TopicConstantClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetTopicConstant returns the requested topic constant in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *TopicConstantClient) GetTopicConstant(ctx context.Context, req *servicespb.GetTopicConstantRequest, opts ...gax.CallOption) (*resourcespb.TopicConstant, error) {
	return c.internalClient.GetTopicConstant(ctx, req, opts...)
}

// topicConstantGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type topicConstantGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TopicConstantClient
	CallOptions **TopicConstantCallOptions

	// The gRPC API client.
	topicConstantClient servicespb.TopicConstantServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTopicConstantClient creates a new topic constant service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch topic constants.
func NewTopicConstantClient(ctx context.Context, opts ...option.ClientOption) (*TopicConstantClient, error) {
	clientOpts := defaultTopicConstantGRPCClientOptions()
	if newTopicConstantClientHook != nil {
		hookOpts, err := newTopicConstantClientHook(ctx, clientHookParams{})
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
	client := TopicConstantClient{CallOptions: defaultTopicConstantCallOptions()}

	c := &topicConstantGRPCClient{
		connPool:            connPool,
		disableDeadlines:    disableDeadlines,
		topicConstantClient: servicespb.NewTopicConstantServiceClient(connPool),
		CallOptions:         &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *topicConstantGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *topicConstantGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *topicConstantGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *topicConstantGRPCClient) GetTopicConstant(ctx context.Context, req *servicespb.GetTopicConstantRequest, opts ...gax.CallOption) (*resourcespb.TopicConstant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTopicConstant[0:len((*c.CallOptions).GetTopicConstant):len((*c.CallOptions).GetTopicConstant)], opts...)
	var resp *resourcespb.TopicConstant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.topicConstantClient.GetTopicConstant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
