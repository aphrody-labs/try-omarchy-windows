//go:build windows

package main

import (
	"os"
	"path/filepath"
)

func msys2QEMUCandidates() []string {
	root := os.Getenv("TRYOMARCHY_MSYS2_ROOT")
	if root == "" {
		root = `C:\msys64`
	}
	name := "qemu-system-x86_64w.exe"
	return []string{
		filepath.Join(root, "ucrt64", "bin", name),
		filepath.Join(root, "clang64", "bin", name),
		filepath.Join(root, "mingw64", "bin", name),
		filepath.Join(root, "usr", "bin", name),
	}
}

func stockQEMUPath() string {
	if explicit := os.Getenv("TRYOMARCHY_QEMU"); explicit != "" {
		if info, err := os.Stat(explicit); err == nil && info.Mode().IsRegular() {
			return explicit
		}
	}
	for _, candidate := range msys2QEMUCandidates() {
		if info, err := os.Stat(candidate); err == nil && info.Mode().IsRegular() {
			return candidate
		}
	}
	return `C:\Program Files\qemu\qemu-system-x86_64w.exe`
}
