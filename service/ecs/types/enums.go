// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AgentUpdateStatus string

// Enum values for AgentUpdateStatus
const (
	AgentUpdateStatusPending  AgentUpdateStatus = "PENDING"
	AgentUpdateStatusStaging  AgentUpdateStatus = "STAGING"
	AgentUpdateStatusStaged   AgentUpdateStatus = "STAGED"
	AgentUpdateStatusUpdating AgentUpdateStatus = "UPDATING"
	AgentUpdateStatusUpdated  AgentUpdateStatus = "UPDATED"
	AgentUpdateStatusFailed   AgentUpdateStatus = "FAILED"
)

// Values returns all known values for AgentUpdateStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AgentUpdateStatus) Values() []AgentUpdateStatus {
	return []AgentUpdateStatus{
		"PENDING",
		"STAGING",
		"STAGED",
		"UPDATING",
		"UPDATED",
		"FAILED",
	}
}

type ApplicationProtocol string

// Enum values for ApplicationProtocol
const (
	ApplicationProtocolHttp  ApplicationProtocol = "http"
	ApplicationProtocolHttp2 ApplicationProtocol = "http2"
	ApplicationProtocolGrpc  ApplicationProtocol = "grpc"
)

// Values returns all known values for ApplicationProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationProtocol) Values() []ApplicationProtocol {
	return []ApplicationProtocol{
		"http",
		"http2",
		"grpc",
	}
}

type AssignPublicIp string

// Enum values for AssignPublicIp
const (
	AssignPublicIpEnabled  AssignPublicIp = "ENABLED"
	AssignPublicIpDisabled AssignPublicIp = "DISABLED"
)

// Values returns all known values for AssignPublicIp. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AssignPublicIp) Values() []AssignPublicIp {
	return []AssignPublicIp{
		"ENABLED",
		"DISABLED",
	}
}

type CapacityProviderField string

// Enum values for CapacityProviderField
const (
	CapacityProviderFieldTags CapacityProviderField = "TAGS"
)

// Values returns all known values for CapacityProviderField. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CapacityProviderField) Values() []CapacityProviderField {
	return []CapacityProviderField{
		"TAGS",
	}
}

type CapacityProviderStatus string

// Enum values for CapacityProviderStatus
const (
	CapacityProviderStatusActive   CapacityProviderStatus = "ACTIVE"
	CapacityProviderStatusInactive CapacityProviderStatus = "INACTIVE"
)

// Values returns all known values for CapacityProviderStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CapacityProviderStatus) Values() []CapacityProviderStatus {
	return []CapacityProviderStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type CapacityProviderUpdateStatus string

// Enum values for CapacityProviderUpdateStatus
const (
	CapacityProviderUpdateStatusDeleteInProgress CapacityProviderUpdateStatus = "DELETE_IN_PROGRESS"
	CapacityProviderUpdateStatusDeleteComplete   CapacityProviderUpdateStatus = "DELETE_COMPLETE"
	CapacityProviderUpdateStatusDeleteFailed     CapacityProviderUpdateStatus = "DELETE_FAILED"
	CapacityProviderUpdateStatusUpdateInProgress CapacityProviderUpdateStatus = "UPDATE_IN_PROGRESS"
	CapacityProviderUpdateStatusUpdateComplete   CapacityProviderUpdateStatus = "UPDATE_COMPLETE"
	CapacityProviderUpdateStatusUpdateFailed     CapacityProviderUpdateStatus = "UPDATE_FAILED"
)

// Values returns all known values for CapacityProviderUpdateStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CapacityProviderUpdateStatus) Values() []CapacityProviderUpdateStatus {
	return []CapacityProviderUpdateStatus{
		"DELETE_IN_PROGRESS",
		"DELETE_COMPLETE",
		"DELETE_FAILED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_COMPLETE",
		"UPDATE_FAILED",
	}
}

type ClusterField string

