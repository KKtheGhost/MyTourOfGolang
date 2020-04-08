//Slice defaults
//When slicing, you may omit the high or low bounds to use their defaults instead.
//
//The default is zero for the low bound and the length of the slice for the high bound.
//
//For the array
//
//var a [10]int
//these slice expressions are equivalent:
//
//a[0:10]
//a[:10]
//a[0:]
//a[:]
//其实这东西的套路和Python也一模一样，或者说几乎所有的语言在切片默认这一点上都是已完全一样的
package main

import "fmt"

var years = []int{2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020}

func main() {
	aa := years[0:11]
	bb := years[:11]
	cc := years[0:]
	dd := years[:]
	fmt.Println(aa, bb, cc, dd)
}
