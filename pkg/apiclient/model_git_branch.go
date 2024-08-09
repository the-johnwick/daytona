/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the GitBranch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitBranch{}

// GitBranch struct for GitBranch
type GitBranch struct {
	Name *string `json:"name,omitempty"`
	Sha *string `json:"sha,omitempty"`
}

// NewGitBranch instantiates a new GitBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitBranch() *GitBranch {
	this := GitBranch{}
	return &this
}

// NewGitBranchWithDefaults instantiates a new GitBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitBranchWithDefaults() *GitBranch {
	this := GitBranch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GitBranch) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitBranch) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GitBranch) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GitBranch) SetName(v string) {
	o.Name = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitBranch) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitBranch) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitBranch) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitBranch) SetSha(v string) {
	o.Sha = &v
}

func (o GitBranch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitBranch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	return toSerialize, nil
}

type NullableGitBranch struct {
	value *GitBranch
	isSet bool
}

func (v NullableGitBranch) Get() *GitBranch {
	return v.value
}

func (v *NullableGitBranch) Set(val *GitBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableGitBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableGitBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitBranch(val *GitBranch) *NullableGitBranch {
	return &NullableGitBranch{value: val, isSet: true}
}

func (v NullableGitBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


