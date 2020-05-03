package main

import (
	"fmt"
	"net/http"

	"github.com/dema501/randomjoke/internal/pkg/request"

	"github.com/dema501/randomjoke/internal/pkg/jokemaker"
	"github.com/dema501/randomjoke/internal/pkg/namegiver"
	"github.com/dema501/randomjoke/internal/pkg/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	rs := rest.Server{
		ServerPort: "5000",
	}

	rs.Run(
		func(rg *gin.RouterGroup) {
			rg.GET("/joke", GetJokeHandler())
		},
	)
}

func MakeJoke(nm namegiver.Doer, jm jokemaker.Doer) (string, error) {
	if err := nm.Generate(); err != nil {
		return "", err
	}

	firstName, lastName := nm.GetName()
	jm.SetName(firstName, lastName)

	if err := jm.Generate(); err != nil {
		return "", err
	}

	return jm.GetJoke(), nil
}

func GetJokeHandler() func(c *gin.Context) {
	sa := request.New()

	return func(c *gin.Context) {
		j, err := MakeJoke(namegiver.New(sa), jokemaker.New(sa))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":  http.StatusInternalServerError,
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": j,
		})
	}
}
