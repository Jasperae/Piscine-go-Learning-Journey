package piscine

func LoafOfBread(str string) string {
	myStr := ""
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "invalid Output\n"
	}
	f := 0
	for _, c := range str {
		if c != ' ' && f != 5 {
			myStr += string(c)
			f++
		} else if f == 5 {
			myStr += " "
			f = 0
		}
	}
	if len(myStr) > 0 && myStr[len(myStr)-1] == ' ' {
		myStr = myStr[:len(myStr)-1]
	}
	return myStr + "\n"
}