// Enum values for ClusterField
const (
	ClusterFieldAttachments    ClusterField = "ATTACHMENTS"
	ClusterFieldConfigurations ClusterField = "CONFIGURATIONS"
	ClusterFieldSettings       ClusterField = "SETTINGS"
	ClusterFieldStatistics     ClusterField = "STATISTICS"
	ClusterFieldTags           ClusterField = "TAGS"
)

// Values returns all known values for ClusterField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClusterField) Values() []ClusterField {
	return []ClusterField{
		"ATTACHMENTS",
		"CONFIGURATIONS",
		"SETTINGS",
		"STATISTICS",
		"TAGS",
	}
}

type ClusterSettingName string

// Enum values for ClusterSettingName
const (
	ClusterSettingNameContainerInsights ClusterSettingName = "containerInsights"
)

// Values returns all known values for ClusterSettingName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClusterSettingName) Values() []ClusterSettingName {
	return []ClusterSettingName{
		"containerInsights",
	}
}

type Compatibility string

// Enum values for Compatibility
const (
	CompatibilityEc2      Compatibility = "EC2"
	CompatibilityFargate  Compatibility = "FARGATE"
	CompatibilityExternal Compatibility = "EXTERNAL"
)

// Values returns all known values for Compatibility. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Compatibility) Values() []Compatibility {
	return []Compatibility{
		"EC2",
		"FARGATE",
		"EXTERNAL",
	}
}

type Connectivity string

// Enum values for Connectivity
const (
	ConnectivityConnected    Connectivity = "CONNECTED"
	ConnectivityDisconnected Connectivity = "DISCONNECTED"
)

// Values returns all known values for Connectivity. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Connectivity) Values() []Connectivity {
	return []Connectivity{
		"CONNECTED",
		"DISCONNECTED",
	}
}

type ContainerCondition string

// Enum values for ContainerCondition
const (
	ContainerConditionStart    ContainerCondition = "START"
	ContainerConditionComplete ContainerCondition = "COMPLETE"
	ContainerConditionSuccess  ContainerCondition = "SUCCESS"
	ContainerConditionHealthy  ContainerCondition = "HEALTHY"
)

// Values returns all known values for ContainerCondition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerCondition) Values() []ContainerCondition {
	return []ContainerCondition{
		"START",
		"COMPLETE",
		"SUCCESS",
		"HEALTHY",
	}
}

type ContainerInstanceField string

// Enum values for ContainerInstanceField
const (
	ContainerInstanceFieldTags                    ContainerInstanceField = "TAGS"
	ContainerInstanceFieldContainerInstanceHealth ContainerInstanceField = "CONTAINER_INSTANCE_HEALTH"
)

// Values returns all known values for ContainerInstanceField. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerInstanceField) Values() []ContainerInstanceField {
	return []ContainerInstanceField{
		"TAGS",
		"CONTAINER_INSTANCE_HEALTH",
	}
}

type ContainerInstanceStatus string

// Enum values for ContainerInstanceStatus
const (
	ContainerInstanceStatusActive             ContainerInstanceStatus = "ACTIVE"
	ContainerInstanceStatusDraining           ContainerInstanceStatus = "DRAINING"
	ContainerInstanceStatusRegistering        ContainerInstanceStatus = "REGISTERING"
	ContainerInstanceStatusDeregistering      ContainerInstanceStatus = "DEREGISTERING"
	ContainerInstanceStatusRegistrationFailed ContainerInstanceStatus = "REGISTRATION_FAILED"
)

// Values returns all known values for ContainerInstanceStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContainerInstanceStatus) Values() []ContainerInstanceStatus {
	return []ContainerInstanceStatus{
		"ACTIVE",
		"DRAINING",
		"REGISTERING",
		"DEREGISTERING",
		"REGISTRATION_FAILED",
	}
}

type CPUArchitecture string

// Enum values for CPUArchitecture
const (
	CPUArchitectureX8664 CPUArchitecture = "X86_64"
	CPUArchitectureArm64 CPUArchitecture = "ARM64"
)

