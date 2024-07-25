/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ServerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConfig{}

// ServerConfig struct for ServerConfig
type ServerConfig struct {
	ApiPort                   *int32      `json:"apiPort,omitempty"`
	BinariesPath              *string     `json:"binariesPath,omitempty"`
	BuildImageNamespace       *string     `json:"buildImageNamespace,omitempty"`
	BuilderImage              *string     `json:"builderImage,omitempty"`
	BuilderRegistryServer     *string     `json:"builderRegistryServer,omitempty"`
	DefaultProjectImage       *string     `json:"defaultProjectImage,omitempty"`
	DefaultProjectUser        *string     `json:"defaultProjectUser,omitempty"`
	Frps                      *FRPSConfig `json:"frps,omitempty"`
	HeadscalePort             *int32      `json:"headscalePort,omitempty"`
	Id                        *string     `json:"id,omitempty"`
	LocalBuilderRegistryImage *string     `json:"localBuilderRegistryImage,omitempty"`
	LocalBuilderRegistryPort  *int32      `json:"localBuilderRegistryPort,omitempty"`
	LogFilePath               *string     `json:"logFilePath,omitempty"`
	ProvidersDir              *string     `json:"providersDir,omitempty"`
	RegistryUrl               *string     `json:"registryUrl,omitempty"`
	ServerDownloadUrl         *string     `json:"serverDownloadUrl,omitempty"`
}

// NewServerConfig instantiates a new ServerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfig() *ServerConfig {
	this := ServerConfig{}
	return &this
}

// NewServerConfigWithDefaults instantiates a new ServerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigWithDefaults() *ServerConfig {
	this := ServerConfig{}
	return &this
}

// GetApiPort returns the ApiPort field value if set, zero value otherwise.
func (o *ServerConfig) GetApiPort() int32 {
	if o == nil || IsNil(o.ApiPort) {
		var ret int32
		return ret
	}
	return *o.ApiPort
}

// GetApiPortOk returns a tuple with the ApiPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetApiPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiPort) {
		return nil, false
	}
	return o.ApiPort, true
}

// HasApiPort returns a boolean if a field has been set.
func (o *ServerConfig) HasApiPort() bool {
	if o != nil && !IsNil(o.ApiPort) {
		return true
	}

	return false
}

// SetApiPort gets a reference to the given int32 and assigns it to the ApiPort field.
func (o *ServerConfig) SetApiPort(v int32) {
	o.ApiPort = &v
}

// GetBinariesPath returns the BinariesPath field value if set, zero value otherwise.
func (o *ServerConfig) GetBinariesPath() string {
	if o == nil || IsNil(o.BinariesPath) {
		var ret string
		return ret
	}
	return *o.BinariesPath
}

// GetBinariesPathOk returns a tuple with the BinariesPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetBinariesPathOk() (*string, bool) {
	if o == nil || IsNil(o.BinariesPath) {
		return nil, false
	}
	return o.BinariesPath, true
}

// HasBinariesPath returns a boolean if a field has been set.
func (o *ServerConfig) HasBinariesPath() bool {
	if o != nil && !IsNil(o.BinariesPath) {
		return true
	}

	return false
}

// SetBinariesPath gets a reference to the given string and assigns it to the BinariesPath field.
func (o *ServerConfig) SetBinariesPath(v string) {
	o.BinariesPath = &v
}

// GetBuildImageNamespace returns the BuildImageNamespace field value if set, zero value otherwise.
func (o *ServerConfig) GetBuildImageNamespace() string {
	if o == nil || IsNil(o.BuildImageNamespace) {
		var ret string
		return ret
	}
	return *o.BuildImageNamespace
}

// GetBuildImageNamespaceOk returns a tuple with the BuildImageNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetBuildImageNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.BuildImageNamespace) {
		return nil, false
	}
	return o.BuildImageNamespace, true
}

// HasBuildImageNamespace returns a boolean if a field has been set.
func (o *ServerConfig) HasBuildImageNamespace() bool {
	if o != nil && !IsNil(o.BuildImageNamespace) {
		return true
	}

	return false
}

// SetBuildImageNamespace gets a reference to the given string and assigns it to the BuildImageNamespace field.
func (o *ServerConfig) SetBuildImageNamespace(v string) {
	o.BuildImageNamespace = &v
}

// GetBuilderImage returns the BuilderImage field value if set, zero value otherwise.
func (o *ServerConfig) GetBuilderImage() string {
	if o == nil || IsNil(o.BuilderImage) {
		var ret string
		return ret
	}
	return *o.BuilderImage
}

