// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	OnecloudClusterResourceKind   = "OnecloudCluster"
	OnecloudClusterResourcePlural = "onecloudclusters"
)

var (
	OnecloudClusterCRDName = OnecloudClusterResourcePlural + "." + GroupName
)

// ComponentType represents component type
type ComponentType string

const (
	// KeystoneComponentType is keystone component type
	KeystoneComponentType ComponentType = "keystone"
	// RegionComponentType is region component type
	RegionComponentType ComponentType = "region"
	// ClimcComponentType is climc component type
	ClimcComponentType ComponentType = "climc"
	// GlanceComponentType is glance component type
	GlanceComponentType ComponentType = "glance"
	// WebconsoleComponentType is webconsole component type
	WebconsoleComponentType ComponentType = "webconsole"
	// SchedulerComponentType is scheduler component type
	SchedulerComponentType ComponentType = "scheduler"
	// LogComponentType is logger service component type
	LoggerComponentType ComponentType = "logger"
	// InfluxdbComponentType is influxdb component type
	InfluxdbComponentType ComponentType = "influxdb"
	// APIGatewayComponentType is apiGateway component type
	APIGatewayComponentType ComponentType = "apigateway"
	// WebComponentType is web frontent component type
	WebComponentType         ComponentType = "web"
	YunionagentComponentType ComponentType = "yunionagent"
	YunionconfComponentType  ComponentType = "yunionconf"
	KubeServerComponentType  ComponentType = "kubeserver"
)

// ComponentPhase is the current state of component
type ComponentPhase string

const (
	// NormalPhase represents normal state of OneCloud cluster.
	NormalPhase ComponentPhase = "Normal"
	// UpgradePhase represents the upgrade state of Onecloud cluster.
	UpgradePhase ComponentPhase = "Upgrade"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnecloudCluster defines the cluster
type OnecloudCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// Spec defines the behavior of a onecloud cluster
	Spec OnecloudClusterSpec `json:"spec"`

	// Most recently observed status of the onecloud cluster
	Status OnecloudClusterStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnecloudClusterList represents a list of Onecloud Clusters
type OnecloudClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnecloudCluster `json:"items"`
}

// OnecloudClusterSpec describes the attributes that a user creates on a onecloud cluster
type OnecloudClusterSpec struct {
	// Etcd holds configuration for etcd
	Etcd Etcd `json:"etcd,omitempty"`
	// Mysql holds configuration for mysql
	Mysql Mysql `json:"mysql"`
	// Version is onecloud components version
	Version string `json:"version"`
	// CertSANs sets extra Subject Alternative Names for the Cluster signing cert.
	CertSANs []string
	// Services list non-headless services type used in OnecloudCluster
	Services []Service `json:"services,omitempty"`
	// ImageRepository defines default image registry
	ImageRepository string `json:"imageRepository"`
	// Region is cluster region
	Region string `json:"region"`
	// Zone is cluster first zone
	Zone string `json:"zone"`
	// Keystone holds configuration for keystone
	Keystone KeystoneSpec `json:"keystone"`
	// RegionServer holds configuration for region
	RegionServer RegionSpec `json:"regionServer"`
	// Scheduler holds configuration for scheduler
	Scheduler DeploymentSpec `json:"scheduler"`
	// Glance holds configuration for glance
	Glance StatefulDeploymentSpec `json:"glance"`
	// Climc holds configuration for climc
	Climc DeploymentSpec `json:"climc"`
	// Webconsole holds configuration for webconsole
	Webconsole DeploymentSpec `json:"webconsole"`
	// Logger holds configuration for log service
	Logger DeploymentSpec `json:"logger"`
	// Yunionconf holds configuration for yunionconf service
	Yunionconf DeploymentSpec `json:"yunionconf"`
	// Yunionagent holds configuration for yunionagent service
	Yunionagent StatefulDeploymentSpec `json:"yunionagent"`
	// Influxdb holds configuration for influxdb
	Influxdb StatefulDeploymentSpec `json:"influxdb"`
	// LoadBalancerEndpoint is upstream loadbalancer virtual ip address or DNS domain
	LoadBalancerEndpoint string `json:"loadBalancerEndpoint"`
	// APIGateway holds configuration for yunoinapi
	APIGateway DeploymentSpec `json:"apiGateway"`
	// Web holds configuration for web
	Web DeploymentSpec `json:"web"`
	// KubeServer holds configuration for kube-server service
	KubeServer DeploymentSpec `json:"kubeserver"`
}

