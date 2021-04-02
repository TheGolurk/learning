package media

import "strings"

type Catalogable interface {
	NewMovie(t string, r Rating, b float32)
	GetTitle() string
	GetRating() string
	GetBox() float32
	SetTitle(t string)
	SetRating(r string)
	SetBox(b float32)
}

type Movie struct {
	title  string
	rating Rating
	box    float32
}

type Rating string

const (
	R    = "R"
	G    = "G"
	PG   = "PG"
	PG13 = "PG13"
	NC17 = "NC17"
)

func (m *Movie) NewMovie(t string, r Rating, b float32) {
	m.title = t
	m.rating = r
	m.box = b
}

func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetRating() Rating {
	return m.rating
}

func (m *Movie) GetBox() float32 {
	return m.box
}

func (m *Movie) SetTitle(t string) {
	m.title = t
}

func (m *Movie) SetRating(r Rating) {
	m.rating = r
}

func (m *Movie) SetBox(b float32) {
	m.box = b
}
