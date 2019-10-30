/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/google/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeGoogleV1alpha1 struct {
	*testing.Fake
}

func (c *FakeGoogleV1alpha1) AppEngineApplications(namespace string) v1alpha1.AppEngineApplicationInterface {
	return &FakeAppEngineApplications{c, namespace}
}

func (c *FakeGoogleV1alpha1) BigqueryDatasets(namespace string) v1alpha1.BigqueryDatasetInterface {
	return &FakeBigqueryDatasets{c, namespace}
}

func (c *FakeGoogleV1alpha1) BigqueryTables(namespace string) v1alpha1.BigqueryTableInterface {
	return &FakeBigqueryTables{c, namespace}
}

func (c *FakeGoogleV1alpha1) BigtableInstances(namespace string) v1alpha1.BigtableInstanceInterface {
	return &FakeBigtableInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) BigtableTables(namespace string) v1alpha1.BigtableTableInterface {
	return &FakeBigtableTables{c, namespace}
}

func (c *FakeGoogleV1alpha1) BillingAccountIamBindings(namespace string) v1alpha1.BillingAccountIamBindingInterface {
	return &FakeBillingAccountIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) BillingAccountIamMembers(namespace string) v1alpha1.BillingAccountIamMemberInterface {
	return &FakeBillingAccountIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) BillingAccountIamPolicies(namespace string) v1alpha1.BillingAccountIamPolicyInterface {
	return &FakeBillingAccountIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) BinaryAuthorizationAttestors(namespace string) v1alpha1.BinaryAuthorizationAttestorInterface {
	return &FakeBinaryAuthorizationAttestors{c, namespace}
}

func (c *FakeGoogleV1alpha1) BinaryAuthorizationPolicies(namespace string) v1alpha1.BinaryAuthorizationPolicyInterface {
	return &FakeBinaryAuthorizationPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) CloudbuildTriggers(namespace string) v1alpha1.CloudbuildTriggerInterface {
	return &FakeCloudbuildTriggers{c, namespace}
}

func (c *FakeGoogleV1alpha1) CloudfunctionsFunctions(namespace string) v1alpha1.CloudfunctionsFunctionInterface {
	return &FakeCloudfunctionsFunctions{c, namespace}
}

