package package_gin

import ("github.com/gin-gonic/gin"
	"net/http"
)









func StartRouter () {

router := gin.Default()



router.GET("/ping", func (c *gin.Context) {


//insertMongo();

c.JSON(200, gin.H{
"message": "pong",
})
})




router.GET("/user/:name", func (c *gin.Context) {
name := c.Param("name")
c.String(http.StatusOK, "Hello %s", name)
})



router.POST("/form_post", func (c *gin.Context) {
message := c.PostForm("message")
nick := c.DefaultPostForm("nick", "anonymous")

c.JSON(200, gin.H{
"status":  "posted",
"message": message,
"nick":    nick,
})
})

//r.Run() // listen and server on 0.0.0.0:8080


router.Run(":8085")

//arg := os.Args


//r.Run(arg[1])


}