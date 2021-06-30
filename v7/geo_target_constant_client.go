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

var newGeoTargetConstantClientHook clientHook

// GeoTargetConstantCallOptions contains the retry settings for each method of GeoTargetConstantClient.
type GeoTargetConstantCallOptions struct {
	GetGeoTargetConstant      []gax.CallOption
	SuggestGeoTargetConstants []gax.CallOption
}

func defaultGeoTargetConstantGRPCClientOptions() []option.ClientOption {
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

func defaultGeoTargetConstantCallOptions() *GeoTargetConstantCallOptions {
	return &GeoTargetConstantCallOptions{
		GetGeoTargetConstant: []gax.CallOption{
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
		SuggestGeoTargetConstants: []gax.CallOption{
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

// internalGeoTargetConstantClient is an interface that defines the methods availaible from Google Ads API.
type internalGeoTargetConstantClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetGeoTargetConstant(context.Context, *servicespb.GetGeoTargetConstantRequest, ...gax.CallOption) (*resourcespb.GeoTargetConstant, error)
	SuggestGeoTargetConstants(context.Context, *servicespb.SuggestGeoTargetConstantsRequest, ...gax.CallOption) (*servicespb.SuggestGeoTargetConstantsResponse, error)
}

// GeoTargetConstantClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch geo target constants.
type GeoTargetConstantClient struct {
	// The internal transport-dependent client.
	internalClient internalGeoTargetConstantClient

	// The call options for this service.
	CallOptions *GeoTargetConstantCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GeoTargetConstantClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GeoTargetConstantClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GeoTargetConstantClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetGeoTargetConstant returns the requested geo target constant in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *GeoTargetConstantClient) GetGeoTargetConstant(ctx context.Context, req *servicespb.GetGeoTargetConstantRequest, opts ...gax.CallOption) (*resourcespb.GeoTargetConstant, error) {
	return c.internalClient.GetGeoTargetConstant(ctx, req, opts...)
}

// SuggestGeoTargetConstants returns GeoTargetConstant suggestions by location name or by resource name.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// GeoTargetConstantSuggestionError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *GeoTargetConstantClient) SuggestGeoTargetConstants(ctx context.Context, req *servicespb.SuggestGeoTargetConstantsRequest, opts ...gax.CallOption) (*servicespb.SuggestGeoTargetConstantsResponse, error) {
	return c.internalClient.SuggestGeoTargetConstants(ctx, req, opts...)
}

// geoTargetConstantGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type geoTargetConstantGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing GeoTargetConstantClient
	CallOptions **GeoTargetConstantCallOptions

	// The gRPC API client.
	geoTargetConstantClient servicespb.GeoTargetConstantServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGeoTargetConstantClient creates a new geo target constant service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch geo target constants.
func NewGeoTargetConstantClient(ctx context.Context, opts ...option.ClientOption) (*GeoTargetConstantClient, error) {
	clientOpts := defaultGeoTargetConstantGRPCClientOptions()
	if newGeoTargetConstantClientHook != nil {
		hookOpts, err := newGeoTargetConstantClientHook(ctx, clientHookParams{})
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
	client := GeoTargetConstantClient{CallOptions: defaultGeoTargetConstantCallOptions()}

	c := &geoTargetConstantGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		geoTargetConstantClient: servicespb.NewGeoTargetConstantServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *geoTargetConstantGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *geoTargetConstantGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *geoTargetConstantGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *geoTargetConstantGRPCClient) GetGeoTargetConstant(ctx context.Context, req *servicespb.GetGeoTargetConstantRequest, opts ...gax.CallOption) (*resourcespb.GeoTargetConstant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetGeoTargetConstant[0:len((*c.CallOptions).GetGeoTargetConstant):len((*c.CallOptions).GetGeoTargetConstant)], opts...)
	var resp *resourcespb.GeoTargetConstant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.geoTargetConstantClient.GetGeoTargetConstant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *geoTargetConstantGRPCClient) SuggestGeoTargetConstants(ctx context.Context, req *servicespb.SuggestGeoTargetConstantsRequest, opts ...gax.CallOption) (*servicespb.SuggestGeoTargetConstantsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SuggestGeoTargetConstants[0:len((*c.CallOptions).SuggestGeoTargetConstants):len((*c.CallOptions).SuggestGeoTargetConstants)], opts...)
	var resp *servicespb.SuggestGeoTargetConstantsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.geoTargetConstantClient.SuggestGeoTargetConstants(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
