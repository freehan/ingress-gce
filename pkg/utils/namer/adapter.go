package namer

import "fmt"

type v1NamerAdapter struct {
	namer *Namer
}

func NewLegacyFrontendNamer(namer *Namer) LoadBalancerFrontendNamer {
	return &v1NamerAdapter{namer: namer}
}

func(n *v1NamerAdapter) ForwardingRule(namespace, name string, protocol NamerProtocol) string {

	return n.namer.ForwardingRule(n.namer.LoadBalancer(keyFunc(namespace,name)), protocol)
}

func(n *v1NamerAdapter) TargetProxy(namespace, name string, protocol NamerProtocol) string {
	return n.namer.TargetProxy(n.namer.LoadBalancer(keyFunc(namespace,name)), protocol)
}

func(n *v1NamerAdapter) UrlMap(namespace, name string) string {
	return n.namer.UrlMap(n.namer.LoadBalancer(keyFunc(namespace,name)))
}

func keyFunc(namespace, name string) string {
	return fmt.Sprintf("%s/%s", namespace, name)
}