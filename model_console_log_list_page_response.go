/*
LogCom API

LogCom Swagger documentation

API version: 1.2.8
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"encoding/json"
)

// ConsoleLogListPageResponse struct for ConsoleLogListPageResponse
type ConsoleLogListPageResponse struct {
	// The actual page number
	CurrentPage *int32 `json:"currentPage,omitempty"`
	// The items
	Items []ConsoleLogDTO `json:"items,omitempty"`
	// The number of items per page
	PageSize *int32 `json:"pageSize,omitempty"`
	// The total count of items
	TotalCount *int32 `json:"totalCount,omitempty"`
	// The total pages
	TotalPages *int32 `json:"totalPages,omitempty"`
}

// NewConsoleLogListPageResponse instantiates a new ConsoleLogListPageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleLogListPageResponse() *ConsoleLogListPageResponse {
	this := ConsoleLogListPageResponse{}
	return &this
}

// NewConsoleLogListPageResponseWithDefaults instantiates a new ConsoleLogListPageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleLogListPageResponseWithDefaults() *ConsoleLogListPageResponse {
	this := ConsoleLogListPageResponse{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *ConsoleLogListPageResponse) GetCurrentPage() int32 {
	if o == nil || isNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogListPageResponse) GetCurrentPageOk() (*int32, bool) {
	if o == nil || isNil(o.CurrentPage) {
    return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *ConsoleLogListPageResponse) HasCurrentPage() bool {
	if o != nil && !isNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *ConsoleLogListPageResponse) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConsoleLogListPageResponse) GetItems() []ConsoleLogDTO {
	if o == nil || isNil(o.Items) {
		var ret []ConsoleLogDTO
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogListPageResponse) GetItemsOk() ([]ConsoleLogDTO, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConsoleLogListPageResponse) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ConsoleLogDTO and assigns it to the Items field.
func (o *ConsoleLogListPageResponse) SetItems(v []ConsoleLogDTO) {
	o.Items = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ConsoleLogListPageResponse) GetPageSize() int32 {
	if o == nil || isNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogListPageResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ConsoleLogListPageResponse) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ConsoleLogListPageResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ConsoleLogListPageResponse) GetTotalCount() int32 {
	if o == nil || isNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogListPageResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCount) {
    return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ConsoleLogListPageResponse) HasTotalCount() bool {
	if o != nil && !isNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ConsoleLogListPageResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *ConsoleLogListPageResponse) GetTotalPages() int32 {
	if o == nil || isNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleLogListPageResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *ConsoleLogListPageResponse) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *ConsoleLogListPageResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

func (o ConsoleLogListPageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !isNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	return json.Marshal(toSerialize)
}

type NullableConsoleLogListPageResponse struct {
	value *ConsoleLogListPageResponse
	isSet bool
}

func (v NullableConsoleLogListPageResponse) Get() *ConsoleLogListPageResponse {
	return v.value
}

func (v *NullableConsoleLogListPageResponse) Set(val *ConsoleLogListPageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleLogListPageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleLogListPageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleLogListPageResponse(val *ConsoleLogListPageResponse) *NullableConsoleLogListPageResponse {
	return &NullableConsoleLogListPageResponse{value: val, isSet: true}
}

func (v NullableConsoleLogListPageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleLogListPageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


