package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Logf("1 == 1")
	}

	/*
	if support receive return many value
	if v, err = someFun() ; err == nil{
		something code
	}else{
		something code
	}
	*/
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("i is not longer between two and three")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		//t.Log(i)
		switch {
		case i%2 == 0:
			t.Log("Even")
			/* go is not need write break if program
			execute here will be auto exit or quit */
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}