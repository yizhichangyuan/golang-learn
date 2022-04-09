package icon_once

import (
	"math/rand"
	"sync"
)

var icons map[string]bool
var once sync.Once

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
	once.Do(loadIcons)
	return icons[name]
}
