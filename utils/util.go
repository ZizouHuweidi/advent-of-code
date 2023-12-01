package utils

import "os"

func GetInput(f string) string {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(data)
}
