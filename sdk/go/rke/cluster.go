// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rke

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides RKE cluster resource. This can be used to create RKE clusters and retrieve their information.
//
// > This content is derived from https://github.com/rancher/terraform-provider-rke/blob/master/website/docs/r/cluster.html.markdown.
type Cluster struct {
	s *pulumi.ResourceState
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["addonJobTimeout"] = nil
		inputs["addons"] = nil
		inputs["addonsIncludes"] = nil
		inputs["authentication"] = nil
		inputs["authorization"] = nil
		inputs["bastionHost"] = nil
		inputs["certDir"] = nil
		inputs["cloudProvider"] = nil
		inputs["clusterName"] = nil
		inputs["customCerts"] = nil
		inputs["delayOnCreation"] = nil
		inputs["dind"] = nil
		inputs["dindDnsServer"] = nil
		inputs["dindStorageDriver"] = nil
		inputs["disablePortCheck"] = nil
		inputs["dns"] = nil
		inputs["ignoreDockerVersion"] = nil
		inputs["ingress"] = nil
		inputs["kubernetesVersion"] = nil
		inputs["monitoring"] = nil
		inputs["network"] = nil
		inputs["nodes"] = nil
		inputs["nodesConfs"] = nil
		inputs["prefixPath"] = nil
		inputs["privateRegistries"] = nil
		inputs["restore"] = nil
		inputs["rotateCertificates"] = nil
		inputs["services"] = nil
		inputs["Services_Etcd"] = nil
		inputs["Services_KubeApi"] = nil
		inputs["Services_KubeController"] = nil
		inputs["Services_Kubelet"] = nil
		inputs["Services_KubeProxy"] = nil
		inputs["Services_KubeScheduler"] = nil
		inputs["sshAgentAuth"] = nil
		inputs["sshCertPath"] = nil
		inputs["sshKeyPath"] = nil
		inputs["systemImages"] = nil
		inputs["updateOnly"] = nil
	} else {
		inputs["addonJobTimeout"] = args.AddonJobTimeout
		inputs["addons"] = args.Addons
		inputs["addonsIncludes"] = args.AddonsIncludes
		inputs["authentication"] = args.Authentication
		inputs["authorization"] = args.Authorization
		inputs["bastionHost"] = args.BastionHost
		inputs["certDir"] = args.CertDir
		inputs["cloudProvider"] = args.CloudProvider
		inputs["clusterName"] = args.ClusterName
		inputs["customCerts"] = args.CustomCerts
		inputs["delayOnCreation"] = args.DelayOnCreation
		inputs["dind"] = args.Dind
		inputs["dindDnsServer"] = args.DindDnsServer
		inputs["dindStorageDriver"] = args.DindStorageDriver
		inputs["disablePortCheck"] = args.DisablePortCheck
		inputs["dns"] = args.Dns
		inputs["ignoreDockerVersion"] = args.IgnoreDockerVersion
		inputs["ingress"] = args.Ingress
		inputs["kubernetesVersion"] = args.KubernetesVersion
		inputs["monitoring"] = args.Monitoring
		inputs["network"] = args.Network
		inputs["nodes"] = args.Nodes
		inputs["nodesConfs"] = args.NodesConfs
		inputs["prefixPath"] = args.PrefixPath
		inputs["privateRegistries"] = args.PrivateRegistries
		inputs["restore"] = args.Restore
		inputs["rotateCertificates"] = args.RotateCertificates
		inputs["services"] = args.Services
		inputs["Services_Etcd"] = args.Services_Etcd
		inputs["Services_KubeApi"] = args.Services_KubeApi
		inputs["Services_KubeController"] = args.Services_KubeController
		inputs["Services_Kubelet"] = args.Services_Kubelet
		inputs["Services_KubeProxy"] = args.Services_KubeProxy
		inputs["Services_KubeScheduler"] = args.Services_KubeScheduler
		inputs["sshAgentAuth"] = args.SshAgentAuth
		inputs["sshCertPath"] = args.SshCertPath
		inputs["sshKeyPath"] = args.SshKeyPath
		inputs["systemImages"] = args.SystemImages
		inputs["updateOnly"] = args.UpdateOnly
	}
	inputs["apiServerUrl"] = nil
	inputs["caCrt"] = nil
	inputs["certificates"] = nil
	inputs["clientCert"] = nil
	inputs["clientKey"] = nil
	inputs["clusterCidr"] = nil
	inputs["clusterDnsServer"] = nil
	inputs["clusterDomain"] = nil
	inputs["controlPlaneHosts"] = nil
	inputs["etcdHosts"] = nil
	inputs["inactiveHosts"] = nil
	inputs["internalKubeConfigYaml"] = nil
	inputs["kubeAdminUser"] = nil
	inputs["kubeConfigYaml"] = nil
	inputs["rkeClusterYaml"] = nil
	inputs["rkeState"] = nil
	inputs["runningSystemImages"] = nil
	inputs["workerHosts"] = nil
	s, err := ctx.RegisterResource("rke:index/cluster:Cluster", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ClusterState, opts ...pulumi.ResourceOpt) (*Cluster, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["addonJobTimeout"] = state.AddonJobTimeout
		inputs["addons"] = state.Addons
		inputs["addonsIncludes"] = state.AddonsIncludes
		inputs["apiServerUrl"] = state.ApiServerUrl
		inputs["authentication"] = state.Authentication
		inputs["authorization"] = state.Authorization
		inputs["bastionHost"] = state.BastionHost
		inputs["caCrt"] = state.CaCrt
		inputs["certDir"] = state.CertDir
		inputs["certificates"] = state.Certificates
		inputs["clientCert"] = state.ClientCert
		inputs["clientKey"] = state.ClientKey
		inputs["cloudProvider"] = state.CloudProvider
		inputs["clusterCidr"] = state.ClusterCidr
		inputs["clusterDnsServer"] = state.ClusterDnsServer
		inputs["clusterDomain"] = state.ClusterDomain
		inputs["clusterName"] = state.ClusterName
		inputs["controlPlaneHosts"] = state.ControlPlaneHosts
		inputs["customCerts"] = state.CustomCerts
		inputs["delayOnCreation"] = state.DelayOnCreation
		inputs["dind"] = state.Dind
		inputs["dindDnsServer"] = state.DindDnsServer
		inputs["dindStorageDriver"] = state.DindStorageDriver
		inputs["disablePortCheck"] = state.DisablePortCheck
		inputs["dns"] = state.Dns
		inputs["etcdHosts"] = state.EtcdHosts
		inputs["ignoreDockerVersion"] = state.IgnoreDockerVersion
		inputs["inactiveHosts"] = state.InactiveHosts
		inputs["ingress"] = state.Ingress
		inputs["internalKubeConfigYaml"] = state.InternalKubeConfigYaml
		inputs["kubeAdminUser"] = state.KubeAdminUser
		inputs["kubeConfigYaml"] = state.KubeConfigYaml
		inputs["kubernetesVersion"] = state.KubernetesVersion
		inputs["monitoring"] = state.Monitoring
		inputs["network"] = state.Network
		inputs["nodes"] = state.Nodes
		inputs["nodesConfs"] = state.NodesConfs
		inputs["prefixPath"] = state.PrefixPath
		inputs["privateRegistries"] = state.PrivateRegistries
		inputs["restore"] = state.Restore
		inputs["rkeClusterYaml"] = state.RkeClusterYaml
		inputs["rkeState"] = state.RkeState
		inputs["rotateCertificates"] = state.RotateCertificates
		inputs["runningSystemImages"] = state.RunningSystemImages
		inputs["services"] = state.Services
		inputs["Services_Etcd"] = state.Services_Etcd
		inputs["Services_KubeApi"] = state.Services_KubeApi
		inputs["Services_KubeController"] = state.Services_KubeController
		inputs["Services_Kubelet"] = state.Services_Kubelet
		inputs["Services_KubeProxy"] = state.Services_KubeProxy
		inputs["Services_KubeScheduler"] = state.Services_KubeScheduler
		inputs["sshAgentAuth"] = state.SshAgentAuth
		inputs["sshCertPath"] = state.SshCertPath
		inputs["sshKeyPath"] = state.SshKeyPath
		inputs["systemImages"] = state.SystemImages
		inputs["updateOnly"] = state.UpdateOnly
		inputs["workerHosts"] = state.WorkerHosts
	}
	s, err := ctx.ReadResource("rke:index/cluster:Cluster", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Cluster{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Cluster) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Cluster) ID() pulumi.IDOutput {
	return r.s.ID()
}

