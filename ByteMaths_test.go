package ByteMaths

import (
	"testing"
)

func TestHumanReadable(t *testing.T) {
	t.Run("Test Kilabytes", func(t *testing.T) {
		oneK := 1024.0
		want := "1 KB"
		got := ToHuman(oneK)
		if(got != want) {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Megabytes", func(t *testing.T) {
		mbs := 999 * 1024.0 * 1024.0
		want := "999 MB"
		got := ToHuman(mbs)
		if(got != want) {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})
}

