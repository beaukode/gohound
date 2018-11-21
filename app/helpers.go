package app

import "github.com/google/uuid"

// GenerateLockUID return an unique id usable to lock probes
func GenerateLockUID() string {
	return uuid.New().String()
}
