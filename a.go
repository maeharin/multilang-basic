package main

import "fmt"

func main() {
	users := []struct {
		id   int
		name string
	}{
		{0, "hoge"},
		{1, "fuge"},
		{2, "piyo"},
	}

	ids := make([]int, len(users))
	for idx, user := range users {
		ids[idx] = user.id
	}
	fmt.Println(ids)

	names := make([]string, len(users))
	for idx, user := range users {
		names[idx] = user.name
	}
	fmt.Println(names)
}
