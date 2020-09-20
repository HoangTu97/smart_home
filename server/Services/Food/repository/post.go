package repository

import (
	"Food/config"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
)

func SavePost(post models.Post) models.Post {
	config.GetDB().Create(&post)
	return post
}

func FindPagePost(pageable pagination.Pageable) page.Page {
	var posts []models.Post

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Preload("User").Preload("Recipe"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &posts)

	return page.From(toInterfacesFromPost(posts), paginator.TotalRecord)
}

func toInterfacesFromPost(posts []models.Post) []interface{} {
	content := make([]interface{}, len(posts))
	for i, v := range posts {
		content[i] = v
	}
	return content
}