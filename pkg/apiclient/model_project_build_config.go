/*
Daytona Server API

Daytona Server API

API version: 0.24.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ProjectBuildConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectBuildConfig{}

// ProjectBuildConfig struct for ProjectBuildConfig
type ProjectBuildConfig struct {
	Devcontainer *DevcontainerConfig `json:"devcontainer,omitempty"`
}

// NewProjectBuildConfig instantiates a new ProjectBuildConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectBuildConfig() *ProjectBuildConfig {
	this := ProjectBuildConfig{}
	return &this
}

// NewProjectBuildConfigWithDefaults instantiates a new ProjectBuildConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectBuildConfigWithDefaults() *ProjectBuildConfig {
	this := ProjectBuildConfig{}
	return &this
}

// GetDevcontainer returns the Devcontainer field value if set, zero value otherwise.
func (o *ProjectBuildConfig) GetDevcontainer() DevcontainerConfig {
	if o == nil || IsNil(o.Devcontainer) {
		var ret DevcontainerConfig
		return ret
	}
	return *o.Devcontainer
}

// GetDevcontainerOk returns a tuple with the Devcontainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectBuildConfig) GetDevcontainerOk() (*DevcontainerConfig, bool) {
	if o == nil || IsNil(o.Devcontainer) {
		return nil, false
	}
	return o.Devcontainer, true
}

// HasDevcontainer returns a boolean if a field has been set.
func (o *ProjectBuildConfig) HasDevcontainer() bool {
	if o != nil && !IsNil(o.Devcontainer) {
		return true
	}

	return false
}

// SetDevcontainer gets a reference to the given DevcontainerConfig and assigns it to the Devcontainer field.
func (o *ProjectBuildConfig) SetDevcontainer(v DevcontainerConfig) {
	o.Devcontainer = &v
}

func (o ProjectBuildConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectBuildConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devcontainer) {
		toSerialize["devcontainer"] = o.Devcontainer
	}
	return toSerialize, nil
}

type NullableProjectBuildConfig struct {
	value *ProjectBuildConfig
	isSet bool
}

func (v NullableProjectBuildConfig) Get() *ProjectBuildConfig {
	return v.value
}

func (v *NullableProjectBuildConfig) Set(val *ProjectBuildConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectBuildConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectBuildConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectBuildConfig(val *ProjectBuildConfig) *NullableProjectBuildConfig {
	return &NullableProjectBuildConfig{value: val, isSet: true}
}

func (v NullableProjectBuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectBuildConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


