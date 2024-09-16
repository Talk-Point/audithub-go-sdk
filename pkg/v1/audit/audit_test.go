package audit

import (
	"os"
	"testing"
	"time"
)

func TestAuditLog(t *testing.T) {
	os.Setenv("ENV", "production")
	entry := AuditLog("test_service", "test_event")

	if entry.Env != "production" {
		t.Errorf("expected Env to be 'production', got %s", entry.Env)
	}
	if entry.Service != "test_service" {
		t.Errorf("expected Service to be 'test_service', got %s", entry.Service)
	}
	if entry.Event != "test_event" {
		t.Errorf("expected Event to be 'test_event', got %s", entry.Event)
	}
	if entry.Timestamp <= 0 {
		t.Errorf("expected Timestamp to be set, got %d", entry.Timestamp)
	}
}

func TestLog(t *testing.T) {
	entry := &AuditEntry{
		Env:       "test",
		Timestamp: time.Now().Unix(),
		Service:   "test_service",
		Event:     "test_event",
		Gids:      []string{"gid1"},
		Labels:    []string{"label1"},
		ByUser:    "user1",
		Metadata:  map[string]string{"key1": "value1"},
	}

	err := entry.Log()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestAddGids(t *testing.T) {
	entry := &AuditEntry{}
	entry.AddGids([]string{"gid1", "gid2"})

	if len(entry.Gids) != 2 {
		t.Errorf("expected 2 GIDs, got %d", len(entry.Gids))
	}
	if entry.Gids[0] != "gid1" || entry.Gids[1] != "gid2" {
		t.Errorf("GIDs not added correctly, got %v", entry.Gids)
	}
}

func TestAddLabels(t *testing.T) {
	entry := &AuditEntry{}
	entry.AddLabels([]string{"label1", "label2"})

	if len(entry.Labels) != 2 {
		t.Errorf("expected 2 labels, got %d", len(entry.Labels))
	}
	if entry.Labels[0] != "label1" || entry.Labels[1] != "label2" {
		t.Errorf("Labels not added correctly, got %v", entry.Labels)
	}
}

func TestBy(t *testing.T) {
	entry := &AuditEntry{}
	entry.By("user1")

	if entry.ByUser != "user1" {
		t.Errorf("expected ByUser to be 'user1', got %s", entry.ByUser)
	}
}

func TestAddMetadata(t *testing.T) {
	entry := &AuditEntry{Metadata: make(map[string]string)}
	entry.AddMetadata("key1", "value1")

	if len(entry.Metadata) != 1 {
		t.Errorf("expected 1 metadata entry, got %d", len(entry.Metadata))
	}
	if entry.Metadata["key1"] != "value1" {
		t.Errorf("Metadata not added correctly, got %v", entry.Metadata)
	}
}
