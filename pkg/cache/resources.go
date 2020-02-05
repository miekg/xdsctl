package cache

import (
	discoverypb "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/gogo/protobuf/proto"
)

// Resource types in xDS v3 (guessing here) and v2
const (
	apiTypePrefix = "type.googleapis.com/envoy.config."
	ClusterType   = apiTypePrefix + "cluster.v3.Cluster"
	ClusterType2  = "type.googleapis.com/envoy.api.v2.Cluster"

	// Note: I've made EndpointType up; it may never exist in envoy.
	EndpointType  = apiTypePrefix + "endpoint.v3.ClusterLoadAssignment"
	EndpointType2 = "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment"

	// AnyType is used only by ADS
	AnyType = ""
)

type Response struct {
	*discoverypb.DiscoveryResponse
	version string
}

type MarshaledResource = []byte

// MarshalResource converts the Resource to MarshaledResource
func MarshalResource(resource proto.Message) (MarshaledResource, error) {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(resource)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
