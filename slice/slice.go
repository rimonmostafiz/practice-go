package _slice

import "fmt"

func CreateSlices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(a, "\nLength:", len(a), "Capacity:", cap(a))
	x := a
	x[10] = 0
	x[11] = 0
	fmt.Println(x)
	fmt.Println(a)

	b := a[:] // slice of all element
	fmt.Println(b)
	c := a[3:] // slice from 4th element to end
	fmt.Println(c)
	d := a[:6] // slice first 6 element
	fmt.Println(d)
	e := a[3:6] // slice the 4th, 5th and 6th element
	fmt.Println(e)

	slc := make([]int, 5)
	fmt.Println(slc, "\nLength:", len(slc), "Capacity:", cap(slc))

	slc1 := make([]int, 5, 100)
	fmt.Println(slc1, "\nLength:", len(slc1), "Capacity:", cap(slc1))

	sc := []int{}
	fmt.Println(sc, "\nLength:", len(sc), "Capacity:", cap(sc))
	sc = append(sc, 1, 2, 3, 4, 5)
	fmt.Println(sc, "\nLength:", len(sc), "Capacity:", cap(sc))

	sc = append(sc, []int{10, 20, 30, 40}...)
	fmt.Println(sc, "\nLength:", len(sc), "Capacity:", cap(sc))

	a = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b = a[1:]
	fmt.Println(b)
	c = a[:len(a)-1]
	fmt.Println(c)
	d = append(a[:2], a[3:]...)
	fmt.Println(d)

}
