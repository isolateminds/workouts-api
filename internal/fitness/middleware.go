package fitness

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SearchWorkouts(ctx *gin.Context) {
	search := ctx.Query("search")
	fmt.Println(search)
	exercises := QueryWorkouts(search)
	ctx.JSON(200, exercises)
}
