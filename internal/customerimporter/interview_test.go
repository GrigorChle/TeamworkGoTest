package customerimporter

import "testing"

func TestCastDomainToSortedStruct(t *testing.T) {
	test_map := map[string]int{
		"yahoo":  10,
		"google": 15,
		"9gag":   3,
		"custom": 2,
	}
	sortedDomains := castDomainToSortedStruct(test_map)
	expectedValue := 15
	if sortedDomains[0].count != expectedValue {
		t.Errorf("Expected google with highest %v count", expectedValue)
	}
}

func TestSplittingDomain(t *testing.T) {
	expectedDomain := "eepurl.com"
	expectedValue := 1

	test_string := "aarnoldv@eepurl.com"
	test_map := map[string]int{}

	fillDomainMap(test_string, test_map)
	v, ok := test_map[expectedDomain]
	if ok != true {
		t.Errorf("Expected domain %v not parsed", expectedDomain)
	}
	if v != 1 {
		t.Errorf("Expected value %v, got %v instead", expectedValue, v)
	}
}

func TestTestSplittingDomainIncrement(t *testing.T) {
	expectedDomain := "eepurl.com"
	expectedValue := 190

	test_string := "aarnoldv@eepurl.com"
	test_map := map[string]int{
		"eepurl.com": 189,
	}

	fillDomainMap(test_string, test_map)
	v := test_map[expectedDomain]
	if v != expectedValue {
		t.Errorf("Expected value %v, got %v instead", expectedValue, v)
	}
}