// Values returns all known values for CPUArchitecture. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CPUArchitecture) Values() []CPUArchitecture {
	return []CPUArchitecture{
		"X86_64",
		"ARM64",
	}
}

type DeploymentControllerType string

// Enum values for DeploymentControllerType
const (
	DeploymentControllerTypeEcs        DeploymentControllerType = "ECS"
	DeploymentControllerTypeCodeDeploy DeploymentControllerType = "CODE_DEPLOY"
	DeploymentControllerTypeExternal   DeploymentControllerType = "EXTERNAL"
)

// Values returns all known values for DeploymentControllerType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentControllerType) Values() []DeploymentControllerType {
	return []DeploymentControllerType{
		"ECS",
		"CODE_DEPLOY",
		"EXTERNAL",
	}
}

type DeploymentRolloutState string

// Enum values for DeploymentRolloutState
const (
	DeploymentRolloutStateCompleted  DeploymentRolloutState = "COMPLETED"
	DeploymentRolloutStateFailed     DeploymentRolloutState = "FAILED"
	DeploymentRolloutStateInProgress DeploymentRolloutState = "IN_PROGRESS"
)

// Values returns all known values for DeploymentRolloutState. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentRolloutState) Values() []DeploymentRolloutState {
	return []DeploymentRolloutState{
		"COMPLETED",
		"FAILED",
		"IN_PROGRESS",
	}
}

type DesiredStatus string

// Enum values for DesiredStatus
const (
	DesiredStatusRunning DesiredStatus = "RUNNING"
	DesiredStatusPending DesiredStatus = "PENDING"
	DesiredStatusStopped DesiredStatus = "STOPPED"
)

// Values returns all known values for DesiredStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DesiredStatus) Values() []DesiredStatus {
	return []DesiredStatus{
		"RUNNING",
		"PENDING",
		"STOPPED",
	}
}

type DeviceCgroupPermission string

// Enum values for DeviceCgroupPermission
const (
	DeviceCgroupPermissionRead  DeviceCgroupPermission = "read"
	DeviceCgroupPermissionWrite DeviceCgroupPermission = "write"
	DeviceCgroupPermissionMknod DeviceCgroupPermission = "mknod"
)

// Values returns all known values for DeviceCgroupPermission. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceCgroupPermission) Values() []DeviceCgroupPermission {
	return []DeviceCgroupPermission{
		"read",
		"write",
		"mknod",
	}
}

type EBSResourceType string

// Enum values for EBSResourceType
const (
	EBSResourceTypeVolume EBSResourceType = "volume"
)

// Values returns all known values for EBSResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EBSResourceType) Values() []EBSResourceType {
	return []EBSResourceType{
		"volume",
	}
}

type EFSAuthorizationConfigIAM string

// Enum values for EFSAuthorizationConfigIAM
const (
	EFSAuthorizationConfigIAMEnabled  EFSAuthorizationConfigIAM = "ENABLED"
	EFSAuthorizationConfigIAMDisabled EFSAuthorizationConfigIAM = "DISABLED"
)

// Values returns all known values for EFSAuthorizationConfigIAM. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EFSAuthorizationConfigIAM) Values() []EFSAuthorizationConfigIAM {
	return []EFSAuthorizationConfigIAM{
		"ENABLED",
		"DISABLED",
	}
}

type EFSTransitEncryption string

// Enum values for EFSTransitEncryption
const (
	EFSTransitEncryptionEnabled  EFSTransitEncryption = "ENABLED"
	EFSTransitEncryptionDisabled EFSTransitEncryption = "DISABLED"
)

// Values returns all known values for EFSTransitEncryption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EFSTransitEncryption) Values() []EFSTransitEncryption {
	return []EFSTransitEncryption{
		"ENABLED",
		"DISABLED",
	}
}

type EnvironmentFileType string

// Enum values for EnvironmentFileType
const (
	EnvironmentFileTypeS3 EnvironmentFileType = "s3"
)

// Values returns all known values for EnvironmentFileType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentFileType) Values() []EnvironmentFileType {
	return []EnvironmentFileType{
		"s3",
	}
}

