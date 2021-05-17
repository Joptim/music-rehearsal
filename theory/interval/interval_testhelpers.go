package interval

import "testing"

func NewTestHelper(code string, t *testing.T) Interval {
	interval, err := New(code)
	if err != nil {
		t.Fatalf("cannot instantiate interval with name %s. Got error %v", code, err)
	}
	return interval
}
