package test

import (
	c "github.com/lianhong2758/RosmBot/ctx"
	"github.com/lianhong2758/RosmBot/web"
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
		Name: "图片",
		Help: "- /随机壁纸" + " | " + "兽耳" + " | " + "星空" + " | " + "白毛" + " | " + "我要涩涩" + " | " + "涩涩达咩" + " | " + "白丝" + " | " + "黑丝" + " | " + "丝袜" + " | " + "随机表情包" + " | " + "cos" + " | " + "盲盒" + " | " + "开盲盒",
	})
	en.AddWord(func(ctx *c.CTX) {
		var url string
		switch ctx.Being.Word[1:] {
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
		data, err := web.RequestDataWith(web.NewDefaultClient(), url2, "", referer, "", nil)
		if err != nil {
			ctx.Send(c.Text("获取图片失败惹"))
			return
		}
		ctx.Send(c.ImageWithText(data, 1080, 2200, 0, "喵~"))
	}, "/随机壁纸", "/兽耳", "/星空", "/白毛", "/我要涩涩", "/涩涩达咩", "/白丝", "/黑丝", "/丝袜", "/随机表情包", "/cos", "/盲盒", "/开盲盒")
}
