package appconfiguration

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
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2019-10-01/appconfiguration"

// APIKey an API key used for authenticating with a configuration store endpoint.
type APIKey struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The key ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; A name for the key describing its usage.
	Name *string `json:"name,omitempty"`
	// Value - READ-ONLY; The value of the key that is used for authentication purposes.
	Value *string `json:"value,omitempty"`
	// ConnectionString - READ-ONLY; A connection string that can be used by supporting clients for authentication.
	ConnectionString *string `json:"connectionString,omitempty"`
	// LastModified - READ-ONLY; The last time any of the key's properties were modified.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// ReadOnly - READ-ONLY; Whether this key can only be used for read operations.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// APIKeyListResult the result of a request to list API keys.
type APIKeyListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection value.
	Value *[]APIKey `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`
}

// APIKeyListResultIterator provides access to a complete listing of APIKey values.
type APIKeyListResultIterator struct {
	i    int
	page APIKeyListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *APIKeyListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeyListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *APIKeyListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter APIKeyListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter APIKeyListResultIterator) Response() APIKeyListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter APIKeyListResultIterator) Value() APIKey {
	if !iter.page.NotDone() {
		return APIKey{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the APIKeyListResultIterator type.
func NewAPIKeyListResultIterator(page APIKeyListResultPage) APIKeyListResultIterator {
	return APIKeyListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (aklr APIKeyListResult) IsEmpty() bool {
	return aklr.Value == nil || len(*aklr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (aklr APIKeyListResult) hasNextLink() bool {
	return aklr.NextLink != nil && len(*aklr.NextLink) != 0
}

// aPIKeyListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (aklr APIKeyListResult) aPIKeyListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !aklr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(aklr.NextLink)))
}

// APIKeyListResultPage contains a page of APIKey values.
type APIKeyListResultPage struct {
	fn   func(context.Context, APIKeyListResult) (APIKeyListResult, error)
	aklr APIKeyListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *APIKeyListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeyListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.aklr)
		if err != nil {
			return err
		}
		page.aklr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *APIKeyListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page APIKeyListResultPage) NotDone() bool {
	return !page.aklr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page APIKeyListResultPage) Response() APIKeyListResult {
	return page.aklr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page APIKeyListResultPage) Values() []APIKey {
	if page.aklr.IsEmpty() {
		return nil
	}
	return *page.aklr.Value
}

// Creates a new instance of the APIKeyListResultPage type.
func NewAPIKeyListResultPage(getNextPage func(context.Context, APIKeyListResult) (APIKeyListResult, error)) APIKeyListResultPage {
	return APIKeyListResultPage{fn: getNextPage}
}

// CheckNameAvailabilityParameters parameters used for checking whether a resource name is available.
type CheckNameAvailabilityParameters struct {
	// Name - The name to check for availability.
	Name *string `json:"name,omitempty"`
	// Type - The resource type to check for name availability.
	Type *string `json:"type,omitempty"`
}

// ConfigurationStore the configuration store along with all resource properties. The Configuration Store will
// have all information to begin utilizing it.
type ConfigurationStore struct {
	autorest.Response `json:"-"`
	// Identity - The managed identity information, if configured.
	Identity *ResourceIdentity `json:"identity,omitempty"`
	// ConfigurationStoreProperties - The properties of a configuration store.
	*ConfigurationStoreProperties `json:"properties,omitempty"`
	// Sku - The sku of the configuration store.
	Sku *Sku `json:"sku,omitempty"`
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ConfigurationStore.
func (cs ConfigurationStore) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cs.Identity != nil {
		objectMap["identity"] = cs.Identity
	}
	if cs.ConfigurationStoreProperties != nil {
		objectMap["properties"] = cs.ConfigurationStoreProperties
	}
	if cs.Sku != nil {
		objectMap["sku"] = cs.Sku
	}
	if cs.Location != nil {
		objectMap["location"] = cs.Location
	}
	if cs.Tags != nil {
		objectMap["tags"] = cs.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ConfigurationStore struct.
func (cs *ConfigurationStore) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "identity":
			if v != nil {
				var identity ResourceIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				cs.Identity = &identity
			}
		case "properties":
			if v != nil {
				var configurationStoreProperties ConfigurationStoreProperties
				err = json.Unmarshal(*v, &configurationStoreProperties)
				if err != nil {
					return err
				}
				cs.ConfigurationStoreProperties = &configurationStoreProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				cs.Sku = &sku
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cs.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cs.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cs.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cs.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cs.Tags = tags
			}
		}
	}

	return nil
}

// ConfigurationStoreListResult the result of a request to list configuration stores.
type ConfigurationStoreListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection value.
	Value *[]ConfigurationStore `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ConfigurationStoreListResultIterator provides access to a complete listing of ConfigurationStore values.
type ConfigurationStoreListResultIterator struct {
	i    int
	page ConfigurationStoreListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ConfigurationStoreListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoreListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ConfigurationStoreListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ConfigurationStoreListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ConfigurationStoreListResultIterator) Response() ConfigurationStoreListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ConfigurationStoreListResultIterator) Value() ConfigurationStore {
	if !iter.page.NotDone() {
		return ConfigurationStore{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ConfigurationStoreListResultIterator type.
func NewConfigurationStoreListResultIterator(page ConfigurationStoreListResultPage) ConfigurationStoreListResultIterator {
	return ConfigurationStoreListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cslr ConfigurationStoreListResult) IsEmpty() bool {
	return cslr.Value == nil || len(*cslr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cslr ConfigurationStoreListResult) hasNextLink() bool {
	return cslr.NextLink != nil && len(*cslr.NextLink) != 0
}

// configurationStoreListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cslr ConfigurationStoreListResult) configurationStoreListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !cslr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cslr.NextLink)))
}

// ConfigurationStoreListResultPage contains a page of ConfigurationStore values.
type ConfigurationStoreListResultPage struct {
	fn   func(context.Context, ConfigurationStoreListResult) (ConfigurationStoreListResult, error)
	cslr ConfigurationStoreListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ConfigurationStoreListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConfigurationStoreListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.cslr)
		if err != nil {
			return err
		}
		page.cslr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ConfigurationStoreListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ConfigurationStoreListResultPage) NotDone() bool {
	return !page.cslr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ConfigurationStoreListResultPage) Response() ConfigurationStoreListResult {
	return page.cslr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ConfigurationStoreListResultPage) Values() []ConfigurationStore {
	if page.cslr.IsEmpty() {
		return nil
	}
	return *page.cslr.Value
}

// Creates a new instance of the ConfigurationStoreListResultPage type.
func NewConfigurationStoreListResultPage(getNextPage func(context.Context, ConfigurationStoreListResult) (ConfigurationStoreListResult, error)) ConfigurationStoreListResultPage {
	return ConfigurationStoreListResultPage{fn: getNextPage}
}

// ConfigurationStoreProperties the properties of a configuration store.
type ConfigurationStoreProperties struct {
	// ProvisioningState - READ-ONLY; The provisioning state of the configuration store. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// CreationDate - READ-ONLY; The creation date of configuration store.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// Endpoint - READ-ONLY; The DNS endpoint where the configuration store API will be available.
	Endpoint *string `json:"endpoint,omitempty"`
}

// ConfigurationStoresCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ConfigurationStoresCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ConfigurationStoresCreateFuture) Result(client ConfigurationStoresClient) (cs ConfigurationStore, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("appconfiguration.ConfigurationStoresCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cs.Response.Response, err = future.GetResult(sender); err == nil && cs.Response.Response.StatusCode != http.StatusNoContent {
		cs, err = client.CreateResponder(cs.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresCreateFuture", "Result", cs.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ConfigurationStoresDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ConfigurationStoresDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ConfigurationStoresDeleteFuture) Result(client ConfigurationStoresClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("appconfiguration.ConfigurationStoresDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// ConfigurationStoresUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ConfigurationStoresUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ConfigurationStoresUpdateFuture) Result(client ConfigurationStoresClient) (cs ConfigurationStore, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("appconfiguration.ConfigurationStoresUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cs.Response.Response, err = future.GetResult(sender); err == nil && cs.Response.Response.StatusCode != http.StatusNoContent {
		cs, err = client.UpdateResponder(cs.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "appconfiguration.ConfigurationStoresUpdateFuture", "Result", cs.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ConfigurationStoreUpdateParameters the parameters for updating a configuration store.
type ConfigurationStoreUpdateParameters struct {
	// Properties - The properties for updating a configuration store.
	Properties interface{} `json:"properties,omitempty"`
	// Identity - The managed identity information for the configuration store.
	Identity *ResourceIdentity `json:"identity,omitempty"`
	// Sku - The SKU of the configuration store.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - The ARM resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ConfigurationStoreUpdateParameters.
func (csup ConfigurationStoreUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if csup.Properties != nil {
		objectMap["properties"] = csup.Properties
	}
	if csup.Identity != nil {
		objectMap["identity"] = csup.Identity
	}
	if csup.Sku != nil {
		objectMap["sku"] = csup.Sku
	}
	if csup.Tags != nil {
		objectMap["tags"] = csup.Tags
	}
	return json.Marshal(objectMap)
}

// Error appConfiguration error object.
type Error struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
}

// KeyValue the result of a request to retrieve a key-value from the specified configuration store.
type KeyValue struct {
	autorest.Response `json:"-"`
	// Key - READ-ONLY; The primary identifier of a key-value.
	// The key is used in unison with the label to uniquely identify a key-value.
	Key *string `json:"key,omitempty"`
	// Label - READ-ONLY; A value used to group key-values.
	// The label is used in unison with the key to uniquely identify a key-value.
	Label *string `json:"label,omitempty"`
	// Value - READ-ONLY; The value of the key-value.
	Value *string `json:"value,omitempty"`
	// ContentType - READ-ONLY; The content type of the key-value's value.
	// Providing a proper content-type can enable transformations of values when they are retrieved by applications.
	ContentType *string `json:"contentType,omitempty"`
	// ETag - READ-ONLY; An ETag indicating the state of a key-value within a configuration store.
	ETag *string `json:"eTag,omitempty"`
	// LastModified - READ-ONLY; The last time a modifying operation was performed on the given key-value.
	LastModified *date.Time `json:"lastModified,omitempty"`
	// Locked - READ-ONLY; A value indicating whether the key-value is locked.
	// A locked key-value may not be modified until it is unlocked.
	Locked *bool `json:"locked,omitempty"`
	// Tags - READ-ONLY; A dictionary of tags that can help identify what a key-value may be applicable for.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for KeyValue.
func (kv KeyValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListKeyValueParameters the parameters used to list a configuration store key-value
type ListKeyValueParameters struct {
	// Key - The key to retrieve.
	Key *string `json:"key,omitempty"`
	// Label - The label of the key.
	Label *string `json:"label,omitempty"`
}

// GetKeyValueParameters the parameters used to list a configuration store key-value
type GetKeyValueParameters struct {
	autorest.Response `json:"-"`
	// Key - The key to retrieve.
	Key *string `json:"key,omitempty"`
	// Value - The key to retrieve.
	Value *string `json:"value,omitempty"`
	// Label - The label of the key.
	Label *string `json:"label,omitempty"`
}

// SetKeyValueParameters the parameters used to list a configuration store key-value
type SetKeyValueParameters struct {
	// Key - The key to create.
	Key *string `json:"key,omitempty"`
	// Value - The value of the key.
	Value *string `json:"value,omitempty"`
	// Label - The label of the key.
	Label *string `json:"label,omitempty"`
	// Secret - The secret of appconfig.
	Secret *string `json:"secret,omitempty"`
	// Credential - The credential of appconfig.
	Credential *string `json:"credential,omitempty"`
}

// NameAvailabilityStatus the result of a request to check the availability of a resource name.
type NameAvailabilityStatus struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; The value indicating whether the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Message - READ-ONLY; If any, the error message that provides more detail for the reason that the name is not available.
	Message *string `json:"message,omitempty"`
	// Reason - READ-ONLY; If any, the reason that the name is not available.
	Reason *string `json:"reason,omitempty"`
}

// OperationDefinition the definition of a configuration store operation.
type OperationDefinition struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The display information for the configuration store operation.
	Display *OperationDefinitionDisplay `json:"display,omitempty"`
}

// OperationDefinitionDisplay the display information for a configuration store operation.
type OperationDefinitionDisplay struct {
	// Provider - READ-ONLY; The resource provider name: Microsoft App Configuration."
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - The operation that users can perform.
	Operation *string `json:"operation,omitempty"`
	// Description - The description for the operation.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationDefinitionDisplay.
func (odd OperationDefinitionDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if odd.Resource != nil {
		objectMap["resource"] = odd.Resource
	}
	if odd.Operation != nil {
		objectMap["operation"] = odd.Operation
	}
	if odd.Description != nil {
		objectMap["description"] = odd.Description
	}
	return json.Marshal(objectMap)
}

// OperationDefinitionListResult the result of a request to list configuration store operations.
type OperationDefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection value.
	Value *[]OperationDefinition `json:"value,omitempty"`
	// NextLink - The URI that can be used to request the next set of paged results.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationDefinitionListResultIterator provides access to a complete listing of OperationDefinition values.
type OperationDefinitionListResultIterator struct {
	i    int
	page OperationDefinitionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationDefinitionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationDefinitionListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationDefinitionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationDefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationDefinitionListResultIterator) Response() OperationDefinitionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationDefinitionListResultIterator) Value() OperationDefinition {
	if !iter.page.NotDone() {
		return OperationDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationDefinitionListResultIterator type.
func NewOperationDefinitionListResultIterator(page OperationDefinitionListResultPage) OperationDefinitionListResultIterator {
	return OperationDefinitionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (odlr OperationDefinitionListResult) IsEmpty() bool {
	return odlr.Value == nil || len(*odlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (odlr OperationDefinitionListResult) hasNextLink() bool {
	return odlr.NextLink != nil && len(*odlr.NextLink) != 0
}

// operationDefinitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (odlr OperationDefinitionListResult) operationDefinitionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !odlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(odlr.NextLink)))
}

// OperationDefinitionListResultPage contains a page of OperationDefinition values.
type OperationDefinitionListResultPage struct {
	fn   func(context.Context, OperationDefinitionListResult) (OperationDefinitionListResult, error)
	odlr OperationDefinitionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationDefinitionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationDefinitionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.odlr)
		if err != nil {
			return err
		}
		page.odlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationDefinitionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationDefinitionListResultPage) NotDone() bool {
	return !page.odlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationDefinitionListResultPage) Response() OperationDefinitionListResult {
	return page.odlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationDefinitionListResultPage) Values() []OperationDefinition {
	if page.odlr.IsEmpty() {
		return nil
	}
	return *page.odlr.Value
}

// Creates a new instance of the OperationDefinitionListResultPage type.
func NewOperationDefinitionListResultPage(getNextPage func(context.Context, OperationDefinitionListResult) (OperationDefinitionListResult, error)) OperationDefinitionListResultPage {
	return OperationDefinitionListResultPage{fn: getNextPage}
}

// RegenerateKeyParameters the parameters used to regenerate an API key.
type RegenerateKeyParameters struct {
	// ID - The id of the key to regenerate.
	ID *string `json:"id,omitempty"`
}

// Resource an Azure resource.
type Resource struct {
	// ID - READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - The location of the resource. This cannot be changed after the resource is created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceIdentity ...
type ResourceIdentity struct {
	// Type - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities. Possible values include: 'None', 'SystemAssigned', 'UserAssigned', 'SystemAssignedUserAssigned'
	Type IdentityType `json:"type,omitempty"`
	// UserAssignedIdentities - The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*UserIdentity `json:"userAssignedIdentities"`
	// PrincipalID - READ-ONLY; The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant id associated with the resource's identity. This property will only be provided for a system-assigned identity.
	TenantID *string `json:"tenantId,omitempty"`
}

// MarshalJSON is the custom marshaler for ResourceIdentity.
func (ri ResourceIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ri.Type != "" {
		objectMap["type"] = ri.Type
	}
	if ri.UserAssignedIdentities != nil {
		objectMap["userAssignedIdentities"] = ri.UserAssignedIdentities
	}
	return json.Marshal(objectMap)
}

// Sku describes a configuration store SKU.
type Sku struct {
	// Name - The SKU name of the configuration store.
	Name *string `json:"name,omitempty"`
}

// UserIdentity ...
type UserIdentity struct {
	// PrincipalID - READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// ClientID - READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string `json:"clientId,omitempty"`
}
