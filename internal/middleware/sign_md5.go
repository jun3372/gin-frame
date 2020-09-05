package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"gin-frame/pkg/errno"
	"gin-frame/pkg/response"
	"gin-frame/pkg/sign"
)

// SignMd5Middleware md5 签名校验中间件
func SignMd5Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sn, err := verifySign(c)
		if err != nil {
			response.Send(c, errno.InternalServerError, nil)
			c.Abort()
			return
		}

		if sn != nil {
			response.Send(c, errno.ErrSignParam, sn)
			c.Abort()
			return
		}

		c.Next()
	}
}

// verifySign 验证签名
func verifySign(c *gin.Context) (map[string]string, error) {
	requestUri := c.Request.RequestURI
	// 创建Verify校验器
	verifier := sign.NewVerifier()
	sn := verifier.GetSign()

	// 假定从RequestUri中读取校验参数
	if err := verifier.ParseQuery(requestUri); nil != err {
		return nil, err
	}

	// 检查时间戳是否超时。
	if err := verifier.CheckTimeStamp(); nil != err {
		return nil, err
	}

	// 验证签名
	localSign := genSign()
	if sn == "" || sn != localSign {
		return nil, errors.New("sign error")
	}

	return nil, nil
}

// genSign 生成签名
func genSign() string {
	// todo: 读取配置
	signer := sign.NewSignerMd5()
	signer.SetAppID("123456")
	signer.SetTimeStamp(time.Now().Unix())
	signer.SetNonceStr("supertempstr")
	signer.SetAppSecretWrapBody("20200711")

	return signer.GetSignature()
}