// GetBuilderImageOk returns a tuple with the BuilderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetBuilderImageOk() (*string, bool) {
	if o == nil || IsNil(o.BuilderImage) {
		return nil, false
	}
	return o.BuilderImage, true
}

// HasBuilderImage returns a boolean if a field has been set.
func (o *ServerConfig) HasBuilderImage() bool {
	if o != nil && !IsNil(o.BuilderImage) {
		return true
	}

	return false
}

// SetBuilderImage gets a reference to the given string and assigns it to the BuilderImage field.
func (o *ServerConfig) SetBuilderImage(v string) {
	o.BuilderImage = &v
}

// GetBuilderRegistryServer returns the BuilderRegistryServer field value if set, zero value otherwise.
func (o *ServerConfig) GetBuilderRegistryServer() string {
	if o == nil || IsNil(o.BuilderRegistryServer) {
		var ret string
		return ret
	}
	return *o.BuilderRegistryServer
}

// GetBuilderRegistryServerOk returns a tuple with the BuilderRegistryServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetBuilderRegistryServerOk() (*string, bool) {
	if o == nil || IsNil(o.BuilderRegistryServer) {
		return nil, false
	}
	return o.BuilderRegistryServer, true
}

// HasBuilderRegistryServer returns a boolean if a field has been set.
func (o *ServerConfig) HasBuilderRegistryServer() bool {
	if o != nil && !IsNil(o.BuilderRegistryServer) {
		return true
	}

	return false
}

// SetBuilderRegistryServer gets a reference to the given string and assigns it to the BuilderRegistryServer field.
func (o *ServerConfig) SetBuilderRegistryServer(v string) {
	o.BuilderRegistryServer = &v
}

// GetDefaultProjectImage returns the DefaultProjectImage field value if set, zero value otherwise.
func (o *ServerConfig) GetDefaultProjectImage() string {
	if o == nil || IsNil(o.DefaultProjectImage) {
		var ret string
		return ret
	}
	return *o.DefaultProjectImage
}

// GetDefaultProjectImageOk returns a tuple with the DefaultProjectImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetDefaultProjectImageOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultProjectImage) {
		return nil, false
	}
	return o.DefaultProjectImage, true
}

// HasDefaultProjectImage returns a boolean if a field has been set.
func (o *ServerConfig) HasDefaultProjectImage() bool {
	if o != nil && !IsNil(o.DefaultProjectImage) {
		return true
	}

	return false
}

// SetDefaultProjectImage gets a reference to the given string and assigns it to the DefaultProjectImage field.
func (o *ServerConfig) SetDefaultProjectImage(v string) {
	o.DefaultProjectImage = &v
}

// GetDefaultProjectUser returns the DefaultProjectUser field value if set, zero value otherwise.
func (o *ServerConfig) GetDefaultProjectUser() string {
	if o == nil || IsNil(o.DefaultProjectUser) {
		var ret string
		return ret
	}
	return *o.DefaultProjectUser
}

// GetDefaultProjectUserOk returns a tuple with the DefaultProjectUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetDefaultProjectUserOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultProjectUser) {
		return nil, false
	}
	return o.DefaultProjectUser, true
}

// HasDefaultProjectUser returns a boolean if a field has been set.
func (o *ServerConfig) HasDefaultProjectUser() bool {
	if o != nil && !IsNil(o.DefaultProjectUser) {
		return true
	}

	return false
}

// SetDefaultProjectUser gets a reference to the given string and assigns it to the DefaultProjectUser field.
func (o *ServerConfig) SetDefaultProjectUser(v string) {
	o.DefaultProjectUser = &v
}

// GetFrps returns the Frps field value if set, zero value otherwise.
func (o *ServerConfig) GetFrps() FRPSConfig {
	if o == nil || IsNil(o.Frps) {
		var ret FRPSConfig
		return ret
	}
	return *o.Frps
}

// GetFrpsOk returns a tuple with the Frps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetFrpsOk() (*FRPSConfig, bool) {
	if o == nil || IsNil(o.Frps) {
		return nil, false
	}
	return o.Frps, true
}

// HasFrps returns a boolean if a field has been set.
func (o *ServerConfig) HasFrps() bool {
	if o != nil && !IsNil(o.Frps) {
		return true
	}

	return false
}

// SetFrps gets a reference to the given FRPSConfig and assigns it to the Frps field.
func (o *ServerConfig) SetFrps(v FRPSConfig) {
	o.Frps = &v
}

