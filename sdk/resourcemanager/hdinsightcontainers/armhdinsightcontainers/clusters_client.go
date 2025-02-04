//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ClustersClient contains the methods for the Clusters group.
// Don't use this type directly, use NewClustersClient() instead.
type ClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClustersClient creates a new instance of ClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClustersClient, error) {
	cl, err := arm.NewClient(moduleName+".ClustersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - hdInsightCluster - The cluster to create.
//   - options - ClustersClientBeginCreateOptions contains the optional parameters for the ClustersClient.BeginCreate method.
func (client *ClustersClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, hdInsightCluster Cluster, options *ClustersClientBeginCreateOptions) (*runtime.Poller[ClustersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterPoolName, clusterName, hdInsightCluster, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClustersClient) create(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, hdInsightCluster Cluster, options *ClustersClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, hdInsightCluster, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ClustersClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, hdInsightCluster Cluster, options *ClustersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hdInsightCluster); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
func (client *ClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*runtime.Poller[ClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterPoolName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
func (client *ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientGetOptions) (ClustersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClustersClient) getHandleResponse(resp *http.Response) (ClustersClientGetResponse, error) {
	result := ClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientGetResponse{}, err
	}
	return result, nil
}

// GetInstanceView - Gets the status of a cluster instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClustersClientGetInstanceViewOptions contains the optional parameters for the ClustersClient.GetInstanceView
//     method.
func (client *ClustersClient) GetInstanceView(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientGetInstanceViewOptions) (ClustersClientGetInstanceViewResponse, error) {
	var err error
	req, err := client.getInstanceViewCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
	if err != nil {
		return ClustersClientGetInstanceViewResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientGetInstanceViewResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientGetInstanceViewResponse{}, err
	}
	resp, err := client.getInstanceViewHandleResponse(httpResp)
	return resp, err
}

// getInstanceViewCreateRequest creates the GetInstanceView request.
func (client *ClustersClient) getInstanceViewCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientGetInstanceViewOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/instanceViews/default"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInstanceViewHandleResponse handles the GetInstanceView response.
func (client *ClustersClient) getInstanceViewHandleResponse(resp *http.Response) (ClustersClientGetInstanceViewResponse, error) {
	result := ClustersClientGetInstanceViewResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterInstanceViewResult); err != nil {
		return ClustersClientGetInstanceViewResponse{}, err
	}
	return result, nil
}

// NewListByClusterPoolNamePager - Lists the HDInsight cluster pools under a resource group.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - options - ClustersClientListByClusterPoolNameOptions contains the optional parameters for the ClustersClient.NewListByClusterPoolNamePager
//     method.
func (client *ClustersClient) NewListByClusterPoolNamePager(resourceGroupName string, clusterPoolName string, options *ClustersClientListByClusterPoolNameOptions) *runtime.Pager[ClustersClientListByClusterPoolNameResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListByClusterPoolNameResponse]{
		More: func(page ClustersClientListByClusterPoolNameResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListByClusterPoolNameResponse) (ClustersClientListByClusterPoolNameResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByClusterPoolNameCreateRequest(ctx, resourceGroupName, clusterPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListByClusterPoolNameResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClustersClientListByClusterPoolNameResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListByClusterPoolNameResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterPoolNameHandleResponse(resp)
		},
	})
}

// listByClusterPoolNameCreateRequest creates the ListByClusterPoolName request.
func (client *ClustersClient) listByClusterPoolNameCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClustersClientListByClusterPoolNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterPoolNameHandleResponse handles the ListByClusterPoolName response.
func (client *ClustersClient) listByClusterPoolNameHandleResponse(resp *http.Response) (ClustersClientListByClusterPoolNameResponse, error) {
	result := ClustersClientListByClusterPoolNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterListResult); err != nil {
		return ClustersClientListByClusterPoolNameResponse{}, err
	}
	return result, nil
}

// NewListInstanceViewsPager - Lists the lists of instance views
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClustersClientListInstanceViewsOptions contains the optional parameters for the ClustersClient.NewListInstanceViewsPager
//     method.
func (client *ClustersClient) NewListInstanceViewsPager(resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientListInstanceViewsOptions) *runtime.Pager[ClustersClientListInstanceViewsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListInstanceViewsResponse]{
		More: func(page ClustersClientListInstanceViewsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListInstanceViewsResponse) (ClustersClientListInstanceViewsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listInstanceViewsCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListInstanceViewsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClustersClientListInstanceViewsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListInstanceViewsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listInstanceViewsHandleResponse(resp)
		},
	})
}

