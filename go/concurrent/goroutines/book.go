package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hobbit",
		Author:        "J.R.R",
		YearPublished: 1937,
	},
	{
		ID:            2,
		Title:         "The Hobbit2",
		Author:        "J.R.R2",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "The Hobbit3",
		Author:        "J.R.R3",
		YearPublished: 1937,
	},
	{
		ID:            4,
		Title:         "The Hobbit4",
		Author:        "J.R.R4",
		YearPublished: 1937,
	},
}
