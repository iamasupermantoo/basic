package user

import (
	"basic/apis/admin/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Relation 用户关系
func Relation(c *fiber.Ctx) error {
	adminId, userId := utils.GetContextClaims(c)

	userTree, _ := users.UserRelationTree(adminId, userId)
	return c.JSON(utils.SuccessJson(userTree))
}
