package main

import "fmt"

func nonempty(string []string) []string {
	i := 0
	for _, s := range string {
		if s != "" {
			string[i] = s
			i++
		}
	}
	return string[:i]
}
