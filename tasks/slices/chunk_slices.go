package main


func ChunkSlice(lst []int, chunk_size int) [][]int {

    lstLen := len(lst)
    resultSlice := make([][]int, 0)

    for i := 0; i < lstLen; i += chunk_size {
        end := min(i + chunk_size, lstLen)
        resultSlice = append(resultSlice, lst[i:end])
    }

    return resultSlice

} 