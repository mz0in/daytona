/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverapiclient

import (
	"encoding/json"
)

// checks if the GitProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitProvider{}

// GitProvider struct for GitProvider
type GitProvider struct {
	BaseApiUrl *string `json:"baseApiUrl,omitempty"`
	Id         *string `json:"id,omitempty"`
	Token      *string `json:"token,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// NewGitProvider instantiates a new GitProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitProvider() *GitProvider {
	this := GitProvider{}
	return &this
}

// NewGitProviderWithDefaults instantiates a new GitProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitProviderWithDefaults() *GitProvider {
	this := GitProvider{}
	return &this
}

// GetBaseApiUrl returns the BaseApiUrl field value if set, zero value otherwise.
func (o *GitProvider) GetBaseApiUrl() string {
	if o == nil || IsNil(o.BaseApiUrl) {
		var ret string
		return ret
	}
	return *o.BaseApiUrl
}

// GetBaseApiUrlOk returns a tuple with the BaseApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitProvider) GetBaseApiUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseApiUrl) {
		return nil, false
	}
	return o.BaseApiUrl, true
}

// HasBaseApiUrl returns a boolean if a field has been set.
func (o *GitProvider) HasBaseApiUrl() bool {
	if o != nil && !IsNil(o.BaseApiUrl) {
		return true
	}

	return false
}

// SetBaseApiUrl gets a reference to the given string and assigns it to the BaseApiUrl field.
func (o *GitProvider) SetBaseApiUrl(v string) {
	o.BaseApiUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GitProvider) SetId(v string) {
	o.Id = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GitProvider) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitProvider) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GitProvider) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GitProvider) SetToken(v string) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GitProvider) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitProvider) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GitProvider) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GitProvider) SetUsername(v string) {
	o.Username = &v
}

func (o GitProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseApiUrl) {
		toSerialize["baseApiUrl"] = o.BaseApiUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableGitProvider struct {
	value *GitProvider
	isSet bool
}

func (v NullableGitProvider) Get() *GitProvider {
	return v.value
}

func (v *NullableGitProvider) Set(val *GitProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableGitProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableGitProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitProvider(val *GitProvider) *NullableGitProvider {
	return &NullableGitProvider{value: val, isSet: true}
}

func (v NullableGitProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
