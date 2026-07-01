package source

import (
	"strings"
	"testing"
)

func TestDecodeAppsUsesExplicitDownloadReleaseAndGatewayMetadata(t *testing.T) {
	src := NewFNOSAppsSource("", "", nil)
	src.platform = "x86"

	apps, err := src.decodeApps([]byte(`{
		"apps": [{
			"appname": "hermes-studio",
			"display_name": "Hermes Studio",
			"version": "0.6.26",
			"fpk_version": "0.6.26",
			"release_tag": "fnos-v0.6.26",
			"file_prefix": "hermes-studio",
			"service_port": 6060,
			"download_url": "https://github.com/MScorpioLee/hermes-studio/releases/download/fnos-v0.6.26/hermes-studio.fpk",
			"release_url": "https://github.com/MScorpioLee/hermes-studio/releases/tag/fnos-v0.6.26",
			"gateway_prefix": "/app/hermes-studio",
			"gateway_socket": "hermes-studio.sock",
			"platforms": ["x86"]
		}]
	}`))
	if err != nil {
		t.Fatalf("decodeApps returned error: %v", err)
	}
	if len(apps) != 1 {
		t.Fatalf("expected one app, got %d", len(apps))
	}

	app := apps[0]
	if app.FpkURL != "https://github.com/MScorpioLee/hermes-studio/releases/download/fnos-v0.6.26/hermes-studio.fpk" {
		t.Fatalf("FpkURL = %q", app.FpkURL)
	}
	if app.ReleaseURL != "https://github.com/MScorpioLee/hermes-studio/releases/tag/fnos-v0.6.26" {
		t.Fatalf("ReleaseURL = %q", app.ReleaseURL)
	}
	if app.GatewayPrefix != "/app/hermes-studio" {
		t.Fatalf("GatewayPrefix = %q", app.GatewayPrefix)
	}
	if app.GatewaySocket != "hermes-studio.sock" {
		t.Fatalf("GatewaySocket = %q", app.GatewaySocket)
	}
}

func TestDecodeAppsKeepsConversunReleasePatternByDefault(t *testing.T) {
	src := NewFNOSAppsSource("", "", nil)
	src.platform = "x86"

	apps, err := src.decodeApps([]byte(`{
		"apps": [{
			"appname": "grafana",
			"display_name": "Grafana",
			"version": "13.0.3",
			"fpk_version": "13.0.3",
			"release_tag": "grafana/v13.0.3",
			"file_prefix": "grafana",
			"platforms": ["x86"]
		}]
	}`))
	if err != nil {
		t.Fatalf("decodeApps returned error: %v", err)
	}
	if len(apps) != 1 {
		t.Fatalf("expected one app, got %d", len(apps))
	}

	if !strings.Contains(apps[0].FpkURL, "github.com/conversun/fnos-apps/releases/download/grafana/v13.0.3/grafana_13.0.3_x86.fpk") {
		t.Fatalf("unexpected default FpkURL: %q", apps[0].FpkURL)
	}
	if apps[0].ReleaseURL != "https://github.com/conversun/fnos-apps/releases/tag/grafana/v13.0.3" {
		t.Fatalf("unexpected default ReleaseURL: %q", apps[0].ReleaseURL)
	}
}
