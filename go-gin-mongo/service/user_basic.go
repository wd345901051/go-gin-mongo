package service

import (
	"fmt"
	helper "im/helpper"
	"im/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserQueryResult struct {
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}

func Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	fmt.Println(account, password)
	ub, err := models.GetUserBasicByAccountPassword(account, helper.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.GenerateToken(ub.Identity, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(c *gin.Context) {
	u, _ := c.Get("user_claims")
	uc := u.(*helper.UserClaims)
	ub, err := models.GetUserBasicByIdentity(uc.Identity)
	if err != nil {
		log.Print("DB ERROR:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": ub,
	})
}

func UserQuery(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	userBasic, err := models.GetUserBasicByAccount(account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	u, _ := c.Get("user_claims")
	uc := u.(*helper.UserClaims)

	data := UserQueryResult{
		Nickname: userBasic.Nickname,
		Sex:      userBasic.Sex,
		Email:    userBasic.Email,
		Avatar:   userBasic.Account,
		IsFriend: false,
	}
	if models.JudgeUserIsFriend(userBasic.Identity, uc.Identity) {
		data.IsFriend = true
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据查询成功",
		"data": data,
	})
}
