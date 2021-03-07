package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/utyosu/rfe/db"
	"math/rand"
)

func isBattlePowerExecute(m *discordgo.MessageCreate) bool {
	return isContainKeywords(m.Content, battlePowerKeyword)
}

func actionBattlePower(m *discordgo.MessageCreate) {
	activities, err := db.FetchTodayActivities(m.Author.ID, db.ActivityKindBattlePower)
	if err != nil {
		sendMessage(m.ChannelID, messageError)
		return
	} else if len(activities) > 0 {
		sendMessage(m.ChannelID, "戦闘力は1日に1回しか測定できません。")
		return
	}

	battlePower := battlePowerList[rand.Intn(len(battlePowerList))]
	sendMessage(
		m.ChannelID,
		fmt.Sprintf(
			"%vさんの戦闘力は…%v！\n【**%v**】と同じくらいだ…！",
			getName(m),
			battlePower.power,
			battlePower.name,
		),
	)
	sendMessage(m.ChannelID, battlePower.url)

	if _, err := db.InsertActivity(m.Author.ID, db.ActivityKindBattlePower); err != nil {
		slackWarning.Post(errors.WithStack(err))
	}
}

type battlePower struct {
	power string
	name  string
	url   string
}

var (
	battlePowerKeyword = []string{
		"戦闘力",
		"強さ",
	}
	battlePowerList = []battlePower{
		{"5", "農夫", "https://img.atwikiimg.com/www30.atwiki.jp/niconicomugen/attach/9291/21657/ossan.PNG"},
		{"416", "ラディッツ編の悟空", "https://i.ytimg.com/vi/u5jo_S6eBMk/maxresdefault.jpg"},
		{"924", "ラディッツ編でかめはめ波を撃つときの悟空", "https://i0.wp.com/jumpmatome2ch.biz/wp-content/uploads/2018/01/409f4550.jpg?fit=481%2C449&ssl=1"},
		{"1330", "ピッコロ", "https://images-na.ssl-images-amazon.com/images/I/51eL0coTGyL._AC_SY450_.jpg"},
		{"1307", "怒りパワー全開の悟飯", "https://lh3.googleusercontent.com/Hom0HXuLFnjSsoWgqhAybMjdP1fSFv6GvWsE05FZlSBSKO1Hy9ESRxULwqTnZduUHjdF8ftPLvTesI37dlMAz62uUxD7G5cQcHadNTDiWo4=w350"},
		{"139", "亀仙人", "https://img.atwikiimg.com/www30.atwiki.jp/niconicomugen/attach/7274/15507/muten.png"},
		{"206", "クリリン", "https://pbs.twimg.com/profile_images/953545034531401729/ZQYq6Dzr_400x400.jpg"},
		{"250", "天津飯", "https://bookvilogger.com/wp-content/uploads/2018/07/tenshin.jpg"},
		{"177", "ヤムチャ", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/026/889/lqip.jpg?1527380473"},
		{"1200", "栽培マン", "https://chie-pctr.c.yimg.jp/dk/iwiz-chie/que-12156205752?w=320&h=320&up=0"},
		{"2800", "ベジータ編の悟飯", "https://stat.ameba.jp/user_images/20110527/21/be-daack/65/94/j/t02200132_0400024011254369201.jpg?caw=800"},
		{"1000", "ナメック星人の若者", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/289/720/lqip.jpg?1555371141"},
		{"4万2千", "ネイル", "https://stat.ameba.jp/user_images/20160729/00/kingintama/b8/45/j/o0443033213709148842.jpg?caw=800"},
		{"0", "気を消した状態のクリリン", "http://blog-imgs-73.fc2.com/w/e/a/wearebottoms/o0710036210254420393.jpg"},
		{"1万8千", "ベジータ", "http://ドラゴンボールネタバレ.net/img/snapcrab_noname_2018-10-14_13-52-56_no-00.jpg"},
		{"18万", "フリーザ編の悟空", "https://i2.wp.com/www.so-ra-no-i-ro.com/wp-content/uploads/hatena/20171103223459.jpg?ssl=1"},
		{"12万", "ギニュー", "https://sukianihaha1253free.up.seesaa.net/image/394.jpg"},
		{"53万", "フリーザ", "https://i.pinimg.com/originals/1b/60/e6/1b60e6a47cd700f728a5e8849ba8d42c.jpg"},
		{"100万", "フリーザ第二形態", "https://2.bp.blogspot.com/-c0LPzHCSYgg/W4lVZlwPR3I/AAAAAAAACd0/uPLAUjxr9EIAOu6v99tMV_LZK_d66p45QCPcBGAYYCw/s1600/Dl50ugmU8AAmdiP.jpg"},
		{"5", "トランクス", "https://i.pinimg.com/474x/a5/3b/58/a53b58311eb9b29e96f3240c1a193f44.jpg"},
		{"1億5千", "超サイヤ人の悟空", "http://ドラゴンボールネタバレ.net/img/snapcrab_noname_2019-2-3_20-25-56_no-00.jpg"},
		{"0.001", "ウミガメ", "https://static.wikia.nocookie.net/dragonball/images/b/bf/Turtles.png/revision/latest?cb=20100506110950&path-prefix=ja"},
		{"1万", "バーダック", "https://static.wikia.nocookie.net/dragonball/images/0/05/BardockAtPM.png/revision/latest?cb=20100413151015&path-prefix=ja"},
		{"2", "赤ん坊のカカロット", "https://i.ytimg.com/vi/650Pv7tAEpU/hqdefault.jpg"},
		{"53億", "ユニバーサル・スタジオ・ジャパンのアトラクションに出てくるフリーザ", "https://出産準備品.com/img/dragonball-usj56.jpg"},
		{"260", "若い頃のピッコロ大魔王", "https://static.wikia.nocookie.net/dragonball/images/d/d8/Kingpiccoloonthrone2.jpg/revision/latest/scale-to-width-down/340?cb=20091106104432"},
		{"1500", "ラディッツ", "https://img.animanch.com/2019/07/bf00e8dd.jpg"},
		{"610", "チャオズ", "https://images.ciatr.jp/2019/12/w_828/T31UqveUXp4N7kQsHEknC1HmzROuq5UovhSHniNN.jpeg"},
		{"4000", "ナッパ", "https://stat.ameba.jp/user_images/20171116/23/don-jara7633/52/a9/j/o0640048014071966582.jpg"},
		{"1万8千", "キュイ", "https://i.ytimg.com/vi/n-S_qqOgQSU/maxresdefault.jpg"},
		{"2万3千", "ザーボン", "https://i.ytimg.com/vi/1U06OmRB4R4/maxresdefault.jpg"},
		{"2万2千", "ドドリア", "https://pbs.twimg.com/media/B45BEFKCcAADetO.jpg"},
		{"130", "チチ", "https://cdn-ak.f.st-hatena.com/images/fotolife/c/catherine_yanagi/20170308/20170308203121.png"},
		{"970", "ヤジロベー", "https://i.ytimg.com/vi/9qmVIAS3su4/maxresdefault.jpg"},
		{"210", "桃白白", "https://gazoku.com/wp-content/uploads/1590095445598.jpg"},
		{"190", "カリン", "https://static.wikia.nocookie.net/dragonball/images/7/7a/KorinMajinBuuSaga.png/revision/latest?cb=20100508123353&path-prefix=ja"},
		{"1030", "ミスターポポ", "https://static.wikia.nocookie.net/dragonball/images/a/a6/Mr.PopoLookOut02.png/revision/latest?cb=20100508144926&path-prefix=ja"},
		{"3500", "界王", "https://static.wikia.nocookie.net/dragonball/images/6/61/KingKaiNV.png/revision/latest?cb=20100509063817&path-prefix=ja"},
		{"25億", "ゴジータ", "https://otakaranet.com/wp-content/uploads/2018/12/maxresdefault-3.jpg"},
		{"14億", "超サイヤ人ブロリー", "https://www.fashion-press.net/img/movies/23018/BB8.jpg"},
	}
)
