package main

import (
	"log"
	"os"

	"bot/telegobot"

	"github.com/joho/godotenv"
)

func main() {

	bot := telegobot.NewBot(300)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	bot.TeleToken = os.Getenv("teleToken")
	bot.SetStartFunction(Start)
	bot.RunLongPolling()

}

func Start(messageType string, messageData string, userID int, messageID int) {

	switch messageType {

	case "botCommand":
		if messageData == "/start" {

		}
		telegobot.SendMessage("Hi Olesia", userID)
	case "text":
		if messageData == "Привет" {
			telegobot.SendMessage("Привет, как тебя зовут?", userID)
		} else {
			telegobot.SendMessage("Тебя зовут "+messageData, userID)
		}
		// telegobot.SendMessage("You say "+messageData, userID)

		// case time.Wednesday:
		//     fmt.Println("Сегодня среда.")

		// case time.Thursday:
		//     fmt.Println("Сегодня четверг.")

		// case time.Friday:
		//     fmt.Println("Сегодня пятница.")

		// case time.Saturday:
		//     fmt.Println("Сегодня суббота.")

		// case time.Sunday:
		//     fmt.Println("Сегодня воскресенье.")
	}
	// }
	// if messageType == "bot_command" {

	// 	inCache := ba.Cache.Get(strconv.Itoa(fromid))
	// 	if inCache == nil {
	// 		// var mess urlstruct.Message
	// 		// mess.ChatID = fromid
	// 		// mess.Text = `Действие не распознано, нажмите /start`
	// 		// ba.SendMessage(mess)
	// 		// return
	// 	} else {
	// 		switch messageData {
	// 		case "СertificateCopyOfTheEmploymentRecord":
	// 			// outgoingMessage.Тип = "КопияТрудовойКнижки"

	// 			newKeyboard := keyboard.GetNewKeyboardByDefault()
	// 			newKeyboard.AddInlineKeyboardButton("1", "QuantityOne")
	// 			newKeyboard.AddInlineKeyboardButton("2", "QuantityTwo")
	// 			newKeyboard.AddInlineKeyboardButton("3", "QuantityThree")

	// 			var mess urlstruct.Message
	// 			mess.ChatID = fromid
	// 			mess.Text = `Укажите количество`
	// 			mess.AddKeyboard(newKeyboard)
	// 			ba.SendMessage(mess)
	// 			// inCache
	// 			iC := inCache.(urlstruct.ReqEmployee)
	// 			iC.Certificate = "КопияТрудовойКнижки"
	// 			ba.Cache.Set(strconv.Itoa(fromid), iC)
	// 			return
	// 			// println(inCache)
	// 			// {
	// 			// 	"GUIDСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
	// 			// 	"НомерСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
	// 			// 	"Тип": "Справка2НДФЛ",
	// 			// 	"ДатаЗаказа": "2021-08-26T12:09:47.840216",
	// 			// 	"GUIDСотрудника": "4d2f4606-b4dd-11e3-af64-005056a702bd",
	// 			// 	"Количество": "2",
	// 			// 	"Комментарий": "Заказ справки с чат-бота",
	// 			// 	"ПериодСправкиДляПосольства": "",
	// 			// 	"ФИОРебенка": "",
	// 			// 	"ДатаРожденияРебенка": "",
	// 			// 	"РасчетныйПериод": "2021",
	// 			// 	"ИсточникЗаказа": "БотТелеграмм"
	// 			// 	}
	// 		case "QuantityOne":
	// 			newUUID := uuid.New().String()
	// 			outgoingMessage := urlstruct.OutgoingMessage{
	// 				GUIDСправки:  newUUID,
	// 				НомерСправки: newUUID,
	// 				Тип:          "КопияТрудовойКнижки",
	// 			}
	// 			println(outgoingMessage)
	// 		}
	// }

	// }

	// var mess urlstruct.Message
	// mess.ChatID = fromid

	// // 	messageText := m.Message.Text
	// if messageData == "/start" {

	// 	newKeyboard := keyboard.Keyboard{}
	// 	newKeyboard.ByDefault()
	// 	newKeyboard.AddButtonRequestContact("Отправить номер телефона")
	// 	// newKeyboard := keyboard.GetNewKeyboardByDefault()
	// 	//
	// 	// keyboard.AddButtonRequestContact(&newKeyboard, text)
	// 	// 		// newKeyboard[]
	// 	// 		// newKeyboard.ByDefault()
	// 	// 		// newKeyboard.AddButtonRequestContact(`Отправить номер`)

	// 	mess.Text = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
	// 	mess.AddKeyboard(newKeyboard)
	// 	// 		ba.SendMessage(mess)
	// } else if messageType == "Contact" {

	// 	// mess.Text = "Ваш номер телефона не найден в системе, добавьте его в личном кабинете или обратитесь в отдел кадров"

	// 	reqEmp := urlstruct.GetЕmployeesData(messageData, strconv.Itoa(fromid))
	// 	ba.Cache.Set(strconv.Itoa(fromid), reqEmp)
	// 	if reqEmp.Data.IO != "" {
	// 		newKeyboard := keyboard.GetNewKeyboardByDefault()
	// 		newKeyboard.AddInlineKeyboardButton("Справка с места работы", "СertificateFromThePlaceOfWork")
	// 		newKeyboard.AddInlineKeyboardButton("Справка 2 НДФЛ", "Сertificate2NDFL")
	// 		newKeyboard.AddInlineKeyboardButton("Справка для посольства", "СertificateForTheEmbassy")
	// 		newKeyboard.AddInlineKeyboardButton("Копия трудовой книжки", "СertificateCopyOfTheEmploymentRecord")
	// 		newKeyboard.AddInlineKeyboardButton("Неполучении пособия до 1,5 лет", "Сertificate15years")
	// 		newKeyboard.AddInlineKeyboardButton("Неполучении пособия до 3 лет", "Сertificate3years")
	// 		newKeyboard.AddInlineKeyboardButton("Неполучении единовременного пособия", "СertificateNonReceipt")
	// 		newKeyboard.AddInlineKeyboardButton("О заработке для расчета пособий", "СertificateForBenefits")

	// 		mess.Text = reqEmp.Data.IO + ", ваш номер телефона подтвержден"
	// 		ba.SendMessage(mess)

	// 		mess.Text = "Выберите тип требуемой справки"
	// 		mess.AddKeyboard(newKeyboard)
	// 	} else {
	// 		mess.Text = "Ваш номер телефона не найден в системе, добавьте его в личном кабинете или обратитесь в отдел кадров"
	// 	}

	// 	println(reqEmp.Data.IO)

	// } else if messageType == "Text" {
	// 	mess.Text = `Действие не распознано, нажмите /start`
	// }

	// ba.SendMessage(mess)

}

