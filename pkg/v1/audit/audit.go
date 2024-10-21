package audit

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type AuditEntry struct {
	Env       string                 `json:"env"`
	Timestamp int64                  `json:"timestamp"`
	Service   string                 `json:"service"`
	Event     string                 `json:"event"`
	Gids      []string               `json:"gids"`
	Labels    []string               `json:"labels"`
	ByUser    string                 `json:"by_user"`
	Metadata  map[string]interface{} `json:"metadata"`
}

func (a *AuditEntry) SetEnv(env string) *AuditEntry {
	a.Env = env
	return a
}

func (a *AuditEntry) SetTimestamp(timestamp int64) *AuditEntry {
	a.Timestamp = timestamp
	return a
}

func (a *AuditEntry) String() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (a *AuditEntry) Log() error {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return err
	}
	fmt.Println("::" + string(jsonData) + "::")
	return nil
}

func (a *AuditEntry) AddGids(gids []string) *AuditEntry {
	a.Gids = append(a.Gids, gids...)
	return a
}

func (a *AuditEntry) AddLabels(labels []string) *AuditEntry {
	a.Labels = append(a.Labels, labels...)
	return a
}

func (a *AuditEntry) By(user string) *AuditEntry {
	a.ByUser = user
	return a
}

func (a *AuditEntry) AddMetadata(key string, value interface{}) *AuditEntry {
	a.Metadata[key] = value
	return a
}

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
		Metadata:  make(map[string]interface{}),
		Labels:    make([]string, 0),
		Gids:      make([]string, 0),
	}
}
