//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
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
	client := &OperationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Lists the supported operations.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// options - OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
func (client *OperationsClient) List(ctx context.Context, options *OperationsClientListOptions) (OperationsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DeploymentManager/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsList); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}
