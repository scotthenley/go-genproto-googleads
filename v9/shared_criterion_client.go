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

var newSharedCriterionClientHook clientHook

// SharedCriterionCallOptions contains the retry settings for each method of SharedCriterionClient.
type SharedCriterionCallOptions struct {
	GetSharedCriterion   []gax.CallOption
	MutateSharedCriteria []gax.CallOption
}

func defaultSharedCriterionGRPCClientOptions() []option.ClientOption {
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

func defaultSharedCriterionCallOptions() *SharedCriterionCallOptions {
	return &SharedCriterionCallOptions{
		GetSharedCriterion: []gax.CallOption{
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
		MutateSharedCriteria: []gax.CallOption{
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

// internalSharedCriterionClient is an interface that defines the methods availaible from Google Ads API.
type internalSharedCriterionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetSharedCriterion(context.Context, *servicespb.GetSharedCriterionRequest, ...gax.CallOption) (*resourcespb.SharedCriterion, error)
	MutateSharedCriteria(context.Context, *servicespb.MutateSharedCriteriaRequest, ...gax.CallOption) (*servicespb.MutateSharedCriteriaResponse, error)
}

// SharedCriterionClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage shared criteria.
type SharedCriterionClient struct {
	// The internal transport-dependent client.
	internalClient internalSharedCriterionClient

	// The call options for this service.
	CallOptions *SharedCriterionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SharedCriterionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SharedCriterionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SharedCriterionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetSharedCriterion returns the requested shared criterion in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *SharedCriterionClient) GetSharedCriterion(ctx context.Context, req *servicespb.GetSharedCriterionRequest, opts ...gax.CallOption) (*resourcespb.SharedCriterion, error) {
	return c.internalClient.GetSharedCriterion(ctx, req, opts...)
}

// MutateSharedCriteria creates or removes shared criteria. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CriterionError (at )
// DatabaseError (at )
// DistinctError (at )
// FieldError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// MutateError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
func (c *SharedCriterionClient) MutateSharedCriteria(ctx context.Context, req *servicespb.MutateSharedCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateSharedCriteriaResponse, error) {
	return c.internalClient.MutateSharedCriteria(ctx, req, opts...)
}

// sharedCriterionGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type sharedCriterionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SharedCriterionClient
	CallOptions **SharedCriterionCallOptions

	// The gRPC API client.
	sharedCriterionClient servicespb.SharedCriterionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSharedCriterionClient creates a new shared criterion service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage shared criteria.
func NewSharedCriterionClient(ctx context.Context, opts ...option.ClientOption) (*SharedCriterionClient, error) {
	clientOpts := defaultSharedCriterionGRPCClientOptions()
	if newSharedCriterionClientHook != nil {
		hookOpts, err := newSharedCriterionClientHook(ctx, clientHookParams{})
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
	client := SharedCriterionClient{CallOptions: defaultSharedCriterionCallOptions()}

	c := &sharedCriterionGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		sharedCriterionClient: servicespb.NewSharedCriterionServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *sharedCriterionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *sharedCriterionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *sharedCriterionGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *sharedCriterionGRPCClient) GetSharedCriterion(ctx context.Context, req *servicespb.GetSharedCriterionRequest, opts ...gax.CallOption) (*resourcespb.SharedCriterion, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSharedCriterion[0:len((*c.CallOptions).GetSharedCriterion):len((*c.CallOptions).GetSharedCriterion)], opts...)
	var resp *resourcespb.SharedCriterion
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sharedCriterionClient.GetSharedCriterion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sharedCriterionGRPCClient) MutateSharedCriteria(ctx context.Context, req *servicespb.MutateSharedCriteriaRequest, opts ...gax.CallOption) (*servicespb.MutateSharedCriteriaResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateSharedCriteria[0:len((*c.CallOptions).MutateSharedCriteria):len((*c.CallOptions).MutateSharedCriteria)], opts...)
	var resp *servicespb.MutateSharedCriteriaResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sharedCriterionClient.MutateSharedCriteria(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
