/*
Copyright 2018 The Kubernetes Authors.
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

package shared

import (
	"fmt"
	api_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/ingress-gce/pkg/annotations"
	"k8s.io/api/networking/v1beta1"
	types2 "k8s.io/ingress-gce/pkg/utils/types"
)

// ServicePort is a helper function that retrieves a port of a Service.
func ServicePort(svc api_v1.Service, port intstr.IntOrString) *api_v1.ServicePort {
	var svcPort *api_v1.ServicePort
PortLoop:
	for _, p := range svc.Spec.Ports {
		np := p
		switch port.Type {
		case intstr.Int:
			if p.Port == port.IntVal {
				svcPort = &np
				break PortLoop
			}
		default:
			if p.Name == port.StrVal {
				svcPort = &np
				break PortLoop
			}
		}
	}
	return svcPort
}



// IsGCEIngress returns true if the Ingress matches the class managed by this
// controller.
func IsGCEIngress(ing *v1beta1.Ingress) bool {
	class := annotations.FromIngress(ing).IngressClass()
	//if flags.F.IngressClass != "" && class == flags.F.IngressClass {
	//	return true
	//}

	switch class {
	case "":
		// Ingress controller does not have any ingress classes that can be
		// specified by spec.IngressClassName. If spec.IngressClassName
		// is nil, then consider GCEIngress.
		return ing.Spec.IngressClassName == nil
	case annotations.GceIngressClass:
		return true
	case annotations.GceL7ILBIngressClass:
		// TODO: (shance) remove flag check for L7-ILB once fully rolled out
		//return flags.F.EnableL7Ilb
		return true
	default:
		return false
	}
}

// IsGCEMultiClusterIngress returns true if the given Ingress has
// ingress.class annotation set to "gce-multi-cluster".
func IsGCEMultiClusterIngress(ing *v1beta1.Ingress) bool {
	class := annotations.FromIngress(ing).IngressClass()
	return class == annotations.GceMultiIngressClass
}

// IsGCEL7ILBIngress returns true if the given Ingress has
// ingress.class annotation set to "gce-l7-ilb"
func IsGCEL7ILBIngress(ing *v1beta1.Ingress) bool {
	class := annotations.FromIngress(ing).IngressClass()
	return class == annotations.GceL7ILBIngressClass
}

// IsGLBCIngress returns true if the given Ingress should be processed by GLBC
func IsGLBCIngress(ing *v1beta1.Ingress) bool {
	return IsGCEIngress(ing) || IsGCEMultiClusterIngress(ing)
}

// TraverseIngressBackends traverse thru all backends specified in the input ingress and call process
// If process return true, then return and stop traversing the backends
func TraverseIngressBackends(ing *v1beta1.Ingress, process func(id types2.ServicePortID) bool) {
	if ing == nil {
		return
	}
	// Check service of default backend
	if ing.Spec.Backend != nil {
		if process(types2.ServicePortID{Service: types.NamespacedName{Namespace: ing.Namespace, Name: ing.Spec.Backend.ServiceName}, Port: ing.Spec.Backend.ServicePort}) {
			return
		}
	}

	// Check the target service for each path rule
	for _, rule := range ing.Spec.Rules {
		if rule.IngressRuleValue.HTTP == nil {
			continue
		}
		for _, p := range rule.IngressRuleValue.HTTP.Paths {
			if process(types2.ServicePortID{Service: types.NamespacedName{Namespace: ing.Namespace, Name: p.Backend.ServiceName}, Port: p.Backend.ServicePort}) {
				return
			}
		}
	}
	return
}

func ServiceKeyFunc(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}
