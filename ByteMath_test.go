package bytemath

import (
	"github.com/spf13/cast"
	"testing"
)

func TestByteMath(t *testing.T) {
	t.Run("Test Kilobytes", func(t *testing.T) {
		var oneK int64 = 1024
		want := "1 KB"
		got := ConvertToHumanReadable(oneK)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Megabytes", func(t *testing.T) {
		mbs := int64(999 * 1024.0 * 1024.0)
		want := "999 MB"
		got := ConvertToHumanReadable(mbs)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	count := cast.ToFloat64(999)
	t.Run("Test convert byte to byte", func(t *testing.T) {

		suffix := B
		want := cast.ToFloat64(999)
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert kilobyte to byte", func(t *testing.T) {

		suffix := KB
		want := 1022976.0000
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert megabyte to byte", func(t *testing.T) {

		suffix := MB
		want := 1047527424.0000
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert gigabyte to byte", func(t *testing.T) {

		suffix := MB
		want := 1047527424.0000
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert gigabyte to byte", func(t *testing.T) {

		suffix := GB
		want := 1072668082176.0000
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert terrabyte to byte", func(t *testing.T) {

		suffix := TB
		want := 1098412116148224.0000
		got := ConvertToBytes(count, suffix)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

}