type ExecuteCommandLogging string

// Enum values for ExecuteCommandLogging
const (
	ExecuteCommandLoggingNone     ExecuteCommandLogging = "NONE"
	ExecuteCommandLoggingDefault  ExecuteCommandLogging = "DEFAULT"
	ExecuteCommandLoggingOverride ExecuteCommandLogging = "OVERRIDE"
)

// Values returns all known values for ExecuteCommandLogging. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecuteCommandLogging) Values() []ExecuteCommandLogging {
	return []ExecuteCommandLogging{
		"NONE",
		"DEFAULT",
		"OVERRIDE",
	}
}

type FirelensConfigurationType string

// Enum values for FirelensConfigurationType
const (
	FirelensConfigurationTypeFluentd   FirelensConfigurationType = "fluentd"
	FirelensConfigurationTypeFluentbit FirelensConfigurationType = "fluentbit"
)

// Values returns all known values for FirelensConfigurationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (FirelensConfigurationType) Values() []FirelensConfigurationType {
	return []FirelensConfigurationType{
		"fluentd",
		"fluentbit",
	}
}

type HealthStatus string

// Enum values for HealthStatus
const (
	HealthStatusHealthy   HealthStatus = "HEALTHY"
	HealthStatusUnhealthy HealthStatus = "UNHEALTHY"
	HealthStatusUnknown   HealthStatus = "UNKNOWN"
)

// Values returns all known values for HealthStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HealthStatus) Values() []HealthStatus {
	return []HealthStatus{
		"HEALTHY",
		"UNHEALTHY",
		"UNKNOWN",
	}
}

type InstanceHealthCheckState string

// Enum values for InstanceHealthCheckState
const (
	InstanceHealthCheckStateOk               InstanceHealthCheckState = "OK"
	InstanceHealthCheckStateImpaired         InstanceHealthCheckState = "IMPAIRED"
	InstanceHealthCheckStateInsufficientData InstanceHealthCheckState = "INSUFFICIENT_DATA"
	InstanceHealthCheckStateInitializing     InstanceHealthCheckState = "INITIALIZING"
)

// Values returns all known values for InstanceHealthCheckState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceHealthCheckState) Values() []InstanceHealthCheckState {
	return []InstanceHealthCheckState{
		"OK",
		"IMPAIRED",
		"INSUFFICIENT_DATA",
		"INITIALIZING",
	}
}

type InstanceHealthCheckType string

// Enum values for InstanceHealthCheckType
const (
	InstanceHealthCheckTypeContainerRuntime InstanceHealthCheckType = "CONTAINER_RUNTIME"
)

// Values returns all known values for InstanceHealthCheckType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceHealthCheckType) Values() []InstanceHealthCheckType {
	return []InstanceHealthCheckType{
		"CONTAINER_RUNTIME",
	}
}

type IpcMode string

// Enum values for IpcMode
const (
	IpcModeHost IpcMode = "host"
	IpcModeTask IpcMode = "task"
	IpcModeNone IpcMode = "none"
)

// Values returns all known values for IpcMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (IpcMode) Values() []IpcMode {
	return []IpcMode{
		"host",
		"task",
		"none",
	}
}

type LaunchType string

// Enum values for LaunchType
const (
	LaunchTypeEc2      LaunchType = "EC2"
	LaunchTypeFargate  LaunchType = "FARGATE"
	LaunchTypeExternal LaunchType = "EXTERNAL"
)

// Values returns all known values for LaunchType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LaunchType) Values() []LaunchType {
	return []LaunchType{
		"EC2",
		"FARGATE",
		"EXTERNAL",
	}
}

type LogDriver string

// Enum values for LogDriver
const (
	LogDriverJsonFile    LogDriver = "json-file"
	LogDriverSyslog      LogDriver = "syslog"
	LogDriverJournald    LogDriver = "journald"
	LogDriverGelf        LogDriver = "gelf"
	LogDriverFluentd     LogDriver = "fluentd"
	LogDriverAwslogs     LogDriver = "awslogs"
	LogDriverSplunk      LogDriver = "splunk"
	LogDriverAwsfirelens LogDriver = "awsfirelens"
)

