package restfight

import "testing"

func TestStatus(t *testing.T) {

	var status = GetStatus()
	if GetStatus() != 2 {
		t.Errorf("Error %d != %d", status, 2)
	}

}
