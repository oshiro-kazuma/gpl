package tempconv

import "testing"

func TestKToC(t *testing.T) {
	want := KToC(Kelvin(0))
	got := Celsius(-273.15)
	if want != got {
		t.Errorf("got: %f, want: %f", got, want)
	}
}
