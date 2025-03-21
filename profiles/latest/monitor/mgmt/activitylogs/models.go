//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package activitylogs

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2020-10-01/activitylogs"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EventLevel = original.EventLevel

const (
	EventLevelCritical      EventLevel = original.EventLevelCritical
	EventLevelError         EventLevel = original.EventLevelError
	EventLevelInformational EventLevel = original.EventLevelInformational
	EventLevelVerbose       EventLevel = original.EventLevelVerbose
	EventLevelWarning       EventLevel = original.EventLevelWarning
)

type ActionGroup = original.ActionGroup
type ActionList = original.ActionList
type AlertResource = original.AlertResource
type AlertRuleAllOfCondition = original.AlertRuleAllOfCondition
type AlertRuleAnyOfOrLeafCondition = original.AlertRuleAnyOfOrLeafCondition
type AlertRuleLeafCondition = original.AlertRuleLeafCondition
type AlertRuleList = original.AlertRuleList
type AlertRuleListIterator = original.AlertRuleListIterator
type AlertRuleListPage = original.AlertRuleListPage
type AlertRulePatchObject = original.AlertRulePatchObject
type AlertRulePatchProperties = original.AlertRulePatchProperties
type AlertRuleProperties = original.AlertRuleProperties
type AlertsClient = original.AlertsClient
type AzureResource = original.AzureResource
type BaseClient = original.BaseClient
type Client = original.Client
type ErrorResponse = original.ErrorResponse
type EventData = original.EventData
type EventDataCollection = original.EventDataCollection
type EventDataCollectionIterator = original.EventDataCollectionIterator
type EventDataCollectionPage = original.EventDataCollectionPage
type HTTPRequestInfo = original.HTTPRequestInfo
type LocalizableString = original.LocalizableString
type SenderAuthorization = original.SenderAuthorization
type TenantActivityLogsClient = original.TenantActivityLogsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAlertRuleListIterator(page AlertRuleListPage) AlertRuleListIterator {
	return original.NewAlertRuleListIterator(page)
}
func NewAlertRuleListPage(cur AlertRuleList, getNextPage func(context.Context, AlertRuleList) (AlertRuleList, error)) AlertRuleListPage {
	return original.NewAlertRuleListPage(cur, getNextPage)
}
func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventDataCollectionIterator(page EventDataCollectionPage) EventDataCollectionIterator {
	return original.NewEventDataCollectionIterator(page)
}
func NewEventDataCollectionPage(cur EventDataCollection, getNextPage func(context.Context, EventDataCollection) (EventDataCollection, error)) EventDataCollectionPage {
	return original.NewEventDataCollectionPage(cur, getNextPage)
}
func NewTenantActivityLogsClient(subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClient(subscriptionID)
}
func NewTenantActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleEventLevelValues() []EventLevel {
	return original.PossibleEventLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
