// package for testing of worldstate, Putin & Crossriver
//


package main
import "testing"

// test WorldState
func TestReturnWorldState(t *testing.T){
	wanted:= "Mann-Korn-Kylling-Rev- \\_ _ _/________ - - - -"
	got := ReturnWorldState()
	if(wanted != got){
		t.Errorf("MakeWorldState Failed, Wanted: %s, Got %s", wanted, got)
	}
}

// test PutIn
func TestPutIn(t *testing.T){
	wanted:= " -Korn-Kylling-Rev- \\_Mann_ _/________ - - - -"
	got := PutIn("Mann")
	if(wanted != got){
		t.Errorf("PutIn Failed, Wanted: %s, Got %s", wanted, got)
	}
}

// test CrossRiver
func TestCrossRiver(t *testing.T){
	wanted := " -Korn-Kylling-Rev-________\\_Mann_ _/   - - - -"
	got := CrossRiver()
	if(wanted != got){
		t.Errorf("CrossRiver Failed, Wanted: %s, Got %s", wanted, got)
	}

}
