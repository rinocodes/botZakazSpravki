package telegobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var pointbot *TeleGoBot

type Message struct {
	Message_id int `json:"message_id"`
	From       struct {
		Id int `json:"id"`
	} `json:"from"`
	Text    string `json:"text"`
	Contact struct {
		Phone_number string `json:"phone_number"`
	} `json:"contact"`
	Entities []struct {
		Type string `json:"type"`
	} `json:"entities"`
}

type IncomingMessages struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Update_id       int     `json:"update_id"`
		Message         Message `json:"message"`
		HandlerFunction struct {
			Name string `json:"data"`
			From struct {
				Id int `json:"id"`
			} `json:"from"`
		} `json:"callback_query"`
		Type string
	} `json:"result"`
}

type TeleGoBot struct {
	LastMessage int
	TeleToken   string
	FuncStart   reflect.Value
	Client      http.Client
}

func NewBot(timeoutinseconds int) *TeleGoBot {

	bot := TeleGoBot{}
	bot.Client.Timeout, _ = time.ParseDuration(fmt.Sprintf("%ds", timeoutinseconds))
	// bot.DefaultText = "default text"
	pointbot = &bot
	return &bot
}

func (bot *TeleGoBot) SetStartFunction(startFunction func(t string, d string, uI int, mI int)) {

	bot.FuncStart = reflect.ValueOf(startFunction)

}

func (bot *TeleGoBot) RunLongPolling() {

	for true {

		incomingMessages := bot.GetUpdates()

		for _, message := range incomingMessages.Result {
			var messageData string
			var userID int
			var messageType string
			messageID := message.Update_id
			if message.Message.Text != "" && strings.HasPrefix(message.Message.Text, "/") {
				messageType = "botCommand"
				messageData = message.Message.Text
				userID = message.Message.From.Id
			} else if message.Message.Text != "" {
				messageType = "text"
				messageData = message.Message.Text
				userID = message.Message.From.Id
			} else if message.Message.Contact.Phone_number != "" {
				messageType = "contact"
				messageData = message.Message.Contact.Phone_number
				userID = message.Message.From.Id
			} else if message.HandlerFunction.Name != "" {
				messageType = "callbackData"
				messageData = message.HandlerFunction.Name
				userID = message.HandlerFunction.From.Id
			}

			// 		if message.HandlerFunction.Name != "" {
			// 			incomingMessages.Result[imess].Type = "CallbackData"
			// 			messageData = message.HandlerFunction.Name
			// 			fromid = message.HandlerFunction.From.Id
			// 		}

			inValue := make([]reflect.Value, 4)
			inValue[0] = reflect.ValueOf(messageType)
			inValue[1] = reflect.ValueOf(messageData)
			inValue[2] = reflect.ValueOf(userID)
			inValue[3] = reflect.ValueOf(messageID)
			go bot.FuncStart.Call(inValue)
			bot.LastMessage = messageID + 1
		}
	}
}

func processMessage(message Message) {

}

func (bot *TeleGoBot) GetUpdates() IncomingMessages {

	urlGetUpdates := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=%g&offset=%s", bot.TeleToken, bot.Client.Timeout.Seconds(), strconv.Itoa(bot.LastMessage))

	resp, err := bot.Client.Get(urlGetUpdates)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var incomingMessages IncomingMessages
	json.Unmarshal([]byte(body), &incomingMessages)

	// incomingMessages.()["result"]
	fmt.Println(string(body))
	return incomingMessages

}

func SendMessage(messageText string, userID int, keyboard interface{}) {

	urlGetUpdates := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", pointbot.TeleToken, userID, messageText)

	if keyboard, ok := keyboard.(Keyboard); ok {
		fmt.Println(keyboard)
		replyMarkup, _ := json.Marshal(keyboard)
		replyMarkupStr := string(replyMarkup)
		urlGetUpdates = urlGetUpdates + "&reply_markup=" + replyMarkupStr
		urlGetUpdates = strings.Replace(urlGetUpdates, `"keyboard":null,`, `"keyboard": [],`, -1)
		urlGetUpdates = strings.Replace(urlGetUpdates, `"inline_keyboard":null,`, `"inline_keyboard": [],`, -1)
	}

	fmt.Println(string(urlGetUpdates))
	resp, err := pointbot.Client.Get(urlGetUpdates)

	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	defer resp.Body.Close()
	// urlSendMessage := ba.Host + ba.TeleToken + ba.SendMessageAddress + strconv.Itoa(m.ChatID)
	// if m.Text != "" {
	// 	urlSendMessage = urlSendMessage + "&text=" + m.Text
	// }
	// if m.ReplyMarkup != "" {
	// 	urlSendMessage = urlSendMessage + "&reply_markup=" + m.ReplyMarkup
	// }
	// resp, err := ba.Client.Get(urlSendMessage)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// body, err := ioutil.ReadAll(resp.Body)
	//

}
