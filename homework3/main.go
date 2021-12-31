package main

import "fmt"

var suggestLevel1 string = "偏瘦"
var suggestLevel2 string = "标准"
var suggestLevel3 string = "偏重"
var suggestLevel4 string = "肥胖"
var suggestLevel5 string = "严重肥胖"

func CalculateBMI(stature float64, weight float64)(float64, error) {
	if stature <= 0 {
		error := fmt.Errorf("身高不能为负数")
		return 0, error
	}

	if weight <= 0 {
		error := fmt.Errorf("体重不能为负数")
		return 0, error
	}

	bmi := weight / (stature * stature)
	return bmi, nil
}

func CalculateFatRate(bmi float64, age int, gender string)(float64, string, error) {

	var suggest string = ""
	var genderType = 0
	// 校验 bmi
	if bmi <= 0 {
		error := fmt.Errorf("bmi录入不能为负数")
		return 0, suggest, error
	}

	// 校验年龄
	if age <= 0 || age > 150 {
		error := fmt.Errorf("年龄不能为负数或者大于150")
		return 0, suggest, error
	}

	// 校验性别
	if "b" == gender || "boy" == gender {
		genderType = 1
	} else if "g" == gender || "girl" == gender {
		genderType = 0
	} else {
		error := fmt.Errorf("性别输入无效 男: b boy  女: g girl")
		return 0, suggest, error
	}

	fatRate := 1.2 * bmi + 0.23 * float64(age) - 5.4 - 10.8 * float64(genderType)
	if genderType == 1 {
		suggest = boySuggest(age, fatRate)
	} else if genderType == 0 {
		suggest = girlSuggest(age, fatRate)
	}
	return fatRate, suggest, nil
}

// Body suggest
func boySuggest(age int, fatRate float64) string {
	var suggest string = ""
	// 男 体脂率建议
	if age >= 18 && age <=39 {
		if fatRate >= 0 && fatRate <= 10 {
			suggest = suggestLevel1
		} else if fatRate > 10 && fatRate <= 16 {
			suggest = suggestLevel2
		} else if fatRate > 16 && fatRate <= 21 {
			suggest = suggestLevel3
		} else if fatRate > 21 && fatRate <= 26 {
			suggest = suggestLevel4
		} else if fatRate > 26 {
			suggest = suggestLevel5
		}

	} else if age >= 40 && age <= 59 {
		if fatRate >= 0 && fatRate <= 11 {
			suggest = suggestLevel1
		} else if fatRate > 11 && fatRate <= 17 {
			suggest = suggestLevel2
		} else if fatRate > 17 && fatRate <= 22 {
			suggest = suggestLevel3
		} else if fatRate > 22 && fatRate <= 27 {
			suggest = suggestLevel4
		} else if fatRate > 27 {
			suggest = suggestLevel5
		}

	} else if age >= 60 {
		if fatRate >= 0 && fatRate <= 13 {
			suggest = suggestLevel1
		} else if fatRate > 13 && fatRate <= 19 {
			suggest = suggestLevel2
		} else if fatRate > 19 && fatRate <= 24 {
			suggest = suggestLevel3
		} else if fatRate > 24 && fatRate <= 29 {
			suggest = suggestLevel4
		} else if fatRate > 29 {
			suggest = suggestLevel5
		}

	}
	return suggest
}


// girl suggest
func girlSuggest(age int, fatRate float64) string {
	var suggest string = ""
	// 女 体脂率 建议
	if age >= 18 && age <= 39 {
		if fatRate >=0 && fatRate <= 20 {
			suggest = suggestLevel1
		} else if fatRate > 20 && fatRate <= 27 {
			suggest = suggestLevel2
		} else if fatRate > 27 && fatRate <= 34 {
			suggest = suggestLevel3
		} else if fatRate > 34 && fatRate <= 39 {
			suggest = suggestLevel4
		} else if fatRate > 39 {
			suggest = suggestLevel5
		}

	} else if age >= 40 && age <= 59 {
		if fatRate >= 0 && fatRate <= 21 {
			suggest = suggestLevel1
		} else if fatRate > 21 && fatRate <= 28 {
			suggest = suggestLevel2
		} else if fatRate > 28 && fatRate <= 35 {
			suggest = suggestLevel3
		} else if fatRate > 35 && fatRate <= 40 {
			suggest = suggestLevel4
		} else if fatRate > 40 {
			suggest = suggestLevel5
		}

	} else if age >= 60 {
		if fatRate >= 0 && fatRate <= 22 {
			suggest = suggestLevel1
		} else if fatRate > 22 && fatRate <= 29 {
			suggest = suggestLevel2
		} else if fatRate > 29 && fatRate <= 36 {
			suggest = suggestLevel3
		} else if fatRate > 36 && fatRate <= 41 {
			suggest = suggestLevel4
		} else if fatRate > 41 {
			suggest = suggestLevel5
		}

	}
	return suggest
}
