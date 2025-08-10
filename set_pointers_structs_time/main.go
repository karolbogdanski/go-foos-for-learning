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
	expire        int
	creation_time time.Time
}

func NewRecord(k string, v string, ex int, crtime time.Time) *record {

	p := record{key: k, value: v, expire: ex, creation_time: crtime}

	return &p
}

func main() {

	var memory []record

	testSet1 := []string{"SET", "konik", "polny"}
	testSet2 := []string{"SET", "świnka", " morska", "px", "1000"}
	//testSet3 := []string{"SET", "hipopotam", "", "px"}

	// testGet1 := []string{"GET", "konik"}
	// testGet2 := []string{"GET", "świnka"}
	// testGet3 := []string{"GET", "hipopotam"}

	//ogarnąć co zwraca time.Now

	// r1 := NewRecord("konik", "polny", 1000, time.Now())
	// r2 := NewRecord("świnka", " morska", 0, time.Now())

	// memory = append(memory, *r1)
	// memory = append(memory, *r2)

	fmt.Println(memory)

	// set, err := SET(testSet2, memory)
	// if err != nil {
	// 	fmt.Println("error setting")
	// }

	// memory = append(memory, set...)

	// set, err = SET(testSet1, memory)
	// if err != nil {
	// 	fmt.Println("error setting")
	// }

	// fmt.Println("Set: ", set)
	// memory = append(memory, set...)

	_, err := SETp(testSet2, &memory)
	if err != nil {
		fmt.Println("error setting")
	}

	// memory = append(memory, set...)

	_, err = SETp(testSet1, &memory)
	if err != nil {
		fmt.Println("error setting")
	}

	// fmt.Println("Set: ", set)
	// memory = append(memory, set...)

	fmt.Println(memory)

}

func SETp(s []string, m *[]record) ([]record, error) {

	// fmt.Println("len: ", len(s))
	// fmt.Println(s)
	if len(s) == 3 {
		*m = append(*m, *NewRecord(s[1], s[2], 0, time.Now()))
		fmt.Println("Appended (first case)")
	} else if len(s) < 3 {
		println("Not enough arguments for SET")
	} else if len(s) == 4 {
		println("Not enough arguments for SET with px or wrong command")
	} else if len(s) == 5 {
		if s[3] == "px" {
			px, err := strconv.Atoi(s[4])
			if err != nil {
				fmt.Println("Error converting string to int")
				return nil, errors.New("wrong argument for px must be a number")
			}

			*m = append(*m, *NewRecord(s[1], s[2], px, time.Now()))
			fmt.Println("Appended (second case)")
		}
	}
	fmt.Println("set return: ", *m)
	return *m, nil
}

//MUSISZ SIE NAUCZYĆ WSKAŹNIKÓW

// func SET(s []string, m []record) ([]record, error) {

// 	// fmt.Println("len: ", len(s))
// 	// fmt.Println(s)
// 	if len(s) == 3 {
// 		m = append(m, *NewRecord(s[1], s[2], 0, time.Now()))
// 		fmt.Println("Appended (first case)")
// 	} else if len(s) < 3 {
// 		println("Not enough arguments for SET")
// 	} else if len(s) == 4 {
// 		println("Not enough arguments for SET with px or wrong command")
// 	} else if len(s) == 5 {
// 		if s[3] == "px" {
// 			px, err := strconv.Atoi(s[4])
// 			if err != nil {
// 				fmt.Println("Error converting string to int")
// 				return nil, errors.New("wrong argument for px must be a number")
// 			}

// 			m = append(m, *NewRecord(s[1], s[2], px, time.Now()))
// 			fmt.Println("Appended (second case)")
// 		}
// 	}
// 	fmt.Println("set return: ", m)
// 	return m, nil
// }
