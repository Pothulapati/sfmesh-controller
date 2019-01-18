package networkapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string) (result network.DNSNameAvailabilityResult, err error)
}

var _ BaseClientAPI = (*network.BaseClient)(nil)

// ApplicationGatewaysClientAPI contains the set of methods on the ApplicationGatewaysClient type.
type ApplicationGatewaysClientAPI interface {
	BackendHealth(ctx context.Context, resourceGroupName string, applicationGatewayName string, expand string) (result network.ApplicationGatewaysBackendHealthFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters network.ApplicationGateway) (result network.ApplicationGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGateway, err error)
	GetSslPredefinedPolicy(ctx context.Context, predefinedPolicyName string) (result network.ApplicationGatewaySslPredefinedPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ApplicationGatewayListResultPage, err error)
	ListAll(ctx context.Context) (result network.ApplicationGatewayListResultPage, err error)
	ListAvailableSslOptions(ctx context.Context) (result network.ApplicationGatewayAvailableSslOptions, err error)
	ListAvailableSslPredefinedPolicies(ctx context.Context) (result network.ApplicationGatewayAvailableSslPredefinedPoliciesPage, err error)
	ListAvailableWafRuleSets(ctx context.Context) (result network.ApplicationGatewayAvailableWafRuleSetsResult, err error)
	Start(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStopFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters network.TagsObject) (result network.ApplicationGatewaysUpdateTagsFuture, err error)
}

var _ ApplicationGatewaysClientAPI = (*network.ApplicationGatewaysClient)(nil)

// ApplicationSecurityGroupsClientAPI contains the set of methods on the ApplicationSecurityGroupsClient type.
type ApplicationSecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters network.ApplicationSecurityGroup) (result network.ApplicationSecurityGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (result network.ApplicationSecurityGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (result network.ApplicationSecurityGroup, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ApplicationSecurityGroupListResultPage, err error)
	ListAll(ctx context.Context) (result network.ApplicationSecurityGroupListResultPage, err error)
}

var _ ApplicationSecurityGroupsClientAPI = (*network.ApplicationSecurityGroupsClient)(nil)

// AvailableEndpointServicesClientAPI contains the set of methods on the AvailableEndpointServicesClient type.
type AvailableEndpointServicesClientAPI interface {
	List(ctx context.Context, location string) (result network.EndpointServicesListResultPage, err error)
}

var _ AvailableEndpointServicesClientAPI = (*network.AvailableEndpointServicesClient)(nil)

// ExpressRouteCircuitAuthorizationsClientAPI contains the set of methods on the ExpressRouteCircuitAuthorizationsClient type.
type ExpressRouteCircuitAuthorizationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters network.ExpressRouteCircuitAuthorization) (result network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorizationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorization, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.AuthorizationListResultPage, err error)
}

var _ ExpressRouteCircuitAuthorizationsClientAPI = (*network.ExpressRouteCircuitAuthorizationsClient)(nil)

// ExpressRouteCircuitPeeringsClientAPI contains the set of methods on the ExpressRouteCircuitPeeringsClient type.
type ExpressRouteCircuitPeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters network.ExpressRouteCircuitPeering) (result network.ExpressRouteCircuitPeeringsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeeringsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeering, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitPeeringListResultPage, err error)
}

var _ ExpressRouteCircuitPeeringsClientAPI = (*network.ExpressRouteCircuitPeeringsClient)(nil)

// ExpressRouteCircuitsClientAPI contains the set of methods on the ExpressRouteCircuitsClient type.
type ExpressRouteCircuitsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters network.ExpressRouteCircuit) (result network.ExpressRouteCircuitsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuit, err error)
	GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitStats, err error)
	GetStats(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitStats, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ExpressRouteCircuitListResultPage, err error)
	ListAll(ctx context.Context) (result network.ExpressRouteCircuitListResultPage, err error)
	ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListArpTableFuture, err error)
	ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListRoutesTableFuture, err error)
	ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListRoutesTableSummaryFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters network.TagsObject) (result network.ExpressRouteCircuitsUpdateTagsFuture, err error)
}

