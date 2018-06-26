package restore_ip_addresses

import (
	"testing"
)

func TestRestoreIPAddresses(t *testing.T) {
	t.Run("25525511135", func(t *testing.T) {
		s := "25525511135"
		t.Log(restoreIpAddresses(s))
	})

	t.Run("010010", func(t *testing.T) {
		s := "010010"
		t.Log(restoreIpAddresses(s))
	})
}