// Values returns all known values for LogDriver. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LogDriver) Values() []LogDriver {
	return []LogDriver{
		"json-file",
		"syslog",
		"journald",
		"gelf",
		"fluentd",
		"awslogs",
		"splunk",
		"awsfirelens",
	}
}

type ManagedAgentName string

// Enum values for ManagedAgentName
const (
	ManagedAgentNameExecuteCommandAgent ManagedAgentName = "ExecuteCommandAgent"
)

// Values returns all known values for ManagedAgentName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManagedAgentName) Values() []ManagedAgentName {
	return []ManagedAgentName{
		"ExecuteCommandAgent",
	}
}

type ManagedDraining string

// Enum values for ManagedDraining
const (
	ManagedDrainingEnabled  ManagedDraining = "ENABLED"
	ManagedDrainingDisabled ManagedDraining = "DISABLED"
)

// Values returns all known values for ManagedDraining. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManagedDraining) Values() []ManagedDraining {
	return []ManagedDraining{
		"ENABLED",
		"DISABLED",
	}
}

type ManagedScalingStatus string

// Enum values for ManagedScalingStatus
const (
	ManagedScalingStatusEnabled  ManagedScalingStatus = "ENABLED"
	ManagedScalingStatusDisabled ManagedScalingStatus = "DISABLED"
)

// Values returns all known values for ManagedScalingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ManagedScalingStatus) Values() []ManagedScalingStatus {
	return []ManagedScalingStatus{
		"ENABLED",
		"DISABLED",
	}
}

type ManagedTerminationProtection string

// Enum values for ManagedTerminationProtection
const (
	ManagedTerminationProtectionEnabled  ManagedTerminationProtection = "ENABLED"
	ManagedTerminationProtectionDisabled ManagedTerminationProtection = "DISABLED"
)

// Values returns all known values for ManagedTerminationProtection. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ManagedTerminationProtection) Values() []ManagedTerminationProtection {
	return []ManagedTerminationProtection{
		"ENABLED",
		"DISABLED",
	}
}

type NetworkMode string

// Enum values for NetworkMode
const (
	NetworkModeBridge NetworkMode = "bridge"
	NetworkModeHost   NetworkMode = "host"
	NetworkModeAwsvpc NetworkMode = "awsvpc"
	NetworkModeNone   NetworkMode = "none"
)

// Values returns all known values for NetworkMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (NetworkMode) Values() []NetworkMode {
	return []NetworkMode{
		"bridge",
		"host",
		"awsvpc",
		"none",
	}
}

type OSFamily string

// Enum values for OSFamily
const (
	OSFamilyWindowsServer2019Full OSFamily = "WINDOWS_SERVER_2019_FULL"
	OSFamilyWindowsServer2019Core OSFamily = "WINDOWS_SERVER_2019_CORE"
	OSFamilyWindowsServer2016Full OSFamily = "WINDOWS_SERVER_2016_FULL"
	OSFamilyWindowsServer2004Core OSFamily = "WINDOWS_SERVER_2004_CORE"
	OSFamilyWindowsServer2022Core OSFamily = "WINDOWS_SERVER_2022_CORE"
	OSFamilyWindowsServer2022Full OSFamily = "WINDOWS_SERVER_2022_FULL"
	OSFamilyWindowsServer20h2Core OSFamily = "WINDOWS_SERVER_20H2_CORE"
	OSFamilyLinux                 OSFamily = "LINUX"
)

// Values returns all known values for OSFamily. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OSFamily) Values() []OSFamily {
	return []OSFamily{
		"WINDOWS_SERVER_2019_FULL",
		"WINDOWS_SERVER_2019_CORE",
		"WINDOWS_SERVER_2016_FULL",
		"WINDOWS_SERVER_2004_CORE",
		"WINDOWS_SERVER_2022_CORE",
		"WINDOWS_SERVER_2022_FULL",
		"WINDOWS_SERVER_20H2_CORE",
		"LINUX",
	}
}

