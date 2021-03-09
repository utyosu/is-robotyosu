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
		{"6", "兎人参化", "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTdzeiTidEl4fgBk3eM09BOFjq_S5njB7FIKw&usqp=CAU"},
		{"60", "牛魔王", "https://ami.animecharactersdatabase.com/uploads/chars/2855-1962955328.png"},
		{"120", "人造人間８号", "https://static.wikia.nocookie.net/dragonball/images/f/f2/Android8.jpg/revision/latest/scale-to-width-down/700?cb=20100517085520&path-prefix=ja"},
		{"110", "ミイラくん", "https://lineup.toei-anim.co.jp/upload/save_image/episode/5912/story_img_1.jpg"},
		{"130", "アックマン", "https://knouprese.com/wp-content/uploads/2020/01/17c2bfc4.png"},
		{"145", "孫悟飯じいちゃん", "https://lineup.toei-anim.co.jp/upload/save_image/episode/5916/story_img_1.jpg"},
		{"170", "タンバリン", "https://renote.jp/uploads/image/file/116804/migtambaling.jpg"},
		{"150", "シンバル", "https://comic-kingdom.jp/wp-content/uploads/2020/06/%E3%82%B7%E3%83%B3%E3%83%90%E3%83%AB8.jpg"},
		{"210", "ドラム", "https://pbs.twimg.com/media/DeLuc49VMAEhDaO?format=jpg&name=small"},
		{"300", "シェン", "https://static.wikia.nocookie.net/dragonball/images/2/21/KamiAsHero.png/revision/latest/scale-to-width-down/700?cb=20100520130456&path-prefix=ja"},
		{"1万8千", "地球襲来時のベジータ", "https://livedoor.blogimg.jp/ponpokonwes/imgs/9/7/97cb68b0-s.jpg"},
		{"2", "デンデ", "https://i2.wp.com/dragonball.littleair.xyz/wordpress/wp-content/uploads/2020/04/t5eseeeeeesg.jpg?w=536&ssl=1"},
		{"1500", "アプール", "https://d2dcan0armyq93.cloudfront.net/photo/odai/600/4bffc055d4dbd2a98ea4d28f459bf459_600.jpg"},
		{"1万", "グルド", "https://stat.ameba.jp/user_images/20140506/16/madoushi-hoi/b0/9a/p/o0449025812932066790.png?caw=800"},
		{"5万4千", "ジース", "https://dragonball-zxk.com/wp-content/uploads/2018/05/1bc1f4abb162b5ab28569109b52fea18.png"},
		{"5万3千", "バータ", "https://livedoor.blogimg.jp/mitsuemon_4649/imgs/b/f/bf77d61d.png"},
		{"5万5千", "リクーム", "https://dragonball-zxk.com/wp-content/uploads/2018/05/ef006ab24dedd2451a00bab49cf4daf2.png"},
		{"2万3千", "悟空と入れ替わったギニュー", "https://lineup.toei-anim.co.jp/upload/save_image/episode/2731/story_img_3.jpg"},
		{"100万", "ネイルと同化したピッコロ", "https://lh3.googleusercontent.com/-IgQnpfyNdX0/W5YgSPp75LI/AAAAAAAAD0w/NNr5EbwLgaQ-fCHfn30weWkxuCz0I1HPwCE0YBhgL/s1024/P_20180910_162633.jpg"},
		{"300万", "フリーザ第三形態", "https://blog-imgs-111-origin.fc2.com/d/r/a/dragonballarekore/db895.jpg"},
		{"6千万", "フリーザ最終形態", "https://www.4gamer.net/games/381/G038174/20180525130/SS/004.jpg"},
		{"1億2千万", "フリーザ最終形態フルパワー", "https://cdn.gamerch.com/contents/wiki/911/entry/1505969267.jpg"},
		{"20億", "人造人間16号", "https://livedoor.blogimg.jp/gourmetmatome/imgs/a/c/acf7abcd-s.png"},
		{"15億", "人造人間17号", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/023/320/original.jpg?1526838992"},
		{"13億", "人造人間18号", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/023/328/original.png?1526840108"},
		{"2億5千万", "人造人間19号", "https://xn--uckwa2arq3e9dsam8d6e.net/img/snapcrab_noname_2018-11-1_17-58-30_no-00.jpg"},
		{"3億", "ドクターゲロ(人造人間20号)", "https://i1.wp.com/dragonball-neta.com/wp-content/uploads/10to20goud.jpg?w=400&ssl=1"},
		{"2億", "メカフリーザ", "https://cdn-ak.f.st-hatena.com/images/fotolife/s/sarugami33/20170828/20170828004111.jpg"},
		{"2億3千万", "コルド大王", "https://cdn.wikiwiki.jp/to/w/animegameex3/%E3%83%9E%E3%83%B3%E3%83%88%E3%82%AD%E3%83%A3%E3%83%A9%E4%B8%80%E8%A6%A7/::ref/koldo_db.jpg?rev=31801637cce24a187e6d698b4323621e&t=20181103162706"},
		{"20億", "神と同化したピッコロ", "https://blogimg.goo.ne.jp/user_image/30/a4/9d2653bddf50d0429385f5f0aa3f689c.jpg"},
		{"20億", "セル第一形態", "https://image.middle-edge.jp/medium/7fd8d624-88d5-48b8-be0c-33a2a0cc7e20.png?1469916825"},
		{"45億", "セル第二形態", "https://stat.ameba.jp/user_images/20190617/20/haikrbo/f2/07/j/o0884063114469862079.jpg?caw=800"},
		{"150億", "セル完全体", "https://blog-imgs-111-origin.fc2.com/d/r/a/dragonballarekore/db954.jpg"},
		{"80億", "セルジュニア", "http://doragonball-ziten.com/wp-content/uploads/2014/10/%E3%82%BB%E3%83%AB%E3%82%B8%E3%83%A5%E3%83%8B%E3%82%A2-300x221.png"},
		{"300億", "パーフェクトセル", "https://xn--uckwa2arq3e9dsam8d6e.net/img/snapcrab_noname_2018-11-1_18-26-43_no-00.jpg"},
		{"120億", "超サイヤ人の悟飯", "https://d2n122k3ikcd8z.cloudfront.net/img/cd8e0070-7a7d-41c8-a8fe-14f83199450b/2477.jpg"},
		{"300億", "超サイヤ人２の悟飯", "https://multimedia.okwave.jp/image/questions/13/137979/137979_L.jpg"},
		{"200万", "孫悟天", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/027/852/original.jpg?1527516740"},
		{"6億", "超サイヤ人の悟天", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/027/850/original.?1527516543"},
		{"250万", "トランクス", "https://stat.ameba.jp/user_images/20170103/17/bluelittlemoon-vagetable/f6/2e/j/o0527038913836963809.jpg?caw=800"},
		{"7億5千万", "超サイヤ人のトランクス", "https://stat.ameba.jp/user_images/20160717/15/bluelittlemoon-vagetable/83/13/j/o0480027113699489588.jpg?caw=800"},
		{"20億", "界王神", "https://xn--uckwa2arq3e9dsam8d6e.net/img/snapcrab_noname_2019-1-19_23-44-35_no-00.jpg"},
		{"2千億", "超サイヤ人３の孫悟空", "https://livedoor.blogimg.jp/ponpokonwes/imgs/7/9/7982b0b3-s.jpg"},
		{"500億", "ゴテンクス", "https://dragonball-t.com/wp-content/uploads/2016/12/0708-500x281.jpg"},
		{"3千億", "超サイヤ人3のゴテンクス", "https://dbz-gt-s.up.seesaa.net/image/story_img_4-thumbnail2.jpg"},
		{"3千億", "ベジット", "https://static.wikia.nocookie.net/dragonball/images/3/3e/VegitoDBZEp268.png/revision/latest/scale-to-width-down/400?cb=20101018162914"},
		{"9千億", "超サイヤ人ベジット", "https://i0.wp.com/manga-ch.com/wp-content/uploads/2019/02/c64ce4cc3667aad77ba3f5d85a9ba0ad.jpg"},
		{"9", "ビーデル", "https://comic-kingdom.jp/wp-content/uploads/2020/10/%E3%83%93%E3%83%BC%E3%83%87%E3%83%AB10.jpg"},
		{"8", "ミスターサタン", "https://livedoor.blogimg.jp/suko_ch-chansoku/imgs/5/a/5aa91ad0-s.jpg"},
		{"300万", "プイプイ", "https://livedoor.blogimg.jp/kakyuusenshi-jqb1ptn8/imgs/3/5/3522e226.png"},
		{"60億", "ヤコン", "https://livedoor.blogimg.jp/redking41_94/imgs/e/1/e1bf16f3-s.jpg"},
		{"250億", "ダーブラ", "https://s3-ap-northeast-1.amazonaws.com/cdn.bibi-star.jp/production/imgs/images/000/412/827/original.png?1571105310"},
		{"350億", "魔人ベジータ", "https://stat.ameba.jp/user_images/20160705/13/bluelittlemoon-vagetable/47/6e/j/o0480027113689911906.jpg?caw=800"},
		{"2万", "バビディ", "https://meigenkakugen.net/wp-content/uploads/Bobbidi.png"},
		{"4千億", "アルティメット悟飯", "https://livedoor.blogimg.jp/suko_ch-chansoku/imgs/3/3/33250608-s.png"},
		{"1200億", "魔人ブウ", "https://livedoor.blogimg.jp/husigiba5503/imgs/8/8/88eafb91.png"},
		{"2千億", "悪の魔人ブウ", "http://iam-publicidad.org/wp-content/uploads/2017/10/20-11-300x204.jpg"},
		{"5千憶", "ゴテンクスを吸収した魔人ブウ", "https://blog-imgs-111-origin.fc2.com/d/r/a/dragonballarekore/db946.png"},
		{"6千憶", "悟飯を吸収した魔人ブウ", "https://i0.wp.com/xn--n9jvd7d3d0ad5cwnpcu694dohxad89g.com/wp-content/uploads/2020/08/3-9.jpg?w=426&ssl=1"},
		{"1500億", "純粋な魔人ブウ", "https://static.wikia.nocookie.net/dragonball/images/a/a6/MajinBuuKidDebutNV.png/revision/latest/scale-to-width-down/180?cb=20150325220800"},
	}
)
