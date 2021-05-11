package natural

import "testing"

func NewTestHelper(name string, t *testing.T) Natural {
	natural, err := New(name)
	if err != nil {
		t.Fatalf("cannot instantiate natural with name %s. Got error %v", name, err)
	}
	return natural
}
