package server

import "strconv"

func parseSize(d []byte) int {
	// Convert the byte slice to a string
	sizeStr := string(d)

	// Parse the string to an integer
	size, _ := strconv.Atoi(sizeStr)

	return size
}

func splitArg(arg []byte) [][]byte {
	var args [][]byte
	start := -1
	for i, b := range arg {
		switch b {
		case ' ', '\t':
			if start >= 0 {
				args = append(args, arg[start:i])
				start = -1
			}
		default:
			if start < 0 {
				start = i
			}
		}
	}
	if start >= 0 {
		args = append(args, arg[start:])
	}
	return args
}
