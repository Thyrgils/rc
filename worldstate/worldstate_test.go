package main

import "testing"

func TestReturnWorldState(t *testing.T){
	wanted:= "Mann-Korn-Kylling-Rev- \\_ _ _/________ - - - -"
	got := ReturnWorldState()
	if(wanted != got){
		t.Errorf("MakeWorldState Failed, Wanted: %s, Got %s", wanted, got)
	}
}
