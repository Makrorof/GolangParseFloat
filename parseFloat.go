package Utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}

func Must(v float64, err error) float64 {
	if err != nil {
		panic(err)
	}
	return v
}

func MustEqual(v1, v2 float64) {
	if v1 != v2 {
		panic(fmt.Errorf("number is not equal, v1=%f, v2=%f", v1, v2))
	} else {
		fmt.Printf("%f = %f\n", v1, v2)
	}
}

//func main() {
//	MustEqual(Must(ParseFloat("1E2")), 100)
//	MustEqual(Must(ParseFloat("1E-5")), 0.00001)
//	MustEqual(Must(ParseFloat("1.6543E2")), 165.43)
//	MustEqual(Must(ParseFloat("0.89E2")), 89)
//	MustEqual(Must(ParseFloat("1.6543E-2")), 0.016543)
//	MustEqual(Must(ParseFloat("156,819,129")), 156819129)
//	MustEqual(Must(ParseFloat("156819129")), 156819129)
//	MustEqual(Must(ParseFloat(".1E0")), 0.1)
//	MustEqual(Must(ParseFloat(".1E1")), 1)
//	MustEqual(Must(ParseFloat("0E1")), 0)
//
//	fmt.Printf("Success\n")
//}
