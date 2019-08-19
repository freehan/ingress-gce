package namer

import "k8s.io/api/networking/v1beta1"

const (
	LegacyNamingScheme = Scheme("v1")
	V2NamingScheme = Scheme("v2")
)

type Scheme string

// per ingress namer for frontend stuff
type IngressFrontendNamer interface {
	ForwardingRule(protocol NamerProtocol) string
	TargetProxy(protocol NamerProtocol) string
	UrlMap() string
	SslCert(secretHash string) string
}


type IngressFrontendNamerFactory interface {
	CreateIngressFrontendNamer(scheme Scheme, ing v1beta1.Ingress) IngressFrontendNamer
}

// LegacyIngressFrontendNamer IngressFrontendNamer with legacy namer
type LegacyIngressFrontendNamer struct {
	ing *v1beta1.Ingress
	namer *Namer
}


type RealIngressFrontendNamerFactory struct {
	namer *Namer

}

func (*RealIngressFrontendNamerFactory) CreateIngressFrontendNamer (scheme Scheme, ing v1beta1.Ingress) IngressFrontendNamer {
	switch(scheme):
		case v1:


}


type newNamer struct {
	cluster-uid string
}




// LegacyIngressFrontendNamer IngressFrontendNamer with legacy namer
type V2IngressFrontendNamer struct {
	namespace string
	name string
	ing ingress
}