// RKE k8s cluster addon deployment timeout in seconds for status check (int)
func (r *Cluster) AddonJobTimeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["addonJobTimeout"])
}

// RKE k8s cluster user addons YAML manifest to be deployed (string)
func (r *Cluster) Addons() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["addons"])
}

// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
func (r *Cluster) AddonsIncludes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["addonsIncludes"])
}

// (Computed) RKE k8s cluster api server url (string)
func (r *Cluster) ApiServerUrl() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["apiServerUrl"])
}

// RKE k8s cluster authentication configuration (list maxitems:1)
func (r *Cluster) Authentication() pulumi.Output {
	return r.s.State["authentication"]
}

// RKE k8s cluster authorization mode configuration (list maxitems:1)
func (r *Cluster) Authorization() pulumi.Output {
	return r.s.State["authorization"]
}

// RKE k8s cluster bastion Host configuration (list maxitems:1)
func (r *Cluster) BastionHost() pulumi.Output {
	return r.s.State["bastionHost"]
}

// (Computed/Sensitive) RKE k8s cluster CA certificate (string)
func (r *Cluster) CaCrt() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["caCrt"])
}

// Specify a certificate dir path (string)
func (r *Cluster) CertDir() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["certDir"])
}

// (Computed/Sensitive) RKE k8s cluster certificates (string)
func (r *Cluster) Certificates() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["certificates"])
}

