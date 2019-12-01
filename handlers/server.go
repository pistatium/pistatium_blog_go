package handlers

import "github.com/pistatium/pistatium_blog_go/repos"

type Server struct {
	Entries repos.EntryRepo
	Photos  repos.PhotoRepo
	Admin   repos.AdminUserRepo
	Conf    *repos.Conf
}
