package main


func FilterEvenNumbers(lst []int) []int {
    

    resultLst := make([]int, 0, 10)

    for _, val := range lst {
        if val % 2 == 0 {
            resultLst = append(resultLst, val)
        }
    }
    return resultLst

}