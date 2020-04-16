package webhandlers

import "github.com/gin-gonic/gin"

var cardJobs = make(chan struct{}, 2)

func LimitingMiddleware(c *gin.Context) {
	select {
	case cardJobs <- struct{}{}:
		{
			defer func() { <-cardJobs }()
			c.Next()
		}
	default:
		{
			c.AbortWithStatus(503)
		}
	}
}
