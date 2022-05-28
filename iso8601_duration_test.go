package iso8601_duration

import "testing"

func TestDurationStrToSeconds(t *testing.T) {
	testcases := []struct {
		in   string
		want uint32
	}{
		{in: "PT15M51S", want: 951},
		{in: "P3Y6M4DT12H30M5S", want: 207491405},
		{in: "P2Y4M2W6DT14H30M2042S", want: 139849442},
	}
	for _, tc := range testcases {
		t.Run(tc.in, func(t *testing.T) {
			actual := DurationStrToSeconds(tc.in)
			if actual != tc.want {
				t.Errorf("Got %d want %d", actual, tc.want)
			}
		})
	}

}
