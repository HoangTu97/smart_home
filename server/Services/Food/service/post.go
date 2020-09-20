package service

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
)

func FindPagePost(pageable pagination.Pageable) page.Page {
	return repository.FindPagePost(pageable)
}