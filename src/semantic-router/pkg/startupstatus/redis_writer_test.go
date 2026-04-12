package startupstatus

import (
	"encoding/json"
	"testing"
)

func TestRedisWriterConfigDefaults(t *testing.T) {
	cfg := RedisWriterConfig{
		Address: "localhost:6379",
	}
	if cfg.TTL != 0 {
		t.Fatalf("expected zero TTL in config (default applied at creation), got %v", cfg.TTL)
	}
}

func TestDefaultStatusKey(t *testing.T) {
	key := DefaultStatusKey()
	if key != "vllm-sr:startup-status" {
		t.Fatalf("DefaultStatusKey() = %q, want %q", key, "vllm-sr:startup-status")
	}
}

func TestStateJSONRoundTrip(t *testing.T) {
	original := State{
		Phase:            "downloading_models",
		Ready:            false,
		Message:          "Downloading model X",
		DownloadingModel: "models/X",
		PendingModels:    []string{"models/X", "models/Y"},
		ReadyModels:      1,
		TotalModels:      3,
		UpdatedAt:        "2025-01-01T00:00:00Z",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var decoded State
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if decoded.Phase != original.Phase {
		t.Errorf("Phase = %q, want %q", decoded.Phase, original.Phase)
	}
	if decoded.Ready != original.Ready {
		t.Errorf("Ready = %v, want %v", decoded.Ready, original.Ready)
	}
	if decoded.DownloadingModel != original.DownloadingModel {
		t.Errorf("DownloadingModel = %q, want %q", decoded.DownloadingModel, original.DownloadingModel)
	}
	if decoded.ReadyModels != original.ReadyModels {
		t.Errorf("ReadyModels = %d, want %d", decoded.ReadyModels, original.ReadyModels)
	}
	if decoded.TotalModels != original.TotalModels {
		t.Errorf("TotalModels = %d, want %d", decoded.TotalModels, original.TotalModels)
	}
	if len(decoded.PendingModels) != len(original.PendingModels) {
		t.Errorf("PendingModels length = %d, want %d", len(decoded.PendingModels), len(original.PendingModels))
	}
}

func TestFileWriterImplementsStatusWriter(t *testing.T) {
	var _ StatusWriter = (*FileWriter)(nil)
}

func TestRedisWriterImplementsStatusWriter(t *testing.T) {
	var _ StatusWriter = (*RedisWriter)(nil)
}
