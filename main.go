package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fnf"
	app.Usage = "Fix filenames"
	app.Authors = []cli.Author{
		cli.Author{Name: "Jimmy Roland", Email: "jimmylroland@gmail.com"},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "R", Usage: "Recursive"},
		cli.BoolFlag{Name: "S", Usage: "Remove Spaces"},
	}

	app.Action = func(c *cli.Context) error {
		var files []string

		root := c.Args().Get(0)
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			fmt.Println(info.Name())
			if string([]rune(info.Name())[0]) == "." || info.IsDir() {
				fmt.Printf("Hidden %q %q\n", path, info.Name())
				return filepath.SkipDir
			}

			files = append(files, path)
			return nil
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Good files!!!")
		for _, file := range files {
			fmt.Println(file)
		}

		fmt.Printf("%#v\n", c.Bool("S"))
		fmt.Println("boom! I say!")
		fmt.Printf("Hello %q", c.Args().Get(0))

		os.Rename("ex.txt", "new.txt")

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
