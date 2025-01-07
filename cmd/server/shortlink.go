package main

type LinksStore interface {
	Get(suffix string) (string, bool)
	Set(link string) string
}

type ShortLink struct {
	store LinksStore
}

func NewShortLink(store LinksStore) *ShortLink {
	return &ShortLink{store: store}
}

func (sl *ShortLink) GetLink(suffix string) (string, bool) {
	return sl.store.Get(suffix)
}

func (sl *ShortLink) Set(link string) string {
	return sl.store.Set(link)
}
