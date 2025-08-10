package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	ping := "+PING\r\n"

	echo := "*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"

	fmt.Println(parser(ping))

	fmt.Println(parser(echo))

}

//The tester will expect to receive $3\r\nhey\r\n
//as a response (that's the string hey encoded as a RESP bulk string.

//The exact bytes your program will receive won't be just ECHO hey,
// you'll receive something like this: *2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n.
// That's ["ECHO", "hey"] encoded using the Redis protocol.

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// elete is O(len(s)-i), so if many items must be deleted, it is better to make a single call
// deleting them all together than to delete one at a time. Delete zeroes the elements s[len(s)-(j-i):len(s)].

func parser(s string) []string {

	if s == "" {
		fmt.Println("empty string niBBa")
		return nil
	}

	data_type := string(s[0])

	var result []string
	switch data_type {

	case "+":
		fmt.Println("simple string")

		s = strings.TrimLeft(s, "+")

	case "*":
		fmt.Println("array")
		num_args, err := strconv.Atoi(string(s[1]))
		if err != nil {
			fmt.Println("error in conversion")
		}
		s = strings.TrimLeft(s, "*")
		s = strings.TrimLeft(s, strconv.Itoa(num_args))

	case "$":
		fmt.Println("bulk string")
		s = strings.TrimLeft(s, "$")
	}
	s = strings.TrimLeft(s, "\r\n")
	s = strings.TrimRight(s, "\r\n")
	fmt.Println(s)
	p := strings.Split(s, "\r\n")

	fmt.Println(p, len(p))

	for _, k := range p {

		if string(k[0]) != "$" {
			result = append(result, k)
		}

	}

	return result
}

// num_args, err := strconv.Atoi(p[0])
// if err != nil {
// 	fmt.Println("error in conversion")
// }