// func main() {

// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalln(err)
// 	}

// 	var botApi urlstruct.BotApi
// 	botApi.TeleToken = os.Getenv("teleToken")
// 	botApi.ByDefault()
// 	botApi.SetStartFunction(Start)
// 	botApi.RunLongPolling()
// }

// func Start(messageType string, messageData string, fromid int, ba urlstruct.BotApi) {

// if messageType == "CallbackData" {

// 	inCache := ba.Cache.Get(strconv.Itoa(fromid))
// 	if inCache == nil {
// 		// var mess urlstruct.Message
// 		// mess.ChatID = fromid
// 		// mess.Text = `Действие не распознано, нажмите /start`
// 		// ba.SendMessage(mess)
// 		// return
// 	} else {
// 		switch messageData {
// 		case "СertificateCopyOfTheEmploymentRecord":
// 			// outgoingMessage.Тип = "КопияТрудовойКнижки"

// 			newKeyboard := keyboard.GetNewKeyboardByDefault()
// 			newKeyboard.AddInlineKeyboardButton("1", "QuantityOne")
// 			newKeyboard.AddInlineKeyboardButton("2", "QuantityTwo")
// 			newKeyboard.AddInlineKeyboardButton("3", "QuantityThree")

// 			var mess urlstruct.Message
// 			mess.ChatID = fromid
// 			mess.Text = `Укажите количество`
// 			mess.AddKeyboard(newKeyboard)
// 			ba.SendMessage(mess)
// 			// inCache
// 			iC := inCache.(urlstruct.ReqEmployee)
// 			iC.Certificate = "КопияТрудовойКнижки"
// 			ba.Cache.Set(strconv.Itoa(fromid), iC)
// 			return
// 			// println(inCache)
// 			// {
// 			// 	"GUIDСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
// 			// 	"НомерСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
// 			// 	"Тип": "Справка2НДФЛ",
// 			// 	"ДатаЗаказа": "2021-08-26T12:09:47.840216",
// 			// 	"GUIDСотрудника": "4d2f4606-b4dd-11e3-af64-005056a702bd",
// 			// 	"Количество": "2",
// 			// 	"Комментарий": "Заказ справки с чат-бота",
// 			// 	"ПериодСправкиДляПосольства": "",
// 			// 	"ФИОРебенка": "",
// 			// 	"ДатаРожденияРебенка": "",
// 			// 	"РасчетныйПериод": "2021",
// 			// 	"ИсточникЗаказа": "БотТелеграмм"
// 			// 	}
// 		case "QuantityOne":
// 			newUUID := uuid.New().String()
// 			outgoingMessage := urlstruct.OutgoingMessage{
// 				GUIDСправки:  newUUID,
// 				НомерСправки: newUUID,
// 				Тип:          "КопияТрудовойКнижки",
// 			}
// 			println(outgoingMessage)
// 		}
// 	}

