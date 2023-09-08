package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/hello-world/apps/user"
	impl2 "github.com/luyasr/hello-world/apps/user/impl"
	"log"
	"strconv"
)

var impl = impl2.NewImpl()

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

type Handler struct {
	svc user.Service
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (h *Handler) HelloUser(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	u, err := impl.GetUser(c.Request.Context(), idInt)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, Response{Code: 10000, Data: u, Msg: "success"})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u user.CreateUserRequest
	err := c.ShouldBindJSON(&u)
	if err != nil {
		log.Println(err)
	}

	createUser, err := impl.CreateUser(c.Request.Context(), &u)
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, Response{Code: 10000, Data: createUser, Msg: "success"})
}
