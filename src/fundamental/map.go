package fundamental

import "fmt"

func MapTest(){
	var countryCapital map[string]string
	countryCapital =make(map[string]string)
	countryCapital["china"] = "beijing"
	countryCapital["usa"] = "new york"
	countryCapital["england"] = "london"
	countryCapital["france"] = "paris"

	for country := range countryCapital {
		fmt.Println("Capital of",country,"is",countryCapital[country])
	}

	captial, ok := countryCapital["england"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	fmt.Println(captial,ok)

}