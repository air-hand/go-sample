package fundamentals

import "testing"

func Test_ChannelSample(t *testing.T) {
	expected_max := 5
	for i := 0; i <= 100; i++ {
		number := ChannelSample(expected_max)
		if number >= expected_max {
			t.Errorf("Expected max %d, but Actual %d", expected_max, number)
		}
	}
}
