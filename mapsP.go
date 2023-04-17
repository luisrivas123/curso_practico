package main

// import "fmt"

// func main() {
// 	v := []interface{}{
// 		map[string]string{"name": "ravi"},
// 		[]string{"art", "coding", "music", "travel"},
// 		map[string]string{"language": "golang"},
// 		map[string]string{"experience": "no"},
// 	}
// 	fmt.Println(v)
// }

import "fmt"

type candidate struct {
	name     string
	language string
}

func main() {
	m := make(map[string]string)
	m["role"] = `"role": "user"`
	m["content"] = `"content": "Say this is a test!"`
	candidates := []candidate{
		{
			name:     m["role"],
			language: m["content"],
		},
	}
	fmt.Println(candidates)

	v := []interface{}{
		map[string]string{"role": "user"},
		map[string]string{"content": "Say this is a test!"},
	}
	fmt.Println(v)
}
