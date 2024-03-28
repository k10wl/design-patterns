package main

import "fmt"

func main() {
	fmt.Printf(">>> Composite demo\n\n")

	fs := Folder{
		Name: "ROOT",
		Content: []Component{
			Folder{
				Name: "bin",
				Content: []Component{
					File{Name: "node"},
					File{Name: "rust"},
					File{Name: "go"},
				},
			},
			File{Name: ".zshrc"},
			Folder{
				Name: "Downloads",
				Content: []Component{
					Folder{
						Name: "Movies",
						Content: []Component{
							File{Name: "work.mov"},
							File{Name: "home.mov"},
							File{Name: "car.mov"},
						},
					},
					File{Name: "README.md"},
				},
			},
		},
	}

	lookup := "work.mov"
	fmt.Printf("Path for `%s`: ", lookup)
	fmt.Print(fmt.Sprintln(fs.Search(lookup)))
	fmt.Printf("\n\nFull path:\n")
	fmt.Println(fs.GetPath(0))
}
