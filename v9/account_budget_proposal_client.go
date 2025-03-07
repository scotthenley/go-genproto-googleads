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

var newAccountBudgetProposalClientHook clientHook

// AccountBudgetProposalCallOptions contains the retry settings for each method of AccountBudgetProposalClient.
type AccountBudgetProposalCallOptions struct {
	GetAccountBudgetProposal    []gax.CallOption
	MutateAccountBudgetProposal []gax.CallOption
}

func defaultAccountBudgetProposalGRPCClientOptions() []option.ClientOption {
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

func defaultAccountBudgetProposalCallOptions() *AccountBudgetProposalCallOptions {
	return &AccountBudgetProposalCallOptions{
		GetAccountBudgetProposal: []gax.CallOption{
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
		MutateAccountBudgetProposal: []gax.CallOption{
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

// internalAccountBudgetProposalClient is an interface that defines the methods availaible from Google Ads API.
type internalAccountBudgetProposalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAccountBudgetProposal(context.Context, *servicespb.GetAccountBudgetProposalRequest, ...gax.CallOption) (*resourcespb.AccountBudgetProposal, error)
	MutateAccountBudgetProposal(context.Context, *servicespb.MutateAccountBudgetProposalRequest, ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error)
}

// AccountBudgetProposalClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing account-level budgets via proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Reads for account-level budgets managed by these proposals will be
// supported in a future version. Until then, please use the
// BudgetOrderService from the AdWords API. Learn more at
// https://developers.google.com/adwords/api/docs/guides/budget-order (at https://developers.google.com/adwords/api/docs/guides/budget-order)
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren’t supported.
// The REMOVE operation cancels a pending proposal.
type AccountBudgetProposalClient struct {
	// The internal transport-dependent client.
	internalClient internalAccountBudgetProposalClient

	// The call options for this service.
	CallOptions *AccountBudgetProposalCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AccountBudgetProposalClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AccountBudgetProposalClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AccountBudgetProposalClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAccountBudgetProposal returns an account-level budget proposal in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *AccountBudgetProposalClient) GetAccountBudgetProposal(ctx context.Context, req *servicespb.GetAccountBudgetProposalRequest, opts ...gax.CallOption) (*resourcespb.AccountBudgetProposal, error) {
	return c.internalClient.GetAccountBudgetProposal(ctx, req, opts...)
}

// MutateAccountBudgetProposal creates, updates, or removes account budget proposals.  Operation statuses
// are returned.
//
// List of thrown errors:
// AccountBudgetProposalError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// DatabaseError (at )
// DateError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
// StringLengthError (at )
func (c *AccountBudgetProposalClient) MutateAccountBudgetProposal(ctx context.Context, req *servicespb.MutateAccountBudgetProposalRequest, opts ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error) {
	return c.internalClient.MutateAccountBudgetProposal(ctx, req, opts...)
}

// accountBudgetProposalGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type accountBudgetProposalGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AccountBudgetProposalClient
	CallOptions **AccountBudgetProposalCallOptions

	// The gRPC API client.
	accountBudgetProposalClient servicespb.AccountBudgetProposalServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAccountBudgetProposalClient creates a new account budget proposal service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing account-level budgets via proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Reads for account-level budgets managed by these proposals will be
// supported in a future version. Until then, please use the
// BudgetOrderService from the AdWords API. Learn more at
// https://developers.google.com/adwords/api/docs/guides/budget-order (at https://developers.google.com/adwords/api/docs/guides/budget-order)
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren’t supported.
// The REMOVE operation cancels a pending proposal.
func NewAccountBudgetProposalClient(ctx context.Context, opts ...option.ClientOption) (*AccountBudgetProposalClient, error) {
	clientOpts := defaultAccountBudgetProposalGRPCClientOptions()
	if newAccountBudgetProposalClientHook != nil {
		hookOpts, err := newAccountBudgetProposalClientHook(ctx, clientHookParams{})
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
	client := AccountBudgetProposalClient{CallOptions: defaultAccountBudgetProposalCallOptions()}

	c := &accountBudgetProposalGRPCClient{
		connPool:                    connPool,
		disableDeadlines:            disableDeadlines,
		accountBudgetProposalClient: servicespb.NewAccountBudgetProposalServiceClient(connPool),
		CallOptions:                 &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *accountBudgetProposalGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *accountBudgetProposalGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *accountBudgetProposalGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *accountBudgetProposalGRPCClient) GetAccountBudgetProposal(ctx context.Context, req *servicespb.GetAccountBudgetProposalRequest, opts ...gax.CallOption) (*resourcespb.AccountBudgetProposal, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAccountBudgetProposal[0:len((*c.CallOptions).GetAccountBudgetProposal):len((*c.CallOptions).GetAccountBudgetProposal)], opts...)
	var resp *resourcespb.AccountBudgetProposal
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.accountBudgetProposalClient.GetAccountBudgetProposal(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *accountBudgetProposalGRPCClient) MutateAccountBudgetProposal(ctx context.Context, req *servicespb.MutateAccountBudgetProposalRequest, opts ...gax.CallOption) (*servicespb.MutateAccountBudgetProposalResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateAccountBudgetProposal[0:len((*c.CallOptions).MutateAccountBudgetProposal):len((*c.CallOptions).MutateAccountBudgetProposal)], opts...)
	var resp *servicespb.MutateAccountBudgetProposalResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.accountBudgetProposalClient.MutateAccountBudgetProposal(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
