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
	servicespb "google.golang.org/genproto/googleapis/ads/googleads/v7/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newConversionUploadClientHook clientHook

// ConversionUploadCallOptions contains the retry settings for each method of ConversionUploadClient.
type ConversionUploadCallOptions struct {
	UploadClickConversions []gax.CallOption
	UploadCallConversions  []gax.CallOption
}

func defaultConversionUploadClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultConversionUploadCallOptions() *ConversionUploadCallOptions {
	return &ConversionUploadCallOptions{
		UploadClickConversions: []gax.CallOption{
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
		UploadCallConversions: []gax.CallOption{
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

// ConversionUploadClient is a client for interacting with Google Ads API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ConversionUploadClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	conversionUploadClient servicespb.ConversionUploadServiceClient

	// The call options for this service.
	CallOptions *ConversionUploadCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewConversionUploadClient creates a new conversion upload service client.
//
// Service to upload conversions.
func NewConversionUploadClient(ctx context.Context, opts ...option.ClientOption) (*ConversionUploadClient, error) {
	clientOpts := defaultConversionUploadClientOptions()

	if newConversionUploadClientHook != nil {
		hookOpts, err := newConversionUploadClientHook(ctx, clientHookParams{})
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
	c := &ConversionUploadClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultConversionUploadCallOptions(),

		conversionUploadClient: servicespb.NewConversionUploadServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ConversionUploadClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ConversionUploadClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ConversionUploadClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// UploadClickConversions processes the given click conversions.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ConversionUploadError (at )
// HeaderError (at )
// InternalError (at )
// PartialFailureError (at )
// QuotaError (at )
// RequestError (at )
func (c *ConversionUploadClient) UploadClickConversions(ctx context.Context, req *servicespb.UploadClickConversionsRequest, opts ...gax.CallOption) (*servicespb.UploadClickConversionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UploadClickConversions[0:len(c.CallOptions.UploadClickConversions):len(c.CallOptions.UploadClickConversions)], opts...)
	var resp *servicespb.UploadClickConversionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionUploadClient.UploadClickConversions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UploadCallConversions processes the given call conversions.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// PartialFailureError (at )
// QuotaError (at )
// RequestError (at )
func (c *ConversionUploadClient) UploadCallConversions(ctx context.Context, req *servicespb.UploadCallConversionsRequest, opts ...gax.CallOption) (*servicespb.UploadCallConversionsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UploadCallConversions[0:len(c.CallOptions.UploadCallConversions):len(c.CallOptions.UploadCallConversions)], opts...)
	var resp *servicespb.UploadCallConversionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.conversionUploadClient.UploadCallConversions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
