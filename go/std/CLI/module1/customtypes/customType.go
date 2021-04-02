package main

import (
	"fmt"

	"github.com/TheGolurk/pluralsight/std/CLI/module1/customtypes/media"
)

func main() {
	fmt.Println("My favorite movie")
	myfav media.Catalogable = &media.Movie{}
	myfav.NewMovie("s", media.R, 12.1)
	fmt.Println(myfav.GetRating(), myfav.GetBox(), myfav.GetTitle())
	myfav.SetTitle("b")
	myfav.SetRating(media.PG)
	myfav.SetBox(9.1)
	fmt.Println(myfav.GetRating(), myfav.GetBox(), myfav.GetTitle())
}
