package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// BackupLongTermRetentionVaultsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type BackupLongTermRetentionVaultsClient struct {
	ManagementClient
}

// NewBackupLongTermRetentionVaultsClient creates an instance of the BackupLongTermRetentionVaultsClient client.
func NewBackupLongTermRetentionVaultsClient(subscriptionID string) BackupLongTermRetentionVaultsClient {
	return NewBackupLongTermRetentionVaultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupLongTermRetentionVaultsClientWithBaseURI creates an instance of the BackupLongTermRetentionVaultsClient
// client.
func NewBackupLongTermRetentionVaultsClientWithBaseURI(baseURI string, subscriptionID string) BackupLongTermRetentionVaultsClient {
	return BackupLongTermRetentionVaultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates a server backup long term retention vault This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. backupLongTermRetentionVaultName is
// the name of the backup long term retention vault parameters is the required parameters to update a backup long term
// retention vault
func (client BackupLongTermRetentionVaultsClient) CreateOrUpdate(resourceGroupName string, serverName string, backupLongTermRetentionVaultName string, parameters BackupLongTermRetentionVault, cancel <-chan struct{}) (<-chan BackupLongTermRetentionVault, <-chan error) {
	resultChan := make(chan BackupLongTermRetentionVault, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BackupLongTermRetentionVaultProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.BackupLongTermRetentionVaultProperties.RecoveryServicesVaultResourceID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "sql.BackupLongTermRetentionVaultsClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result BackupLongTermRetentionVault
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, serverName, backupLongTermRetentionVaultName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BackupLongTermRetentionVaultsClient) CreateOrUpdatePreparer(resourceGroupName string, serverName string, backupLongTermRetentionVaultName string, parameters BackupLongTermRetentionVault, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupLongTermRetentionVaultName": autorest.Encode("path", backupLongTermRetentionVaultName),
		"resourceGroupName":                autorest.Encode("path", resourceGroupName),
		"serverName":                       autorest.Encode("path", serverName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/backupLongTermRetentionVaults/{backupLongTermRetentionVaultName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BackupLongTermRetentionVaultsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client),
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BackupLongTermRetentionVaultsClient) CreateOrUpdateResponder(resp *http.Response) (result BackupLongTermRetentionVault, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a server backup long term retention vault
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. backupLongTermRetentionVaultName is
// the name of the Azure SQL Server backup LongTermRetention vault
func (client BackupLongTermRetentionVaultsClient) Get(resourceGroupName string, serverName string, backupLongTermRetentionVaultName string) (result BackupLongTermRetentionVault, err error) {
	req, err := client.GetPreparer(resourceGroupName, serverName, backupLongTermRetentionVaultName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.BackupLongTermRetentionVaultsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BackupLongTermRetentionVaultsClient) GetPreparer(resourceGroupName string, serverName string, backupLongTermRetentionVaultName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupLongTermRetentionVaultName": autorest.Encode("path", backupLongTermRetentionVaultName),
		"resourceGroupName":                autorest.Encode("path", resourceGroupName),
		"serverName":                       autorest.Encode("path", serverName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/backupLongTermRetentionVaults/{backupLongTermRetentionVaultName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BackupLongTermRetentionVaultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BackupLongTermRetentionVaultsClient) GetResponder(resp *http.Response) (result BackupLongTermRetentionVault, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
