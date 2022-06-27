package user

import (
	"encoding/json"
	"fmt"
	"github.com/David10238/socialmedia-api/db"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"net/http"
	"strconv"
)

func CheckAuthentication(c *fiber.Ctx) (uint, error) {
	j, _ := json.Marshal(c.GetReqHeaders())
	fmt.Println(string(j))
	userIdStr := c.Get("User-Id", "")
	userKey := c.Get("User-Key", "")

	if len(userIdStr) == 0 || len(userKey) == 0 {
		return 0, c.SendStatus(http.StatusForbidden)
	}

	userId := uint(0)
	if err := json.Unmarshal([]byte(userIdStr), &userId); err != nil {
		return 0, c.SendStatus(http.StatusForbidden)
	}

	user := db.User{}
	notFound := db.Conn.First(&user, userId).RowsAffected == 0

	if notFound || user.Key != userKey {
		return 0, c.SendStatus(http.StatusForbidden)
	}

	return userId, nil
}

func newKey() string {
	return strconv.Itoa(rand.Int())
}