type PidMode string

// Enum values for PidMode
const (
	PidModeHost PidMode = "host"
	PidModeTask PidMode = "task"
)

// Values returns all known values for PidMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (PidMode) Values() []PidMode {
	return []PidMode{
		"host",
		"task",
	}
}

type PlacementConstraintType string

// Enum values for PlacementConstraintType
const (
	PlacementConstraintTypeDistinctInstance PlacementConstraintType = "distinctInstance"
	PlacementConstraintTypeMemberOf         PlacementConstraintType = "memberOf"
)

// Values returns all known values for PlacementConstraintType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementConstraintType) Values() []PlacementConstraintType {
	return []PlacementConstraintType{
		"distinctInstance",
		"memberOf",
	}
}

type PlacementStrategyType string

// Enum values for PlacementStrategyType
const (
	PlacementStrategyTypeRandom  PlacementStrategyType = "random"
	PlacementStrategyTypeSpread  PlacementStrategyType = "spread"
	PlacementStrategyTypeBinpack PlacementStrategyType = "binpack"
)

// Values returns all known values for PlacementStrategyType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlacementStrategyType) Values() []PlacementStrategyType {
	return []PlacementStrategyType{
		"random",
		"spread",
		"binpack",
	}
}

type PlatformDeviceType string

// Enum values for PlatformDeviceType
const (
	PlatformDeviceTypeGpu PlatformDeviceType = "GPU"
)

// Values returns all known values for PlatformDeviceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PlatformDeviceType) Values() []PlatformDeviceType {
	return []PlatformDeviceType{
		"GPU",
	}
}

type PropagateTags string

// Enum values for PropagateTags
const (
	PropagateTagsTaskDefinition PropagateTags = "TASK_DEFINITION"
	PropagateTagsService        PropagateTags = "SERVICE"
	PropagateTagsNone           PropagateTags = "NONE"
)

// Values returns all known values for PropagateTags. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PropagateTags) Values() []PropagateTags {
	return []PropagateTags{
		"TASK_DEFINITION",
		"SERVICE",
		"NONE",
	}
}

type ProxyConfigurationType string

// Enum values for ProxyConfigurationType
const (
	ProxyConfigurationTypeAppmesh ProxyConfigurationType = "APPMESH"
)

// Values returns all known values for ProxyConfigurationType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProxyConfigurationType) Values() []ProxyConfigurationType {
	return []ProxyConfigurationType{
		"APPMESH",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGpu                  ResourceType = "GPU"
	ResourceTypeInferenceAccelerator ResourceType = "InferenceAccelerator"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"GPU",
		"InferenceAccelerator",
	}
}

type ScaleUnit string

// Enum values for ScaleUnit
const (
	ScaleUnitPercent ScaleUnit = "PERCENT"
)

// Values returns all known values for ScaleUnit. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScaleUnit) Values() []ScaleUnit {
	return []ScaleUnit{
		"PERCENT",
	}
}

type SchedulingStrategy string

// Enum values for SchedulingStrategy
const (
	SchedulingStrategyReplica SchedulingStrategy = "REPLICA"
	SchedulingStrategyDaemon  SchedulingStrategy = "DAEMON"
)

// Values returns all known values for SchedulingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SchedulingStrategy) Values() []SchedulingStrategy {
	return []SchedulingStrategy{
		"REPLICA",
		"DAEMON",
	}
}

type Scope string

// Enum values for Scope
const (
	ScopeTask   Scope = "task"
	ScopeShared Scope = "shared"
)

// Values returns all known values for Scope. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Scope) Values() []Scope {
	return []Scope{
		"task",
		"shared",
	}
}

type ServiceField string

// Enum values for ServiceField
const (
	ServiceFieldTags ServiceField = "TAGS"
)

// Values returns all known values for ServiceField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceField) Values() []ServiceField {
	return []ServiceField{
		"TAGS",
	}
}

