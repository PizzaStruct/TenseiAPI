package handlers

type IGenreHandler interface {
}

type GenreHandler struct {
}

func NewGenreHandler() IGenreHandler {
	return new(GenreHandler)
}
