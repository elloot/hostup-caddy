package hostup

import (
	"testing"

	"github.com/caddyserver/caddy/v2"
)

func TestModuleRegistered(t *testing.T) {
	m, err := caddy.GetModule("dns.providers.hostup")
	if err != nil {
		t.Fatalf("module dns.providers.hostup not registered: %v", err)
	}
	if m.ID != "dns.providers.hostup" {
		t.Fatalf("unexpected module ID: %v", m.ID)
	}
	if m.New() == nil {
		t.Fatal("module constructor returned nil")
	}
}
