
package main 

import "fmt"

type Date struct {
	d,m,y int
}


func countleap(d Date) int {
	year := d.y
	if d.m <= 2 {
		year--
	}
	return (year/4 - year/100 + year/400)
}

func isleapyear(y int) bool {
	if y % 400 == 0 {
		return true
	}
	if y % 100 == 0 {
		return false
	}
	if y % 4 == 0 {
		return true
	}
	return false
}

func daysofmonth(d Date) int{
	if d.m == 2 {
		if isleapyear(d.y) {
			return 29
		}
		return 28
	}
	if d.m ==4 || d.m ==6 || d.m ==9 || d.m ==11 {
		return 30
	}
	return 31
}

func (d Date) addDate() Date { 
	if d.d < daysofmonth(d){
		d.d += 1
		return Date{d.d, d.m,d.y}
	}
	if d.m < 12 {
		d.m += 1
		return Date{1,d.m,d.y}
	}
	d.y += 1
	return Date{1, 1,d.y}
}

func main() {
	d1 := Date{d:31, m:3, y:1992}
	d2 := Date{d:2, m:3, y:2016}
	var count int
	for {
		d1 = d1.addDate()
		count++
		if d1 == d2 {
			break
		}
	}
	//fmt.Println(countleap(d1))
	fmt.Println(count)
}

