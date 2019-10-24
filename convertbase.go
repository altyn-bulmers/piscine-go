package piscine

func ConvertNbrBase(n int, base string) string {
	result := ""
	length := Length(base)

	for n >= length {
		result = result + string(base[(n%length)])
		n = n / length
	}
	result = result + string(base[n])

	resultReversed := ""
	for _, s := range result {
		resultReversed = string(s) + resultReversed
	}

	return (resultReversed)

}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	resultIntermediary := AtoiBase(nbr, baseFrom)

	resultFinal := ConvertNbrBase(resultIntermediary, baseTo)

	return resultFinal

}
