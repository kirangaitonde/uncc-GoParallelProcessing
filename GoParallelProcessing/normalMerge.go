
// Merge sort without controlling the number of processors
// Reads the input from an xlsx file
// @ Kiran Gaitonde
// References :  
//  https://golang.org/
//  http://www.golang-book.com/
//  http://codereview.stackexchange.com/
//  https://github.com/tealeg/xlsx
//


package main

import (
    "fmt"
	"xlsx"
	"strconv"
	"time"
)

func main() {
    
	var X [1000500]int  //give size here
	size := len(X)
	
	//getting the inputs from excel file and storing in an array
    xl := "C:/Users/K G/Desktop/array.xlsx"
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
    
	//sort and display time taken to sort
	A := X[0:size] 
	//fmt.Println(A)
	start := time.Now()
	mergeSort(A)
	time := time.Since(start)
	fmt.Println("time %s", time)
	//fmt.Println(B)
    //fmt.Println(mergeSort(A))
}


// merge Function
func merge(m1, m2 []int) []int {
    mergeArr := make([]int, len(m1) + len(m2))
    
    j, k := 0, 0
    for i := 0; i < len(mergeArr); i++ {
        
        if j >= len(m1) {
            mergeArr[i] = m2[k]
            k++
            continue
        } else if k >= len(m2) {
            mergeArr[i] = m1[j]
            j++
            continue
        }		
        
        if m1[j] > m2[k] {
            mergeArr[i] = m2[k]
            k++
        } else {
            mergeArr[i] = m1[j]
            j++
        }
    }

    return mergeArr
}


//  mergeSort Function
func mergeSort(A []int) []int {
    if len(A) <= 1 {
        return A
    }

    l := A[0:len(A) / 2]
	r := A[len(A) / 2:]
    return merge(mergeSort(l), mergeSort(r))
}
