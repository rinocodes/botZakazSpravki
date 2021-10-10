package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	zkgurequest "bot/ZKGURequest"
	"bot/cache"
	"bot/telegobot"

	"github.com/joho/godotenv"
)

type orderDetails struct {
	typeOfCertificate string `json:"Тип"`
}

func main() {

	runtime.GOMAXPROCS(2)

	//Initialize the cache
	cache.InitUsersCacheByDefault()

	//Create a new bot
	bot := telegobot.NewBot(300)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	bot.TeleToken = os.Getenv("teleToken")
	bot.SetStartFunction(Start)
	bot.RunLongPolling()

}

func Start(messageType string, messageData string, userID int, messageID int, messageIDCallback int) {

	userCache := cache.GetUserCache(userID)

	switch messageType {

	case "botCommand":

		if messageData == "/start" {

			if userCache == nil {

				keyboard := telegobot.NewKeyboard()
				keyboard.AddInlineButtonBelow("НИТУ «МИСиС»", "OrganizationMISIS")
				keyboard.AddInlineButtonBelow("СТИ НИТУ «МИСиС»", "OrganizationSTIMISIS")

				telegobot.SendMessage("Выберите место работы:", userID, keyboard)

			} else {
				telegobot.SendMessage("Выберите тип справки:", userID, getKeyboardSelectionСertificate())
			}

		}

	case "text":

		if messageData == "Привет" {
			telegobot.SendMessage("Привет, как тебя зовут?", userID, nil)
		} else {
			telegobot.SendMessage("Тебя зовут "+messageData, userID, nil)
		}

	case "callbackData":

		if strings.Contains(messageData, "Organization") {
			OrganizationStr := strings.Replace(messageData, "Organization", "", 1)
			userCache := cache.SetUserCache(userID)
			userCache.UserData.Organization = OrganizationStr
			cache.SaveUserCache(userID, userCache)
			messageText := `Добрый день, уважаемые коллеги!` +
				`%0AДля получения доступа к функциям чат-бота, потвердите личность, нажав на кнопку "Отправить номер телефона"` +
				`%0A%0AКнопка находится под строкой ввода сообщения` +
				`%0A%F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87 %F0%9F%91%87`
			keyboard := telegobot.NewKeyboard()
			keyboard.AddButtonRequestContact("Отправить номер телефона")
			telegobot.SendMessage(messageText, userID, keyboard)
			telegobot.DeleteMessage(userID, messageIDCallback)
			return
		}
		if userCache == nil {
			telegobot.SendMessage("Ваша сессия окончена, для заказа справки нажмите /start", userID, nil)
			return
		} else if strings.Contains(messageData, "Сertificate") {
			СertificateStr := strings.Replace(messageData, "Сertificate", "", 1)
			userCache.UserData.Certificate = СertificateStr
			cache.SaveUserCache(userID, userCache)
			keyboard := telegobot.NewKeyboard()
			keyboard.AddInlineButtonBelow("1", "Quantity1")
			keyboard.AddInlineButtonBelow("2", "Quantity2")
			keyboard.AddInlineButtonBelow("3", "Quantity3")

			telegobot.EditMessage("Выберите количество:", userID, keyboard, messageIDCallback)
		} else if strings.Contains(messageData, "Quantity") {
			QuantityStr := strings.Replace(messageData, "Quantity", "", 1)
			userCache.UserData.Quantity, _ = strconv.Atoi(QuantityStr)
			cache.SaveUserCache(userID, userCache)
			keyboard := telegobot.NewKeyboard()
			keyboard.AddInlineButtonBelow("Да", "СonfirmationYes")
			keyboard.AddInlineButtonBelow("Нет", "СonfirmationNo")

			PresentationOfTheCertificateName := getPresentationOfTheCertificateName(userCache.UserData.Certificate)
			messageText := fmt.Sprintf("Вы заказали: %s%s (%d шт.)", "%0A", PresentationOfTheCertificateName, userCache.UserData.Quantity)
			telegobot.EditMessage(messageText, userID, keyboard, messageIDCallback)
		} else if strings.Contains(messageData, "Сonfirmation") {
			СonfirmationStr := strings.Replace(messageData, "Сonfirmation", "", 1)
			switch СonfirmationStr {
			case "Yes":
				sendOrderingOfCertificate(userCache)
			case "No":
				messageText := "Робот НИТУ «МИСиС» желает Вам хорошего дня!%0A%0AЧтобы заказать новую справку нажмите /start"
				telegobot.EditMessage(messageText, userID, nil, messageIDCallback)
			}
		}
	case "contact":
		if userCache != nil {
			userCache.UserData.PhoneNumber = messageData

			cache.SaveUserCache(userID, userCache)
			details, ok := zkgurequest.SetEmployeeDataFromZKGU(userID, userCache.UserData.PhoneNumber)
			userCache := cache.GetUserCache(userID)
			if !ok {
				telegobot.SendMessage(details, userID, nil)
				return
			}

			telegobot.SendMessage(userCache.UserData.IO+", Ваш номер телефона подтвержден.", userID, nil)
			telegobot.SendMessage("Выберите тип справки:", userID, getKeyboardSelectionСertificate())

		} else {
			telegobot.SendMessage("Ваша сообщение не распознано, для заказа справки нажмите /start", userID, nil)
		}
	}

}

