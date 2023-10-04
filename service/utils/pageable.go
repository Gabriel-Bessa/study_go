package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"study_go/models/entity"
)

func ExtractPageable(c *gin.Context) entity.Pageable {
	defaultSize := "10"
	defaultPage := "0"
	pageParam := c.DefaultQuery("page", defaultPage)
	sizeParam := c.DefaultQuery("size", defaultSize)
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		panic("fsdfs")
	}
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		panic("defaultPage")
	}
	return entity.Pageable{Page: page, Size: size, OffSet: page * size}
}
