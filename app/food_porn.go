package app

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/utyosu/rfe/db"
	"github.com/utyosu/rfe/env"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func isFoodPornExecute(m *discordgo.MessageCreate) bool {
	return isContainKeywords(m.Content, foodPornKeywords)
}

func actionFoodPorn(m *discordgo.MessageCreate) {
	baseWord := foodPornBaseWords[rand.Intn(len(foodPornBaseWords))]
	foodWord := foodPornSearchWords[rand.Intn(len(foodPornSearchWords))]
	searchWord := fmt.Sprintf("%v %v", baseWord, foodWord)
	page := strconv.Itoa(rand.Intn(10) + 1)

	request, err := http.NewRequest("GET", foodPornGoogleSearchUrl, nil)
	if err != nil {
		failedFoodPorn(m, errors.WithStack(err))
		return
	}

	params := request.URL.Query()
	params.Add("key", env.GoogleSearchApiKey)
	params.Add("cx", env.GoogleSearchApiCx)
	params.Add("q", searchWord)
	params.Add("num", "1")
	params.Add("start", page)
	params.Add("searchType", "image")
	request.URL.RawQuery = params.Encode()

	log.Println(request.URL.String())

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		failedFoodPorn(m, errors.WithStack(err))
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		failedFoodPorn(m, errors.WithStack(err))
		return
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		failedFoodPorn(m, errors.WithStack(err))
		return
	}
	if !jsonParsed.Exists("items", "0", "link") {
		err := errors.New(
			fmt.Sprintf(
				"結果を取得できませんでした。\nRequest: %v\nResponse: %v",
				request.URL.String(),
				string(body),
			),
		)
		failedFoodPorn(m, err)
		return
	}
	linkRaw := jsonParsed.Search("items", "0", "link").String()

	// 結果にダブルクォートが入っているので外しておく
	linkUrl := strings.Replace(linkRaw, "\"", "", -1)

	sendMessage(m.ChannelID, linkUrl)

	if _, err := db.InsertActivity(m.Author.ID, db.ActivityKindFoodPorn); err != nil {
		postSlackWarning(errors.WithStack(err))
	}
}

func failedFoodPorn(m *discordgo.MessageCreate, err error) {
	postSlackWarning(err)
	sendMessage(m.ChannelID, "ごちそうの取得に失敗しました(´・ω・`)ｼｮﾎﾞﾝ")
}

const (
	foodPornGoogleSearchUrl = "https://www.googleapis.com/customsearch/v1"
)

var (
	foodPornKeywords = []string{
		"ごちそう",
		"ご馳走",
	}
	foodPornBaseWords = []string{
		"飯テロ",
		"美味しい",
		"インスタ映え",
	}
	foodPornSearchWords = []string{
		"アイスクリーム",
		"インド料理",
		"うどん",
		"おでん",
		"おにぎり",
		"おばんざい",
		"オムライス",
		"お好み焼き",
		"かき氷",
		"かに料理",
		"からあげ",
		"カレーライス",
		"キャラ弁当",
		"クレープ",
		"ケーキ",
		"サムギョプサル",
		"サンドイッチ",
		"シュークリーム",
		"シュラスコ",
		"しらす丼",
		"ジンギスカン",
		"すき焼き",
		"ステーキ",
		"そば",
		"たこ焼き",
		"タピオカ",
		"チーズタッカルビ",
		"チーズフォンデュ",
		"チャーハン",
		"ちゃんこ鍋",
		"チョコレート",
		"つけ麺",
		"ドーナッツ",
		"トルコ料理",
		"とんかつ",
		"パスタ",
		"パフェ",
		"ハワイ料理",
		"パンケーキ",
		"ハンバーガー",
		"ハンバーグ",
		"ピザ",
		"ひつまぶし",
		"ふぐ料理",
		"フランス料理",
		"プリン",
		"フルコース",
		"フレンチトースト",
		"もつ鍋",
		"もんじゃ焼き",
		"ラーメン",
		"ろばた焼き",
		"沖縄料理",
		"家系ラーメン",
		"懐石料理",
		"釜飯",
		"牛丼",
		"刺身",
		"寿司",
		"焼き鳥",
		"焼き飯",
		"焼肉",
		"親子丼",
		"担々麺",
		"中華料理",
		"天ぷら",
		"天丼",
		"二郎系ラーメン",
		"麻婆豆腐",
		"油そば",
		"和食",
	}
)