func (c *FakeGoogleV1alpha1) CloudiotRegistries(namespace string) v1alpha1.CloudiotRegistryInterface {
	return &FakeCloudiotRegistries{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComposerEnvironments(namespace string) v1alpha1.ComposerEnvironmentInterface {
	return &FakeComposerEnvironments{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeAddresses(namespace string) v1alpha1.ComputeAddressInterface {
	return &FakeComputeAddresses{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeAttachedDisks(namespace string) v1alpha1.ComputeAttachedDiskInterface {
	return &FakeComputeAttachedDisks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeAutoscalers(namespace string) v1alpha1.ComputeAutoscalerInterface {
	return &FakeComputeAutoscalers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeBackendBuckets(namespace string) v1alpha1.ComputeBackendBucketInterface {
	return &FakeComputeBackendBuckets{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeBackendServices(namespace string) v1alpha1.ComputeBackendServiceInterface {
	return &FakeComputeBackendServices{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeDisks(namespace string) v1alpha1.ComputeDiskInterface {
	return &FakeComputeDisks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeFirewalls(namespace string) v1alpha1.ComputeFirewallInterface {
	return &FakeComputeFirewalls{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeForwardingRules(namespace string) v1alpha1.ComputeForwardingRuleInterface {
	return &FakeComputeForwardingRules{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeGlobalAddresses(namespace string) v1alpha1.ComputeGlobalAddressInterface {
	return &FakeComputeGlobalAddresses{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeGlobalForwardingRules(namespace string) v1alpha1.ComputeGlobalForwardingRuleInterface {
	return &FakeComputeGlobalForwardingRules{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeHTTPHealthChecks(namespace string) v1alpha1.ComputeHTTPHealthCheckInterface {
	return &FakeComputeHTTPHealthChecks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeHTTPSHealthChecks(namespace string) v1alpha1.ComputeHTTPSHealthCheckInterface {
	return &FakeComputeHTTPSHealthChecks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeHealthChecks(namespace string) v1alpha1.ComputeHealthCheckInterface {
	return &FakeComputeHealthChecks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeImages(namespace string) v1alpha1.ComputeImageInterface {
	return &FakeComputeImages{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInstances(namespace string) v1alpha1.ComputeInstanceInterface {
	return &FakeComputeInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInstanceFromTemplates(namespace string) v1alpha1.ComputeInstanceFromTemplateInterface {
	return &FakeComputeInstanceFromTemplates{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInstanceGroups(namespace string) v1alpha1.ComputeInstanceGroupInterface {
	return &FakeComputeInstanceGroups{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInstanceGroupManagers(namespace string) v1alpha1.ComputeInstanceGroupManagerInterface {
	return &FakeComputeInstanceGroupManagers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInstanceTemplates(namespace string) v1alpha1.ComputeInstanceTemplateInterface {
	return &FakeComputeInstanceTemplates{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeInterconnectAttachments(namespace string) v1alpha1.ComputeInterconnectAttachmentInterface {
	return &FakeComputeInterconnectAttachments{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeNetworks(namespace string) v1alpha1.ComputeNetworkInterface {
	return &FakeComputeNetworks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeNetworkPeerings(namespace string) v1alpha1.ComputeNetworkPeeringInterface {
	return &FakeComputeNetworkPeerings{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeProjectMetadatas(namespace string) v1alpha1.ComputeProjectMetadataInterface {
	return &FakeComputeProjectMetadatas{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeProjectMetadataItems(namespace string) v1alpha1.ComputeProjectMetadataItemInterface {
	return &FakeComputeProjectMetadataItems{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRegionAutoscalers(namespace string) v1alpha1.ComputeRegionAutoscalerInterface {
	return &FakeComputeRegionAutoscalers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRegionBackendServices(namespace string) v1alpha1.ComputeRegionBackendServiceInterface {
	return &FakeComputeRegionBackendServices{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRegionDisks(namespace string) v1alpha1.ComputeRegionDiskInterface {
	return &FakeComputeRegionDisks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRegionInstanceGroupManagers(namespace string) v1alpha1.ComputeRegionInstanceGroupManagerInterface {
	return &FakeComputeRegionInstanceGroupManagers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRoutes(namespace string) v1alpha1.ComputeRouteInterface {
	return &FakeComputeRoutes{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRouters(namespace string) v1alpha1.ComputeRouterInterface {
	return &FakeComputeRouters{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRouterInterfaces(namespace string) v1alpha1.ComputeRouterInterfaceInterface {
	return &FakeComputeRouterInterfaces{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRouterNATs(namespace string) v1alpha1.ComputeRouterNATInterface {
	return &FakeComputeRouterNATs{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeRouterPeers(namespace string) v1alpha1.ComputeRouterPeerInterface {
	return &FakeComputeRouterPeers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSSLCertificates(namespace string) v1alpha1.ComputeSSLCertificateInterface {
	return &FakeComputeSSLCertificates{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSSLPolicies(namespace string) v1alpha1.ComputeSSLPolicyInterface {
	return &FakeComputeSSLPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSecurityPolicies(namespace string) v1alpha1.ComputeSecurityPolicyInterface {
	return &FakeComputeSecurityPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSharedVpcHostProjects(namespace string) v1alpha1.ComputeSharedVpcHostProjectInterface {
	return &FakeComputeSharedVpcHostProjects{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSharedVpcServiceProjects(namespace string) v1alpha1.ComputeSharedVpcServiceProjectInterface {
	return &FakeComputeSharedVpcServiceProjects{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSnapshots(namespace string) v1alpha1.ComputeSnapshotInterface {
	return &FakeComputeSnapshots{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSubnetworks(namespace string) v1alpha1.ComputeSubnetworkInterface {
	return &FakeComputeSubnetworks{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSubnetworkIamBindings(namespace string) v1alpha1.ComputeSubnetworkIamBindingInterface {
	return &FakeComputeSubnetworkIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSubnetworkIamMembers(namespace string) v1alpha1.ComputeSubnetworkIamMemberInterface {
	return &FakeComputeSubnetworkIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeSubnetworkIamPolicies(namespace string) v1alpha1.ComputeSubnetworkIamPolicyInterface {
	return &FakeComputeSubnetworkIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeTargetHTTPProxies(namespace string) v1alpha1.ComputeTargetHTTPProxyInterface {
	return &FakeComputeTargetHTTPProxies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeTargetHTTPSProxies(namespace string) v1alpha1.ComputeTargetHTTPSProxyInterface {
	return &FakeComputeTargetHTTPSProxies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeTargetPools(namespace string) v1alpha1.ComputeTargetPoolInterface {
	return &FakeComputeTargetPools{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeTargetSSLProxies(namespace string) v1alpha1.ComputeTargetSSLProxyInterface {
	return &FakeComputeTargetSSLProxies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeTargetTcpProxies(namespace string) v1alpha1.ComputeTargetTcpProxyInterface {
	return &FakeComputeTargetTcpProxies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeURLMaps(namespace string) v1alpha1.ComputeURLMapInterface {
	return &FakeComputeURLMaps{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeVPNGateways(namespace string) v1alpha1.ComputeVPNGatewayInterface {
	return &FakeComputeVPNGateways{c, namespace}
}

func (c *FakeGoogleV1alpha1) ComputeVPNTunnels(namespace string) v1alpha1.ComputeVPNTunnelInterface {
	return &FakeComputeVPNTunnels{c, namespace}
}

func (c *FakeGoogleV1alpha1) ContainerAnalysisNotes(namespace string) v1alpha1.ContainerAnalysisNoteInterface {
	return &FakeContainerAnalysisNotes{c, namespace}
}

func (c *FakeGoogleV1alpha1) ContainerClusters(namespace string) v1alpha1.ContainerClusterInterface {
	return &FakeContainerClusters{c, namespace}
}

func (c *FakeGoogleV1alpha1) ContainerNodePools(namespace string) v1alpha1.ContainerNodePoolInterface {
	return &FakeContainerNodePools{c, namespace}
}

func (c *FakeGoogleV1alpha1) DataflowJobs(namespace string) v1alpha1.DataflowJobInterface {
	return &FakeDataflowJobs{c, namespace}
}

func (c *FakeGoogleV1alpha1) DataprocClusters(namespace string) v1alpha1.DataprocClusterInterface {
	return &FakeDataprocClusters{c, namespace}
}

func (c *FakeGoogleV1alpha1) DataprocJobs(namespace string) v1alpha1.DataprocJobInterface {
	return &FakeDataprocJobs{c, namespace}
}

func (c *FakeGoogleV1alpha1) DnsManagedZones(namespace string) v1alpha1.DnsManagedZoneInterface {
	return &FakeDnsManagedZones{c, namespace}
}

func (c *FakeGoogleV1alpha1) DnsRecordSets(namespace string) v1alpha1.DnsRecordSetInterface {
	return &FakeDnsRecordSets{c, namespace}
}

func (c *FakeGoogleV1alpha1) EndpointsServices(namespace string) v1alpha1.EndpointsServiceInterface {
	return &FakeEndpointsServices{c, namespace}
}

func (c *FakeGoogleV1alpha1) FilestoreInstances(namespace string) v1alpha1.FilestoreInstanceInterface {
	return &FakeFilestoreInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) Folders(namespace string) v1alpha1.FolderInterface {
	return &FakeFolders{c, namespace}
}

func (c *FakeGoogleV1alpha1) FolderIamBindings(namespace string) v1alpha1.FolderIamBindingInterface {
	return &FakeFolderIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) FolderIamMembers(namespace string) v1alpha1.FolderIamMemberInterface {
	return &FakeFolderIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) FolderIamPolicies(namespace string) v1alpha1.FolderIamPolicyInterface {
	return &FakeFolderIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) FolderOrganizationPolicies(namespace string) v1alpha1.FolderOrganizationPolicyInterface {
	return &FakeFolderOrganizationPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsCryptoKeys(namespace string) v1alpha1.KmsCryptoKeyInterface {
	return &FakeKmsCryptoKeys{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsCryptoKeyIamBindings(namespace string) v1alpha1.KmsCryptoKeyIamBindingInterface {
	return &FakeKmsCryptoKeyIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsCryptoKeyIamMembers(namespace string) v1alpha1.KmsCryptoKeyIamMemberInterface {
	return &FakeKmsCryptoKeyIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsKeyRings(namespace string) v1alpha1.KmsKeyRingInterface {
	return &FakeKmsKeyRings{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsKeyRingIamBindings(namespace string) v1alpha1.KmsKeyRingIamBindingInterface {
	return &FakeKmsKeyRingIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsKeyRingIamMembers(namespace string) v1alpha1.KmsKeyRingIamMemberInterface {
	return &FakeKmsKeyRingIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) KmsKeyRingIamPolicies(namespace string) v1alpha1.KmsKeyRingIamPolicyInterface {
	return &FakeKmsKeyRingIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingBillingAccountExclusions(namespace string) v1alpha1.LoggingBillingAccountExclusionInterface {
	return &FakeLoggingBillingAccountExclusions{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingBillingAccountSinks(namespace string) v1alpha1.LoggingBillingAccountSinkInterface {
	return &FakeLoggingBillingAccountSinks{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingFolderExclusions(namespace string) v1alpha1.LoggingFolderExclusionInterface {
	return &FakeLoggingFolderExclusions{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingFolderSinks(namespace string) v1alpha1.LoggingFolderSinkInterface {
	return &FakeLoggingFolderSinks{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingOrganizationExclusions(namespace string) v1alpha1.LoggingOrganizationExclusionInterface {
	return &FakeLoggingOrganizationExclusions{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingOrganizationSinks(namespace string) v1alpha1.LoggingOrganizationSinkInterface {
	return &FakeLoggingOrganizationSinks{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingProjectExclusions(namespace string) v1alpha1.LoggingProjectExclusionInterface {
	return &FakeLoggingProjectExclusions{c, namespace}
}

func (c *FakeGoogleV1alpha1) LoggingProjectSinks(namespace string) v1alpha1.LoggingProjectSinkInterface {
	return &FakeLoggingProjectSinks{c, namespace}
}

func (c *FakeGoogleV1alpha1) MonitoringAlertPolicies(namespace string) v1alpha1.MonitoringAlertPolicyInterface {
	return &FakeMonitoringAlertPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) MonitoringGroups(namespace string) v1alpha1.MonitoringGroupInterface {
	return &FakeMonitoringGroups{c, namespace}
}

func (c *FakeGoogleV1alpha1) MonitoringNotificationChannels(namespace string) v1alpha1.MonitoringNotificationChannelInterface {
	return &FakeMonitoringNotificationChannels{c, namespace}
}

func (c *FakeGoogleV1alpha1) MonitoringUptimeCheckConfigs(namespace string) v1alpha1.MonitoringUptimeCheckConfigInterface {
	return &FakeMonitoringUptimeCheckConfigs{c, namespace}
}

func (c *FakeGoogleV1alpha1) OrganizationIamBindings(namespace string) v1alpha1.OrganizationIamBindingInterface {
	return &FakeOrganizationIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) OrganizationIamCustomRoles(namespace string) v1alpha1.OrganizationIamCustomRoleInterface {
	return &FakeOrganizationIamCustomRoles{c, namespace}
}

func (c *FakeGoogleV1alpha1) OrganizationIamMembers(namespace string) v1alpha1.OrganizationIamMemberInterface {
	return &FakeOrganizationIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) OrganizationIamPolicies(namespace string) v1alpha1.OrganizationIamPolicyInterface {
	return &FakeOrganizationIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) OrganizationPolicies(namespace string) v1alpha1.OrganizationPolicyInterface {
	return &FakeOrganizationPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) Projects(namespace string) v1alpha1.ProjectInterface {
	return &FakeProjects{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectIamBindings(namespace string) v1alpha1.ProjectIamBindingInterface {
	return &FakeProjectIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectIamCustomRoles(namespace string) v1alpha1.ProjectIamCustomRoleInterface {
	return &FakeProjectIamCustomRoles{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectIamMembers(namespace string) v1alpha1.ProjectIamMemberInterface {
	return &FakeProjectIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectIamPolicies(namespace string) v1alpha1.ProjectIamPolicyInterface {
	return &FakeProjectIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectOrganizationPolicies(namespace string) v1alpha1.ProjectOrganizationPolicyInterface {
	return &FakeProjectOrganizationPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectServices(namespace string) v1alpha1.ProjectServiceInterface {
	return &FakeProjectServices{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectServiceses(namespace string) v1alpha1.ProjectServicesInterface {
	return &FakeProjectServiceses{c, namespace}
}

func (c *FakeGoogleV1alpha1) ProjectUsageExportBuckets(namespace string) v1alpha1.ProjectUsageExportBucketInterface {
	return &FakeProjectUsageExportBuckets{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubSubscriptions(namespace string) v1alpha1.PubsubSubscriptionInterface {
	return &FakePubsubSubscriptions{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubSubscriptionIamBindings(namespace string) v1alpha1.PubsubSubscriptionIamBindingInterface {
	return &FakePubsubSubscriptionIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubSubscriptionIamMembers(namespace string) v1alpha1.PubsubSubscriptionIamMemberInterface {
	return &FakePubsubSubscriptionIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubSubscriptionIamPolicies(namespace string) v1alpha1.PubsubSubscriptionIamPolicyInterface {
	return &FakePubsubSubscriptionIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubTopics(namespace string) v1alpha1.PubsubTopicInterface {
	return &FakePubsubTopics{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubTopicIamBindings(namespace string) v1alpha1.PubsubTopicIamBindingInterface {
	return &FakePubsubTopicIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubTopicIamMembers(namespace string) v1alpha1.PubsubTopicIamMemberInterface {
	return &FakePubsubTopicIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) PubsubTopicIamPolicies(namespace string) v1alpha1.PubsubTopicIamPolicyInterface {
	return &FakePubsubTopicIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) RedisInstances(namespace string) v1alpha1.RedisInstanceInterface {
	return &FakeRedisInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) ResourceManagerLiens(namespace string) v1alpha1.ResourceManagerLienInterface {
	return &FakeResourceManagerLiens{c, namespace}
}

func (c *FakeGoogleV1alpha1) RuntimeconfigConfigs(namespace string) v1alpha1.RuntimeconfigConfigInterface {
	return &FakeRuntimeconfigConfigs{c, namespace}
}

func (c *FakeGoogleV1alpha1) RuntimeconfigVariables(namespace string) v1alpha1.RuntimeconfigVariableInterface {
	return &FakeRuntimeconfigVariables{c, namespace}
}

func (c *FakeGoogleV1alpha1) ServiceAccounts(namespace string) v1alpha1.ServiceAccountInterface {
	return &FakeServiceAccounts{c, namespace}
}

func (c *FakeGoogleV1alpha1) ServiceAccountIamBindings(namespace string) v1alpha1.ServiceAccountIamBindingInterface {
	return &FakeServiceAccountIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) ServiceAccountIamMembers(namespace string) v1alpha1.ServiceAccountIamMemberInterface {
	return &FakeServiceAccountIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) ServiceAccountIamPolicies(namespace string) v1alpha1.ServiceAccountIamPolicyInterface {
	return &FakeServiceAccountIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) ServiceAccountKeys(namespace string) v1alpha1.ServiceAccountKeyInterface {
	return &FakeServiceAccountKeys{c, namespace}
}

func (c *FakeGoogleV1alpha1) SourcerepoRepositories(namespace string) v1alpha1.SourcerepoRepositoryInterface {
	return &FakeSourcerepoRepositories{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerDatabases(namespace string) v1alpha1.SpannerDatabaseInterface {
	return &FakeSpannerDatabases{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerDatabaseIamBindings(namespace string) v1alpha1.SpannerDatabaseIamBindingInterface {
	return &FakeSpannerDatabaseIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerDatabaseIamMembers(namespace string) v1alpha1.SpannerDatabaseIamMemberInterface {
	return &FakeSpannerDatabaseIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerDatabaseIamPolicies(namespace string) v1alpha1.SpannerDatabaseIamPolicyInterface {
	return &FakeSpannerDatabaseIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerInstances(namespace string) v1alpha1.SpannerInstanceInterface {
	return &FakeSpannerInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerInstanceIamBindings(namespace string) v1alpha1.SpannerInstanceIamBindingInterface {
	return &FakeSpannerInstanceIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerInstanceIamMembers(namespace string) v1alpha1.SpannerInstanceIamMemberInterface {
	return &FakeSpannerInstanceIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) SpannerInstanceIamPolicies(namespace string) v1alpha1.SpannerInstanceIamPolicyInterface {
	return &FakeSpannerInstanceIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) SqlDatabases(namespace string) v1alpha1.SqlDatabaseInterface {
	return &FakeSqlDatabases{c, namespace}
}

func (c *FakeGoogleV1alpha1) SqlDatabaseInstances(namespace string) v1alpha1.SqlDatabaseInstanceInterface {
	return &FakeSqlDatabaseInstances{c, namespace}
}

func (c *FakeGoogleV1alpha1) SqlSSLCerts(namespace string) v1alpha1.SqlSSLCertInterface {
	return &FakeSqlSSLCerts{c, namespace}
}

func (c *FakeGoogleV1alpha1) SqlUsers(namespace string) v1alpha1.SqlUserInterface {
	return &FakeSqlUsers{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBuckets(namespace string) v1alpha1.StorageBucketInterface {
	return &FakeStorageBuckets{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBucketACLs(namespace string) v1alpha1.StorageBucketACLInterface {
	return &FakeStorageBucketACLs{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBucketIamBindings(namespace string) v1alpha1.StorageBucketIamBindingInterface {
	return &FakeStorageBucketIamBindings{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBucketIamMembers(namespace string) v1alpha1.StorageBucketIamMemberInterface {
	return &FakeStorageBucketIamMembers{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBucketIamPolicies(namespace string) v1alpha1.StorageBucketIamPolicyInterface {
	return &FakeStorageBucketIamPolicies{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageBucketObjects(namespace string) v1alpha1.StorageBucketObjectInterface {
	return &FakeStorageBucketObjects{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageDefaultObjectACLs(namespace string) v1alpha1.StorageDefaultObjectACLInterface {
	return &FakeStorageDefaultObjectACLs{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageDefaultObjectAccessControls(namespace string) v1alpha1.StorageDefaultObjectAccessControlInterface {
	return &FakeStorageDefaultObjectAccessControls{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageNotifications(namespace string) v1alpha1.StorageNotificationInterface {
	return &FakeStorageNotifications{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageObjectACLs(namespace string) v1alpha1.StorageObjectACLInterface {
	return &FakeStorageObjectACLs{c, namespace}
}

func (c *FakeGoogleV1alpha1) StorageObjectAccessControls(namespace string) v1alpha1.StorageObjectAccessControlInterface {
	return &FakeStorageObjectAccessControls{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGoogleV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
