package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	//method 对应着SigningMethodHMAC struct{}, 其中包含 SigningMethodHS256, SigningMethodHS384, SigningMethodHS512三种cryoto.Hash(加密哈希）方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//内部生成签名字符串，在用于获取完整、已签名对token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//func (p *Parser) ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	//func (m MapClaims) Valid() 验证基于时间的声明 exp, iat, nbf, 如果没有任何声明在令牌中，仍在会被认为是有效的，并且对于时区偏差没有计算方法。
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
