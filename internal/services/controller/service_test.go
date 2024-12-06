package controller

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func TestServer_CreateVolume(t *testing.T) {
	server := Server{}
	req := csi.CreateVolumeRequest{
		Name: "some-volume",
	}
	volRes, err := server.CreateVolume(nil, &req)

	if err != nil {
		t.Error("Create volume should not error")
	}

	volName := volRes.Volume.VolumeId

	if volName != req.Name {
		t.Errorf("Expected volume name to be %s but was %s", req.Name, volName)
	}
}

func TestServer_DeleteVolume(t *testing.T) {
	server := Server{}
	req := csi.DeleteVolumeRequest{
		VolumeId: "some-volume",
	}
	_, err := server.DeleteVolume(nil, &req)

	if err != nil {
		t.Error("Delete volume should never error")
	}
}
