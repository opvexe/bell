package middleware

import (
	"github.com/gin-gonic/gin"
)

type Recover struct {
}

func NewRecover() *Recover {
	return &Recover{}
}

func (r *Recover) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
