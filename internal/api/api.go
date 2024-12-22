package api

import (
	"example/internal/business_logic"

	"github.com/gin-gonic/gin"
)

// у апи нет интерфейса - он не нужен, от него ничего не зависит
// в некоторых случаях его делают, чтобы мокать в тестах
type API struct {
	logic business_logic.Logic

	r *gin.Engine
}

func (api *API) get(c *gin.Context) {
	// тут происходит анмаршлинг реквест боди или квери параметров
	// я не буду это все приводить для краткости

	api.logic.Get("id")
}

func (api *API) set(c *gin.Context) {
	api.logic.Set("id", "name")
}

func Run(
	logic business_logic.Logic,
) {
	// все что тут происходит не то чтобы по красоте, но в целом норм
	r := gin.New()

	api := API{
		r:     r,
		logic: logic,
	}

	r.GET("/get", api.get)
	r.GET("/set", api.set)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
