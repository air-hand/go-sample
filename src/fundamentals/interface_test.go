package fundamentals

import "testing"

func Test_InterfaceDogSample(t *testing.T) {
	dog := Dog{
		Name:  "Bull",
		Breed: "Bulldog",
	}

	dog_info := AnimalInfo(&dog)
	dog_expected := "This animal roar 'Bark' and has 4 legs."
	if dog_info != dog_expected {
		t.Logf("Expected %s, Actual %s", dog_expected, dog_info)
		t.FailNow()
		return
	}
}