var _ ExpressRouteCircuitsClientAPI = (*network.ExpressRouteCircuitsClient)(nil)

// ExpressRouteServiceProvidersClientAPI contains the set of methods on the ExpressRouteServiceProvidersClient type.
type ExpressRouteServiceProvidersClientAPI interface {
	List(ctx context.Context) (result network.ExpressRouteServiceProviderListResultPage, err error)
}

var _ ExpressRouteServiceProvidersClientAPI = (*network.ExpressRouteServiceProvidersClient)(nil)

// LoadBalancersClientAPI contains the set of methods on the LoadBalancersClient type.
type LoadBalancersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer) (result network.LoadBalancersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LoadBalancerListResultPage, err error)
	ListAll(ctx context.Context) (result network.LoadBalancerListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters network.TagsObject) (result network.LoadBalancersUpdateTagsFuture, err error)
}

var _ LoadBalancersClientAPI = (*network.LoadBalancersClient)(nil)

// LoadBalancerBackendAddressPoolsClientAPI contains the set of methods on the LoadBalancerBackendAddressPoolsClient type.
type LoadBalancerBackendAddressPoolsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (result network.BackendAddressPool, err error)
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancerBackendAddressPoolListResultPage, err error)
}

var _ LoadBalancerBackendAddressPoolsClientAPI = (*network.LoadBalancerBackendAddressPoolsClient)(nil)

// LoadBalancerFrontendIPConfigurationsClientAPI contains the set of methods on the LoadBalancerFrontendIPConfigurationsClient type.
type LoadBalancerFrontendIPConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (result network.FrontendIPConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancerFrontendIPConfigurationListResultPage, err error)
}

var _ LoadBalancerFrontendIPConfigurationsClientAPI = (*network.LoadBalancerFrontendIPConfigurationsClient)(nil)

// InboundNatRulesClientAPI contains the set of methods on the InboundNatRulesClient type.
type InboundNatRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters network.InboundNatRule) (result network.InboundNatRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string) (result network.InboundNatRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, expand string) (result network.InboundNatRule, err error)
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.InboundNatRuleListResultPage, err error)
}

var _ InboundNatRulesClientAPI = (*network.InboundNatRulesClient)(nil)

// LoadBalancerLoadBalancingRulesClientAPI contains the set of methods on the LoadBalancerLoadBalancingRulesClient type.
type LoadBalancerLoadBalancingRulesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (result network.LoadBalancingRule, err error)
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancerLoadBalancingRuleListResultPage, err error)
}

var _ LoadBalancerLoadBalancingRulesClientAPI = (*network.LoadBalancerLoadBalancingRulesClient)(nil)

// LoadBalancerNetworkInterfacesClientAPI contains the set of methods on the LoadBalancerNetworkInterfacesClient type.
type LoadBalancerNetworkInterfacesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.InterfaceListResultPage, err error)
}

var _ LoadBalancerNetworkInterfacesClientAPI = (*network.LoadBalancerNetworkInterfacesClient)(nil)

// LoadBalancerProbesClientAPI contains the set of methods on the LoadBalancerProbesClient type.
type LoadBalancerProbesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, probeName string) (result network.Probe, err error)
	List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancerProbeListResultPage, err error)
}

var _ LoadBalancerProbesClientAPI = (*network.LoadBalancerProbesClient)(nil)

