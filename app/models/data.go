package models

type Data struct {
	Posts []Posts
	Comment []Comment
	Category []Category
	CategoryReport []Category
	IsLoggedIn bool
	Role string
	StatusReq string
	ModRequests []ModReq
	ReportedPosts []Posts
}