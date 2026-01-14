package values

import (
	"strings"

	"github.com/go-logr/logr"
	"github.com/loft-sh/utils/pkg/helm"
)

var K8SAPIVersionMap = map[string]string{
	"1.34": "registry.k8s.io/kube-apiserver:v1.34.3",
	"1.33": "registry.k8s.io/kube-apiserver:v1.33.7",
	"1.32": "registry.k8s.io/kube-apiserver:v1.32.10",
	"1.31": "registry.k8s.io/kube-apiserver:v1.31.14",
	"1.30": "registry.k8s.io/kube-apiserver:v1.30.14",
	"1.29": "registry.k8s.io/kube-apiserver:v1.29.15",
	"1.28": "registry.k8s.io/kube-apiserver:v1.28.15",
	"1.27": "registry.k8s.io/kube-apiserver:v1.27.16",
	"1.26": "registry.k8s.io/kube-apiserver:v1.26.20",
	"1.25": "registry.k8s.io/kube-apiserver:v1.25.25",
}

var K8SControllerVersionMap = map[string]string{
	"1.34": "registry.k8s.io/kube-controller-manager:v1.34.3",
	"1.33": "registry.k8s.io/kube-controller-manager:v1.33.7",
	"1.32": "registry.k8s.io/kube-controller-manager:v1.32.10",
	"1.31": "registry.k8s.io/kube-controller-manager:v1.31.14",
	"1.30": "registry.k8s.io/kube-controller-manager:v1.30.14",
	"1.29": "registry.k8s.io/kube-controller-manager:v1.29.15",
	"1.28": "registry.k8s.io/kube-controller-manager:v1.28.15",
	"1.27": "registry.k8s.io/kube-controller-manager:v1.27.16",
	"1.26": "registry.k8s.io/kube-controller-manager:v1.26.20",
	"1.25": "registry.k8s.io/kube-controller-manager:v1.25.25",
}

var K8SSchedulerVersionMap = map[string]string{
	"1.34": "registry.k8s.io/kube-scheduler:v1.34.3",
	"1.33": "registry.k8s.io/kube-scheduler:v1.33.7",
	"1.32": "registry.k8s.io/kube-scheduler:v1.32.10",
	"1.31": "registry.k8s.io/kube-scheduler:v1.31.14",
	"1.30": "registry.k8s.io/kube-scheduler:v1.30.14",
	"1.29": "registry.k8s.io/kube-scheduler:v1.29.15",
	"1.28": "registry.k8s.io/kube-scheduler:v1.28.15",
	"1.27": "registry.k8s.io/kube-scheduler:v1.27.16",
	"1.26": "registry.k8s.io/kube-scheduler:v1.26.20",
	"1.25": "registry.k8s.io/kube-scheduler:v1.25.25",
}

var K8SEtcdVersionMap = map[string]string{
	"1.34": "registry.k8s.io/etcd:3.5.10-0",
	"1.33": "registry.k8s.io/etcd:3.5.9-0",
	"1.32": "registry.k8s.io/etcd:3.5.9-0",
	"1.31": "registry.k8s.io/etcd:3.5.9-0",
	"1.30": "registry.k8s.io/etcd:3.5.9-0",
	"1.29": "registry.k8s.io/etcd:3.5.9-0",
	"1.28": "registry.k8s.io/etcd:3.5.9-0",
	"1.27": "registry.k8s.io/etcd:3.5.7-0",
	"1.26": "registry.k8s.io/etcd:3.5.6-0",
	"1.25": "registry.k8s.io/etcd:3.5.6-0",
}

func getDefaultK8SReleaseValues(chartOptions *helm.ChartOptions, log logr.Logger) (string, error) {
	apiImage := ""
	controllerImage := ""
	etcdImage := ""
	schedulerImage := ""
	if chartOptions.KubernetesVersion.Major != "" && chartOptions.KubernetesVersion.Minor != "" {
		serverVersionString := GetKubernetesVersion(chartOptions.KubernetesVersion)
		serverMinorInt, err := GetKubernetesMinorVersion(chartOptions.KubernetesVersion)
		if err != nil {
			return "", err
		}

		var ok bool
		apiImage = K8SAPIVersionMap[serverVersionString]
		controllerImage = K8SControllerVersionMap[serverVersionString]
		schedulerImage = K8SSchedulerVersionMap[serverVersionString]
		etcdImage, ok = K8SEtcdVersionMap[serverVersionString]
		if !ok {
			if serverMinorInt > 34 {
				log.Info("officially unsupported host server version, will fallback to virtual cluster version v1.34", "serverVersion", serverVersionString)
				apiImage = K8SAPIVersionMap["1.34"]
				controllerImage = K8SControllerVersionMap["1.34"]
				schedulerImage = K8SSchedulerVersionMap["1.34"]
				etcdImage = K8SEtcdVersionMap["1.34"]
			} else {
				log.Info("officially unsupported host server version, will fallback to virtual cluster version v1.25", "serverVersion", serverVersionString)
				apiImage = K8SAPIVersionMap["1.25"]
				controllerImage = K8SControllerVersionMap["1.25"]
				schedulerImage = K8SSchedulerVersionMap["1.25"]
				etcdImage = K8SEtcdVersionMap["1.25"]
			}
		}
	}

	// build values
	values := ""
	if apiImage != "" {
		values = `api:
  image: ##API_IMAGE##
scheduler:
  image: ##SCHEDULER_IMAGE##
controller:
  image: ##CONTROLLER_IMAGE##
etcd:
  image: ##ETCD_IMAGE##
`
		values = strings.ReplaceAll(values, "##API_IMAGE##", apiImage)
		values = strings.ReplaceAll(values, "##CONTROLLER_IMAGE##", controllerImage)
		values = strings.ReplaceAll(values, "##SCHEDULER_IMAGE##", schedulerImage)
		values = strings.ReplaceAll(values, "##ETCD_IMAGE##", etcdImage)
	}
	return addCommonReleaseValues(values, chartOptions)
}
