package lib

import "regexp"
import "fmt"

func max(a int, b int) int{
	if a > b {
		return a;
	} else {
		return b;
	}
}


func Lcs(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)
	L := make([][]int, m+1)   
	
	result := 0
	for i := 0; i < m+1; i ++ {
		L[i] = make([]int, n+1)  
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0;
			} else if str1[i-1] == str2[j-1] {
				L[i][j] = L[i-1][j-1] + 1
				result = max(result,L[i][j]);
			} else {
				L[i][j] = 0;
			}
		}
	}

	var percentage float64;
	percentage = float64(result) / float64(n)

	return int(percentage * 100);
}

func Regex(str1 string) bool {

	if (len(str1) == 0) {
		return false
	}

	var regex, err = regexp.Compile(`^[AGTC]*$`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return regex.MatchString(str1) 
}

func RegexSearch(str1 string, str2 string) bool {

	var regex, err = regexp.Compile(`.*` + str1 + `.*`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return regex.MatchString(str2) 
}

func RegexTanggal(str1 string) bool {

	var regex, err = regexp.Compile(`^[0-9][0-9][0-9][0-9]-[0-9][0-9]?-[0-9][0-9]?$`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return regex.MatchString(str1) 
}