// OnecloudClusterStatus
type OnecloudClusterStatus struct {
	ClusterID    string           `json:"clusterID,omitempty"`
	Keystone     KeystoneStatus   `json:"keystone,omitempty"`
	RegionServer RegionStatus     `json:"region,omitempty"`
	Glance       GlanceStatus     `json:"glance,omitempty"`
	Scheduler    DeploymentStatus `json:"scheduler,omitempty"`
	Webconsole   DeploymentStatus `json:"webconsole,omitempty"`
	Influxdb     DeploymentStatus `json:"influxdb,omitempty"`
	Logger       DeploymentStatus `json:"logger,omitempty"`
	APIGateway   DeploymentStatus `json:"apiGateway,omitempty"`
	Web          DeploymentStatus `json:"web,omitempty"`
	Yunionconf   DeploymentStatus `json:"yunionconf,omitempty"`
	Yunionagent  DeploymentStatus `json:"yunionagent,omitempty"`
	KubeServer   DeploymentStatus `json:"kubeserver,omitempty"`
}

// Etcd describes an etcd cluster
type Etcd struct {
	// Endpoints of etcd members
	Endpoints []string `json:"endpoints"`
	// CA is an SSL Certificate Authority data used to secure etcd communication
	CA string `json:"ca"`
	// CertFile is an SSL certification data used to secure etcd communication
	Cert string `json:"cert"`
	// Key is an SSL key data used to secure etcd communication
	Key string `json:"key"`
}

// Mysql describes an mysql server
type Mysql struct {
	// Host is mysql ip address of hostname
	Host string `json:"host"`
	// Port is mysql listen port
	Port int32 `json:"port"`
	// Username is mysql username
	Username string `json:"username"`
	// Password is mysql user password
	Password string `json:"password"`
}

// DeploymentSpec constains defails of deployment resource service
type DeploymentSpec struct {
	ContainerSpec
	Disable      bool                `json:"disable"`
	Replicas     int32               `json:"replicas"`
	Affinity     *corev1.Affinity    `json:"affinity,omitempty"`
	NodeSelector map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations  []corev1.Toleration `json:"tolerations,omitempty"`
	Annotations  map[string]string   `json:"annotations,omitempty"`
}

type StatefulDeploymentSpec struct {
	DeploymentSpec
	StorageClassName string `json:"storageClassName,omitempty"`
}

// KeystoneSpec contains details of keystone service
type KeystoneSpec struct {
	DeploymentSpec
	BootstrapPassword string `json:"bootstrapPassword"`
}

type DeploymentStatus struct {
	Phase      ComponentPhase         `json:"phase,omitempty"`
	Deployment *apps.DeploymentStatus `json:"deployment,omitempty"`
}

// KeystoneStatus is Keystone status
type KeystoneStatus struct {
	DeploymentStatus
}

type RegionStatus struct {
	DeploymentStatus
}

type GlanceStatus struct {
	DeploymentStatus
}

type WebconsoleStatus struct {
	DeploymentStatus
}

type RegionSpec struct {
	DeploymentSpec
}

// ContainerSpec is the container spec of a pod
type ContainerSpec struct {
	Image           string               `json:"image"`
	Repository      string               `json:"repository,omitempty"`
	ImageName       string               `json:"imageName,omitempty"`
	Tag             string               `json:"tag,omitempty"`
	ImagePullPolicy corev1.PullPolicy    `json:"imagePullPolicy,omitempty"`
	Requests        *ResourceRequirement `json:"requests,omitempty"`
	Limits          *ResourceRequirement `json:"limits,omitempty"`
}

// Service represent service type used in OnecloudCluster
type Service struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type ResourceRequirement struct {
	// CPU is how many cores a pod requires
	CPU string `json:"cpu,omitempty"`
	// Memory is how much memory a pod requires
	Memory string `json:"memory,omitempty"`
	// Storage is storage size a pod requires
	Storage string `json:"storage,omitempty"`
}

type DBConfig struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CloudUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ServiceBaseConfig struct {
	Port int `json:"port"`
}

type ServiceCommonOptions struct {
	ServiceBaseConfig

	CloudUser
}

type KeystoneConfig struct {
	ServiceBaseConfig

	DB DBConfig `json:"db"`
}

type ServiceDBCommonOptions struct {
	ServiceCommonOptions

	DB DBConfig `json:"db"`
}

type RegionConfig struct {
	ServiceDBCommonOptions
}

type GlanceConfig struct {
	ServiceDBCommonOptions
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OnecloudClusterConfig struct {
	metav1.TypeMeta

	Keystone     KeystoneConfig         `json:"keystone"`
	RegionServer RegionConfig           `json:"region"`
	Glance       GlanceConfig           `json:"glance"`
	Webconsole   ServiceCommonOptions   `json:"webconsole"`
	Logger       ServiceDBCommonOptions `json:"logger"`
	Yunionconf   ServiceDBCommonOptions `json:"yunionconf"`
	Yunionagent  ServiceDBCommonOptions `json:"yunionagent"`
	KubeServer   ServiceDBCommonOptions `json:"kubeserver"`
	APIGateway   ServiceCommonOptions   `json:"apiGateway"`
}
