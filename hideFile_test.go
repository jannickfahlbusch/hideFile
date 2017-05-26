package hideFile

import (
	"reflect"
	"testing"
)

func TestNewHider(t *testing.T) {
	hider := NewHider()

	if reflect.TypeOf(hider) != reflect.TypeOf(&Hider{}) {
		t.Fatalf("Returned the wrong type")
	}

}

func TestHider_Convert(t *testing.T) {
	// @ToDo
}

func TestHider_Deconvert(t *testing.T) {
	// @ToDO
}

func TestHider_GetType(t *testing.T) {
	hider := NewHider()

	testMagicNumber := magicNumber{
		Extension:   "test",
		Name:        "TEST",
		Number:      []byte{0x00},
		Description: "Test magic number",
		Offset: offset{
			Count: 1,
			Value: []byte{0x00},
		},
	}

	magicNumberList = []magicNumber{
		testMagicNumber,
	}

	returnedMagicNumber, err := hider.GetType(testMagicNumber.Name)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(returnedMagicNumber, testMagicNumber) {
		t.Fatalf("Didn't got the correct magicNumber")
	}
}

func TestHider_GetTypelist(t *testing.T) {
	hider := NewHider()

	lenGetTypeList := len(hider.GetTypelist())
	lenMagicNumberList := len(magicNumberList)

	if lenMagicNumberList != lenGetTypeList {
		t.Fatalf("Didn't got all types. Got %d types, expected %d.", lenGetTypeList, lenMagicNumberList)
	}
}
