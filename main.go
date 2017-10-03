package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"math/rand"
	"time"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					imageURL := getResMessage(message.Text)
					postMessage := linebot.NewImageMessage(imageURL, imageURL)
					if _, err = bot.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func getResMessage(reqMessage string) (message string) {
	// resMessages := [3]string{"わかるわかる", "それで？それで？", "からの〜？"}

	rand.Seed(time.Now().UnixNano())
	math := rand.Intn(4)
	switch math {
	case 0:
		// message = resMessages[math]
		message = "https://images-na.ssl-images-amazon.com/images/I/513WLTl9xRL._AC_UL320_SR234,320_.jpg"
	case 1:
		// message = reqMessage + "じゃねーよw"
		message = "https://i2.wp.com/fatesoku.com/wp-content/uploads/DK9ymaiVoAA9GN6.jpg"
	case 2:
		message = "https://pbs.twimg.com/media/DKe7aJgUEAAVPpW?format=jpg"
	case 3:
		message = "https://i.ytimg.com/vi/nodfeKY5Fes/maxresdefault.jpg"
	}
	// imageURL := "https://img.atwikiimg.com/www9.atwiki.jp/f_go/attach/497/179/070-d3.png"
	// message := linebot.NewImageMessage(imageURL, imageURL)
	// message = "https://img.atwikiimg.com/www9.atwiki.jp/f_go/attach/497/179/070-d3.png"
	return
}
