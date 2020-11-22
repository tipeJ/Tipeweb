package main

type Page struct {
	Title string
}

type Project struct {
	Title           string
	DescriptionName string

	GithubLink     *string
	GooglePlayLink *string
}