// (Computed/Sensitive) RKE k8s cluster client certificate (string)
func (r *Cluster) ClientCert() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientCert"])
}

// (Computed/Sensitive) RKE k8s cluster client key (string)
func (r *Cluster) ClientKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clientKey"])
}

// Calico cloud provider (string)
func (r *Cluster) CloudProvider() pulumi.Output {
	return r.s.State["cloudProvider"]
}

// Cluster CIDR option for kube controller service (string)
func (r *Cluster) ClusterCidr() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterCidr"])
}

// Cluster DNS Server option for kubelet service (string)
func (r *Cluster) ClusterDnsServer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterDnsServer"])
}

// Cluster Domain option for kubelet service. Default `cluster.local` (string)
func (r *Cluster) ClusterDomain() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterDomain"])
}

// RKE k8s cluster name used in the kube config (string)
func (r *Cluster) ClusterName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["clusterName"])
}

// (Computed) RKE k8s cluster control plane nodes (list)
func (r *Cluster) ControlPlaneHosts() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["controlPlaneHosts"])
}

// Use custom certificates from a cert dir (string)
func (r *Cluster) CustomCerts() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["customCerts"])
}

// RKE k8s cluster delay on creation (int)
func (r *Cluster) DelayOnCreation() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["delayOnCreation"])
}

// Deploy RKE cluster on a dind environment. Default: `false` (bool)
func (r *Cluster) Dind() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["dind"])
}

// DinD RKE cluster dns (string)
func (r *Cluster) DindDnsServer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dindDnsServer"])
}

// DinD RKE cluster storage driver (string)
func (r *Cluster) DindStorageDriver() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dindStorageDriver"])
}

// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
func (r *Cluster) DisablePortCheck() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disablePortCheck"])
}

// RKE k8s cluster DNS Config (list maxitems:1)
func (r *Cluster) Dns() pulumi.Output {
	return r.s.State["dns"]
}

// (Computed) RKE k8s cluster etcd nodes (list)
func (r *Cluster) EtcdHosts() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["etcdHosts"])
}

// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
func (r *Cluster) IgnoreDockerVersion() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["ignoreDockerVersion"])
}

// (Computed) RKE k8s cluster inactive nodes (list)
func (r *Cluster) InactiveHosts() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["inactiveHosts"])
}

// Docker image for ingress (string)
func (r *Cluster) Ingress() pulumi.Output {
	return r.s.State["ingress"]
}

// (Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)
func (r *Cluster) InternalKubeConfigYaml() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["internalKubeConfigYaml"])
}

// (Computed) RKE k8s cluster admin user (string)
func (r *Cluster) KubeAdminUser() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kubeAdminUser"])
}

// (Computed/Sensitive) RKE k8s cluster kube config yaml (string)
func (r *Cluster) KubeConfigYaml() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kubeConfigYaml"])
}

// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
func (r *Cluster) KubernetesVersion() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kubernetesVersion"])
}

// RKE k8s cluster monitoring Config (list maxitems:1)
func (r *Cluster) Monitoring() pulumi.Output {
	return r.s.State["monitoring"]
}

// (list maxitems:1)
func (r *Cluster) Network() pulumi.Output {
	return r.s.State["network"]
}

// RKE k8s cluster nodes (list)
func (r *Cluster) Nodes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["nodes"])
}

// RKE k8s cluster nodes (YAML | JSON)
func (r *Cluster) NodesConfs() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["nodesConfs"])
}

// RKE k8s directory path (string)
func (r *Cluster) PrefixPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["prefixPath"])
}

// RKE k8s cluster private docker registries (list)
func (r *Cluster) PrivateRegistries() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["privateRegistries"])
}

// Restore cluster. Default `false` (bool)
func (r *Cluster) Restore() pulumi.Output {
	return r.s.State["restore"]
}

