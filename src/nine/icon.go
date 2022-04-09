package main

import (
	"math/rand"
	"sync"
)

var icons map[string]bool
var mut sync.RWMutex

func loadIcons() {
	icons = map[string]bool{
		"spades.png":   loadIcon("spade.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(s string) bool {
	x := rand.Intn(2)
	if x >= 1 {
		return true
	}
	return false
}

func Icon(name string) bool {
	mut.RLock()
	if icons != nil {
		mut.RUnlock()
		return icons[name]
	}
	mut.RUnlock()

	mut.Lock()
	if icons == nil {
		loadIcons()
	}
	mut.Unlock()
	return icons[name]
}
