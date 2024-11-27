package identity

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	csiName    = "simple-csi.grant.goose"
	csiVersion = "0.0.1"
)

type Service struct {
	csi.UnimplementedIdentityServer
}

// TODO log the calls to the service

func (Service) GetPluginInfo(context.Context, *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	response := csi.GetPluginInfoResponse{}
	response.Name = csiName
	response.VendorVersion = csiVersion
	return &response, nil
}

func (Service) GetPluginCapabilities(context.Context, *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	cap := []*csi.PluginCapability{
		{
			Type: &csi.PluginCapability_Service_{
				Service: &csi.PluginCapability_Service{
					Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
				},
			},
		},
	}

	return &csi.GetPluginCapabilitiesResponse{Capabilities: cap}, nil
}

func (Service) Probe(context.Context, *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	response := csi.ProbeResponse{}
	response.Ready = wrapperspb.Bool(true)
	return &response, nil
}
