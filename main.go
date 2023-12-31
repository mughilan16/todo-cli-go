package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Option string

type cliArugment struct {
	option   Option
	arugment string
}

type todo struct {
	id      int
	text    string
	checked bool
}

const (
	LIST      Option = "list"
	ADD       Option = "add"
	REMOVE    Option = "remove"
	DONE      Option = "done"
	FILE_NAME string = "data.csv"
)

func main() {
  args := getArgument()
  option := args.option;
  value := args.arugment;
	file := openFile(FILE_NAME)
	data := readData(file)
	todos := parseData(data)
	writer := csv.NewWriter(file)
	switch option {
	case LIST:
		for _, t := range todos {
			t.display()
		}
		break
	case ADD:
		todos = append(todos, todo{text: value, checked: false})
		data = parseCSV(todos)
		writer.WriteAll(data)
		writer.Flush()
		break
	case REMOVE:
		break
	case DONE:
		break
	default:
		fmt.Println("Invalid Option. use -h to see the options")
		os.Exit(1)
	}
}

func parseCSV(todos []todo) [][]string {
	data := make([][]string, len(todos))
	for i, t := range todos {
		data[i][0] = t.text
		if t.checked {
			data[i][1] = "1"
		} else {
			data[i][1] = "0"
		}
	}
	return data
}

func (t todo) display() {
	fmt.Printf("%d.", t.id)
	fmt.Printf(" %s ", t.text)
	if t.checked {
		fmt.Printf("✅\n")
	} else {
		fmt.Printf("❌\n")
	}
}

func parseData(data [][]string) []todo {
	todos := make([]todo, len(data))
	for i, p := range data {
		todos[i].id = i + 1
		todos[i].text = p[0]
		if p[1] == "1" {
			todos[i].checked = true
		} else {
			todos[i].checked = false
		}
	}
	return todos
}

func readData(file *os.File) [][]string {
	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		exit("Can't Read the data")
	}
	return data
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		exit("Can't Open the data file")
	}
	return file
}

func getArgument() cliArugment {
	list := flag.Bool("list", true, "List the Todo List")
	add := flag.String("add", "", "Add Todo to the list")
	remove := flag.String("remove", "", "Remove Todo from the list")
	done := flag.String("done", "", "Mark Todo as Done")
	flag.Parse()
  var args cliArugment;
	if *list {
    args = cliArugment{
      option: LIST,
      arugment: "",
    }
	} else if *add != "" {
    args = cliArugment{
      option: ADD,
      arugment: *add,
    }
	} else if *remove != "" {
    args = cliArugment{
      option: REMOVE,
      arugment: *remove,
    }
	} else if *done != "" {
    args = cliArugment{
      option: DONE,
      arugment: *done,
    }
	}
	return args
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
