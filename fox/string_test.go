package fox

import "testing"

func TestPadc(t *testing.T) {
	t.Log(Padc("1", 3, "0"))
	t.Log(Padr("1", 3, "0"))
	t.Log(Padl("1", 3, "0"))
}
