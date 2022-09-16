package bytemath

import (
	"testing"
)

var (
	byte     = 1
	kilobyte = byte * 1024
	megabyte = kilobyte * 1024
	gigabyte = megabyte * 1024
	terabyte = gigabyte * 1024
	petabyte = terabyte * 1024
)

func TestHumanReadable(t *testing.T) {

	t.Run("Test Bytes to Human-Readable Bytes", func(t *testing.T) {
		want := "1 B"
		got := ConvertBytesToHumanReadable(int64(byte))
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Bytes to Human-Readable Kilobytes", func(t *testing.T) {
		want := "1 KB"
		got := ConvertBytesToHumanReadable(int64(kilobyte))

		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Bytes to Human-Readable Megabytes", func(t *testing.T) {
		mbs := int64(megabyte)
		want := "1 MB"
		got := ConvertBytesToHumanReadable(mbs)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Bytes to Human-Readable Gigabytes", func(t *testing.T) {
		gbs := int64(gigabyte)
		want := "1 GB"
		got := ConvertBytesToHumanReadable(gbs)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Bytes to Human-Readable Terabytes", func(t *testing.T) {
		tbs := int64(terabyte)
		want := "1 TB"
		got := ConvertBytesToHumanReadable(tbs)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})

	t.Run("Test Bytes to Human-Readable Petabytes", func(t *testing.T) {
		pbs := int64(petabyte)
		want := "1 PB"
		got := ConvertBytesToHumanReadable(pbs)
		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})
}

func TestConvertToBytes(t *testing.T) {

	//test convert to bytes function
	t.Run("Test Convert Bytes to Bytes", func(t *testing.T) {
		want := float64(byte)
		got := ConvertToBytes(float64(byte), B)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test Convert Kilobytes to Bytes", func(t *testing.T) {
		want := float64(kilobyte)
		got := ConvertToBytes(float64(1), KB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Megabytes to Bytes", func(t *testing.T) {
		want := float64(megabyte)
		got := ConvertToBytes(float64(1), MB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Gigabytes to Bytes", func(t *testing.T) {
		want := float64(gigabyte)
		got := ConvertToBytes(float64(1), GB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Terabytes to Bytes", func(t *testing.T) {
		want := float64(terabyte)
		got := ConvertToBytes(float64(1), TB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test Convert Petabytes to Bytes", func(t *testing.T) {
		want := float64(petabyte)
		got := ConvertToBytes(float64(1), PB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})
}

func TestConvertBytesToOther(t *testing.T) {

	t.Run("Test convert Bytes to Bytes", func(t *testing.T) {
		var want = float64(byte)
		got := ConvertBytes(int64(byte), B)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Bytes to Kilobytes", func(t *testing.T) {
		var want float64 = 1.00
		got := ConvertBytes(int64(kilobyte), KB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Byte to Megabytes", func(t *testing.T) {
		var want float64 = 1.00
		got := ConvertBytes(int64(megabyte), MB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Bytes to Gigabytes", func(t *testing.T) {
		var want float64 = 1.00
		got := ConvertBytes(int64(gigabyte), GB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Bytes to Terabytes", func(t *testing.T) {
		var want float64 = 1.00
		got := ConvertBytes(int64(terabyte), TB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})

	t.Run("Test convert Bytes to Petabytes", func(t *testing.T) {
		var want float64 = 1.00
		got := ConvertBytes(int64(petabyte), PB)

		if got != want {
			t.Errorf("Wanted %f got %f", want, got)
		}
	})
}

func TestGetSuffixByValue(t *testing.T) {
	t.Run("Test Getting Suffix by String Value", func(t *testing.T) {
		suffixString := "KB"
		want := KB
		got := *GetSuffixByString(suffixString)
		if got != want {
			t.Errorf("Wanted %v got %v", want, got)
		}
	})
}
