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
	resourcespb "google.golang.org/genproto/googleapis/ads/googleads/v8/resources"
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v8/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newMerchantCenterLinkClientHook clientHook

// MerchantCenterLinkCallOptions contains the retry settings for each method of MerchantCenterLinkClient.
type MerchantCenterLinkCallOptions struct {
	ListMerchantCenterLinks  []gax.CallOption
	GetMerchantCenterLink    []gax.CallOption
	MutateMerchantCenterLink []gax.CallOption
}

func defaultMerchantCenterLinkGRPCClientOptions() []option.ClientOption {
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

func defaultMerchantCenterLinkCallOptions() *MerchantCenterLinkCallOptions {
	return &MerchantCenterLinkCallOptions{
		ListMerchantCenterLinks: []gax.CallOption{
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
		GetMerchantCenterLink: []gax.CallOption{
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
		MutateMerchantCenterLink: []gax.CallOption{
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

// internalMerchantCenterLinkClient is an interface that defines the methods availaible from Google Ads API.
type internalMerchantCenterLinkClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListMerchantCenterLinks(context.Context, *servicespb.ListMerchantCenterLinksRequest, ...gax.CallOption) (*servicespb.ListMerchantCenterLinksResponse, error)
	GetMerchantCenterLink(context.Context, *servicespb.GetMerchantCenterLinkRequest, ...gax.CallOption) (*resourcespb.MerchantCenterLink, error)
	MutateMerchantCenterLink(context.Context, *servicespb.MutateMerchantCenterLinkRequest, ...gax.CallOption) (*servicespb.MutateMerchantCenterLinkResponse, error)
}

// MerchantCenterLinkClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of links between Google Ads and Google
// Merchant Center.
type MerchantCenterLinkClient struct {
	// The internal transport-dependent client.
	internalClient internalMerchantCenterLinkClient

	// The call options for this service.
	CallOptions *MerchantCenterLinkCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *MerchantCenterLinkClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *MerchantCenterLinkClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *MerchantCenterLinkClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListMerchantCenterLinks returns Merchant Center links available for this customer.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *MerchantCenterLinkClient) ListMerchantCenterLinks(ctx context.Context, req *servicespb.ListMerchantCenterLinksRequest, opts ...gax.CallOption) (*servicespb.ListMerchantCenterLinksResponse, error) {
	return c.internalClient.ListMerchantCenterLinks(ctx, req, opts...)
}

// GetMerchantCenterLink returns the Merchant Center link in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *MerchantCenterLinkClient) GetMerchantCenterLink(ctx context.Context, req *servicespb.GetMerchantCenterLinkRequest, opts ...gax.CallOption) (*resourcespb.MerchantCenterLink, error) {
	return c.internalClient.GetMerchantCenterLink(ctx, req, opts...)
}

// MutateMerchantCenterLink updates status or removes a Merchant Center link.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *MerchantCenterLinkClient) MutateMerchantCenterLink(ctx context.Context, req *servicespb.MutateMerchantCenterLinkRequest, opts ...gax.CallOption) (*servicespb.MutateMerchantCenterLinkResponse, error) {
	return c.internalClient.MutateMerchantCenterLink(ctx, req, opts...)
}

// merchantCenterLinkGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type merchantCenterLinkGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing MerchantCenterLinkClient
	CallOptions **MerchantCenterLinkCallOptions

	// The gRPC API client.
	merchantCenterLinkClient servicespb.MerchantCenterLinkServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewMerchantCenterLinkClient creates a new merchant center link service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of links between Google Ads and Google
// Merchant Center.
func NewMerchantCenterLinkClient(ctx context.Context, opts ...option.ClientOption) (*MerchantCenterLinkClient, error) {
	clientOpts := defaultMerchantCenterLinkGRPCClientOptions()
	if newMerchantCenterLinkClientHook != nil {
		hookOpts, err := newMerchantCenterLinkClientHook(ctx, clientHookParams{})
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
	client := MerchantCenterLinkClient{CallOptions: defaultMerchantCenterLinkCallOptions()}

	c := &merchantCenterLinkGRPCClient{
		connPool:                 connPool,
		disableDeadlines:         disableDeadlines,
		merchantCenterLinkClient: servicespb.NewMerchantCenterLinkServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *merchantCenterLinkGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *merchantCenterLinkGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *merchantCenterLinkGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *merchantCenterLinkGRPCClient) ListMerchantCenterLinks(ctx context.Context, req *servicespb.ListMerchantCenterLinksRequest, opts ...gax.CallOption) (*servicespb.ListMerchantCenterLinksResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListMerchantCenterLinks[0:len((*c.CallOptions).ListMerchantCenterLinks):len((*c.CallOptions).ListMerchantCenterLinks)], opts...)
	var resp *servicespb.ListMerchantCenterLinksResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.merchantCenterLinkClient.ListMerchantCenterLinks(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *merchantCenterLinkGRPCClient) GetMerchantCenterLink(ctx context.Context, req *servicespb.GetMerchantCenterLinkRequest, opts ...gax.CallOption) (*resourcespb.MerchantCenterLink, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetMerchantCenterLink[0:len((*c.CallOptions).GetMerchantCenterLink):len((*c.CallOptions).GetMerchantCenterLink)], opts...)
	var resp *resourcespb.MerchantCenterLink
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.merchantCenterLinkClient.GetMerchantCenterLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *merchantCenterLinkGRPCClient) MutateMerchantCenterLink(ctx context.Context, req *servicespb.MutateMerchantCenterLinkRequest, opts ...gax.CallOption) (*servicespb.MutateMerchantCenterLinkResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateMerchantCenterLink[0:len((*c.CallOptions).MutateMerchantCenterLink):len((*c.CallOptions).MutateMerchantCenterLink)], opts...)
	var resp *servicespb.MutateMerchantCenterLinkResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.merchantCenterLinkClient.MutateMerchantCenterLink(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
