package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Employee struct {
	Person
	ManageID int
}
type Contractor struct {
	Person
	CompanyID int
}
type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

func struct12() {
	//arrayXp() //arrInit() //arrSpec() //twoArr() //threeArr() //slice1() //removeE() //slice2() //slice3()
	//slice4() //map1() //map2() //map3() //map4() //map5()
	//struct1() struct2()
	//var slice = fi(1)
	//fmt.Println(slice)
	//fmt.Println(luomaNumber("MCLX"))
	str := "123sg"
	for index, c := range str {
		fmt.Println(index, c)
	}
}

func romanToArabic(numerial string) int {
	//romanMap:= map[rune]int{'M':1000, 'D':500, 'C':100, 'L':50, 'X':10, 'V':5, 'I':1}

	return 0
}

func luomaNumber(str string) int {
	maps := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	split := strings.Split(str, "")
	var before, sum int
	for i := 0; i < len(split); i++ {
		value, exist := maps[split[i]]
		if exist {
			if i > 0 {
				if before > value {
					sum += value
				} else {
					sum += value - 2*before
				}
			} else {
				sum = value
			}
			before = value
		} else {
			panic("不存在的罗马字符")
		}
	}
	return sum
}

func fibonacci(size int) []int {
	if size < 2 {
		return make([]int, 0)
	}
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		if i < 2 {
			ints[i] = 1
		} else {
			ints[i] = ints[i-1] + ints[i-2]
		}
	}
	return ints
}

func struct2() {
	employees := []Employee{
		{Person: Person{LastName: "Doe", FirstName: "John"}},
		{Person: Person{LastName: "Campbell", FirstName: "David"}},
	}
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)
	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}

func struct1() {
	employee := Employee{Person: Person{FirstName: "john"}}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}
func map5() {
	studentAge := make(map[string]int)
	studentAge["john"] = 32
	studentAge["bob"] = 31
	for name, age := range studentAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
func map4() {
	studentAge := make(map[string]int)
	studentAge["john"] = 31
	studentAge["bob"] = 32
	fmt.Println(studentAge)
	age, exist := studentAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	//delete(studentAge,"john")
	delete(studentAge, "christy")
	fmt.Println(studentAge)

}

func map3() {
	// 不要给 nil 映射  会报错
	var studentAge map[string]int
	studentAge["john"] = 32
	studentAge["bob"] = 31
	fmt.Println(studentAge)
}

func map2() {
	studentAge := make(map[string]int)
	studentAge["john"] = 32
	studentAge["bob"] = 31
	fmt.Println(studentAge)

}
func map1() {
	studentAge := map[string]int{
		"John": 32,
		"bob":  31,
	}
	fmt.Println(studentAge)
}

func slice4() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)
	slice1 := letters[0:2]
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])
	slice1[1] = "Z"
	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

}

func slice3() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)
	slice1 := letters[0:2]
	slice2 := letters[1:4]
	slice1[1] = "Z"
	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

}

func slice2() {
	letters := []string{"A", "B", "C", "D", "E"}
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])
	fmt.Println(slice2)

}

func removeE() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	fmt.Println("before", letters)
	letters[remove] = letters[len(letters)-1]
	letters = letters[:len(letters)-1]
	fmt.Println("after", letters)
}

func slice1() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

func threeArr() {
	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)
}

func twoArr() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)
}

func arrSpec() {
	numbers := [...]int{99: -1}
	fmt.Println("first position:", numbers[0])
	fmt.Println("last position", numbers[99])
	fmt.Println("Length", len(numbers))

}

func arrInit() {
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	strings := cities[1:]
	fmt.Println(strings)

}

func arrayXp() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

}
