package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	Info("git clone https://github.com/lindsayloughlin/node-proxy")

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/lindsayloughlin/node-proxy",
	})

	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		iterator, err := c.Files()
		CheckIfError(err)

		fileitem, err := iterator.Next()
		CheckIfError(err)
		for fileitem != nil {
			fileitem, err := iterator.Next()
			CheckIfError(err)
			if bin, _ := fileitem.IsBinary(); !bin {
				fmt.Println(fileitem.Name)
				data, err := fileitem.Contents()
				CheckIfError(err)
				// Do some parsing to see what is being modified.
				fmt.Print(data)
			}
		}

		return nil
	})
	CheckIfError(err)

}
