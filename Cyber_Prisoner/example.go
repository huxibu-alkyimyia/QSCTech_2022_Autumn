package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"fmt"
)

func main() {

	var jointQSCTech bool = false
	var jointZJUer bool = false

	router := gin.Default()
	router.LoadHTMLGlob("./*")
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "pong",
		})
	})

	router.GET("/join", func(context *gin.Context) {
		context.HTML(http.StatusOK, "place.html", nil)
	})

	router.POST("/login", func(context *gin.Context) {
		user := context.PostForm("place")
		switch user {
		case "QSCTech":
			if jointQSCTech {
				context.JSON(http.StatusOK, gin.H{
					"code": 100,
					"msg":  "Place QSCTech is already occupied!",
				})
			} else {
				context.SetCookie("c", "QSCTech", 3600, "/", "localhost", false, true)
				jointQSCTech = true
				context.HTML(http.StatusOK, "lsq.html", nil)
			}
		case "ZJUer":
			if jointZJUer {
				context.JSON(http.StatusOK, gin.H{
					"code": 100,
					"msg":  "Place ZJUer is already occupied!",
				})
			} else {
				context.SetCookie("c", "ZJUer", 3600, "/", "localhost", false, true)
				jointZJUer = true
				context.HTML(http.StatusOK, "lsz.html", nil)
			}
		default:
			context.JSON(http.StatusOK, gin.H{
				"code": 101,
				"msg":  "Incorrect place!",
			})
		}
	})

	var modeSet bool = false
	var round int = 0
	var mode string
	var modeint int
	var _ error
	var result = [2]int{0, 0}

	router.GET("/playround", func(context *gin.Context) {
		modeint, _ = strconv.Atoi(mode)
		user, cerr := context.Cookie("c")
		if cerr == nil {
			switch user {
			case "QSCTech":
				if modeSet {
					if jointZJUer {
						if round < modeint {
							context.HTML(http.StatusOK, "choose.html", gin.H{
								"mode":    mode,
								"round":   round,
								"yresult": result[0],
								"oresult": result[1],
							})
							round += 1
						}
					} else {
						context.JSON(http.StatusOK, gin.H{
							"code": 102,
							"msg":  "Opponent is disconnected, please refresh the page later",
						})
					}
				} else {
					context.HTML(http.StatusOK, "setmode.html", nil)
				}
			case "ZJUer":
				if jointQSCTech {
					if modeSet {
						context.HTML(http.StatusOK, "choose.html", gin.H{
							"mode":    mode,
							"round":   round,
							"yresult": result[1],
							"oresult": result[0],
						})
					} else {
						context.JSON(http.StatusOK, gin.H{
							"code": 102,
							"msg":  "Opponent has not yet chosen gamemode, please refresh the page later",
						})
					}
				} else {
					context.JSON(http.StatusOK, gin.H{
						"code": 102,
						"msg":  "Opponent is disconnected, please refresh the page later",
					})
				}
			}
		} else {
			fmt.Println(cerr)
		}
	})

	router.POST("/modeset", func(context *gin.Context) {
		modeSet = true
		mode = context.PostForm("mode")
		context.HTML(http.StatusOK, "modeset.html", nil)
	})

	router.Run(":50000")
}
