package telegobot

type KeyboardButton struct {
	Text           string `json:"text"`
	RequestContact bool   `json:"request_contact"`
}
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
type Keyboard struct {
	ResizeKeyboard        bool                     `json:"resize_keyboard"`
	OneTimeKeyboard       bool                     `json:"one_time_keyboard"`
	KeyboardButtons       [][]KeyboardButton       `json:"keyboard"`
	InlineKeyboardButtons [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (k *Keyboard) AddInlineButtonBelow(text, callbackdata string) {

	newInlineButton := InlineKeyboardButton{}
	newInlineButton.Text = text
	newInlineButton.CallbackData = callbackdata

	newInlineButtonSlice := []InlineKeyboardButton{}
	newInlineButtonSlice = append(newInlineButtonSlice, newInlineButton)

	k.InlineKeyboardButtons = append(k.InlineKeyboardButtons, newInlineButtonSlice)

}

func (k *Keyboard) AddButtonRequestContact(text string) {

	newButton := KeyboardButton{}
	newButton.Text = text
	newButton.RequestContact = true
	k.ResizeKeyboard = false

	newButtonSlice := []KeyboardButton{}
	newButtonSlice = append(newButtonSlice, newButton)

	k.KeyboardButtons = append(k.KeyboardButtons, newButtonSlice)

}

func (k *Keyboard) Add() {

}
func NewKeyboard() Keyboard {

	newKeyboard := Keyboard{}
	newKeyboard.OneTimeKeyboard = true
	newKeyboard.ResizeKeyboard = true

	return newKeyboard

}

// type KeyboardButton struct {
// 	Text            string `json:"text"`
// 	Request_contact bool   `json:"request_contact"`
// }

// type InlineKeyboardButton struct {
// 	Text          string `json:"text"`
// 	Callback_data string `json:"callback_data"`
// }

// type Keyboard struct {
// 	button string
// }

// type Keyboard map[string]interface{}
// type KeyboardButton map[string]interface{}
// type InlineKeyboardButton map[string]interface{}

// KeyboardButt [][]interface{} `json:"keyboard"`
// KeyboardButtonArr         [][]KeyboardButton       `json:"keyboard"`
// KeyboardButtonArray       [][]KeyboardButton       `json:"keyboard"`
// InlineKeyboardButtonArray [][]InlineKeyboardButton `json:"inline_keyboard"`
// // 	Resize_keyboard   bool `json:"resize_keyboard"`
// // 	One_time_keyboard bool `json:"one_time_keyboard"`
// // }
// func GetNewKeyboardByDefault() Keyboard {
// 	keyboard := Keyboard{}
// 	keyboard["resize_keyboard"] = true
// 	keyboard["one_time_keyboard"] = true
// 	return keyboard
// }

// func (k *Keyboard) ByDefault() {

// 	(*k)["resize_keyboard"] = true
// 	(*k)["one_time_keyboard"] = true

// }

// func (k *Keyboard) AddButtonRequestContact(text string) {

// 	newKeyboardButton := KeyboardButton{}
// 	var newKeyboardButtonArray []KeyboardButton
// 	var newKeyboardButtonArrayArray [][]KeyboardButton

// 	newKeyboardButtonArray = append(newKeyboardButtonArray, newKeyboardButton)
// 	newKeyboardButtonArrayArray = append(newKeyboardButtonArrayArray, newKeyboardButtonArray)
// 	(*k)["keyboard"] = newKeyboardButtonArrayArray

// }

// func (k *Keyboard) AddInlineKeyboardButton(text string, callbackData string) {

// 	newInlineKeyboardButton := KeyboardButton{}
// 	var newInlineKeyboardButtonArray []KeyboardButton
// 	var newInlineKeyboardButtonArrayArray [][]KeyboardButton

// 	newInlineKeyboardButton["text"] = text
// 	newInlineKeyboardButton["callback_data"] = callbackData

// 	newInlineKeyboardButtonArray = append(newInlineKeyboardButtonArray, newInlineKeyboardButton)
// 	newInlineKeyboardButtonArrayArray = append(newInlineKeyboardButtonArrayArray, newInlineKeyboardButtonArray)

// 	// F := (*k)["inline_keyboard"].(data)
// 	// if  > 0 {
// 	// iKS := (*k)["one_time_keyboard"].([][]KeyboardButton)
// 	if _, ok := (*k)["inline_keyboard"]; ok {
// 		kb := (*k)["inline_keyboard"].([][]KeyboardButton)
// 		kb = append(kb, newInlineKeyboardButtonArray)
// 		(*k)["inline_keyboard"] = kb
// 	} else {
// 		(*k)["inline_keyboard"] = newInlineKeyboardButtonArrayArray
// 	}

// }

// type InlineKeyboard struct {
// 	ButtonType [][]InlineKeyboardButton `json:"inline_keyboard"`
// }

// func AddKeyboardButton(text string, request_contact bool) {

// }
