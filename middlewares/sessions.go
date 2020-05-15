package middlewares

import (
	"awesomeProject/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedisSessionStore() redis.Store {
	address := config.LoadConfig().Get("redis_setting.host")
	token := config.LoadConfig().Get("redis_setting.token")
	secret := config.LoadConfig().Get("redis_setting.secret")

	store, err := redis.NewStore(2, "tcp", address.(string), token.(string), []byte(secret.(string)))
	if err != nil {
		panic(err)
	}

	return store

}

func ValidataAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/sys/login" {
			c.Next()
			return
		}
		//这一部分可以替换成从session/cookie中获取，
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return // return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "身份验证成功"})
			c.Next() //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		}
	}
}