# ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockSpec**](ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockSpec.md) |  | [optional] 
**Status** | Pointer to [**ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockStatus**](ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockStatus.md) |  | [optional] 

## Methods

### NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock

`func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock() *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock`

NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockWithDefaults

`func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock`

NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetSpec() ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetSpecOk() (*ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) SetSpec(v ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetStatus() ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) GetStatusOk() (*ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) SetStatus(v ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClockStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1TestClock) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


