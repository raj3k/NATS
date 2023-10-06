package server

import "strconv"

func parseSize(d []byte) int {
	// Convert the byte slice to a string
	sizeStr := string(d)

	// Parse the string to an integer
	size, _ := strconv.Atoi(sizeStr)

	return size
}
