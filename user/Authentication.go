package user

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func CheckAuthentication(c *fiber.Ctx) (int, error) {
	j, _ := json.Marshal(c.GetReqHeaders())
	fmt.Println(string(j))
	userIdStr := c.Get("User-Id", "")
	userKeyStr := c.Get("User-Key", "")

	if len(userIdStr) == 0 || len(userKeyStr) == 0 {
		return 0, c.SendStatus(http.StatusForbidden)
	}

	userId := 0
	userKey := 0
	idErr := json.Unmarshal([]byte(userIdStr), &userId)
	keyErr := json.Unmarshal([]byte(userKeyStr), &userKey)

	if idErr != nil || keyErr != nil {
		return 0, c.SendStatus(http.StatusForbidden)
	}

	//todo check in db
	return userId, nil
}
