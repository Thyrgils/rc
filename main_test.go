// package for testing of worldstate, Putin & Crossriver
//


package main
import "testing"

// test WorldState
func TestReturnWorldState(t *testing.T){
	wanted:= "Menneske-Korn-Kylling-Rev- \\_ _ _/________ - - - -"
	got := ReturnWorldState()
	if(wanted != got){
		t.Errorf("MakeWorldState Failed, Wanted: %s, Got %s", wanted, got)
	}
}

// test PutIn
func TestPutIn(t *testing.T){
	wanted:= " -Korn-Kylling-Rev- \\_Menneske_ _/________ - - - -"
	got := PutIn("Menneske")
	if(wanted != got){
		t.Errorf("PutIn Failed, Wanted: %s, Got %s", wanted, got)
	}
}

// test CrossRiver
func TestCrossRiver(t *testing.T){
	wanted := " -Korn-Kylling-Rev-________\\_Menneske_ _/   - - - -"
	got := CrossRiver()
	if(wanted != got){
		t.Errorf("CrossRiver Failed, Wanted: %s, Got %s", wanted, got)
	}

}
