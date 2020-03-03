package component

import (
	"fmt"

	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"yunion.io/x/onecloud-operator/pkg/apis/constants"
	"yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"
	"yunion.io/x/onecloud-operator/pkg/controller"
	"yunion.io/x/onecloud-operator/pkg/manager"
	"yunion.io/x/onecloud/pkg/cloudcommon/options"
)

type hostManager struct {
	*ComponentManager
}

func newHostManager(man *ComponentManager) manager.Manager {
	return &hostManager{man}
}

func (m *hostManager) Sync(oc *v1alpha1.OnecloudCluster) error {
	return syncComponent(m, oc, oc.Spec.HostAgent.Disable)
}

func (m *hostManager) getCloudUser(cfg *v1alpha1.OnecloudClusterConfig) *v1alpha1.CloudUser {
	return &cfg.HostAgent.CloudUser
}

func (m *hostManager) getPhaseControl(man controller.ComponentManager) controller.PhaseControl {
	return controller.NewRegisterServiceComponent(man, constants.ServiceNameHost, constants.ServiceTypeHost)
}

func (m *hostManager) getConfigMap(
	oc *v1alpha1.OnecloudCluster,
	cfg *v1alpha1.OnecloudClusterConfig,
) (*corev1.ConfigMap, error) {
	commonOpt := new(options.CommonOptions)
	// opt := &options.HostOptions
	if err := SetOptionsDefault(commonOpt, ""); err != nil {
		return nil, err
	}
	config := cfg.HostAgent
	SetOptionsServiceTLS(&commonOpt.BaseOptions)
	SetServiceCommonOptions(commonOpt, oc, config.ServiceCommonOptions)
	commonOpt.Port = constants.HostPort

	return m.newServiceConfigMap(v1alpha1.HostComponentType, oc, commonOpt), nil
}

func (m *hostManager) getDaemonSet(
	oc *v1alpha1.OnecloudCluster,
	cfg *v1alpha1.OnecloudClusterConfig,
) (*apps.DaemonSet, error) {
	return m.newHostPrivilegedDaemonSet(v1alpha1.HostComponentType, oc, cfg)
}

func (m *hostManager) newHostPrivilegedDaemonSet(
	cType v1alpha1.ComponentType,
	oc *v1alpha1.OnecloudCluster,
	cfg *v1alpha1.OnecloudClusterConfig,
) (*apps.DaemonSet, error) {
	var (
		privileged  = true
		dsSpec      = oc.Spec.HostAgent
		configMap   = controller.ComponentConfigMapName(oc, cType)
		containersF = func(volMounts []corev1.VolumeMount) []corev1.Container {
			return []corev1.Container{
				{
					Name:  cType.String(),
					Image: dsSpec.Image,
					Command: []string{
						fmt.Sprintf("/opt/yunion/bin/%s", cType.String()),
						"--common-config-file",
						"/etc/yunion/common/common.conf",
					},
					ImagePullPolicy: dsSpec.ImagePullPolicy,
					VolumeMounts:    volMounts,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
					WorkingDir: "/opt/cloud",
				},
				{
					Name:  fmt.Sprintf("%s-image", cType.String()),
					Image: dsSpec.Image,
					Command: []string{
						fmt.Sprintf("/opt/yunion/bin/%s-image", cType.String()),
						"--common-config-file",
						"/etc/yunion/common/common.conf",
					},
					ImagePullPolicy: dsSpec.ImagePullPolicy,
					VolumeMounts:    volMounts,
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
					WorkingDir: "/opt/cloud",
				},
			}
		}
	)
	if dsSpec.NodeSelector == nil {
		dsSpec.NodeSelector = make(map[string]string)
	}
	dsSpec.NodeSelector[constants.OnecloudEnableHostLabelKey] = "enable"
	ds, err := m.newDaemonSet(cType, oc, cfg,
		NewHostVolume(cType, oc, configMap), dsSpec, "", nil, containersF)
	if err != nil {
		return nil, err
	}

	// add pod label for pod affinity
	if ds.Spec.Template.ObjectMeta.Labels == nil {
		ds.Spec.Template.ObjectMeta.Labels = make(map[string]string)
	}
	ds.Spec.Template.ObjectMeta.Labels[constants.OnecloudHostDeployerLabelKey] = ""
	if ds.Spec.Selector == nil {
		ds.Spec.Selector = &metav1.LabelSelector{}
	}
	if ds.Spec.Selector.MatchLabels == nil {
		ds.Spec.Selector.MatchLabels = make(map[string]string)
	}
	ds.Spec.Selector.MatchLabels[constants.OnecloudHostDeployerLabelKey] = ""
	return ds, nil
}