// GetHeadscalePort returns the HeadscalePort field value if set, zero value otherwise.
func (o *ServerConfig) GetHeadscalePort() int32 {
	if o == nil || IsNil(o.HeadscalePort) {
		var ret int32
		return ret
	}
	return *o.HeadscalePort
}

// GetHeadscalePortOk returns a tuple with the HeadscalePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetHeadscalePortOk() (*int32, bool) {
	if o == nil || IsNil(o.HeadscalePort) {
		return nil, false
	}
	return o.HeadscalePort, true
}

// HasHeadscalePort returns a boolean if a field has been set.
func (o *ServerConfig) HasHeadscalePort() bool {
	if o != nil && !IsNil(o.HeadscalePort) {
		return true
	}

	return false
}

// SetHeadscalePort gets a reference to the given int32 and assigns it to the HeadscalePort field.
func (o *ServerConfig) SetHeadscalePort(v int32) {
	o.HeadscalePort = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerConfig) SetId(v string) {
	o.Id = &v
}

// GetLocalBuilderRegistryImage returns the LocalBuilderRegistryImage field value if set, zero value otherwise.
func (o *ServerConfig) GetLocalBuilderRegistryImage() string {
	if o == nil || IsNil(o.LocalBuilderRegistryImage) {
		var ret string
		return ret
	}
	return *o.LocalBuilderRegistryImage
}

// GetLocalBuilderRegistryImageOk returns a tuple with the LocalBuilderRegistryImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetLocalBuilderRegistryImageOk() (*string, bool) {
	if o == nil || IsNil(o.LocalBuilderRegistryImage) {
		return nil, false
	}
	return o.LocalBuilderRegistryImage, true
}

// HasLocalBuilderRegistryImage returns a boolean if a field has been set.
func (o *ServerConfig) HasLocalBuilderRegistryImage() bool {
	if o != nil && !IsNil(o.LocalBuilderRegistryImage) {
		return true
	}

	return false
}

// SetLocalBuilderRegistryImage gets a reference to the given string and assigns it to the LocalBuilderRegistryImage field.
func (o *ServerConfig) SetLocalBuilderRegistryImage(v string) {
	o.LocalBuilderRegistryImage = &v
}

// GetLocalBuilderRegistryPort returns the LocalBuilderRegistryPort field value if set, zero value otherwise.
func (o *ServerConfig) GetLocalBuilderRegistryPort() int32 {
	if o == nil || IsNil(o.LocalBuilderRegistryPort) {
		var ret int32
		return ret
	}
	return *o.LocalBuilderRegistryPort
}

// GetLocalBuilderRegistryPortOk returns a tuple with the LocalBuilderRegistryPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetLocalBuilderRegistryPortOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalBuilderRegistryPort) {
		return nil, false
	}
	return o.LocalBuilderRegistryPort, true
}

// HasLocalBuilderRegistryPort returns a boolean if a field has been set.
func (o *ServerConfig) HasLocalBuilderRegistryPort() bool {
	if o != nil && !IsNil(o.LocalBuilderRegistryPort) {
		return true
	}

	return false
}

// SetLocalBuilderRegistryPort gets a reference to the given int32 and assigns it to the LocalBuilderRegistryPort field.
func (o *ServerConfig) SetLocalBuilderRegistryPort(v int32) {
	o.LocalBuilderRegistryPort = &v
}

// GetLogFilePath returns the LogFilePath field value if set, zero value otherwise.
func (o *ServerConfig) GetLogFilePath() string {
	if o == nil || IsNil(o.LogFilePath) {
		var ret string
		return ret
	}
	return *o.LogFilePath
}

// GetLogFilePathOk returns a tuple with the LogFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetLogFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.LogFilePath) {
		return nil, false
	}
	return o.LogFilePath, true
}

// HasLogFilePath returns a boolean if a field has been set.
func (o *ServerConfig) HasLogFilePath() bool {
	if o != nil && !IsNil(o.LogFilePath) {
		return true
	}

	return false
}

// SetLogFilePath gets a reference to the given string and assigns it to the LogFilePath field.
func (o *ServerConfig) SetLogFilePath(v string) {
	o.LogFilePath = &v
}

// GetProvidersDir returns the ProvidersDir field value if set, zero value otherwise.
func (o *ServerConfig) GetProvidersDir() string {
	if o == nil || IsNil(o.ProvidersDir) {
		var ret string
		return ret
	}
	return *o.ProvidersDir
}