// InterfacesClientAPI contains the set of methods on the InterfacesClient type.
type InterfacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters network.Interface) (result network.InterfacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error)
	GetEffectiveRouteTable(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesGetEffectiveRouteTableFuture, err error)
	GetVirtualMachineScaleSetIPConfiguration(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, expand string) (result network.InterfaceIPConfiguration, err error)
	GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error)
	List(ctx context.Context, resourceGroupName string) (result network.InterfaceListResultPage, err error)
	ListAll(ctx context.Context) (result network.InterfaceListResultPage, err error)
	ListEffectiveNetworkSecurityGroups(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesListEffectiveNetworkSecurityGroupsFuture, err error)
	ListVirtualMachineScaleSetIPConfigurations(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.InterfaceIPConfigurationListResultPage, err error)
	ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.InterfaceListResultPage, err error)
	ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result network.InterfaceListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters network.TagsObject) (result network.InterfacesUpdateTagsFuture, err error)
}

var _ InterfacesClientAPI = (*network.InterfacesClient)(nil)

// InterfaceIPConfigurationsClientAPI contains the set of methods on the InterfaceIPConfigurationsClient type.
type InterfaceIPConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, IPConfigurationName string) (result network.InterfaceIPConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfaceIPConfigurationListResultPage, err error)
}

var _ InterfaceIPConfigurationsClientAPI = (*network.InterfaceIPConfigurationsClient)(nil)

// InterfaceLoadBalancersClientAPI contains the set of methods on the InterfaceLoadBalancersClient type.
type InterfaceLoadBalancersClientAPI interface {
	List(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfaceLoadBalancerListResultPage, err error)
}

var _ InterfaceLoadBalancersClientAPI = (*network.InterfaceLoadBalancersClient)(nil)

// SecurityGroupsClientAPI contains the set of methods on the SecurityGroupsClient type.
type SecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup) (result network.SecurityGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error)
	List(ctx context.Context, resourceGroupName string) (result network.SecurityGroupListResultPage, err error)
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters network.TagsObject) (result network.SecurityGroupsUpdateTagsFuture, err error)
}

var _ SecurityGroupsClientAPI = (*network.SecurityGroupsClient)(nil)

// SecurityRulesClientAPI contains the set of methods on the SecurityRulesClient type.
type SecurityRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters network.SecurityRule) (result network.SecurityRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRule, err error)
	List(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityRuleListResultPage, err error)
}

var _ SecurityRulesClientAPI = (*network.SecurityRulesClient)(nil)

// DefaultSecurityRulesClientAPI contains the set of methods on the DefaultSecurityRulesClient type.
type DefaultSecurityRulesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string) (result network.SecurityRule, err error)
	List(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityRuleListResultPage, err error)
}

var _ DefaultSecurityRulesClientAPI = (*network.DefaultSecurityRulesClient)(nil)

// WatchersClientAPI contains the set of methods on the WatchersClient type.
type WatchersClientAPI interface {
	CheckConnectivity(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.ConnectivityParameters) (result network.WatchersCheckConnectivityFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.Watcher) (result network.Watcher, err error)
	Delete(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.WatchersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.Watcher, err error)
	GetAzureReachabilityReport(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.AzureReachabilityReportParameters) (result network.WatchersGetAzureReachabilityReportFuture, err error)
	GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogStatusParameters) (result network.WatchersGetFlowLogStatusFuture, err error)
	GetNextHop(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.NextHopParameters) (result network.WatchersGetNextHopFuture, err error)
	GetTopology(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.TopologyParameters) (result network.Topology, err error)
	GetTroubleshooting(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.TroubleshootingParameters) (result network.WatchersGetTroubleshootingFuture, err error)
	GetTroubleshootingResult(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.QueryTroubleshootingParameters) (result network.WatchersGetTroubleshootingResultFuture, err error)
	GetVMSecurityRules(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.SecurityGroupViewParameters) (result network.WatchersGetVMSecurityRulesFuture, err error)
	List(ctx context.Context, resourceGroupName string) (result network.WatcherListResult, err error)
	ListAll(ctx context.Context) (result network.WatcherListResult, err error)
	ListAvailableProviders(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.AvailableProvidersListParameters) (result network.WatchersListAvailableProvidersFuture, err error)
	SetFlowLogConfiguration(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogInformation) (result network.WatchersSetFlowLogConfigurationFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.TagsObject) (result network.Watcher, err error)
	VerifyIPFlow(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.VerificationIPFlowParameters) (result network.WatchersVerifyIPFlowFuture, err error)
}

