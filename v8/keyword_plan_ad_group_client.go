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

var newKeywordPlanAdGroupClientHook clientHook

// KeywordPlanAdGroupCallOptions contains the retry settings for each method of KeywordPlanAdGroupClient.
type KeywordPlanAdGroupCallOptions struct {
	GetKeywordPlanAdGroup     []gax.CallOption
	MutateKeywordPlanAdGroups []gax.CallOption
}

func defaultKeywordPlanAdGroupGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordPlanAdGroupCallOptions() *KeywordPlanAdGroupCallOptions {
	return &KeywordPlanAdGroupCallOptions{
		GetKeywordPlanAdGroup: []gax.CallOption{
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
		MutateKeywordPlanAdGroups: []gax.CallOption{
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

// internalKeywordPlanAdGroupClient is an interface that defines the methods availaible from Google Ads API.
type internalKeywordPlanAdGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetKeywordPlanAdGroup(context.Context, *servicespb.GetKeywordPlanAdGroupRequest, ...gax.CallOption) (*resourcespb.KeywordPlanAdGroup, error)
	MutateKeywordPlanAdGroups(context.Context, *servicespb.MutateKeywordPlanAdGroupsRequest, ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupsResponse, error)
}

// KeywordPlanAdGroupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage Keyword Plan ad groups.
type KeywordPlanAdGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordPlanAdGroupClient

	// The call options for this service.
	CallOptions *KeywordPlanAdGroupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordPlanAdGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordPlanAdGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *KeywordPlanAdGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetKeywordPlanAdGroup returns the requested Keyword Plan ad group in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordPlanAdGroupClient) GetKeywordPlanAdGroup(ctx context.Context, req *servicespb.GetKeywordPlanAdGroupRequest, opts ...gax.CallOption) (*resourcespb.KeywordPlanAdGroup, error) {
	return c.internalClient.GetKeywordPlanAdGroup(ctx, req, opts...)
}

// MutateKeywordPlanAdGroups creates, updates, or removes Keyword Plan ad groups. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// KeywordPlanAdGroupError (at )
// KeywordPlanError (at )
// MutateError (at )
// NewResourceCreationError (at )
// QuotaError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
func (c *KeywordPlanAdGroupClient) MutateKeywordPlanAdGroups(ctx context.Context, req *servicespb.MutateKeywordPlanAdGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupsResponse, error) {
	return c.internalClient.MutateKeywordPlanAdGroups(ctx, req, opts...)
}

// keywordPlanAdGroupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordPlanAdGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing KeywordPlanAdGroupClient
	CallOptions **KeywordPlanAdGroupCallOptions

	// The gRPC API client.
	keywordPlanAdGroupClient servicespb.KeywordPlanAdGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewKeywordPlanAdGroupClient creates a new keyword plan ad group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage Keyword Plan ad groups.
func NewKeywordPlanAdGroupClient(ctx context.Context, opts ...option.ClientOption) (*KeywordPlanAdGroupClient, error) {
	clientOpts := defaultKeywordPlanAdGroupGRPCClientOptions()
	if newKeywordPlanAdGroupClientHook != nil {
		hookOpts, err := newKeywordPlanAdGroupClientHook(ctx, clientHookParams{})
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
	client := KeywordPlanAdGroupClient{CallOptions: defaultKeywordPlanAdGroupCallOptions()}

	c := &keywordPlanAdGroupGRPCClient{
		connPool:                 connPool,
		disableDeadlines:         disableDeadlines,
		keywordPlanAdGroupClient: servicespb.NewKeywordPlanAdGroupServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *keywordPlanAdGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordPlanAdGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordPlanAdGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordPlanAdGroupGRPCClient) GetKeywordPlanAdGroup(ctx context.Context, req *servicespb.GetKeywordPlanAdGroupRequest, opts ...gax.CallOption) (*resourcespb.KeywordPlanAdGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKeywordPlanAdGroup[0:len((*c.CallOptions).GetKeywordPlanAdGroup):len((*c.CallOptions).GetKeywordPlanAdGroup)], opts...)
	var resp *resourcespb.KeywordPlanAdGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanAdGroupClient.GetKeywordPlanAdGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordPlanAdGroupGRPCClient) MutateKeywordPlanAdGroups(ctx context.Context, req *servicespb.MutateKeywordPlanAdGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateKeywordPlanAdGroupsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateKeywordPlanAdGroups[0:len((*c.CallOptions).MutateKeywordPlanAdGroups):len((*c.CallOptions).MutateKeywordPlanAdGroups)], opts...)
	var resp *servicespb.MutateKeywordPlanAdGroupsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordPlanAdGroupClient.MutateKeywordPlanAdGroups(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
