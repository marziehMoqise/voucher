package utils

import (
	"github.com/ttacon/libphonenumber"
	"strconv"
	"strings"
)

func MobileValidate(mobile string) (bool, string) {
	mobile = ConvertFaToEn(mobile)
	mobileToInt, _ := ToInt64(mobile)
	if mobile == "" || mobileToInt == 0 {
		return false, ""
	}

	var (
		areaCode            int64
		foundedCountryCodes []string
		err                 error
	)

	for mobile[0] == '0' {
		mobile = mobile[1:]
	}

	if mobile[0] == '+' {
		mobile = mobile[1:]
	}
	if mobile[0] == '9' && len(mobile) == 10 {
		foundedCountryCodes = append(foundedCountryCodes, "IR")
	} else {
		for i := 1; i < 4 && len(mobile) > i; i++ {
			areaCode, err = ToInt64(mobile[:i])
			if err == nil {
				foundedCountryCodes = libphonenumber.CountryCodeToRegion[int(areaCode)]
				if foundedCountryCodes != nil {
					break
				}
			}
			if i == 3 {
				return false, ""
			}
		}
	}

	for _, code := range foundedCountryCodes {
		num, err := libphonenumber.Parse(mobile, code)
		if err == nil {
			valid := libphonenumber.IsValidNumber(num)
			if !valid {
				continue
			}
			numberType := libphonenumber.GetNumberType(num)
			if numberType == libphonenumber.FIXED_LINE {
				continue
			}
			formattedNum := libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
			formattedNum = strings.ReplaceAll(formattedNum, " ", "")
			formattedNum = strings.ReplaceAll(formattedNum, "-", "")
			return true, formattedNum
		}
	}

	return false, ""
}

func ConvertFaToEn(strNum string) string {
	persianReplacer := strings.NewReplacer("۰", "0", "۱", "1", "۲", "2", "۳", "3", "۴", "4", "۵", "5", "۶", "6", "۷", "7", "۸", "8", "۹", "9")
	enStrNum := persianReplacer.Replace(strNum)
	arabicReplacer := strings.NewReplacer("٤", "4", "٥", "5", "٦", "6")
	return arabicReplacer.Replace(enStrNum)
}

func ToInt64(i string) (int64, error) {
	v, err := strconv.ParseInt(i, 10, 64)
	if err == nil {
		return v, nil
	}
	return 0, err
}
