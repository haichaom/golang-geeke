package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type data struct {
	name string
}

/* type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/

type Person struct {
	Name   string
	Sexual string
	Age    int
}

func PrintPerson(p *Person) {
	fmt.Printf("PrintPerson: Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}

func (p *Person) Print() {
	fmt.Printf("Print: Name=%s, Sexual=%s, Age=%d\n", p.Name, p.Sexual, p.Age)
}

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr() {
	fmt.Println(c.Name)
}

type WithName struct {
	Name string
}

type CountryW struct {
	WithName
}
type CityW struct {
	WithName
}

func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

type Stringable interface {
	ToString() string
}

func (c Country) ToString() string {
	return "Country = " + c.Name
}

func (c City) ToString() string {
	return "City = " + c.Name
}

func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

type Shape interface {
	Slides() int
	Area() int
}

type Square struct {
	len int
}

func (s *Square) Slides() int {
	return 4
}

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	fmt.Println("foo", foo)
	bar := foo[1:4]
	bar[1] = 99
	fmt.Println("bar", bar)

	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println("a", a, cap(a), len(a))
	fmt.Println("b", b, cap(b), len(b))

	path := []byte("AAAA/BBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	//dir1 := path[:sepIndex]
	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1), cap(dir1), len(dir1))
	fmt.Println("dir2 =>", string(dir2), cap(dir2), len(dir2))
	dir1 = append(dir1, "suffix"...)
	fmt.Println("dir1 =>", string(dir1), cap(dir1), len(dir1))
	fmt.Println("dir2 =>", string(dir2), cap(dir2), len(dir2))

	v1 := data{"a"}
	v2 := data{"b"}
	fmt.Println("v1 === v2", reflect.DeepEqual(v1, v2))

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2", reflect.DeepEqual(m1, m2))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2", reflect.DeepEqual(s1, s2))

	p := Person{
		Name:   "Haichao Ma",
		Sexual: "Male",
		Age:    44,
	}
	PrintPerson(&p)
	p.Print()

	c1 := Country{"China"}
	c2 := City{"Beijing"}
	c1.PrintStr()
	c2.PrintStr()

	c1w := CountryW{WithName{"Chinaw"}}
	c2w := CityW{}
	c2w.Name = "Beijingw"
	c1w.PrintStr()
	c2w.PrintStr()

	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)

	s := Square{len: 5}
	fmt.Printf("%d", s.Slides())
	//var _ Shape = (*Square)(nil)
}
