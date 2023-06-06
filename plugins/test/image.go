package test

import (
	"io"
	"net/http"
	"os"
	"time"

	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
	"github.com/lianhong2758/RosmBot/zero"
)

const (
	referer  = "https://weibo.com/"
	shouer   = "https://iw233.cn/api.php?sort=cat&referer"
	baisi    = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=ergbskjhebrgkjlhkerjsbkbregsbg"
	heisi    = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=rsetbgsekbjlghelkrabvfgheiv"
	siwa     = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=dsrgvkbaergfvyagvbkjavfwe"
	bizhi    = "https://iw233.cn/api.php?sort=iw233"
	baimao   = "https://iw233.cn/api.php?sort=yin"
	xing     = "https://iw233.cn/api.php?sort=xing"
	sese     = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php?sort=qwuydcuqwgbvwgqefvbwgueahvbfkbegh"
	biaoqing = "https://iw233.cn/api.php?sort=img"
	cos      = "http://aikohfiosehgairl.fgimax2.fgnwctvip.com/uyfvnuvhgbuiesbrghiuudvbfkllsgdhngvbhsdfklbghdfsjksdhnvfgkhdfkslgvhhrjkdshgnverhbgkrthbklg.php/?sort=cos"
	manghe   = "https://iw233.cn/api.php?sort=random"
)

func init() {
	en := c.Register("image", &c.PluginData{
		Name:       "图片",
		Help:       "随机壁纸" + "兽耳" + "星空" + "白毛" + "我要涩涩" + "涩涩达咩" + "白丝" + "黑丝" + "丝袜" + "随机表情包" + "cos" + "盲盒" + "开盲盒",
		DataFolder: "image",
	})
	en.AddWord(func(ctx *c.CTX) {
		var url string
		switch ctx.Being.Word {
		case "兽耳":
			url = shouer
		case "随机壁纸":
			url = bizhi
		case "白毛":
			url = baimao
		case "星空":
			url = xing
		case "我要涩涩", "涩涩达咩":
			url = sese
		case "白丝":
			url = baisi
		case "黑丝":
			url = heisi
		case "丝袜":
			url = siwa
		case "随机表情包":
			url = biaoqing
		case "cos":
			url = cos
		case "盲盒", "开盲盒":
			url = manghe
		default:
			return
		}
		url2, err := web.GetRealURL(url)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		var client = &http.Client{}
		request, err := http.NewRequest(http.MethodGet, url2, nil)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		// 增加header选项
		if referer != "" {
			request.Header.Add("Referer", referer)
		}
		response, err := client.Do(request)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}

		timestamp := time.Now().Format("20060102150405")
		// 构造文件名
		imageFileName := "data/image/" + timestamp + ".jpg"
		file2, err := os.Create(imageFileName)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		defer file2.Close()

		_, err = io.Copy(file2, response.Body)
		if err != nil {
			ctx.Send(c.Text("ERROR: ", err))
			return
		}
		ctx.Send(c.ImageWithText(zero.MYSconfig.Host+"/file/image/"+timestamp+".jpg", 0, 0, 0, "喵~"))
	}, "随机壁纸", "兽耳", "星空", "白毛", "我要涩涩", "涩涩达咩", "白丝", "黑丝", "丝袜", "随机表情包", "cos", "盲盒", "开盲盒")
}