var _ WatchersClientAPI = (*network.WatchersClient)(nil)

// PacketCapturesClientAPI contains the set of methods on the PacketCapturesClient type.
type PacketCapturesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters network.PacketCapture) (result network.PacketCapturesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCaptureResult, err error)
	GetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesGetStatusFuture, err error)
	List(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.PacketCaptureListResult, err error)
	Stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesStopFuture, err error)
}

var _ PacketCapturesClientAPI = (*network.PacketCapturesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result network.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*network.OperationsClient)(nil)

// PublicIPAddressesClientAPI contains the set of methods on the PublicIPAddressesClient type.
type PublicIPAddressesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (result network.PublicIPAddressesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	List(ctx context.Context, resourceGroupName string) (result network.PublicIPAddressListResultPage, err error)
	ListAll(ctx context.Context) (result network.PublicIPAddressListResultPage, err error)
	ListVirtualMachineScaleSetPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.PublicIPAddressListResultPage, err error)
	ListVirtualMachineScaleSetVMPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result network.PublicIPAddressListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters network.TagsObject) (result network.PublicIPAddressesUpdateTagsFuture, err error)
}

var _ PublicIPAddressesClientAPI = (*network.PublicIPAddressesClient)(nil)

// RouteFiltersClientAPI contains the set of methods on the RouteFiltersClient type.
type RouteFiltersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters network.RouteFilter) (result network.RouteFiltersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeFilterName string) (result network.RouteFiltersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, expand string) (result network.RouteFilter, err error)
	List(ctx context.Context) (result network.RouteFilterListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result network.RouteFilterListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters network.PatchRouteFilter) (result network.RouteFiltersUpdateFuture, err error)
}

var _ RouteFiltersClientAPI = (*network.RouteFiltersClient)(nil)

// RouteFilterRulesClientAPI contains the set of methods on the RouteFilterRulesClient type.
type RouteFilterRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters network.RouteFilterRule) (result network.RouteFilterRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result network.RouteFilterRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result network.RouteFilterRule, err error)
	ListByRouteFilter(ctx context.Context, resourceGroupName string, routeFilterName string) (result network.RouteFilterRuleListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters network.PatchRouteFilterRule) (result network.RouteFilterRulesUpdateFuture, err error)
}

var _ RouteFilterRulesClientAPI = (*network.RouteFilterRulesClient)(nil)

// RouteTablesClientAPI contains the set of methods on the RouteTablesClient type.
type RouteTablesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters network.RouteTable) (result network.RouteTablesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteTablesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error)
	List(ctx context.Context, resourceGroupName string) (result network.RouteTableListResultPage, err error)
	ListAll(ctx context.Context) (result network.RouteTableListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, routeTableName string, parameters network.TagsObject) (result network.RouteTablesUpdateTagsFuture, err error)
}

var _ RouteTablesClientAPI = (*network.RouteTablesClient)(nil)

// RoutesClientAPI contains the set of methods on the RoutesClient type.
type RoutesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route) (result network.RoutesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.RoutesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.Route, err error)
	List(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteListResultPage, err error)
}

var _ RoutesClientAPI = (*network.RoutesClient)(nil)

// BgpServiceCommunitiesClientAPI contains the set of methods on the BgpServiceCommunitiesClient type.
type BgpServiceCommunitiesClientAPI interface {
	List(ctx context.Context) (result network.BgpServiceCommunityListResultPage, err error)
}

var _ BgpServiceCommunitiesClientAPI = (*network.BgpServiceCommunitiesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result network.UsagesListResultPage, err error)
}

var _ UsagesClientAPI = (*network.UsagesClient)(nil)

