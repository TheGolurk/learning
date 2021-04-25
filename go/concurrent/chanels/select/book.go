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
	{
		ID:            5,
		Title:         "The Hobbit5",
		Author:        "J.R.R5",
		YearPublished: 1937,
	},
	{
		ID:            6,
		Title:         "The Hobbit6",
		Author:        "J.R.R6",
		YearPublished: 1937,
	},
	{
		ID:            7,
		Title:         "The Hobbit7",
		Author:        "J.R.R7",
		YearPublished: 1937,
	},
	{
		ID:            8,
		Title:         "The Hobbit8",
		Author:        "J.R.R8",
		YearPublished: 1937,
	},
	{
		ID:            9,
		Title:         "The Hobbit9",
		Author:        "J.R.R9",
		YearPublished: 1937,
	},
	{
		ID:            10,
		Title:         "The Hobbit10",
		Author:        "J.R.R10",
		YearPublished: 1937,
	},
}
