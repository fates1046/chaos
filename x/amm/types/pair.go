package types

func SortDenoms(denomA, denomB string) (denom0, denom1 string) {
	if denomA < denomB {
		return denomA, denomB
	}

	return denomB, denomA
}