type SettingName string

// Enum values for SettingName
const (
	SettingNameServiceLongArnFormat            SettingName = "serviceLongArnFormat"
	SettingNameTaskLongArnFormat               SettingName = "taskLongArnFormat"
	SettingNameContainerInstanceLongArnFormat  SettingName = "containerInstanceLongArnFormat"
	SettingNameAwsvpcTrunking                  SettingName = "awsvpcTrunking"
	SettingNameContainerInsights               SettingName = "containerInsights"
	SettingNameFargateFipsMode                 SettingName = "fargateFIPSMode"
	SettingNameTagResourceAuthorization        SettingName = "tagResourceAuthorization"
	SettingNameFargateTaskRetirementWaitPeriod SettingName = "fargateTaskRetirementWaitPeriod"
	SettingNameGuardDutyActivate               SettingName = "guardDutyActivate"
)

// Values returns all known values for SettingName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SettingName) Values() []SettingName {
	return []SettingName{
		"serviceLongArnFormat",
		"taskLongArnFormat",
		"containerInstanceLongArnFormat",
		"awsvpcTrunking",
		"containerInsights",
		"fargateFIPSMode",
		"tagResourceAuthorization",
		"fargateTaskRetirementWaitPeriod",
		"guardDutyActivate",
	}
}

type SettingType string

// Enum values for SettingType
const (
	SettingTypeUser       SettingType = "user"
	SettingTypeAwsManaged SettingType = "aws_managed"
)

// Values returns all known values for SettingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SettingType) Values() []SettingType {
	return []SettingType{
		"user",
		"aws_managed",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASC",
		"DESC",
	}
}

type StabilityStatus string

// Enum values for StabilityStatus
const (
	StabilityStatusSteadyState StabilityStatus = "STEADY_STATE"
	StabilityStatusStabilizing StabilityStatus = "STABILIZING"
)

// Values returns all known values for StabilityStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StabilityStatus) Values() []StabilityStatus {
	return []StabilityStatus{
		"STEADY_STATE",
		"STABILIZING",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeContainerInstance TargetType = "container-instance"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"container-instance",
	}
}

type TaskDefinitionFamilyStatus string

// Enum values for TaskDefinitionFamilyStatus
const (
	TaskDefinitionFamilyStatusActive   TaskDefinitionFamilyStatus = "ACTIVE"
	TaskDefinitionFamilyStatusInactive TaskDefinitionFamilyStatus = "INACTIVE"
	TaskDefinitionFamilyStatusAll      TaskDefinitionFamilyStatus = "ALL"
)

// Values returns all known values for TaskDefinitionFamilyStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionFamilyStatus) Values() []TaskDefinitionFamilyStatus {
	return []TaskDefinitionFamilyStatus{
		"ACTIVE",
		"INACTIVE",
		"ALL",
	}
}

type TaskDefinitionField string

// Enum values for TaskDefinitionField
const (
	TaskDefinitionFieldTags TaskDefinitionField = "TAGS"
)

// Values returns all known values for TaskDefinitionField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionField) Values() []TaskDefinitionField {
	return []TaskDefinitionField{
		"TAGS",
	}
}

type TaskDefinitionPlacementConstraintType string

// Enum values for TaskDefinitionPlacementConstraintType
const (
	TaskDefinitionPlacementConstraintTypeMemberOf TaskDefinitionPlacementConstraintType = "memberOf"
)

// Values returns all known values for TaskDefinitionPlacementConstraintType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TaskDefinitionPlacementConstraintType) Values() []TaskDefinitionPlacementConstraintType {
	return []TaskDefinitionPlacementConstraintType{
		"memberOf",
	}
}

type TaskDefinitionStatus string

// Enum values for TaskDefinitionStatus
const (
	TaskDefinitionStatusActive           TaskDefinitionStatus = "ACTIVE"
	TaskDefinitionStatusInactive         TaskDefinitionStatus = "INACTIVE"
	TaskDefinitionStatusDeleteInProgress TaskDefinitionStatus = "DELETE_IN_PROGRESS"
)

