package main

import "os"

func putStr(s string) {
	for i := 0; i < len(s); i++ {
		os.Stdout.Write([]byte{s[i]})
	}
	os.Stdout.Write([]byte{'\n'})
}

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	var n int64
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i++
	} else if s[0] == '+' {
		i++
	}
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int64(s[i]-'0')
	}
	if neg {
		n = -n
	}
	return n, true
}

func itoa(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var b [20]byte
	i := 20
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	a, ok1 := atoi(os.Args[1])
	b, ok2 := atoi(os.Args[3])
	op := os.Args[2]
	if !ok1 || !ok2 {
		return
	}
	switch op {
	case "+":
		res := a + b
		if (b > 0 && res < a) || (b < 0 && res > a) {
			return
		}
		putStr(itoa(res))
	case "-":
		res := a - b
		if (b < 0 && res < a) || (b > 0 && res > a) {
			return
		}
		putStr(itoa(res))
	case "*":
		res := a * b
		if a != 0 && res/a != b {
			return
		}
		putStr(itoa(res))
	case "/":
		if b == 0 {
			putStr("No division by 0")
			return
		}
		putStr(itoa(a / b))
	case "%":
		if b == 0 {
			putStr("No modulo by 0")
			return
		}
		putStr(itoa(a % b))
	}
}