// (Computed/Sensitive) RKE k8s cluster config yaml (string)
func (r *Cluster) RkeClusterYaml() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rkeClusterYaml"])
}

// (Computed/Sensitive) RKE k8s cluster state (string)
func (r *Cluster) RkeState() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["rkeState"])
}

// RKE k8s cluster rotate certificates configuration (list maxitems:1)
func (r *Cluster) RotateCertificates() pulumi.Output {
	return r.s.State["rotateCertificates"]
}

// (Computed) RKE k8s cluster running system images list (list)
func (r *Cluster) RunningSystemImages() pulumi.Output {
	return r.s.State["runningSystemImages"]
}

// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
func (r *Cluster) Services() pulumi.Output {
	return r.s.State["services"]
}

// Use services.etcd instead (list maxitems:1)
func (r *Cluster) Services_Etcd() pulumi.Output {
	return r.s.State["Services_Etcd"]
}

// Use services.kube_api instead (list maxitems:1)
func (r *Cluster) Services_KubeApi() pulumi.Output {
	return r.s.State["Services_KubeApi"]
}

// Use services.kube_controller instead (list maxitems:1)
func (r *Cluster) Services_KubeController() pulumi.Output {
	return r.s.State["Services_KubeController"]
}

// Use services.kubelet instead (list maxitems:1)
func (r *Cluster) Services_Kubelet() pulumi.Output {
	return r.s.State["Services_Kubelet"]
}

// Use services.kubeproxy instead (list maxitems:1)
func (r *Cluster) Services_KubeProxy() pulumi.Output {
	return r.s.State["Services_KubeProxy"]
}

// Use services.scheduler instead (list maxitems:1)
func (r *Cluster) Services_KubeScheduler() pulumi.Output {
	return r.s.State["Services_KubeScheduler"]
}

// SSH Agent Auth enable (bool)
func (r *Cluster) SshAgentAuth() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["sshAgentAuth"])
}

// SSH Certificate path (string)
func (r *Cluster) SshCertPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sshCertPath"])
}

// SSH Private Key path (string)
func (r *Cluster) SshKeyPath() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sshKeyPath"])
}

// RKE k8s cluster system images list (list maxitems:1)
func (r *Cluster) SystemImages() pulumi.Output {
	return r.s.State["systemImages"]
}

// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
func (r *Cluster) UpdateOnly() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["updateOnly"])
}

// (Computed) RKE k8s cluster worker nodes (list)
func (r *Cluster) WorkerHosts() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["workerHosts"])
}

