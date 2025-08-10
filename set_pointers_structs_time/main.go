package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type record struct {
	key           string
	value         string
	expire        int64
	creation_time time.Time
}

func NewRecord(k string, v string, ex int64, crtime time.Time) *record {

	p := record{key: k, value: v, expire: ex, creation_time: crtime}

	return &p
}

func main() {

	var memory []record

	testSet1 := []string{"SET", "konik", "polny"}
	testSet2 := []string{"SET", "świnka", "morska", "px", "1000"}
	//testSet3 := []string{"SET", "hipopotam", "", "px"}

	testGet1 := []string{"GET", "konik"}
	testGet2 := []string{"GET", "świnka"}
	// testGet3 := []string{"GET", "hipopotam"}

	fmt.Println(memory)

	_, err := SET(testSet2, &memory)
	if err != nil {
		fmt.Println("error setting")
	}

	_, err = SET(testSet1, &memory)
	if err != nil {
		fmt.Println("error setting")
	}

	fmt.Println(memory)

	g, err := GET(testGet1, memory)
	if err != nil {
		fmt.Println("error getting")
	}

	println(g)
	time.Sleep(100 * time.Millisecond)
	g, err = GET(testGet2, memory)
	if err != nil {
		fmt.Println("error getting")
	}

	println(g)
}

func SET(s []string, m *[]record) ([]record, error) {

	// fmt.Println("len: ", len(s))
	// fmt.Println(s)
	if len(s) == 3 {
		*m = append(*m, *NewRecord(s[1], s[2], 0, time.Now()))
		// fmt.Println("Appended (first case)")
	} else if len(s) < 3 {
		println("Not enough arguments for SET")
	} else if len(s) == 4 {
		println("Not enough arguments for SET with px or wrong command")
	} else if len(s) == 5 {
		if s[3] == "px" {
			px_int, err := strconv.Atoi(s[4])
			px := int64(px_int)
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil, errors.New("wrong argument for px must be a number")
			}

			*m = append(*m, *NewRecord(s[1], s[2], px, time.Now()))
			// fmt.Println("Appended (second case)")
		}
	}
	// fmt.Println("set return: ", *m)
	return *m, nil
}

func GET(s []string, m []record) (string, error) {

	nbs := "$-1\r\n" //null bulk string
	for _, r := range m {

		if r.key == s[1] {

			if r.expire == 0 {
				return ConvToRedisBulkString(r.value), nil

			}

			t := time.Now()

			//		fmt.Println("time now: ", t, "creation time: ", r.creation_time)
			elapsed := t.Sub(r.creation_time)
			//		fmt.Println("elapsed: ", elapsed)

			fmt.Println("elapsed milis: ", elapsed.Milliseconds(), "px", r.expire)
			is_expired := elapsed.Milliseconds() > r.expire
			if !is_expired {
				return ConvToRedisBulkString(r.value), nil
			} else {
				return nbs, nil
			}
		}

	}

	return "", nil
}

func ConvToRedisBulkString(s string) string {

	r := fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)

	return r
}
