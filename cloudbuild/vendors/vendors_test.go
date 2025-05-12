package vendors_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra"
	kinit "github.com/openconfig/ondatra/knebind/init"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, kinit.Init)
}

func testConfigPush(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("config push", func(t *testing.T) {
		tmpl := "hostname %s"
		switch dut.Vendor() {
		case ondatra.JUNIPER:
			// TODO: Figure out Juniper config.
		case ondatra.NOKIA:
			tmpl = "host-name %s"
		}
		cfg := fmt.Sprintf(tmpl, dut.Name())
		dut.Config().New().WithText(cfg).Append(t)
		t.Logf("Successfully pushed config for DUT %s", dut)
	})
}

func testGNMI(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("gnmi", func(t *testing.T) {
		dut.RawAPIs().GNMI(t)
		t.Logf("Got GNMI client for DUT %s", dut)
	})
}

func testGNOI(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("gnoi", func(t *testing.T) {
		dut.RawAPIs().GNOI(t)
		t.Logf("Got GNOI client for DUT %s", dut)
	})
}

func testGNSI(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("gnsi", func(t *testing.T) {
		dut.RawAPIs().GNSI(t)
		t.Logf("Got GNSI client for DUT %s", dut)
	})
}

func testGRIBI(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("gribi", func(t *testing.T) {
		dut.RawAPIs().GRIBI(t)
		t.Logf("Got GRIBI client for DUT %s", dut)
	})
}

func testP4RT(t *testing.T, dut *ondatra.DUTDevice) {
	t.Helper()
	t.Run("p4rt", func(t *testing.T) {
		dut.RawAPIs().P4RT(t)
		t.Logf("Got P4RT client for DUT %s", dut)
	})
}

func TestXRD(t *testing.T) {
	dut := ondatra.DUT(t, "xrd")
	testConfigPush(t, dut)
	testGNMI(t, dut)
	testGNOI(t, dut)
	testGNSI(t, dut)
	testGRIBI(t, dut)
	testP4RT(t, dut)
}
