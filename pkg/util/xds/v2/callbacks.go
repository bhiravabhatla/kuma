package v2

import (
	"context"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_xds "github.com/envoyproxy/go-control-plane/pkg/server/v2"
	"github.com/golang/protobuf/proto"

	"github.com/kumahq/kuma/pkg/util/xds"
)

type adapterCallbacks struct {
	callbacks xds.Callbacks
}

func NewAdapterCallbacks(callbacks xds.Callbacks) envoy_xds.Callbacks {
	return &adapterCallbacks{
		callbacks: callbacks,
	}
}

var _ envoy_xds.Callbacks = &adapterCallbacks{}

func (a *adapterCallbacks) OnFetchRequest(ctx context.Context, request *envoy_api_v2.DiscoveryRequest) error {
	panic("implement me")
}

func (a *adapterCallbacks) OnFetchResponse(request *envoy_api_v2.DiscoveryRequest, response *envoy_api_v2.DiscoveryResponse) {
	panic("implement me")
}

func (a *adapterCallbacks) OnStreamOpen(ctx context.Context, streamID int64, typeURL string) error {
	return a.callbacks.OnStreamOpen(ctx, streamID, typeURL)
}

func (a *adapterCallbacks) OnStreamClosed(streamID int64) {
	a.callbacks.OnStreamClosed(streamID)
}

func (a *adapterCallbacks) OnStreamRequest(streamID int64, request *envoy_api_v2.DiscoveryRequest) error {
	return a.callbacks.OnStreamRequest(streamID, &discoveryRequest{request})
}

func (a *adapterCallbacks) OnStreamResponse(streamID int64, request *envoy_api_v2.DiscoveryRequest, response *envoy_api_v2.DiscoveryResponse) {
	a.callbacks.OnStreamResponse(streamID, &discoveryRequest{request}, &discoveryResponse{response})
}

type discoveryRequest struct {
	*envoy_api_v2.DiscoveryRequest
}

func (d *discoveryRequest) NodeId() string {
	return d.GetNode().GetId()
}

func (d *discoveryRequest) HasErrors() bool {
	return d.ErrorDetail != nil
}

func (d *discoveryRequest) Proto() proto.Message {
	return d
}

var _ xds.DiscoveryRequest = &discoveryRequest{}

type discoveryResponse struct {
	*envoy_api_v2.DiscoveryResponse
}