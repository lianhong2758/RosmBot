package ctx

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"net/url"
	"strings"
	"strconv"

	"github.com/lianhong2758/RosmBot/zero"
)

// 是否是主人权限
func IsMaster(userID string) bool {
	for _, v := range zero.MYSconfig.BotToken.Master {
		if v == userID {
			return true
		}
	}
	return false
}

func (ctx *CTX) IsMaster() bool {
	for _, v := range zero.MYSconfig.BotToken.Master {
		if v == ctx.Being.User.ID {
			return true
		}
	}
	return false
}

func (ctx *CTX) IntUserID() int { x, _ := strconv.Atoi(ctx.Being.User.ID); return x }

// 签名验证
func verify(sign, body, botSecret, pubKey string) bool {
	signArg, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	str := url.Values{
		"body":   {strings.TrimSpace(body)},
		"secret": {botSecret},
	}.Encode()

	hashedOrigin := sha256.Sum256([]byte(str))

	// 将 PEM 格式的公钥解码为 DER 格式的数据
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		log.Println("failed to decode PEM block containing public key")
		return false
	}
	derBytes := block.Bytes

	// 解析 DER 格式的公钥，得到 *rsa.PublicKey 类型的变量
	publicKey, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		log.Printf("failed to parse DER encoded public key: %v", err)
		return false
	}
	//log.Println(publicKey, crypto.SHA256, hashedOrigin[:], signArg)
	if err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashedOrigin[:], signArg); err != nil {
		log.Println("[info-err] (接收消息): 签名错误")
		return false
	}
	return true
}
