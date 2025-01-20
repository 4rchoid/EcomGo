package dummyRouter

import (
	"github.com/gin-gonic/gin"

	"myproject/dummy/handlers"
)

func RouterInit(r *gin.Engine){
	dummyGroup := r.Group("/user")
	{
		dummyGroup.GET("", dummyHandler.SayHello)
	
	
	}


}