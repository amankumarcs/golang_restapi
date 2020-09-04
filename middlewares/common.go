package middleware

import(
  "github.com/gin-gonic/gin"
  "fmt"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
  
	// Pass on to the next-in-chain
	c.Next();
}