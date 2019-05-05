package limit

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//Limit function is gin middleware to limit current requests
func Limit(max int) gin.HandlerFunc {
	if max <= 0 {
		log.Panic("max must be more than 0")
	}
	sema := make(chan struct{}, max)
	return func(c *gin.Context) {
		select {
		case sema <- struct{}{}:
			c.Next()
			<-sema
		}
	}
}
