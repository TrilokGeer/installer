// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/compute/mgmt/compute"
	network "github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/network/mgmt/network"
	resources "github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/resources/mgmt/resources"
	subscriptions "github.com/Azure/azure-sdk-for-go/profiles/2018-03-01/resources/mgmt/subscriptions"
	compute0 "github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// AreMarketplaceImageTermsAccepted mocks base method.
func (m *MockAPI) AreMarketplaceImageTermsAccepted(ctx context.Context, publisher, offer, sku string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreMarketplaceImageTermsAccepted", ctx, publisher, offer, sku)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AreMarketplaceImageTermsAccepted indicates an expected call of AreMarketplaceImageTermsAccepted.
func (mr *MockAPIMockRecorder) AreMarketplaceImageTermsAccepted(ctx, publisher, offer, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreMarketplaceImageTermsAccepted", reflect.TypeOf((*MockAPI)(nil).AreMarketplaceImageTermsAccepted), ctx, publisher, offer, sku)
}

// GetAvailabilityZones mocks base method.
func (m *MockAPI) GetAvailabilityZones(ctx context.Context, region, instanceType string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilityZones", ctx, region, instanceType)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailabilityZones indicates an expected call of GetAvailabilityZones.
func (mr *MockAPIMockRecorder) GetAvailabilityZones(ctx, region, instanceType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilityZones", reflect.TypeOf((*MockAPI)(nil).GetAvailabilityZones), ctx, region, instanceType)
}

// GetComputeSubnet mocks base method.
func (m *MockAPI) GetComputeSubnet(ctx context.Context, resourceGroupName, virtualNetwork, subnet string) (*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComputeSubnet", ctx, resourceGroupName, virtualNetwork, subnet)
	ret0, _ := ret[0].(*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComputeSubnet indicates an expected call of GetComputeSubnet.
func (mr *MockAPIMockRecorder) GetComputeSubnet(ctx, resourceGroupName, virtualNetwork, subnet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComputeSubnet", reflect.TypeOf((*MockAPI)(nil).GetComputeSubnet), ctx, resourceGroupName, virtualNetwork, subnet)
}

// GetControlPlaneSubnet mocks base method.
func (m *MockAPI) GetControlPlaneSubnet(ctx context.Context, resourceGroupName, virtualNetwork, subnet string) (*network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlPlaneSubnet", ctx, resourceGroupName, virtualNetwork, subnet)
	ret0, _ := ret[0].(*network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetControlPlaneSubnet indicates an expected call of GetControlPlaneSubnet.
func (mr *MockAPIMockRecorder) GetControlPlaneSubnet(ctx, resourceGroupName, virtualNetwork, subnet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlPlaneSubnet", reflect.TypeOf((*MockAPI)(nil).GetControlPlaneSubnet), ctx, resourceGroupName, virtualNetwork, subnet)
}

// GetDiskEncryptionSet mocks base method.
func (m *MockAPI) GetDiskEncryptionSet(ctx context.Context, subscriptionID, groupName, diskEncryptionSetName string) (*compute0.DiskEncryptionSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskEncryptionSet", ctx, subscriptionID, groupName, diskEncryptionSetName)
	ret0, _ := ret[0].(*compute0.DiskEncryptionSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskEncryptionSet indicates an expected call of GetDiskEncryptionSet.
func (mr *MockAPIMockRecorder) GetDiskEncryptionSet(ctx, subscriptionID, groupName, diskEncryptionSetName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskEncryptionSet", reflect.TypeOf((*MockAPI)(nil).GetDiskEncryptionSet), ctx, subscriptionID, groupName, diskEncryptionSetName)
}

// GetDiskSkus mocks base method.
func (m *MockAPI) GetDiskSkus(ctx context.Context, region string) ([]compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskSkus", ctx, region)
	ret0, _ := ret[0].([]compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskSkus indicates an expected call of GetDiskSkus.
func (mr *MockAPIMockRecorder) GetDiskSkus(ctx, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskSkus", reflect.TypeOf((*MockAPI)(nil).GetDiskSkus), ctx, region)
}

// GetGroup mocks base method.
func (m *MockAPI) GetGroup(ctx context.Context, groupName string) (*resources.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, groupName)
	ret0, _ := ret[0].(*resources.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockAPIMockRecorder) GetGroup(ctx, groupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockAPI)(nil).GetGroup), ctx, groupName)
}

// GetHyperVGenerationVersion mocks base method.
func (m *MockAPI) GetHyperVGenerationVersion(ctx context.Context, instanceType, region string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHyperVGenerationVersion", ctx, instanceType, region)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHyperVGenerationVersion indicates an expected call of GetHyperVGenerationVersion.
func (mr *MockAPIMockRecorder) GetHyperVGenerationVersion(ctx, instanceType, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHyperVGenerationVersion", reflect.TypeOf((*MockAPI)(nil).GetHyperVGenerationVersion), ctx, instanceType, region)
}

// GetMarketplaceImage mocks base method.
func (m *MockAPI) GetMarketplaceImage(ctx context.Context, region, publisher, offer, sku, version string) (compute0.VirtualMachineImage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketplaceImage", ctx, region, publisher, offer, sku, version)
	ret0, _ := ret[0].(compute0.VirtualMachineImage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketplaceImage indicates an expected call of GetMarketplaceImage.
func (mr *MockAPIMockRecorder) GetMarketplaceImage(ctx, region, publisher, offer, sku, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketplaceImage", reflect.TypeOf((*MockAPI)(nil).GetMarketplaceImage), ctx, region, publisher, offer, sku, version)
}

// GetResourcesProvider mocks base method.
func (m *MockAPI) GetResourcesProvider(ctx context.Context, resourceProviderNamespace string) (*resources.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesProvider", ctx, resourceProviderNamespace)
	ret0, _ := ret[0].(*resources.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesProvider indicates an expected call of GetResourcesProvider.
func (mr *MockAPIMockRecorder) GetResourcesProvider(ctx, resourceProviderNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesProvider", reflect.TypeOf((*MockAPI)(nil).GetResourcesProvider), ctx, resourceProviderNamespace)
}

// GetStorageEndpointSuffix mocks base method.
func (m *MockAPI) GetStorageEndpointSuffix(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageEndpointSuffix", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageEndpointSuffix indicates an expected call of GetStorageEndpointSuffix.
func (mr *MockAPIMockRecorder) GetStorageEndpointSuffix(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageEndpointSuffix", reflect.TypeOf((*MockAPI)(nil).GetStorageEndpointSuffix), ctx)
}

// GetVMCapabilities mocks base method.
func (m *MockAPI) GetVMCapabilities(ctx context.Context, instanceType, region string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMCapabilities", ctx, instanceType, region)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVMCapabilities indicates an expected call of GetVMCapabilities.
func (mr *MockAPIMockRecorder) GetVMCapabilities(ctx, instanceType, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMCapabilities", reflect.TypeOf((*MockAPI)(nil).GetVMCapabilities), ctx, instanceType, region)
}

// GetVirtualMachineSku mocks base method.
func (m *MockAPI) GetVirtualMachineSku(ctx context.Context, name, region string) (*compute.ResourceSku, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineSku", ctx, name, region)
	ret0, _ := ret[0].(*compute.ResourceSku)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineSku indicates an expected call of GetVirtualMachineSku.
func (mr *MockAPIMockRecorder) GetVirtualMachineSku(ctx, name, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineSku", reflect.TypeOf((*MockAPI)(nil).GetVirtualMachineSku), ctx, name, region)
}

// GetVirtualNetwork mocks base method.
func (m *MockAPI) GetVirtualNetwork(ctx context.Context, resourceGroupName, virtualNetwork string) (*network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNetwork", ctx, resourceGroupName, virtualNetwork)
	ret0, _ := ret[0].(*network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualNetwork indicates an expected call of GetVirtualNetwork.
func (mr *MockAPIMockRecorder) GetVirtualNetwork(ctx, resourceGroupName, virtualNetwork interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNetwork", reflect.TypeOf((*MockAPI)(nil).GetVirtualNetwork), ctx, resourceGroupName, virtualNetwork)
}

// ListLocations mocks base method.
func (m *MockAPI) ListLocations(ctx context.Context) (*[]subscriptions.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocations", ctx)
	ret0, _ := ret[0].(*[]subscriptions.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocations indicates an expected call of ListLocations.
func (mr *MockAPIMockRecorder) ListLocations(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocations", reflect.TypeOf((*MockAPI)(nil).ListLocations), ctx)
}

// ListResourceIDsByGroup mocks base method.
func (m *MockAPI) ListResourceIDsByGroup(ctx context.Context, groupName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceIDsByGroup", ctx, groupName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceIDsByGroup indicates an expected call of ListResourceIDsByGroup.
func (mr *MockAPIMockRecorder) ListResourceIDsByGroup(ctx, groupName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceIDsByGroup", reflect.TypeOf((*MockAPI)(nil).ListResourceIDsByGroup), ctx, groupName)
}
