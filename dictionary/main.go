package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/dictionary/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./bagster")
	handleErr(err)
	defer d.Close()

	flag.Parse()

	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n:", *action)
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Print("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' was removed form the dictionary\n", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Sprintf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
