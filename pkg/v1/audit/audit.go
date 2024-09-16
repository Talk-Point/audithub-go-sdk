package audit

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type AuditEntry struct {
	Env       string            `json:"env"`
	Timestamp int64             `json:"timestamp"`
	Service   string            `json:"service"`
	Event     string            `json:"event"`
	Gids      []string          `json:"gids"`
	Labels    []string          `json:"labels"`
	ByUser    string            `json:"by_user"`
	Metadata  map[string]string `json:"metadata"`
}

// Log outputs the audit entry as a JSON string.
func (a *AuditEntry) Log() error {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return err
	}
	fmt.Println("::" + string(jsonData) + "::")
	return nil
}

// AddGids adds GIDs to the audit entry.
func (a *AuditEntry) AddGids(gids []string) *AuditEntry {
	a.Gids = append(a.Gids, gids...)
	return a
}

// AddLabels adds labels to the audit entry.
func (a *AuditEntry) AddLabels(labels []string) *AuditEntry {
	a.Labels = append(a.Labels, labels...)
	return a
}

// By sets the user in the audit entry.
func (a *AuditEntry) By(user string) *AuditEntry {
	a.ByUser = user
	return a
}

// AddMetadata adds a key-value pair to the metadata.
func (a *AuditEntry) AddMetadata(key, value string) *AuditEntry {
	a.Metadata[key] = value
	return a
}

// AuditLog initializes a new AuditEntry.
func AuditLog(service string, event string) *AuditEntry {
	// Example: Set the environment from an environment variable.
	env := os.Getenv("ENV")
	if env == "" {
		env = "test"
	}

	return &AuditEntry{
		Env:       env,
		Timestamp: time.Now().Unix(),
		Service:   service,
		Event:     event,
		Metadata:  make(map[string]string),
		Labels:    make([]string, 0),
		Gids:      make([]string, 0),
	}
}
