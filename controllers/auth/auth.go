package auth

import (
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ApiLogin(r *gin.Context) {
	userInfo := UserInfo{}
	if err := r.ShouldBindJSON(&userInfo); err != nil { //把获取的url的json请求体绑定到userInfo
		logs.Warn(nil, "json未绑定成功")
		r.JSON(400, gin.H{"message": err.Error()})
	} else {
		// r.JSON(200, gin.H{"message": "登录成功"})
		logs.Info(map[string]interface{}{"用户名": userInfo.Username, "密码": userInfo.Password}, "登录成功")
	}
	//验证用户的账号和密码
	if userInfo.Username == "llkk" && userInfo.Password == "aabbcc" {
		if ss, err := jwtutil.GenToken(userInfo.Username); err != nil {
			logs.Error(nil, "生成token失败")
			r.JSON(200, gin.H{"message": "生成token失败",
				"status": 400})
		} else {
			logs.Info(map[string]interface{}{"用户名": userInfo.Username, "token": ss}, "用户登录成功")
			data := make(map[string]interface{})
			data["token"] = ss
			r.JSON(200, gin.H{"message": "用户登录成功",
				"status": 200,
				"data":   data,
			})
		}
	} else {
		r.JSON(400, gin.H{"message": "用户名或密码错误",
			"status": 400})
	}
}

func ApiLogout(g *gin.Context) {
	g.JSON(200, gin.H{"message": "登出成功",
		"status": 200})
	logs.Info(nil, "用户登出")
}
