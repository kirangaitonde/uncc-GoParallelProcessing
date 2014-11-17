
// Parallel merge sort running on diffrent number of processors
// Reads the input from an xlsx file
// @ Kiran Gaitonde
// References :  
//  https://golang.org/
//  http://www.golang-book.com/
//  https://groups.google.com/forum/#!topic/golang-nuts
//  https://github.com/tealeg/xlsx
//


package main

import (
    "sync"
	"time"	
	"strconv"    
	"fmt"
	"xlsx"
	"runtime"
	"sort"	
	
)

func main() {
    
	
    var X [1000500]int  //give size here
	size := len(X) 
	
	//getting the inputs from excel file and storing in an array
    xl := "C:/Users/K G/Desktop/ip.xlsx"
    xlFile, err := xlsx.OpenFile(xl)
    if err != nil {
        
    }
	
	var i int = 0
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
			    n,err :=  strconv.Atoi(cell.String())
				if err != nil {
        
				}
                //fmt.Println("%v",n)
				X[i] = n
				i++
            }
        }
    } 
	
	A := X[0:size] 
	//fmt.Println(A)
	B := make([]int, size)
	C := make([]int, size)
	

	//sort using diffrent number of processors and display time taken to sort
	fmt.Printf("No_of_CPU\tTime_taken\n")
	for noOfCPU := 1; noOfCPU <= runtime.NumCPU(); noOfCPU++ {
		copy(B, A)
		runtime.GC()
		runtime.GOMAXPROCS(noOfCPU)
		start := time.Now()
		parallelSort(B, C, 0, len(B))
		//fmt.Println(B)
		duration := time.Since(start)		
		t := duration.Seconds()
		//fmt.Println(D)
		
		fmt.Printf("%d\t\t%.3f\n", noOfCPU, t)
		
		if !sort.IntsAreSorted(B) {
			fmt.Printf("multisort(%d): FAIL!\n", noOfCPU)
		}
	}
}

//parallel sort
func parallelSort(b, c []int, p, q int) {
	
	// normal sort
	if q-p <= 4096 {
		sort.Ints(b[p:q])
		return
	}

	// Divide the input into four parts.
	n0 := p + (q-p)/4
	n1 := n0 + (q-p)/4
	n2 := n1 + (q-p)/4

    //  sort four parts parallelly
    var wg sync.WaitGroup
	wg.Add(4)
	go func() { parallelSort(b, c, p, n0); wg.Done() }()
	go func() { parallelSort(b, c, n0, n1); wg.Done() }()
	go func() { parallelSort(b, c, n1, n2); wg.Done() }()
	go func() { parallelSort(b, c, n2, q); wg.Done() }()
	wg.Wait()
    
	
	wg.Add(2)
	go func() { parallelMerge(b[p:n0], b[n0:n1], c[p:n1]); wg.Done() }()
	go func() { parallelMerge(b[n1:n2], b[n2:q], c[n1:q]); wg.Done() }()
	wg.Wait()

	parallelMerge(c[p:n1], c[n1:q], b[p:q])
}


//parallel merge
func parallelMerge(m1, m2 []int, m []int) {
	
	if len(m2) > len(m1) {
		m1, m2 = m2, m1
	}

	// Normal merge
	if len(m1) <= 4096 {
		merge(m1, m2, m)
		return	
		
	}	
	x := len(m1) / 2
	y := sort.SearchInts(m2, m1[x])

    var wg sync.WaitGroup
	wg.Add(2)
	go func() { parallelMerge(m1[:x], m2[:y], m[:x+y]); wg.Done() }()
	go func() { parallelMerge(m1[x:], m2[y:], m[x+y:]); wg.Done() }()
	wg.Wait()
}

//normal merge
func merge(m1, m2 []int, m []int) {		
		i, j, k := 0, 0, 0
		for ; i < len(m1) && j < len(m2); k++ {
			if m1[i] < m2[j] {
				m[k] = m1[i]
				i++
			} else {
				m[k] = m2[j]
				j++
			}
		}
		if i < len(m1) {
			copy(m[k:], m1[i:])
		} else {
			copy(m[k:], m2[j:])
		}
	}






