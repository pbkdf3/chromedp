package storage

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventCacheStorageListUpdated a cache has been added/deleted.
type EventCacheStorageListUpdated struct {
	Origin string `json:"origin"` // Origin to update.
}

// EventCacheStorageContentUpdated a cache's contents have been modified.
type EventCacheStorageContentUpdated struct {
	Origin    string `json:"origin"`    // Origin to update.
	CacheName string `json:"cacheName"` // Name of cache in origin.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventStorageCacheStorageListUpdated,
	cdp.EventStorageCacheStorageContentUpdated,
}
