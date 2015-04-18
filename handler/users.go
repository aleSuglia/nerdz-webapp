package handler

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"log"
	"net/http"
	"strconv"
)

func GetUserBoard(c *echo.Context) {
	userId, errParse := strconv.ParseUint(c.Param("id"), 10, 64)

	if errParse != nil {
		c.JSON(http.StatusBadRequest, &NerdzResponse{
			Message:      errParse.Error(),
			HumanMessage: "Invalid user identifier specificied in the request",
		})

		return
	}

	user, err := nerdz.NewUser(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, &NerdzResponse{
			Message:      err.Error(),
			HumanMessage: "Invalid user identifier specificied in the request",
		})

		return
	}

	var postOpt *nerdz.PostlistOptions
	var n string

	if n = c.Request.FormValue("n"); n == "" {
		postOpt = &nerdz.PostlistOptions{
			User:      true,
			Following: true,
		}

		log.Panic(err)
	} else {

		n, errPN := strconv.ParseUint(n, 10, 8)

		if errPN != nil {
			postOpt = &nerdz.PostlistOptions{
				User:      true,
				Following: true,
			}

			log.Panic(err)
		} else {
			postOpt = &nerdz.PostlistOptions{
				User:      true,
				Following: true,
				N:         uint8(n),
			}

			log.Print("n=", n)
		}
	}

	posts := user.UserHome(postOpt)

	if posts == nil {
		c.JSON(http.StatusBadRequest, &NerdzResponse{
			Message:      "error user board",
			HumanMessage: "Unable to get board information for the specified user",
		})

		return

	}

	c.JSON(http.StatusOK, &NerdzResponse{
		Message:      "Request ok",
		HumanMessage: "Correctly retrieved board data for user",
		Data:         *posts,
	})

}

func GetUser(c *echo.Context) {
	userId, errParse := strconv.ParseUint(c.Param("id"), 10, 64)

	if errParse != nil {
		c.JSON(http.StatusBadRequest, &NerdzResponse{
			Message:      errParse.Error(),
			HumanMessage: "Invalid user identifier specificied in the request",
		})

		return
	}

	user, err := nerdz.NewUser(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, &NerdzResponse{
			Message:      err.Error(),
			HumanMessage: "Invalid user identifier specificied in the request",
		})

		return
	}

	c.JSON(http.StatusOK, &NerdzResponse{
		Message:      "Request ok",
		HumanMessage: "Correctly retrieved data for user",
		Data:         user,
	})

}