func sendOrderingOfCertificate(userCache *cache.UserCache) {

	orderDetails := orderDetails{}
	orderDetails.typeOfCertificate = getTypeOfCertificate(userCache.UserData.Certificate)
	// "GUIDСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
	// "НомерСправки": "b8d07ce7-4033-4b67-a057-1ded697b5f28",
	// "Тип": "Справка2НДФЛ",
	// "ДатаЗаказа": "2021-08-26T12:09:47.840216",
	// "GUIDСотрудника": "4d2f4606-b4dd-11e3-af64-005056a702bd",
	// "Количество": "2",
	// "Комментарий": "Заказ справки с чат-бота",
	// "ПериодСправкиДляПосольства": "",
	// "ФИОРебенка": "",
	// "ДатаРожденияРебенка": "",
	// "РасчетныйПериод": "2021",
	// "ИсточникЗаказа": "БотТелеграмм"
	// }

}

func getTypeOfCertificate(presentationInEnglish string) string {

	var presentationInRussian string
	switch presentationInEnglish {
	case "FromThePlaceOfWork":
		presentationInRussian = "СправкаСМестаРаботы"
	case "CopyOfTheEmploymentRecord":
		presentationInRussian = "КопияТрудовойКнижки"
	}

	return presentationInRussian
}

func getPresentationOfTheCertificateName(Certificate string) string {

	var presentation string
	switch Certificate {
	case "FromThePlaceOfWork":
		presentation = "Справка с места работы"
	}

	return presentation
}

func getKeyboardSelectionСertificate() telegobot.Keyboard {

	keyboard := telegobot.NewKeyboard()
	keyboard.AddInlineButtonBelow("Справка с места работы", "СertificateFromThePlaceOfWork")
	keyboard.AddInlineButtonBelow("Копия трудовой книжки", "СertificateCopyOfTheEmploymentRecord")
	keyboard.AddInlineButtonBelow("Справка 2 НДФЛ", "Сertificate2NDFL")
	keyboard.AddInlineButtonBelow("Справка для посольства", "СertificateForTheEmbassy")
	keyboard.AddInlineButtonBelow("Неполучении пособия до 1,5 лет", "Сertificate15years")
	keyboard.AddInlineButtonBelow("Неполучении пособия до 3 лет", "Сertificate3years")
	keyboard.AddInlineButtonBelow("Неполучении единовременного пособия", "СertificateNonReceipt")
	keyboard.AddInlineButtonBelow("О заработке для расчета пособий", "СertificateForBenefits")

	return keyboard
}
