package main

import "fmt"
func CamelToSnakeCase(s string) string {
	var result string

	for i := range s {
		if i > 0 && 'A' <= s[i] && s[i] <= 'Z' {
			if 'A' <= s[i-1] && s[i-1] <= 'Z' {
				result = s

			} else if i == len(s)-1 && ('A' <= s[i] && s[i] <= 'Z') {
				result = s
			} else {
				result += "_" + string(s[i])
			}

		} else {
			result += string(s[i])
		}
	}

	return result
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	fmt.Println(CamelToSnakeCase("heyT"))

}