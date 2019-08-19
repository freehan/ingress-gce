package namer

import (
	"fmt"
	"strings"
)

type frontEndNamer struct {
	clusterUid  string
}

func NewV1FrontEndNamer(uid string) LoadBalancerFrontendNamer {
	return &frontEndNamer{clusterUid: uid}
}

// need a util to calcuate max length
const maxLength = 40

func(n *frontEndNamer) ForwardingRule(namespace, name string, protocol NamerProtocol) string {
	return fmt.Sprintf("k8s-%v-%v-%v", n.clusterUid, protocol, strings.Join(TrimFieldsEvenly(maxLength, namespace, name), ""))
}

func(n *frontEndNamer) TargetProxy(namespace, name string, protocol NamerProtocol) string {
	return fmt.Sprintf("k8s-%v-%v-%v", n.clusterUid, protocol, strings.Join(TrimFieldsEvenly(maxLength, namespace, name), ""))
}

func(n *frontEndNamer) UrlMap(namespace, name string) string {
	return fmt.Sprintf("k8s-%v-%v", n.clusterUid, strings.Join(TrimFieldsEvenly(maxLength, namespace, name), ""))
}