// Values returns all known values for TaskDefinitionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskDefinitionStatus) Values() []TaskDefinitionStatus {
	return []TaskDefinitionStatus{
		"ACTIVE",
		"INACTIVE",
		"DELETE_IN_PROGRESS",
	}
}

type TaskField string

// Enum values for TaskField
const (
	TaskFieldTags TaskField = "TAGS"
)

// Values returns all known values for TaskField. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskField) Values() []TaskField {
	return []TaskField{
		"TAGS",
	}
}

type TaskFilesystemType string

// Enum values for TaskFilesystemType
const (
	TaskFilesystemTypeExt3 TaskFilesystemType = "ext3"
	TaskFilesystemTypeExt4 TaskFilesystemType = "ext4"
	TaskFilesystemTypeXfs  TaskFilesystemType = "xfs"
)

// Values returns all known values for TaskFilesystemType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskFilesystemType) Values() []TaskFilesystemType {
	return []TaskFilesystemType{
		"ext3",
		"ext4",
		"xfs",
	}
}

type TaskSetField string

// Enum values for TaskSetField
const (
	TaskSetFieldTags TaskSetField = "TAGS"
)

// Values returns all known values for TaskSetField. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskSetField) Values() []TaskSetField {
	return []TaskSetField{
		"TAGS",
	}
}

type TaskStopCode string

// Enum values for TaskStopCode
const (
	TaskStopCodeTaskFailedToStart         TaskStopCode = "TaskFailedToStart"
	TaskStopCodeEssentialContainerExited  TaskStopCode = "EssentialContainerExited"
	TaskStopCodeUserInitiated             TaskStopCode = "UserInitiated"
	TaskStopCodeServiceSchedulerInitiated TaskStopCode = "ServiceSchedulerInitiated"
	TaskStopCodeSpotInterruption          TaskStopCode = "SpotInterruption"
	TaskStopCodeTerminationNotice         TaskStopCode = "TerminationNotice"
)

// Values returns all known values for TaskStopCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TaskStopCode) Values() []TaskStopCode {
	return []TaskStopCode{
		"TaskFailedToStart",
		"EssentialContainerExited",
		"UserInitiated",
		"ServiceSchedulerInitiated",
		"SpotInterruption",
		"TerminationNotice",
	}
}

type TransportProtocol string

// Enum values for TransportProtocol
const (
	TransportProtocolTcp TransportProtocol = "tcp"
	TransportProtocolUdp TransportProtocol = "udp"
)

// Values returns all known values for TransportProtocol. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransportProtocol) Values() []TransportProtocol {
	return []TransportProtocol{
		"tcp",
		"udp",
	}
}

type UlimitName string

// Enum values for UlimitName
const (
	UlimitNameCore       UlimitName = "core"
	UlimitNameCpu        UlimitName = "cpu"
	UlimitNameData       UlimitName = "data"
	UlimitNameFsize      UlimitName = "fsize"
	UlimitNameLocks      UlimitName = "locks"
	UlimitNameMemlock    UlimitName = "memlock"
	UlimitNameMsgqueue   UlimitName = "msgqueue"
	UlimitNameNice       UlimitName = "nice"
	UlimitNameNofile     UlimitName = "nofile"
	UlimitNameNproc      UlimitName = "nproc"
	UlimitNameRss        UlimitName = "rss"
	UlimitNameRtprio     UlimitName = "rtprio"
	UlimitNameRttime     UlimitName = "rttime"
	UlimitNameSigpending UlimitName = "sigpending"
	UlimitNameStack      UlimitName = "stack"
)

// Values returns all known values for UlimitName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UlimitName) Values() []UlimitName {
	return []UlimitName{
		"core",
		"cpu",
		"data",
		"fsize",
		"locks",
		"memlock",
		"msgqueue",
		"nice",
		"nofile",
		"nproc",
		"rss",
		"rtprio",
		"rttime",
		"sigpending",
		"stack",
	}
}
