package zshhist

import (
	"testing"
)

func TestMetafy(t *testing.T) {
	testCases := []struct {
		str    string
		length int
	}{
		{"Hello, world!", 13},
		{"echo こんにちは、世界", 35},
		{"echo '😊'", 14},
	}
	for _, testCase := range testCases {
		metafied := Metafy(testCase.str)
		if len(metafied) != testCase.length {
			t.Errorf("metafied string should have length %d but got: %d", testCase.length, len(metafied))
		}
		got := Unmetafy(metafied)
		if testCase.str != got {
			t.Errorf("%q should be %q", got, testCase.str)
		}
	}
}
