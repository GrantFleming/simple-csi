package identity

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
)

func TestService_GetPluginInfo(t *testing.T) {
	server := Service{}

	pluginInfo, err := server.GetPluginInfo(nil, nil)

	if err != nil {
		t.Fatal("GetPluginInfo should never error")
	}

	expectedName := "simple-csi.grant.goose"
	actualName := pluginInfo.Name
	if expectedName != actualName {
		t.Errorf("Expected plugin name to be %s but was actually %s", expectedName, actualName)
	}

	expectedVersion := "0.0.1"
	actualVersion := pluginInfo.VendorVersion
	if expectedVersion != actualVersion {
		t.Errorf("Expected plugin version to be %s but was actually %s", expectedVersion, actualVersion)
	}

}

func TestService_Probe(t *testing.T) {
	server := Service{}
	probe, err := server.Probe(nil, nil)

	if err != nil {
		t.Fatal("Probe should never error")
	}

	if !probe.GetReady().Value {
		t.Errorf("Expected probe to declare plugin is ready, but it returned false")
	}
}

func TestService_GetPluginCapabilities(t *testing.T) {
	server := Service{}
	capResponse, err := server.GetPluginCapabilities(nil, nil)

	if err != nil {
		t.Fatal("Capabilities should never error")
	}

	caps := capResponse.Capabilities
	if len(caps) != 1 {
		t.Errorf("Expected the service to have only 1 capability but it had %v", len(caps))
	}

	serviceType := caps[0].GetService().Type
	if serviceType != csi.PluginCapability_Service_CONTROLLER_SERVICE {
		t.Errorf("Expected the plugin to have a controller service")
	}
}