// Input properties used for looking up and filtering Cluster resources.
type ClusterState struct {
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout interface{}
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons interface{}
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes interface{}
	// (Computed) RKE k8s cluster api server url (string)
	ApiServerUrl interface{}
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication interface{}
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization interface{}
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost interface{}
	// (Computed/Sensitive) RKE k8s cluster CA certificate (string)
	CaCrt interface{}
	// Specify a certificate dir path (string)
	CertDir interface{}
	// (Computed/Sensitive) RKE k8s cluster certificates (string)
	Certificates interface{}
	// (Computed/Sensitive) RKE k8s cluster client certificate (string)
	ClientCert interface{}
	// (Computed/Sensitive) RKE k8s cluster client key (string)
	ClientKey interface{}
	// Calico cloud provider (string)
	CloudProvider interface{}
	// Cluster CIDR option for kube controller service (string)
	ClusterCidr interface{}
	// Cluster DNS Server option for kubelet service (string)
	ClusterDnsServer interface{}
	// Cluster Domain option for kubelet service. Default `cluster.local` (string)
	ClusterDomain interface{}
	// RKE k8s cluster name used in the kube config (string)
	ClusterName interface{}
	// (Computed) RKE k8s cluster control plane nodes (list)
	ControlPlaneHosts interface{}
	// Use custom certificates from a cert dir (string)
	CustomCerts interface{}
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation interface{}
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind interface{}
	// DinD RKE cluster dns (string)
	DindDnsServer interface{}
	// DinD RKE cluster storage driver (string)
	DindStorageDriver interface{}
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck interface{}
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns interface{}
	// (Computed) RKE k8s cluster etcd nodes (list)
	EtcdHosts interface{}
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion interface{}
	// (Computed) RKE k8s cluster inactive nodes (list)
	InactiveHosts interface{}
	// Docker image for ingress (string)
	Ingress interface{}
	// (Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)
	InternalKubeConfigYaml interface{}
	// (Computed) RKE k8s cluster admin user (string)
	KubeAdminUser interface{}
	// (Computed/Sensitive) RKE k8s cluster kube config yaml (string)
	KubeConfigYaml interface{}
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion interface{}
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring interface{}
	// (list maxitems:1)
	Network interface{}
	// RKE k8s cluster nodes (list)
	Nodes interface{}
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs interface{}
	// RKE k8s directory path (string)
	PrefixPath interface{}
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries interface{}
	// Restore cluster. Default `false` (bool)
	Restore interface{}
	// (Computed/Sensitive) RKE k8s cluster config yaml (string)
	RkeClusterYaml interface{}
	// (Computed/Sensitive) RKE k8s cluster state (string)
	RkeState interface{}
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates interface{}
	// (Computed) RKE k8s cluster running system images list (list)
	RunningSystemImages interface{}
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services interface{}
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd interface{}
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi interface{}
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController interface{}
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet interface{}
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy interface{}
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler interface{}
	// SSH Agent Auth enable (bool)
	SshAgentAuth interface{}
	// SSH Certificate path (string)
	SshCertPath interface{}
	// SSH Private Key path (string)
	SshKeyPath interface{}
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages interface{}
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly interface{}
	// (Computed) RKE k8s cluster worker nodes (list)
	WorkerHosts interface{}
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// RKE k8s cluster addon deployment timeout in seconds for status check (int)
	AddonJobTimeout interface{}
	// RKE k8s cluster user addons YAML manifest to be deployed (string)
	Addons interface{}
	// RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)
	AddonsIncludes interface{}
	// RKE k8s cluster authentication configuration (list maxitems:1)
	Authentication interface{}
	// RKE k8s cluster authorization mode configuration (list maxitems:1)
	Authorization interface{}
	// RKE k8s cluster bastion Host configuration (list maxitems:1)
	BastionHost interface{}
	// Specify a certificate dir path (string)
	CertDir interface{}
	// Calico cloud provider (string)
	CloudProvider interface{}
	// RKE k8s cluster name used in the kube config (string)
	ClusterName interface{}
	// Use custom certificates from a cert dir (string)
	CustomCerts interface{}
	// RKE k8s cluster delay on creation (int)
	DelayOnCreation interface{}
	// Deploy RKE cluster on a dind environment. Default: `false` (bool)
	Dind interface{}
	// DinD RKE cluster dns (string)
	DindDnsServer interface{}
	// DinD RKE cluster storage driver (string)
	DindStorageDriver interface{}
	// Enable/Disable RKE k8s cluster port checking. Default `false` (bool)
	DisablePortCheck interface{}
	// RKE k8s cluster DNS Config (list maxitems:1)
	Dns interface{}
	// Enable/Disable RKE k8s cluster strict docker version checking. Default `false` (bool)
	IgnoreDockerVersion interface{}
	// Docker image for ingress (string)
	Ingress interface{}
	// K8s version to deploy (if kubernetes image is specified, image version takes precedence) (string)
	KubernetesVersion interface{}
	// RKE k8s cluster monitoring Config (list maxitems:1)
	Monitoring interface{}
	// (list maxitems:1)
	Network interface{}
	// RKE k8s cluster nodes (list)
	Nodes interface{}
	// RKE k8s cluster nodes (YAML | JSON)
	NodesConfs interface{}
	// RKE k8s directory path (string)
	PrefixPath interface{}
	// RKE k8s cluster private docker registries (list)
	PrivateRegistries interface{}
	// Restore cluster. Default `false` (bool)
	Restore interface{}
	// RKE k8s cluster rotate certificates configuration (list maxitems:1)
	RotateCertificates interface{}
	// Services to rotate their certs. `etcd`, `kubelet`, `kube-apiserver`, `kube-proxy`, `kube-scheduler` and `kube-controller-manager` are supported (list)
	Services interface{}
	// Use services.etcd instead (list maxitems:1)
	Services_Etcd interface{}
	// Use services.kube_api instead (list maxitems:1)
	Services_KubeApi interface{}
	// Use services.kube_controller instead (list maxitems:1)
	Services_KubeController interface{}
	// Use services.kubelet instead (list maxitems:1)
	Services_Kubelet interface{}
	// Use services.kubeproxy instead (list maxitems:1)
	Services_KubeProxy interface{}
	// Use services.scheduler instead (list maxitems:1)
	Services_KubeScheduler interface{}
	// SSH Agent Auth enable (bool)
	SshAgentAuth interface{}
	// SSH Certificate path (string)
	SshCertPath interface{}
	// SSH Private Key path (string)
	SshKeyPath interface{}
	// RKE k8s cluster system images list (list maxitems:1)
	SystemImages interface{}
	// Skip idempotent deployment of control and etcd plane. Default `false` (bool)
	UpdateOnly interface{}
}
