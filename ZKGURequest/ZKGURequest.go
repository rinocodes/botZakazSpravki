package zkgurequest

import (
	"bot/cache"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type responseMessage struct {
	Status      string
	Description string
	Data        struct {
		IO               string
		ArrayOfEmployees []struct {
			GUIDEmployee     string
			TypeOfEmployment string
		}
	}
	// OrderedDocuments []struct {
	// }

}

func SetEmployeeDataFromZKGU(userID int, phoneNumber string) (string, bool) {

	login := os.Getenv("loginZKGU")
	password := os.Getenv("passwordZKGU")
	url := os.Getenv("urlZKGU")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(login, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Phonenumber", phoneNumber)
	req.Header.Set("userid", strconv.Itoa(userID))
	resp, _ := client.Do(req)

	if resp.StatusCode != 200 {
		return "Сервис авторизации временно недоступен, попробуйте немного позже", false
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var responseMessage responseMessage
	json.Unmarshal([]byte(body), &responseMessage)

	if responseMessage.Status != "Ok" {
		return responseMessage.Description, false
	}

	userCache := cache.GetUserCache(userID)
	userCache.UserData.IO = responseMessage.Data.IO
	cache.SaveUserCache(userID, userCache)
	for _, employee := range responseMessage.Data.ArrayOfEmployees {
		userCache.UserData.EmployeeGUID = employee.GUIDEmployee
		if employee.TypeOfEmployment == "Основное место работы" {
			break
		}
	}

	fmt.Println(string(body))

	return "ok", true

}

// {
// 	"Status": "Ok",
// 	"Description": "",
// 	"Data": {
// 	"FIO": "Филатов Ярослав Андреевич",
// 	"DateOfBirth": "1993-10-04T00:00:00",
// 	"OrderedDocuments": [],
// 	"GUIDPhysicalPerson": null,
// 	"ArrayOfEmployees": [
// 	{
// 	"GUIDEmployee": "4cb0e659-1da5-11ea-9d2b-005056981c46",
// 	"TypeOfEmployment": "Основное место работы"
// 	}
// 	],
// 	"IO": "Ярослав Андреевич"
// 	}
// 	}