// }

// var mess urlstruct.Message
// mess.ChatID = fromid

// // 	messageText := m.Message.Text
// if messageData == "/start" {

// 	newKeyboard := keyboard.Keyboard{}
// 	newKeyboard.ByDefault()
// 	newKeyboard.AddButtonRequestContact("Отправить номер телефона")
// 	// newKeyboard := keyboard.GetNewKeyboardByDefault()
// 	//
// 	// keyboard.AddButtonRequestContact(&newKeyboard, text)
// 	// 		// newKeyboard[]
// 	// 		// newKeyboard.ByDefault()
// 	// 		// newKeyboard.AddButtonRequestContact(`Отправить номер`)

// 	mess.Text = `Добрый день, уважаемые коллеги! Для получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"`
// 	mess.AddKeyboard(newKeyboard)
// 	// 		ba.SendMessage(mess)
// } else if messageType == "Contact" {

// 	// mess.Text = "Ваш номер телефона не найден в системе, добавьте его в личном кабинете или обратитесь в отдел кадров"

// 	reqEmp := urlstruct.GetЕmployeesData(messageData, strconv.Itoa(fromid))
// 	ba.Cache.Set(strconv.Itoa(fromid), reqEmp)
// 	if reqEmp.Data.IO != "" {
// 		newKeyboard := keyboard.GetNewKeyboardByDefault()
// 		newKeyboard.AddInlineKeyboardButton("Справка с места работы", "СertificateFromThePlaceOfWork")
// 		newKeyboard.AddInlineKeyboardButton("Справка 2 НДФЛ", "Сertificate2NDFL")
// 		newKeyboard.AddInlineKeyboardButton("Справка для посольства", "СertificateForTheEmbassy")
// 		newKeyboard.AddInlineKeyboardButton("Копия трудовой книжки", "СertificateCopyOfTheEmploymentRecord")
// 		newKeyboard.AddInlineKeyboardButton("Неполучении пособия до 1,5 лет", "Сertificate15years")
// 		newKeyboard.AddInlineKeyboardButton("Неполучении пособия до 3 лет", "Сertificate3years")
// 		newKeyboard.AddInlineKeyboardButton("Неполучении единовременного пособия", "СertificateNonReceipt")
// 		newKeyboard.AddInlineKeyboardButton("О заработке для расчета пособий", "СertificateForBenefits")

// 		mess.Text = reqEmp.Data.IO + ", ваш номер телефона подтвержден"
// 		ba.SendMessage(mess)

// 		mess.Text = "Выберите тип требуемой справки"
// 		mess.AddKeyboard(newKeyboard)
// 	} else {
// 		mess.Text = "Ваш номер телефона не найден в системе, добавьте его в личном кабинете или обратитесь в отдел кадров"
// 	}

// 	println(reqEmp.Data.IO)

// } else if messageType == "Text" {
// 	mess.Text = `Действие не распознано, нажмите /start`
// }

// ba.SendMessage(mess)

// }
