package main

func RemoveEmptyElems(lst []string) []string {

	write := 0

	for _, v := range lst {
		if v != "" {
			lst[write] = v
			write++
		}
	}

	return lst[:write]
}
