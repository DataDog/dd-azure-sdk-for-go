//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// ProductPolicyClient contains the methods for the ProductPolicy group.
// Don't use this type directly, use NewProductPolicyClient() instead.
type ProductPolicyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProductPolicyClient creates a new instance of ProductPolicyClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProductPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductPolicyClient, error) {
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
	client := &ProductPolicyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates policy configuration for the Product.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// parameters - The policy contents to apply.
// options - ProductPolicyClientCreateOrUpdateOptions contains the optional parameters for the ProductPolicyClient.CreateOrUpdate
// method.
func (client *ProductPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, parameters PolicyContract, options *ProductPolicyClientCreateOrUpdateOptions) (ProductPolicyClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, productID, policyID, parameters, options)
	if err != nil {
		return ProductPolicyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductPolicyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProductPolicyClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductPolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, parameters PolicyContract, options *ProductPolicyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProductPolicyClient) createOrUpdateHandleResponse(resp *http.Response) (ProductPolicyClientCreateOrUpdateResponse, error) {
	result := ProductPolicyClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return ProductPolicyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the policy configuration at the Product.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - ProductPolicyClientDeleteOptions contains the optional parameters for the ProductPolicyClient.Delete method.
func (client *ProductPolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, ifMatch string, options *ProductPolicyClientDeleteOptions) (ProductPolicyClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, productID, policyID, ifMatch, options)
	if err != nil {
		return ProductPolicyClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductPolicyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProductPolicyClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProductPolicyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductPolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, ifMatch string, options *ProductPolicyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the policy configuration at the Product level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// options - ProductPolicyClientGetOptions contains the optional parameters for the ProductPolicyClient.Get method.
func (client *ProductPolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, options *ProductPolicyClientGetOptions) (ProductPolicyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, productID, policyID, options)
	if err != nil {
		return ProductPolicyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductPolicyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProductPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, options *ProductPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("format", string(*options.Format))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProductPolicyClient) getHandleResponse(resp *http.Response) (ProductPolicyClientGetResponse, error) {
	result := ProductPolicyClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return ProductPolicyClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Get the ETag of the policy configuration at the Product level.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// policyID - The identifier of the Policy.
// options - ProductPolicyClientGetEntityTagOptions contains the optional parameters for the ProductPolicyClient.GetEntityTag
// method.
func (client *ProductPolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, options *ProductPolicyClientGetEntityTagOptions) (ProductPolicyClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, productID, policyID, options)
	if err != nil {
		return ProductPolicyClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductPolicyClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductPolicyClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *ProductPolicyClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID PolicyIDName, options *ProductPolicyClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *ProductPolicyClient) getEntityTagHandleResponse(resp *http.Response) (ProductPolicyClientGetEntityTagResponse, error) {
	result := ProductPolicyClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// ListByProduct - Get the policy configuration at the Product level.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-01
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// productID - Product identifier. Must be unique in the current API Management service instance.
// options - ProductPolicyClientListByProductOptions contains the optional parameters for the ProductPolicyClient.ListByProduct
// method.
func (client *ProductPolicyClient) ListByProduct(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductPolicyClientListByProductOptions) (ProductPolicyClientListByProductResponse, error) {
	req, err := client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
	if err != nil {
		return ProductPolicyClientListByProductResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProductPolicyClientListByProductResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductPolicyClientListByProductResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByProductHandleResponse(resp)
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ProductPolicyClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductPolicyClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *ProductPolicyClient) listByProductHandleResponse(resp *http.Response) (ProductPolicyClientListByProductResponse, error) {
	result := ProductPolicyClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyCollection); err != nil {
		return ProductPolicyClientListByProductResponse{}, err
	}
	return result, nil
}
