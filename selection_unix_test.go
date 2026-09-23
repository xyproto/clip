//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly
// +build freebsd linux netbsd openbsd solaris dragonfly

package clip

import (
	"strings"
	"testing"
)

func sameArgs(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestPrimarySelectionKeepsClipboardArgs checks that asking for the primary
// selection does not modify the arguments used for the regular clipboard.
func TestPrimarySelectionKeepsClipboardArgs(t *testing.T) {
	if Unsupported {
		t.Skip("no clipboard utilities available")
	}

	pasteBefore := append([]string{}, getPasteCommand(false).Args...)
	getPasteCommand(true)
	if pasteAfter := getPasteCommand(false).Args; !sameArgs(pasteBefore, pasteAfter) {
		t.Errorf("paste arguments changed after a primary selection read: %q, want %q", strings.Join(pasteAfter, " "), strings.Join(pasteBefore, " "))
	}

	copyBefore := append([]string{}, getCopyCommand(false).Args...)
	getCopyCommand(true)
	if copyAfter := getCopyCommand(false).Args; !sameArgs(copyBefore, copyAfter) {
		t.Errorf("copy arguments changed after a primary selection write: %q, want %q", strings.Join(copyAfter, " "), strings.Join(copyBefore, " "))
	}
}
