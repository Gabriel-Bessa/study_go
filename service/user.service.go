package service

import (
	"github.com/gin-gonic/gin"
	"study_go/models/dto"
	"study_go/models/entity"
)
import "study_go/service/utils"

func List(c *gin.Context) {
	pageable := utils.ExtractPageable(c)
	var users []entity.User
	entity.DB.Offset(pageable.OffSet).Limit(pageable.Size).Find(&users)
	var usersDto []dto.UserDTO

	parseToDto := func(n entity.User) dto.UserDTO {
		return dto.UserDTO{
			Id:   n.Id,
			Name: n.Name,
		}
	}

	for _, n := range users {
		usersDto = append(usersDto, parseToDto(n))
	}

	c.JSON(200, entity.Page{
		Content:  usersDto,
		Pageable: pageable,
	})
}
