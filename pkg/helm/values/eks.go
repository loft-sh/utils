package values

import (
	"strings"

	"github.com/go-logr/logr"
	"github.com/loft-sh/utils/pkg/helm"
)

var EKSAPIVersionMap = map[string]string{
	"1.34": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.34.3-eks-1-34-15",
	"1.33": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.33.7-eks-1-33-20",
	"1.32": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.32.10-eks-1-32-25",
	"1.31": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.31.14-eks-1-31-30",
	"1.30": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.30.14-eks-1-30-35",
	"1.29": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.29.15-eks-1-29-40",
	"1.28": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.28.15-eks-1-28-45",
	"1.27": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.27.16-eks-1-27-50",
	"1.26": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.26.20-eks-1-26-55",
	"1.25": "public.ecr.aws/eks-distro/kubernetes/kube-apiserver:v1.25.25-eks-1-25-60",
}

var EKSControllerVersionMap = map[string]string{
	"1.34": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.34.3-eks-1-34-15",
	"1.33": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.33.7-eks-1-33-20",
	"1.32": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.32.10-eks-1-32-25",
	"1.31": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.31.14-eks-1-31-30",
	"1.30": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.30.14-eks-1-30-35",
	"1.29": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.29.15-eks-1-29-40",
	"1.28": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.28.15-eks-1-28-45",
	"1.27": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.27.16-eks-1-27-50",
	"1.26": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.26.20-eks-1-26-55",
	"1.25": "public.ecr.aws/eks-distro/kubernetes/kube-controller-manager:v1.25.25-eks-1-25-60",
}

var EKSEtcdVersionMap = map[string]string{
	"1.34": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.10-eks-1-34-15",
	"1.33": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-33-20",
	"1.32": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-32-25",
	"1.31": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-31-30",
	"1.30": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-30-35",
	"1.29": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-29-40",
	"1.28": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.9-eks-1-28-45",
	"1.27": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.8-eks-1-27-50",
	"1.26": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.8-eks-1-26-55",
	"1.25": "public.ecr.aws/eks-distro/etcd-io/etcd:v3.5.8-eks-1-25-60",
}

var EKSCoreDNSVersionMap = map[string]string{
	"1.34": "public.ecr.aws/eks-distro/coredns/coredns:v1.11.1-eks-1-34-15",
	"1.33": "public.ecr.aws/eks-distro/coredns/coredns:v1.11.1-eks-1-33-20",
	"1.32": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-32-25",
	"1.31": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-31-30",
	"1.30": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-30-35",
	"1.29": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-29-40",
	"1.28": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-28-45",
	"1.27": "public.ecr.aws/eks-distro/coredns/coredns:v1.10.1-eks-1-27-50",
	"1.26": "public.ecr.aws/eks-distro/coredns/coredns:v1.9.3-eks-1-26-55",
	"1.25": "public.ecr.aws/eks-distro/coredns/coredns:v1.9.3-eks-1-25-60",
}

func getDefaultEKSReleaseValues(chartOptions *helm.ChartOptions, log logr.Logger) (string, error) {
	apiImage := ""
	controllerImage := ""
	etcdImage := ""
	corednsImage := ""
	if chartOptions.KubernetesVersion.Major != "" && chartOptions.KubernetesVersion.Minor != "" {
		serverVersionString := GetKubernetesVersion(chartOptions.KubernetesVersion)
		serverMinorInt, err := GetKubernetesMinorVersion(chartOptions.KubernetesVersion)
		if err != nil {
			return "", err
		}

		var ok bool
		apiImage = EKSAPIVersionMap[serverVersionString]
		controllerImage = EKSControllerVersionMap[serverVersionString]
		etcdImage = EKSEtcdVersionMap[serverVersionString]
		corednsImage, ok = EKSCoreDNSVersionMap[serverVersionString]
		if !ok {
			if serverMinorInt > 34 {
				log.Info("officially unsupported host server version, will fallback to virtual cluster version v1.34", "serverVersion", serverVersionString)
				apiImage = EKSAPIVersionMap["1.34"]
				controllerImage = EKSControllerVersionMap["1.34"]
				etcdImage = EKSEtcdVersionMap["1.34"]
				corednsImage = EKSCoreDNSVersionMap["1.34"]
			} else {
				log.Info("officially unsupported host server version, will fallback to virtual cluster version v1.25", "serverVersion", serverVersionString)
				apiImage = EKSAPIVersionMap["1.25"]
				controllerImage = EKSControllerVersionMap["1.25"]
				etcdImage = EKSEtcdVersionMap["1.25"]
				corednsImage = EKSCoreDNSVersionMap["1.25"]
			}
		}
	}

	// build values
	values := ""
	if apiImage != "" {
		values = `api:
  image: ##API_IMAGE##
controller:
  image: ##CONTROLLER_IMAGE##
etcd:
  image: ##ETCD_IMAGE##
coredns:
  image: ##COREDNS_IMAGE##
`
		values = strings.ReplaceAll(values, "##API_IMAGE##", apiImage)
		values = strings.ReplaceAll(values, "##CONTROLLER_IMAGE##", controllerImage)
		values = strings.ReplaceAll(values, "##ETCD_IMAGE##", etcdImage)
		values = strings.ReplaceAll(values, "##COREDNS_IMAGE##", corednsImage)
	}
	return addCommonReleaseValues(values, chartOptions)
}
