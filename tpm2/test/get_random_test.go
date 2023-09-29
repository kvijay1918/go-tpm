package tpm2test

import (
	"testing"

	. "github.com/kvijay1918/go-tpm/tpm2"
	"github.com/kvijay1918/go-tpm/tpm2/transport/simulator"
)

func TestGetRandom(t *testing.T) {
	thetpm, err := simulator.OpenSimulator()
	if err != nil {
		t.Fatalf("could not connect to TPM simulator: %v", err)
	}
	defer thetpm.Close()

	grc := GetRandom{
		BytesRequested: 16,
	}

	if _, err := grc.Execute(thetpm); err != nil {
		t.Fatalf("GetRandom failed: %v", err)
	}
}