// VirtualNetworksClientAPI contains the set of methods on the VirtualNetworksClient type.
type VirtualNetworksClientAPI interface {
	CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, IPAddress string) (result network.IPAddressAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters network.VirtualNetwork) (result network.VirtualNetworksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand string) (result network.VirtualNetwork, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkListResultPage, err error)
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
	ListUsage(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkListUsageResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters network.TagsObject) (result network.VirtualNetworksUpdateTagsFuture, err error)
}

var _ VirtualNetworksClientAPI = (*network.VirtualNetworksClient)(nil)

// SubnetsClientAPI contains the set of methods on the SubnetsClient type.
type SubnetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet) (result network.SubnetsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result network.SubnetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error)
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.SubnetListResultPage, err error)
}

var _ SubnetsClientAPI = (*network.SubnetsClient)(nil)

// VirtualNetworkPeeringsClientAPI contains the set of methods on the VirtualNetworkPeeringsClient type.
type VirtualNetworkPeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters network.VirtualNetworkPeering) (result network.VirtualNetworkPeeringsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result network.VirtualNetworkPeeringsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result network.VirtualNetworkPeering, err error)
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkPeeringListResultPage, err error)
}

var _ VirtualNetworkPeeringsClientAPI = (*network.VirtualNetworkPeeringsClient)(nil)

// VirtualNetworkGatewaysClientAPI contains the set of methods on the VirtualNetworkGatewaysClient type.
type VirtualNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VirtualNetworkGateway) (result network.VirtualNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysDeleteFuture, err error)
	Generatevpnclientpackage(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VpnClientParameters) (result network.VirtualNetworkGatewaysGeneratevpnclientpackageFuture, err error)
	GenerateVpnProfile(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VpnClientParameters) (result network.VirtualNetworkGatewaysGenerateVpnProfileFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGateway, err error)
	GetAdvertisedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result network.VirtualNetworkGatewaysGetAdvertisedRoutesFuture, err error)
	GetBgpPeerStatus(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result network.VirtualNetworkGatewaysGetBgpPeerStatusFuture, err error)
	GetLearnedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysGetLearnedRoutesFuture, err error)
	GetVpnProfilePackageURL(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysGetVpnProfilePackageURLFuture, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultPage, err error)
	ListConnections(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewayListConnectionsResultPage, err error)
	Reset(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, gatewayVip string) (result network.VirtualNetworkGatewaysResetFuture, err error)
	SupportedVpnDevices(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.String, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.TagsObject) (result network.VirtualNetworkGatewaysUpdateTagsFuture, err error)
	VpnDeviceConfigurationScript(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.VpnDeviceScriptParameters) (result network.String, err error)
}

var _ VirtualNetworkGatewaysClientAPI = (*network.VirtualNetworkGatewaysClient)(nil)

// VirtualNetworkGatewayConnectionsClientAPI contains the set of methods on the VirtualNetworkGatewayConnectionsClient type.
type VirtualNetworkGatewayConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.VirtualNetworkGatewayConnection) (result network.VirtualNetworkGatewayConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnection, err error)
	GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.ConnectionSharedKey, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayConnectionListResultPage, err error)
	ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionResetSharedKey) (result network.VirtualNetworkGatewayConnectionsResetSharedKeyFuture, err error)
	SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionSharedKey) (result network.VirtualNetworkGatewayConnectionsSetSharedKeyFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.TagsObject) (result network.VirtualNetworkGatewayConnectionsUpdateTagsFuture, err error)
}

var _ VirtualNetworkGatewayConnectionsClientAPI = (*network.VirtualNetworkGatewayConnectionsClient)(nil)

// LocalNetworkGatewaysClientAPI contains the set of methods on the LocalNetworkGatewaysClient type.
type LocalNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters network.LocalNetworkGateway) (result network.LocalNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LocalNetworkGatewayListResultPage, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters network.TagsObject) (result network.LocalNetworkGatewaysUpdateTagsFuture, err error)
}

var _ LocalNetworkGatewaysClientAPI = (*network.LocalNetworkGatewaysClient)(nil)
