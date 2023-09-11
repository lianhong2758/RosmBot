package ctx

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"net/url"
	"strconv"
	"strings"

	"github.com/lianhong2758/RosmBot/zero"
	log "github.com/sirupsen/logrus"
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

func OnlyMaster(ctx *CTX) bool {
	for _, v := range zero.MYSconfig.BotToken.Master {
		if v == ctx.Being.User.ID {
			return true
		}
	}
	return false
}

func (ctx *CTX) IntUserID() int { x, _ := strconv.Atoi(ctx.Being.User.ID); return x }
func (m *Matcher) RulePass(ctx *CTX) bool {
	for _, v := range m.Rules {
		if !v(ctx) {
			return false
		}
	}
	return true
}

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
		log.Errorln("failed to decode PEM block containing public key")
		return false
	}
	derBytes := block.Bytes

	// 解析 DER 格式的公钥，得到 *rsa.PublicKey 类型的变量
	publicKey, err := x509.ParsePKIXPublicKey(derBytes)
	if err != nil {
		log.Errorf("failed to parse DER encoded public key: %v", err)
		return false
	}
	//log.Println(publicKey, crypto.SHA256, hashedOrigin[:], signArg)
	if err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashedOrigin[:], signArg); err != nil {
		log.Errorln("[err] (接收消息): 签名错误")
		return false
	}
	log.Debugln("回调签名通过,sign:", sign)
	return true
}
