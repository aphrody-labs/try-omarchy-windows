//go:build windows

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStockQEMUUsesMSYS2UCRT64(t *testing.T) {
	root := t.TempDir()
	qemu := filepath.Join(root, "ucrt64", "bin", "qemu-system-x86_64w.exe")
	if err := os.MkdirAll(filepath.Dir(qemu), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(qemu, []byte("fixture"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TRYOMARCHY_MSYS2_ROOT", root)
	t.Setenv("TRYOMARCHY_QEMU", "")
	if got := stockQEMUPath(); got != qemu {
		t.Fatalf("stockQEMUPath() = %q, want %q", got, qemu)
	}
}

func TestExplicitQEMUWinsOverMSYS2(t *testing.T) {
	qemu := filepath.Join(t.TempDir(), "qemu-system-x86_64w.exe")
	if err := os.WriteFile(qemu, []byte("fixture"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TRYOMARCHY_QEMU", qemu)
	if got := stockQEMUPath(); got != qemu {
		t.Fatalf("stockQEMUPath() = %q, want %q", got, qemu)
	}
}
