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

var newCarrierConstantClientHook clientHook

// CarrierConstantCallOptions contains the retry settings for each method of CarrierConstantClient.
type CarrierConstantCallOptions struct {
	GetCarrierConstant []gax.CallOption
}

func defaultCarrierConstantGRPCClientOptions() []option.ClientOption {
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

func defaultCarrierConstantCallOptions() *CarrierConstantCallOptions {
	return &CarrierConstantCallOptions{
		GetCarrierConstant: []gax.CallOption{
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

// internalCarrierConstantClient is an interface that defines the methods availaible from Google Ads API.
type internalCarrierConstantClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCarrierConstant(context.Context, *servicespb.GetCarrierConstantRequest, ...gax.CallOption) (*resourcespb.CarrierConstant, error)
}

// CarrierConstantClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch carrier constants.
type CarrierConstantClient struct {
	// The internal transport-dependent client.
	internalClient internalCarrierConstantClient

	// The call options for this service.
	CallOptions *CarrierConstantCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CarrierConstantClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CarrierConstantClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CarrierConstantClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCarrierConstant returns the requested carrier constant in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CarrierConstantClient) GetCarrierConstant(ctx context.Context, req *servicespb.GetCarrierConstantRequest, opts ...gax.CallOption) (*resourcespb.CarrierConstant, error) {
	return c.internalClient.GetCarrierConstant(ctx, req, opts...)
}

// carrierConstantGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type carrierConstantGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CarrierConstantClient
	CallOptions **CarrierConstantCallOptions

	// The gRPC API client.
	carrierConstantClient servicespb.CarrierConstantServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCarrierConstantClient creates a new carrier constant service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch carrier constants.
func NewCarrierConstantClient(ctx context.Context, opts ...option.ClientOption) (*CarrierConstantClient, error) {
	clientOpts := defaultCarrierConstantGRPCClientOptions()
	if newCarrierConstantClientHook != nil {
		hookOpts, err := newCarrierConstantClientHook(ctx, clientHookParams{})
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
	client := CarrierConstantClient{CallOptions: defaultCarrierConstantCallOptions()}

	c := &carrierConstantGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		carrierConstantClient: servicespb.NewCarrierConstantServiceClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *carrierConstantGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *carrierConstantGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *carrierConstantGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *carrierConstantGRPCClient) GetCarrierConstant(ctx context.Context, req *servicespb.GetCarrierConstantRequest, opts ...gax.CallOption) (*resourcespb.CarrierConstant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCarrierConstant[0:len((*c.CallOptions).GetCarrierConstant):len((*c.CallOptions).GetCarrierConstant)], opts...)
	var resp *resourcespb.CarrierConstant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.carrierConstantClient.GetCarrierConstant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
