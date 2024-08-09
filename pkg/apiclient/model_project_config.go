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

// checks if the ProjectConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectConfig{}

// ProjectConfig struct for ProjectConfig
type ProjectConfig struct {
	BuildConfig *ProjectBuildConfig `json:"buildConfig,omitempty"`
	Default *bool `json:"default,omitempty"`
	EnvVars *map[string]string `json:"envVars,omitempty"`
	Image *string `json:"image,omitempty"`
	Name *string `json:"name,omitempty"`
	Repository *GitRepository `json:"repository,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewProjectConfig instantiates a new ProjectConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectConfig() *ProjectConfig {
	this := ProjectConfig{}
	return &this
}

// NewProjectConfigWithDefaults instantiates a new ProjectConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectConfigWithDefaults() *ProjectConfig {
	this := ProjectConfig{}
	return &this
}

// GetBuildConfig returns the BuildConfig field value if set, zero value otherwise.
func (o *ProjectConfig) GetBuildConfig() ProjectBuildConfig {
	if o == nil || IsNil(o.BuildConfig) {
		var ret ProjectBuildConfig
		return ret
	}
	return *o.BuildConfig
}

// GetBuildConfigOk returns a tuple with the BuildConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetBuildConfigOk() (*ProjectBuildConfig, bool) {
	if o == nil || IsNil(o.BuildConfig) {
		return nil, false
	}
	return o.BuildConfig, true
}

// HasBuildConfig returns a boolean if a field has been set.
func (o *ProjectConfig) HasBuildConfig() bool {
	if o != nil && !IsNil(o.BuildConfig) {
		return true
	}

	return false
}

// SetBuildConfig gets a reference to the given ProjectBuildConfig and assigns it to the BuildConfig field.
func (o *ProjectConfig) SetBuildConfig(v ProjectBuildConfig) {
	o.BuildConfig = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ProjectConfig) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ProjectConfig) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *ProjectConfig) SetDefault(v bool) {
	o.Default = &v
}

// GetEnvVars returns the EnvVars field value if set, zero value otherwise.
func (o *ProjectConfig) GetEnvVars() map[string]string {
	if o == nil || IsNil(o.EnvVars) {
		var ret map[string]string
		return ret
	}
	return *o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvVars) {
		return nil, false
	}
	return o.EnvVars, true
}

// HasEnvVars returns a boolean if a field has been set.
func (o *ProjectConfig) HasEnvVars() bool {
	if o != nil && !IsNil(o.EnvVars) {
		return true
	}

	return false
}

// SetEnvVars gets a reference to the given map[string]string and assigns it to the EnvVars field.
func (o *ProjectConfig) SetEnvVars(v map[string]string) {
	o.EnvVars = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ProjectConfig) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ProjectConfig) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ProjectConfig) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectConfig) SetName(v string) {
	o.Name = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ProjectConfig) GetRepository() GitRepository {
	if o == nil || IsNil(o.Repository) {
		var ret GitRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetRepositoryOk() (*GitRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ProjectConfig) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GitRepository and assigns it to the Repository field.
func (o *ProjectConfig) SetRepository(v GitRepository) {
	o.Repository = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ProjectConfig) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectConfig) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ProjectConfig) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *ProjectConfig) SetUser(v string) {
	o.User = &v
}

func (o ProjectConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildConfig) {
		toSerialize["buildConfig"] = o.BuildConfig
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.EnvVars) {
		toSerialize["envVars"] = o.EnvVars
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableProjectConfig struct {
	value *ProjectConfig
	isSet bool
}

func (v NullableProjectConfig) Get() *ProjectConfig {
	return v.value
}

func (v *NullableProjectConfig) Set(val *ProjectConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectConfig(val *ProjectConfig) *NullableProjectConfig {
	return &NullableProjectConfig{value: val, isSet: true}
}

func (v NullableProjectConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


