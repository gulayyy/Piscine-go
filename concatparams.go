package piscine

func ConcatParams(args []string) string {
	s := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			s += args[i]
		} else {
			s += args[i] + "\n"
		}
	}
	return s
}
