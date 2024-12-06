package controller

import (
	"context"
	"log"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

type Server struct {
	csi.UnimplementedControllerServer
}

func (Server) CreateVolume(_ context.Context, req *csi.CreateVolumeRequest) (_ *csi.CreateVolumeResponse, _ error) {
	log.Printf("Volume created: %s", req.Name)
	res := csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId: req.Name,
		},
	}
	return &res, nil
}

func (Server) DeleteVolume(_ context.Context, req *csi.DeleteVolumeRequest) (_ *csi.DeleteVolumeResponse, _ error) {
	log.Printf("Volume deleted: %s", req.VolumeId)
	res := csi.DeleteVolumeResponse{}
	return &res, nil
}

func (Server) ControllerPublishVolume(_ context.Context, _ *csi.ControllerPublishVolumeRequest) (_ *csi.ControllerPublishVolumeResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Server) ControllerUnpublishVolume(_ context.Context, _ *csi.ControllerUnpublishVolumeRequest) (_ *csi.ControllerUnpublishVolumeResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Server) ValidateVolumeCapabilities(_ context.Context, _ *csi.ValidateVolumeCapabilitiesRequest) (_ *csi.ValidateVolumeCapabilitiesResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Server) ControllerGetCapabilities(_ context.Context, _ *csi.ControllerGetCapabilitiesRequest) (_ *csi.ControllerGetCapabilitiesResponse, _ error) {
	panic("not implemented") // TODO: Implement
	// In this one I will say that you can create/delete volumes and public/unpublish volumes
}
