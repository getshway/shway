package homebrew

import (
	"github.com/getshway/shway/sidekick/shell"
)

// Cleanup cleans up caches and downloaded files associated with homebrew
func Cleanup() (err error) {
	return shell.Run("brew", "cleanup")
}

// Install a package by homebrew
func Install(name string) (err error) {
	return shell.Run("brew", "install", name)
}

// InstallCask a package by homebrew-cask
func InstallCask(name string) (err error) {
	return shell.Run("brew", "cask", "install", name)
}

// Update homebrew information
func Update() (err error) {
	return shell.Run("brew", "update")
}

// Upgrade installed packages
func Upgrade() (err error) {
	return shell.Run("brew", "upgrade")
}

// Tap a homebrew formula repository
func Tap(name string) (err error) {
	return shell.Run("brew", "tap", name)
}
