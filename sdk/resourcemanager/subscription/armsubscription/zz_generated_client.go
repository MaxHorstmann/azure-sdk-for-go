//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Subscription group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host string
	pl   runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginAcceptOwnership - Accept subscription ownership.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// subscriptionID - Subscription Id.
// options - ClientBeginAcceptOwnershipOptions contains the optional parameters for the Client.BeginAcceptOwnership method.
func (client *Client) BeginAcceptOwnership(ctx context.Context, subscriptionID string, body AcceptOwnershipRequest, options *ClientBeginAcceptOwnershipOptions) (*runtime.Poller[ClientAcceptOwnershipResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.acceptOwnership(ctx, subscriptionID, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ClientAcceptOwnershipResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ClientAcceptOwnershipResponse](options.ResumeToken, client.pl, nil)
	}
}

// AcceptOwnership - Accept subscription ownership.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
func (client *Client) acceptOwnership(ctx context.Context, subscriptionID string, body AcceptOwnershipRequest, options *ClientBeginAcceptOwnershipOptions) (*http.Response, error) {
	req, err := client.acceptOwnershipCreateRequest(ctx, subscriptionID, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// acceptOwnershipCreateRequest creates the AcceptOwnership request.
func (client *Client) acceptOwnershipCreateRequest(ctx context.Context, subscriptionID string, body AcceptOwnershipRequest, options *ClientBeginAcceptOwnershipOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/subscriptions/{subscriptionId}/acceptOwnership"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// AcceptOwnershipStatus - Accept subscription ownership status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// subscriptionID - Subscription Id.
// options - ClientAcceptOwnershipStatusOptions contains the optional parameters for the Client.AcceptOwnershipStatus method.
func (client *Client) AcceptOwnershipStatus(ctx context.Context, subscriptionID string, options *ClientAcceptOwnershipStatusOptions) (ClientAcceptOwnershipStatusResponse, error) {
	req, err := client.acceptOwnershipStatusCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return ClientAcceptOwnershipStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientAcceptOwnershipStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientAcceptOwnershipStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.acceptOwnershipStatusHandleResponse(resp)
}

// acceptOwnershipStatusCreateRequest creates the AcceptOwnershipStatus request.
func (client *Client) acceptOwnershipStatusCreateRequest(ctx context.Context, subscriptionID string, options *ClientAcceptOwnershipStatusOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/subscriptions/{subscriptionId}/acceptOwnershipStatus"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// acceptOwnershipStatusHandleResponse handles the AcceptOwnershipStatus response.
func (client *Client) acceptOwnershipStatusHandleResponse(resp *http.Response) (ClientAcceptOwnershipStatusResponse, error) {
	result := ClientAcceptOwnershipStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AcceptOwnershipStatusResponse); err != nil {
		return ClientAcceptOwnershipStatusResponse{}, err
	}
	return result, nil
}

// Cancel - The operation to cancel a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// subscriptionID - Subscription Id.
// options - ClientCancelOptions contains the optional parameters for the Client.Cancel method.
func (client *Client) Cancel(ctx context.Context, subscriptionID string, options *ClientCancelOptions) (ClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return ClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *Client) cancelCreateRequest(ctx context.Context, subscriptionID string, options *ClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/cancel"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *Client) cancelHandleResponse(resp *http.Response) (ClientCancelResponse, error) {
	result := ClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CanceledSubscriptionID); err != nil {
		return ClientCancelResponse{}, err
	}
	return result, nil
}

// Enable - The operation to enable a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// subscriptionID - Subscription Id.
// options - ClientEnableOptions contains the optional parameters for the Client.Enable method.
func (client *Client) Enable(ctx context.Context, subscriptionID string, options *ClientEnableOptions) (ClientEnableResponse, error) {
	req, err := client.enableCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return ClientEnableResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientEnableResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientEnableResponse{}, runtime.NewResponseError(resp)
	}
	return client.enableHandleResponse(resp)
}

// enableCreateRequest creates the Enable request.
func (client *Client) enableCreateRequest(ctx context.Context, subscriptionID string, options *ClientEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/enable"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// enableHandleResponse handles the Enable response.
func (client *Client) enableHandleResponse(resp *http.Response) (ClientEnableResponse, error) {
	result := ClientEnableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnabledSubscriptionID); err != nil {
		return ClientEnableResponse{}, err
	}
	return result, nil
}

// Rename - The operation to rename a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-01
// subscriptionID - Subscription Id.
// body - Subscription Name
// options - ClientRenameOptions contains the optional parameters for the Client.Rename method.
func (client *Client) Rename(ctx context.Context, subscriptionID string, body Name, options *ClientRenameOptions) (ClientRenameResponse, error) {
	req, err := client.renameCreateRequest(ctx, subscriptionID, body, options)
	if err != nil {
		return ClientRenameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClientRenameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientRenameResponse{}, runtime.NewResponseError(resp)
	}
	return client.renameHandleResponse(resp)
}

// renameCreateRequest creates the Rename request.
func (client *Client) renameCreateRequest(ctx context.Context, subscriptionID string, body Name, options *ClientRenameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/rename"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// renameHandleResponse handles the Rename response.
func (client *Client) renameHandleResponse(resp *http.Response) (ClientRenameResponse, error) {
	result := ClientRenameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RenamedSubscriptionID); err != nil {
		return ClientRenameResponse{}, err
	}
	return result, nil
}
