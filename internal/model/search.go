package model

type SearchMapResult struct {
}

type SearchVectorResult struct {
	BookName string
	Sums int32
}

type SearchAnsResult struct {
	BookName string
	Title string
	Content string
}

type SearchHistory struct {
	Id uint
	History string
}