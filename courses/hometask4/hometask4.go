package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

//  ====== First task - basic structure
type Person struct {
	firstName string
	lastName string
	birthDay time.Time
}

type People []Person

func (p People) Len() int{
	return len(p)
}

func (p People) Less(i, j int) bool{

	// Compare by birth day
	if p[i].birthDay != p[j].birthDay{
		return p[i].birthDay.After(p[j].birthDay)
	}

	// Compare by first name
	if less:=strings.Compare(p[i].firstName, p[j].firstName); less != 0{
		switch less{
		    case -1: return true
		    case 1: return false
		}
	} // end if less ...

	// Compare by second name
	switch less:=strings.Compare(p[i].lastName, p[j].lastName); less{
	case -1: return true
	case 1: return false
		}

// They are absolutely equal
return false

} // end func Less

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// For proper time conversion
func birthDay (layout, value string) time.Time {

	date, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return date
}

// ======= Second task - Figure interface implementation

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {

	side float64
}

type Circle struct {

	radii float64
}

func (s Square) area() float64{

	return s.side*s.side
}

func (s Square) perimeter() float64{
	return s.side*4
}

func (c Circle) area() float64{

	return math.Pow(c.radii, 2.0)*math.Pi
}

func (c Circle) perimeter() float64{

	return 2*c.radii*math.Pi
}

func main() {

	// === First task - implementation of the Interface interface from Sort package
	layout := "2006-Jan-02"
	p := People{
		{"Ivan", "Ivanov", birthDay(layout, "2005-Aug-10")},
		{"Ivan", "Ivanov", birthDay(layout, "2003-Aug-05")},
		{"Tomas", "Geller", birthDay(layout, "2001-Sep-29")},
		{"Artiom", "Ivanov", birthDay(layout, "2005-Aug-10")},
		{"Hans", "Zimmer", birthDay(layout, "1995-Feb-12")},
		{"Alan", "Moore", birthDay(layout, "1992-Nov-17")},
		{"Adam", "Horovitz", birthDay(layout, "2009-Apr-25")},
	}

	sort.Sort(p)

	// Print the People elements in readable manner
	for _, value:= range p{
		fmt.Println(value.firstName, value.lastName, value.birthDay.Format(layout))
	}

	// === Second task - implementation of the Figure interface
	var s Figure = Square{10}
	var c Figure = Circle{10}

	fmt.Println("")
	fmt.Println("Square:", "area:", s.area(), "perimeter:", s.perimeter())
	fmt.Println("Circle: ", "area:", c.area(), "perimeter:", c.perimeter())

	// === Third task - Hello world with emoji
	fmt.Println("")
	emoji.Println("Hello world :smile:!")

}