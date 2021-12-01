package handlers

type IMangaHandler interface {
}

type MangaHandler struct {
}

func NewMangaHandler() IMangaHandler {
	return new(MangaHandler)
}
