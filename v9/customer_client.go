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

var newCustomerClientHook clientHook

// CustomerCallOptions contains the retry settings for each method of CustomerClient.
type CustomerCallOptions struct {
	GetCustomer             []gax.CallOption
	MutateCustomer          []gax.CallOption
	ListAccessibleCustomers []gax.CallOption
	CreateCustomerClient    []gax.CallOption
}

func defaultCustomerGRPCClientOptions() []option.ClientOption {
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

func defaultCustomerCallOptions() *CustomerCallOptions {
	return &CustomerCallOptions{
		GetCustomer: []gax.CallOption{
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
		MutateCustomer: []gax.CallOption{
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
		ListAccessibleCustomers: []gax.CallOption{
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
		CreateCustomerClient: []gax.CallOption{
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

// internalCustomerClient is an interface that defines the methods availaible from Google Ads API.
type internalCustomerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCustomer(context.Context, *servicespb.GetCustomerRequest, ...gax.CallOption) (*resourcespb.Customer, error)
	MutateCustomer(context.Context, *servicespb.MutateCustomerRequest, ...gax.CallOption) (*servicespb.MutateCustomerResponse, error)
	ListAccessibleCustomers(context.Context, *servicespb.ListAccessibleCustomersRequest, ...gax.CallOption) (*servicespb.ListAccessibleCustomersResponse, error)
	CreateCustomerClient(context.Context, *servicespb.CreateCustomerClientRequest, ...gax.CallOption) (*servicespb.CreateCustomerClientResponse, error)
}

// CustomerClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customers.
type CustomerClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomerClient

	// The call options for this service.
	CallOptions *CustomerCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CustomerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCustomer returns the requested customer in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerClient) GetCustomer(ctx context.Context, req *servicespb.GetCustomerRequest, opts ...gax.CallOption) (*resourcespb.Customer, error) {
	return c.internalClient.GetCustomer(ctx, req, opts...)
}

// MutateCustomer updates a customer. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
// UrlFieldError (at )
func (c *CustomerClient) MutateCustomer(ctx context.Context, req *servicespb.MutateCustomerRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerResponse, error) {
	return c.internalClient.MutateCustomer(ctx, req, opts...)
}

// ListAccessibleCustomers returns resource names of customers directly accessible by the
// user authenticating the call.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomerClient) ListAccessibleCustomers(ctx context.Context, req *servicespb.ListAccessibleCustomersRequest, opts ...gax.CallOption) (*servicespb.ListAccessibleCustomersResponse, error) {
	return c.internalClient.ListAccessibleCustomers(ctx, req, opts...)
}

// CreateCustomerClient creates a new client under manager. The new client customer is returned.
//
// List of thrown errors:
// AccessInvitationError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CurrencyCodeError (at )
// HeaderError (at )
// InternalError (at )
// ManagerLinkError (at )
// QuotaError (at )
// RequestError (at )
// StringLengthError (at )
// TimeZoneError (at )
func (c *CustomerClient) CreateCustomerClient(ctx context.Context, req *servicespb.CreateCustomerClientRequest, opts ...gax.CallOption) (*servicespb.CreateCustomerClientResponse, error) {
	return c.internalClient.CreateCustomerClient(ctx, req, opts...)
}

// customerGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CustomerClient
	CallOptions **CustomerCallOptions

	// The gRPC API client.
	customerClient servicespb.CustomerServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomerClient creates a new customer service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customers.
func NewCustomerClient(ctx context.Context, opts ...option.ClientOption) (*CustomerClient, error) {
	clientOpts := defaultCustomerGRPCClientOptions()
	if newCustomerClientHook != nil {
		hookOpts, err := newCustomerClientHook(ctx, clientHookParams{})
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
	client := CustomerClient{CallOptions: defaultCustomerCallOptions()}

	c := &customerGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		customerClient:   servicespb.NewCustomerServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *customerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customerGRPCClient) GetCustomer(ctx context.Context, req *servicespb.GetCustomerRequest, opts ...gax.CallOption) (*resourcespb.Customer, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCustomer[0:len((*c.CallOptions).GetCustomer):len((*c.CallOptions).GetCustomer)], opts...)
	var resp *resourcespb.Customer
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerClient.GetCustomer(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customerGRPCClient) MutateCustomer(ctx context.Context, req *servicespb.MutateCustomerRequest, opts ...gax.CallOption) (*servicespb.MutateCustomerResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomer[0:len((*c.CallOptions).MutateCustomer):len((*c.CallOptions).MutateCustomer)], opts...)
	var resp *servicespb.MutateCustomerResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerClient.MutateCustomer(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customerGRPCClient) ListAccessibleCustomers(ctx context.Context, req *servicespb.ListAccessibleCustomersRequest, opts ...gax.CallOption) (*servicespb.ListAccessibleCustomersResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListAccessibleCustomers[0:len((*c.CallOptions).ListAccessibleCustomers):len((*c.CallOptions).ListAccessibleCustomers)], opts...)
	var resp *servicespb.ListAccessibleCustomersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerClient.ListAccessibleCustomers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customerGRPCClient) CreateCustomerClient(ctx context.Context, req *servicespb.CreateCustomerClientRequest, opts ...gax.CallOption) (*servicespb.CreateCustomerClientResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCustomerClient[0:len((*c.CallOptions).CreateCustomerClient):len((*c.CallOptions).CreateCustomerClient)], opts...)
	var resp *servicespb.CreateCustomerClientResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customerClient.CreateCustomerClient(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
