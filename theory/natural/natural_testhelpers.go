package natural

import "testing"

func NewNaturalTestHelper(name string, t *testing.T) Natural {
	natural, err := NewNatural(name)
	if err != nil {
		t.Fatalf("cannot instantiate natural with name %s. Got error %v", name, err)
	}
	return natural
}
