# ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer**](ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ListMeta**](V1ListMeta.md) |  | [optional] 

## Methods

### NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList

`func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList(items []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer, ) *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList`

NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferListWithDefaults

`func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferListWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList`

NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferListWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetItems() []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetItemsOk() (*[]ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) SetItems(v []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOffer)`

SetItems sets Items field to given value.


### GetKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetMetadata() V1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) GetMetadataOk() (*V1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) SetMetadata(v V1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PrivateOfferList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


