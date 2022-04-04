package event

import (
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/asaskevich/EventBus"
	"github.com/olup/kobowriter/utils"
)

type KeyEvent struct {
	IsCtrl      bool
	IsAlt       bool
	IsAltGr     bool
	IsShift     bool
	IsShiftLock bool
	KeyCode     int
	IsChar      bool
	KeyChar     string
	KeyValue    string
}

var KeyboardLang string = utils.AZERTY

func BindKeyEvent(k *keylogger.KeyLogger, b EventBus.Bus, lang string) {
	event := KeyEvent{
		IsShift:     false,
		IsShiftLock: false,
		IsAltGr:     false,
		IsAlt:       false,
		IsCtrl:      false,
	}

	if lang != KeyboardLang {
		KeyboardLang = lang
	}

	var currentLang string = strings.Clone(KeyboardLang)

	keyMapMaj := GetKeyMapMaj(currentLang)
	keyMapAltGr := GetKeyMapAltGr(currentLang)
	keyMap := GetKeyMap(currentLang)

	events := k.Read()
	for e := range events {
		// Check if the keyboard language has been changed
		if currentLang != KeyboardLang {
			currentLang = strings.Clone(KeyboardLang)

			keyMapMaj = GetKeyMapMaj(currentLang)
			keyMapAltGr = GetKeyMapAltGr(currentLang)
			keyMap = GetKeyMap(currentLang)
		}

		if e.Type == keylogger.EvKey {

			keyValue := keyMap[int(e.Code)]
			if keyValue == "" {
				continue
			}

			event.KeyChar = ""
			event.IsChar = false
			event.KeyCode = int(e.Code)
			event.KeyValue = keyValue

			if e.KeyPress() {
				switch keyValue {
				case "KEY_L_SHIFT", "KEY_R_SHIFT":
					event.IsShift = true
				case "KEY_CAPSLOCK":
					event.IsShiftLock = !event.IsShiftLock
				case "KEY_ALT_GR":
					event.IsAltGr = true
				case "KEY_L_ALT":
					event.IsAlt = true
				case "KEY_L_CTRL", "KEY_R_CTRL":
					event.IsCtrl = true

				}
			}

			if e.KeyRelease() {
				switch keyValue {
				case "KEY_L_SHIFT", "KEY_R_SHIFT":
					event.IsShift = false
				case "KEY_ALT_GR":
					event.IsAltGr = false
				case "KEY_L_GR":
					event.IsAlt = false
				case "KEY_L_CTRL", "KEY_R_CTRL":
					event.IsCtrl = false
				}
			} else {

				// letters
				if utils.IsLetter(keyValue) {
					event.IsChar = true
					if event.IsShift || event.IsShiftLock {
						event.KeyChar = keyMapMaj[int(e.Code)]
					} else if event.IsAltGr {
						event.KeyChar = keyMapAltGr[int(e.Code)]
					} else {
						event.KeyChar = keyMap[int(e.Code)]
					}
				}

				b.Publish("KEY", event)
			}

		}
	}
	println("lost keyboadr")
	b.Publish("REQUIRE_KEYBOARD")
}
