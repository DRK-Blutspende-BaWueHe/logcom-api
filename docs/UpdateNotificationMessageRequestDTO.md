# UpdateNotificationMessageRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the message (only the \&quot;SEEN\&quot; status is supported) | [optional] 

## Methods

### NewUpdateNotificationMessageRequestDTO

`func NewUpdateNotificationMessageRequestDTO() *UpdateNotificationMessageRequestDTO`

NewUpdateNotificationMessageRequestDTO instantiates a new UpdateNotificationMessageRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationMessageRequestDTOWithDefaults

`func NewUpdateNotificationMessageRequestDTOWithDefaults() *UpdateNotificationMessageRequestDTO`

NewUpdateNotificationMessageRequestDTOWithDefaults instantiates a new UpdateNotificationMessageRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateNotificationMessageRequestDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNotificationMessageRequestDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNotificationMessageRequestDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateNotificationMessageRequestDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