// GetProvidersDirOk returns a tuple with the ProvidersDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetProvidersDirOk() (*string, bool) {
	if o == nil || IsNil(o.ProvidersDir) {
		return nil, false
	}
	return o.ProvidersDir, true
}

// HasProvidersDir returns a boolean if a field has been set.
func (o *ServerConfig) HasProvidersDir() bool {
	if o != nil && !IsNil(o.ProvidersDir) {
		return true
	}

	return false
}

// SetProvidersDir gets a reference to the given string and assigns it to the ProvidersDir field.
func (o *ServerConfig) SetProvidersDir(v string) {
	o.ProvidersDir = &v
}

// GetRegistryUrl returns the RegistryUrl field value if set, zero value otherwise.
func (o *ServerConfig) GetRegistryUrl() string {
	if o == nil || IsNil(o.RegistryUrl) {
		var ret string
		return ret
	}
	return *o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetRegistryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryUrl) {
		return nil, false
	}
	return o.RegistryUrl, true
}

// HasRegistryUrl returns a boolean if a field has been set.
func (o *ServerConfig) HasRegistryUrl() bool {
	if o != nil && !IsNil(o.RegistryUrl) {
		return true
	}

	return false
}

// SetRegistryUrl gets a reference to the given string and assigns it to the RegistryUrl field.
func (o *ServerConfig) SetRegistryUrl(v string) {
	o.RegistryUrl = &v
}

// GetServerDownloadUrl returns the ServerDownloadUrl field value if set, zero value otherwise.
func (o *ServerConfig) GetServerDownloadUrl() string {
	if o == nil || IsNil(o.ServerDownloadUrl) {
		var ret string
		return ret
	}
	return *o.ServerDownloadUrl
}

// GetServerDownloadUrlOk returns a tuple with the ServerDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfig) GetServerDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServerDownloadUrl) {
		return nil, false
	}
	return o.ServerDownloadUrl, true
}

// HasServerDownloadUrl returns a boolean if a field has been set.
func (o *ServerConfig) HasServerDownloadUrl() bool {
	if o != nil && !IsNil(o.ServerDownloadUrl) {
		return true
	}

	return false
}

// SetServerDownloadUrl gets a reference to the given string and assigns it to the ServerDownloadUrl field.
func (o *ServerConfig) SetServerDownloadUrl(v string) {
	o.ServerDownloadUrl = &v
}

func (o ServerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiPort) {
		toSerialize["apiPort"] = o.ApiPort
	}
	if !IsNil(o.BinariesPath) {
		toSerialize["binariesPath"] = o.BinariesPath
	}
	if !IsNil(o.BuildImageNamespace) {
		toSerialize["buildImageNamespace"] = o.BuildImageNamespace
	}
	if !IsNil(o.BuilderImage) {
		toSerialize["builderImage"] = o.BuilderImage
	}
	if !IsNil(o.BuilderRegistryServer) {
		toSerialize["builderRegistryServer"] = o.BuilderRegistryServer
	}
	if !IsNil(o.DefaultProjectImage) {
		toSerialize["defaultProjectImage"] = o.DefaultProjectImage
	}
	if !IsNil(o.DefaultProjectUser) {
		toSerialize["defaultProjectUser"] = o.DefaultProjectUser
	}
	if !IsNil(o.Frps) {
		toSerialize["frps"] = o.Frps
	}
	if !IsNil(o.HeadscalePort) {
		toSerialize["headscalePort"] = o.HeadscalePort
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LocalBuilderRegistryImage) {
		toSerialize["localBuilderRegistryImage"] = o.LocalBuilderRegistryImage
	}
	if !IsNil(o.LocalBuilderRegistryPort) {
		toSerialize["localBuilderRegistryPort"] = o.LocalBuilderRegistryPort
	}
	if !IsNil(o.LogFilePath) {
		toSerialize["logFilePath"] = o.LogFilePath
	}
	if !IsNil(o.ProvidersDir) {
		toSerialize["providersDir"] = o.ProvidersDir
	}
	if !IsNil(o.RegistryUrl) {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
	if !IsNil(o.ServerDownloadUrl) {
		toSerialize["serverDownloadUrl"] = o.ServerDownloadUrl
	}
	return toSerialize, nil
}

type NullableServerConfig struct {
	value *ServerConfig
	isSet bool
}

func (v NullableServerConfig) Get() *ServerConfig {
	return v.value
}

func (v *NullableServerConfig) Set(val *ServerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfig(val *ServerConfig) *NullableServerConfig {
	return &NullableServerConfig{value: val, isSet: true}
}

func (v NullableServerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