// listInstanceViewsCreateRequest creates the ListInstanceViews request.
func (client *ClustersClient) listInstanceViewsCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientListInstanceViewsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/instanceViews"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listInstanceViewsHandleResponse handles the ListInstanceViews response.
func (client *ClustersClient) listInstanceViewsHandleResponse(resp *http.Response) (ClustersClientListInstanceViewsResponse, error) {
	result := ClustersClientListInstanceViewsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterInstanceViewsResult); err != nil {
		return ClustersClientListInstanceViewsResponse{}, err
	}
	return result, nil
}

// NewListServiceConfigsPager - Lists the config dump of all services running in cluster.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClustersClientListServiceConfigsOptions contains the optional parameters for the ClustersClient.NewListServiceConfigsPager
//     method.
func (client *ClustersClient) NewListServiceConfigsPager(resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientListServiceConfigsOptions) *runtime.Pager[ClustersClientListServiceConfigsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListServiceConfigsResponse]{
		More: func(page ClustersClientListServiceConfigsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListServiceConfigsResponse) (ClustersClientListServiceConfigsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listServiceConfigsCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClustersClientListServiceConfigsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClustersClientListServiceConfigsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClustersClientListServiceConfigsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listServiceConfigsHandleResponse(resp)
		},
	})
}

// listServiceConfigsCreateRequest creates the ListServiceConfigs request.
func (client *ClustersClient) listServiceConfigsCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClustersClientListServiceConfigsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/serviceConfigs"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listServiceConfigsHandleResponse handles the ListServiceConfigs response.
func (client *ClustersClient) listServiceConfigsHandleResponse(resp *http.Response) (ClustersClientListServiceConfigsResponse, error) {
	result := ClustersClientListServiceConfigsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceConfigListResult); err != nil {
		return ClustersClientListServiceConfigsResponse{}, err
	}
	return result, nil
}

// BeginResize - Resize an existing Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - clusterResizeRequest - Resize a cluster.
//   - options - ClustersClientBeginResizeOptions contains the optional parameters for the ClustersClient.BeginResize method.
func (client *ClustersClient) BeginResize(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterResizeRequest ClusterResizeData, options *ClustersClientBeginResizeOptions) (*runtime.Poller[ClustersClientResizeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resize(ctx, resourceGroupName, clusterPoolName, clusterName, clusterResizeRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientResizeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientResizeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Resize - Resize an existing Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClustersClient) resize(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterResizeRequest ClusterResizeData, options *ClustersClientBeginResizeOptions) (*http.Response, error) {
	var err error
	req, err := client.resizeCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, clusterResizeRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resizeCreateRequest creates the Resize request.
func (client *ClustersClient) resizeCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterResizeRequest ClusterResizeData, options *ClustersClientBeginResizeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/resize"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterResizeRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Updates an existing Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - clusterPatchRequest - Patch a cluster.
//   - options - ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
func (client *ClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterPatchRequest ClusterPatch, options *ClustersClientBeginUpdateOptions) (*runtime.Poller[ClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterPoolName, clusterName, clusterPatchRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClustersClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing Cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClustersClient) update(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterPatchRequest ClusterPatch, options *ClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, clusterPatchRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterPatchRequest ClusterPatch, options *ClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterPatchRequest); err != nil {
		return nil, err
	}
	return req, nil
}
