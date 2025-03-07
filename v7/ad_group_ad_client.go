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

var newAdGroupAdClientHook clientHook

// AdGroupAdCallOptions contains the retry settings for each method of AdGroupAdClient.
type AdGroupAdCallOptions struct {
	GetAdGroupAd     []gax.CallOption
	MutateAdGroupAds []gax.CallOption
}

func defaultAdGroupAdGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupAdCallOptions() *AdGroupAdCallOptions {
	return &AdGroupAdCallOptions{
		GetAdGroupAd: []gax.CallOption{
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
		MutateAdGroupAds: []gax.CallOption{
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

// internalAdGroupAdClient is an interface that defines the methods availaible from Google Ads API.
type internalAdGroupAdClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAdGroupAd(context.Context, *servicespb.GetAdGroupAdRequest, ...gax.CallOption) (*resourcespb.AdGroupAd, error)
	MutateAdGroupAds(context.Context, *servicespb.MutateAdGroupAdsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error)
}

// AdGroupAdClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ads in an ad group.
type AdGroupAdClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupAdClient

	// The call options for this service.
	CallOptions *AdGroupAdCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupAdClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupAdClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AdGroupAdClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAdGroupAd returns the requested ad in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupAdClient) GetAdGroupAd(ctx context.Context, req *servicespb.GetAdGroupAdRequest, opts ...gax.CallOption) (*resourcespb.AdGroupAd, error) {
	return c.internalClient.GetAdGroupAd(ctx, req, opts...)
}

// MutateAdGroupAds creates, updates, or removes ads. Operation statuses are returned.
//
// List of thrown errors:
// AdCustomizerError (at )
// AdError (at )
// AdGroupAdError (at )
// AdSharingError (at )
// AdxError (at )
// AssetError (at )
// AssetLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// ContextError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FeedAttributeReferenceError (at )
// FieldError (at )
// FieldMaskError (at )
// FunctionError (at )
// FunctionParsingError (at )
// HeaderError (at )
// IdError (at )
// ImageError (at )
// InternalError (at )
// ListOperationError (at )
// MediaBundleError (at )
// MediaFileError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperationAccessDeniedError (at )
// OperatorError (at )
// PolicyFindingError (at )
// PolicyValidationParameterError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdGroupAdClient) MutateAdGroupAds(ctx context.Context, req *servicespb.MutateAdGroupAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error) {
	return c.internalClient.MutateAdGroupAds(ctx, req, opts...)
}

// adGroupAdGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupAdGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AdGroupAdClient
	CallOptions **AdGroupAdCallOptions

	// The gRPC API client.
	adGroupAdClient servicespb.AdGroupAdServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAdGroupAdClient creates a new ad group ad service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ads in an ad group.
func NewAdGroupAdClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupAdClient, error) {
	clientOpts := defaultAdGroupAdGRPCClientOptions()
	if newAdGroupAdClientHook != nil {
		hookOpts, err := newAdGroupAdClientHook(ctx, clientHookParams{})
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
	client := AdGroupAdClient{CallOptions: defaultAdGroupAdCallOptions()}

	c := &adGroupAdGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		adGroupAdClient:  servicespb.NewAdGroupAdServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *adGroupAdGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupAdGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupAdGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupAdGRPCClient) GetAdGroupAd(ctx context.Context, req *servicespb.GetAdGroupAdRequest, opts ...gax.CallOption) (*resourcespb.AdGroupAd, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAdGroupAd[0:len((*c.CallOptions).GetAdGroupAd):len((*c.CallOptions).GetAdGroupAd)], opts...)
	var resp *resourcespb.AdGroupAd
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupAdClient.GetAdGroupAd(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adGroupAdGRPCClient) MutateAdGroupAds(ctx context.Context, req *servicespb.MutateAdGroupAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAdGroupAds[0:len((*c.CallOptions).MutateAdGroupAds):len((*c.CallOptions).MutateAdGroupAds)], opts...)
	var resp *servicespb.MutateAdGroupAdsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupAdClient.MutateAdGroupAds(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
