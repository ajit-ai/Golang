package main

import (
	"testing"
	"time"
)

func TestCacheSetGet(t *testing.T) {
	cache := NewCache()
	cache.SetValue("name", "ajit kumar", time.Minute)
	if got := cache.GetObject("name"); got != "ajit kumar" {
		t.Errorf("GetObject after SetValue = %q, want %q", got, "ajit kumar")
	}
}

func TestCacheExpiredEntryReturnsEmptyAndIsDeleted(t *testing.T) {
	cache := NewCache()
	cache.SetValue("temp", "value", -time.Second)
	if got := cache.GetObject("temp"); got != "" {
		t.Errorf("GetObject with expired TTL = %q, want empty", got)
	}
	cache.mutex.RLock()
	_, exists := cache.objects["temp"]
	cache.mutex.RUnlock()
	if exists {
		t.Error("expired entry was not deleted from cache")
	}
}

func TestCacheMissingKeyReturnsEmpty(t *testing.T) {
	cache := NewCache()
	if got := cache.GetObject("missing"); got != "" {
		t.Errorf("GetObject for missing key = %q, want empty", got)
	}
}

func TestCacheObjectIfExpired(t *testing.T) {
	tests := []struct {
		name string
		ttl  int64
		want bool
	}{
		{"zero ttl never expires", 0, false},
		{"past timestamp expired", time.Now().Add(-time.Minute).UnixNano(), true},
		{"future timestamp valid", time.Now().Add(time.Minute).UnixNano(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := CacheObject{Value: "v", TimeToLive: tt.ttl}
			if got := o.IfExpired(); got != tt.want {
				t.Errorf("IfExpired(TimeToLive=%d) = %v, want %v", tt.ttl, got, tt.want)
			}
		})
	}
}

func TestCacheManagementMainSmoke(t *testing.T) {
	CacheManagementMain()
}
