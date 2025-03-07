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

var newRecommendationClientHook clientHook

// RecommendationCallOptions contains the retry settings for each method of RecommendationClient.
type RecommendationCallOptions struct {
	GetRecommendation     []gax.CallOption
	ApplyRecommendation   []gax.CallOption
	DismissRecommendation []gax.CallOption
}

func defaultRecommendationGRPCClientOptions() []option.ClientOption {
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

func defaultRecommendationCallOptions() *RecommendationCallOptions {
	return &RecommendationCallOptions{
		GetRecommendation: []gax.CallOption{
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
		ApplyRecommendation: []gax.CallOption{
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
		DismissRecommendation: []gax.CallOption{
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

// internalRecommendationClient is an interface that defines the methods availaible from Google Ads API.
type internalRecommendationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetRecommendation(context.Context, *servicespb.GetRecommendationRequest, ...gax.CallOption) (*resourcespb.Recommendation, error)
	ApplyRecommendation(context.Context, *servicespb.ApplyRecommendationRequest, ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error)
	DismissRecommendation(context.Context, *servicespb.DismissRecommendationRequest, ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error)
}

// RecommendationClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage recommendations.
type RecommendationClient struct {
	// The internal transport-dependent client.
	internalClient internalRecommendationClient

	// The call options for this service.
	CallOptions *RecommendationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RecommendationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RecommendationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RecommendationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetRecommendation returns the requested recommendation in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *RecommendationClient) GetRecommendation(ctx context.Context, req *servicespb.GetRecommendationRequest, opts ...gax.CallOption) (*resourcespb.Recommendation, error) {
	return c.internalClient.GetRecommendation(ctx, req, opts...)
}

// ApplyRecommendation applies given recommendations with corresponding apply parameters.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RecommendationError (at )
// RequestError (at )
// UrlFieldError (at )
func (c *RecommendationClient) ApplyRecommendation(ctx context.Context, req *servicespb.ApplyRecommendationRequest, opts ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error) {
	return c.internalClient.ApplyRecommendation(ctx, req, opts...)
}

// DismissRecommendation dismisses given recommendations.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RecommendationError (at )
// RequestError (at )
func (c *RecommendationClient) DismissRecommendation(ctx context.Context, req *servicespb.DismissRecommendationRequest, opts ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error) {
	return c.internalClient.DismissRecommendation(ctx, req, opts...)
}

// recommendationGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type recommendationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing RecommendationClient
	CallOptions **RecommendationCallOptions

	// The gRPC API client.
	recommendationClient servicespb.RecommendationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRecommendationClient creates a new recommendation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage recommendations.
func NewRecommendationClient(ctx context.Context, opts ...option.ClientOption) (*RecommendationClient, error) {
	clientOpts := defaultRecommendationGRPCClientOptions()
	if newRecommendationClientHook != nil {
		hookOpts, err := newRecommendationClientHook(ctx, clientHookParams{})
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
	client := RecommendationClient{CallOptions: defaultRecommendationCallOptions()}

	c := &recommendationGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		recommendationClient: servicespb.NewRecommendationServiceClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *recommendationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *recommendationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *recommendationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *recommendationGRPCClient) GetRecommendation(ctx context.Context, req *servicespb.GetRecommendationRequest, opts ...gax.CallOption) (*resourcespb.Recommendation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetRecommendation[0:len((*c.CallOptions).GetRecommendation):len((*c.CallOptions).GetRecommendation)], opts...)
	var resp *resourcespb.Recommendation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.GetRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *recommendationGRPCClient) ApplyRecommendation(ctx context.Context, req *servicespb.ApplyRecommendationRequest, opts ...gax.CallOption) (*servicespb.ApplyRecommendationResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ApplyRecommendation[0:len((*c.CallOptions).ApplyRecommendation):len((*c.CallOptions).ApplyRecommendation)], opts...)
	var resp *servicespb.ApplyRecommendationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.ApplyRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *recommendationGRPCClient) DismissRecommendation(ctx context.Context, req *servicespb.DismissRecommendationRequest, opts ...gax.CallOption) (*servicespb.DismissRecommendationResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DismissRecommendation[0:len((*c.CallOptions).DismissRecommendation):len((*c.CallOptions).DismissRecommendation)], opts...)
	var resp *servicespb.DismissRecommendationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.recommendationClient.DismissRecommendation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
