package engine

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"seckill/api"
	myjwt "seckill/middleware/jwt"
	"seckill/mysql"
	"strconv"
	"time"
)

var a, b int

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
func JoinUs(ctx *gin.Context) {
	_ = Verify(-1)
	ctx.HTML(http.StatusOK, "join.html", gin.H{
		"A": a,
		"B": b,
	})
}
func Verify(answer int) bool {
	if answer == a+b {
		a = rand.Intn(100)
		b = rand.Intn(100)
		return true
	} else {
		a = rand.Intn(100)
		b = rand.Intn(100)
		return false
	}

}
func JoinUsOk(ctx *gin.Context) {

	answer, _ := strconv.Atoi(ctx.PostForm("Verify_result"))

	if Verify(answer) {

		/*join*/

		//join := initt.JoinInit(ctx.PostForm("Name"), ctx.PostForm("Id"), ctx.PostForm("Email"), ctx.PostForm("Tel"), ctx.PostForm("pass1"), ctx.PostForm("pass2"))
		var join api.Join
		join.Pass1 = ctx.PostForm("Pass1")
		join.Pass2 = ctx.PostForm("Pass2")
		fmt.Println(join.Pass1, join.Pass2)
		if join.Pass1 != join.Pass2 {

			ctx.HTML(http.StatusOK, "wrong.html", gin.H{

				"Why": "两次密码不一致",
				"A":   a, //rand.Intn(100),
				"B":   b, //rand.Intn(100),
			})

		} else {
			//写入mysql
			join.Name = ctx.PostForm("Name")
			join.Id = ctx.PostForm("Id")
			join.Email = ctx.PostForm("Email")
			join.Tel = ctx.PostForm("Tel")
			mysql.Insert(join.Name, join.Id, join.Email, join.Tel, join.Pass1)

			ctx.HTML(http.StatusOK, "login.html", gin.H{})
		}
	} else {
		fmt.Println(a+b, answer)

		ctx.HTML(http.StatusOK, "wrong.html", gin.H{
			"Why": "验证码错误",
			"A":   a,
			"B":   b,
		})

	}

}
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
func LoginOK(ctx *gin.Context) {

	var login api.Login
	login.Tel = ctx.PostForm("tel")
	fmt.Println(login.Tel)
	login.Passwd = ctx.PostForm("pwd")
	fmt.Println(login.Tel)
	login1, err := mysql.Query(login.Tel)
	fmt.Println(login, login1)
	if err != nil {
		//错误处理
		return
	}
	if login1.Tel != login.Tel || login1.Tel == "" {
		//去注册
		ctx.HTML(http.StatusOK, "join.html", gin.H{
			"status": "",
		})
	} else if login1.Passwd != login.Passwd {
		//密码错误
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"status": "用户名或密码错误",
		})
	} else if login.Passwd == login1.Passwd {
		//登录成功

		//进入秒杀页面

		//返回token
		var user api.Login
		generateToken(ctx, user)

	}

}
func generateToken(ctx *gin.Context, user api.Login) {
	j := myjwt.NewJWT()
	claims := myjwt.CustomClaims{
		Username: user.Tel,
		Password: user.Passwd,

		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    myjwt.Issuer,                    //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			//kindKey:   user.Kind,
			"ErrMsgKey": err,
		})
		return
	}

	//log.Println(token)
	ctx.Header("Authorization", token)
	//ctx.JSON(http.StatusOK, gin.H{
	//	//kindKey:   user.Kind,
	//	"ErrMsgKey": "",
	//})
	return

}
