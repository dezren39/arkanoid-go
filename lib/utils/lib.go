package utils

import (
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/hajimehoshi/ebiten"
)

// LogError prints error and exits if error is not nil
func LogError(err error) {
	if err != nil {
		log.Printf("%v\n\n", err)
		log.Println("Original error line:")

		stackLines := strings.Split(string(debug.Stack()), "\n")
		for iLine, line := range stackLines {
			if strings.Contains(line, "utils.LogError") {
				stackLines = stackLines[iLine+2 : iLine+4]
				break
			}
		}

		for _, line := range stackLines {
			log.Println(line)
		}
		os.Exit(1)
	}
}

// KeyMap is a US keyboard mapping
var KeyMap = map[string]ebiten.Key{
	"0":            ebiten.Key0,
	"1":            ebiten.Key1,
	"2":            ebiten.Key2,
	"3":            ebiten.Key3,
	"4":            ebiten.Key4,
	"5":            ebiten.Key5,
	"6":            ebiten.Key6,
	"7":            ebiten.Key7,
	"8":            ebiten.Key8,
	"9":            ebiten.Key9,
	"A":            ebiten.KeyA,
	"B":            ebiten.KeyB,
	"C":            ebiten.KeyC,
	"D":            ebiten.KeyD,
	"E":            ebiten.KeyE,
	"F":            ebiten.KeyF,
	"G":            ebiten.KeyG,
	"H":            ebiten.KeyH,
	"I":            ebiten.KeyI,
	"J":            ebiten.KeyJ,
	"K":            ebiten.KeyK,
	"L":            ebiten.KeyL,
	"M":            ebiten.KeyM,
	"N":            ebiten.KeyN,
	"O":            ebiten.KeyO,
	"P":            ebiten.KeyP,
	"Q":            ebiten.KeyQ,
	"R":            ebiten.KeyR,
	"S":            ebiten.KeyS,
	"T":            ebiten.KeyT,
	"U":            ebiten.KeyU,
	"V":            ebiten.KeyV,
	"W":            ebiten.KeyW,
	"X":            ebiten.KeyX,
	"Y":            ebiten.KeyY,
	"Z":            ebiten.KeyZ,
	"Apostrophe":   ebiten.KeyApostrophe,
	"Backslash":    ebiten.KeyBackslash,
	"Backspace":    ebiten.KeyBackspace,
	"CapsLock":     ebiten.KeyCapsLock,
	"Comma":        ebiten.KeyComma,
	"Delete":       ebiten.KeyDelete,
	"Down":         ebiten.KeyDown,
	"End":          ebiten.KeyEnd,
	"Enter":        ebiten.KeyEnter,
	"Equal":        ebiten.KeyEqual,
	"Escape":       ebiten.KeyEscape,
	"F1":           ebiten.KeyF1,
	"F2":           ebiten.KeyF2,
	"F3":           ebiten.KeyF3,
	"F4":           ebiten.KeyF4,
	"F5":           ebiten.KeyF5,
	"F6":           ebiten.KeyF6,
	"F7":           ebiten.KeyF7,
	"F8":           ebiten.KeyF8,
	"F9":           ebiten.KeyF9,
	"F10":          ebiten.KeyF10,
	"F11":          ebiten.KeyF11,
	"F12":          ebiten.KeyF12,
	"GraveAccent":  ebiten.KeyGraveAccent,
	"Home":         ebiten.KeyHome,
	"Insert":       ebiten.KeyInsert,
	"KP0":          ebiten.KeyKP0,
	"KP1":          ebiten.KeyKP1,
	"KP2":          ebiten.KeyKP2,
	"KP3":          ebiten.KeyKP3,
	"KP4":          ebiten.KeyKP4,
	"KP5":          ebiten.KeyKP5,
	"KP6":          ebiten.KeyKP6,
	"KP7":          ebiten.KeyKP7,
	"KP8":          ebiten.KeyKP8,
	"KP9":          ebiten.KeyKP9,
	"KPAdd":        ebiten.KeyKPAdd,
	"KPDecimal":    ebiten.KeyKPDecimal,
	"KPDivide":     ebiten.KeyKPDivide,
	"KPEnter":      ebiten.KeyKPEnter,
	"KPEqual":      ebiten.KeyKPEqual,
	"KPMultiply":   ebiten.KeyKPMultiply,
	"KPSubtract":   ebiten.KeyKPSubtract,
	"Left":         ebiten.KeyLeft,
	"LeftBracket":  ebiten.KeyLeftBracket,
	"Menu":         ebiten.KeyMenu,
	"Minus":        ebiten.KeyMinus,
	"NumLock":      ebiten.KeyNumLock,
	"PageDown":     ebiten.KeyPageDown,
	"PageUp":       ebiten.KeyPageUp,
	"Pause":        ebiten.KeyPause,
	"Period":       ebiten.KeyPeriod,
	"PrintScreen":  ebiten.KeyPrintScreen,
	"Right":        ebiten.KeyRight,
	"RightBracket": ebiten.KeyRightBracket,
	"ScrollLock":   ebiten.KeyScrollLock,
	"Semicolon":    ebiten.KeySemicolon,
	"Slash":        ebiten.KeySlash,
	"Space":        ebiten.KeySpace,
	"Tab":          ebiten.KeyTab,
	"Up":           ebiten.KeyUp,
	"Alt":          ebiten.KeyAlt,
	"Control":      ebiten.KeyControl,
	"Shift":        ebiten.KeyShift,
}

// GamepadButtonMap is a gamepad button mapping
var GamepadButtonMap = map[string]ebiten.GamepadButton{
	"GamepadButton0":  ebiten.GamepadButton0,
	"GamepadButton1":  ebiten.GamepadButton1,
	"GamepadButton2":  ebiten.GamepadButton2,
	"GamepadButton3":  ebiten.GamepadButton3,
	"GamepadButton4":  ebiten.GamepadButton4,
	"GamepadButton5":  ebiten.GamepadButton5,
	"GamepadButton6":  ebiten.GamepadButton6,
	"GamepadButton7":  ebiten.GamepadButton7,
	"GamepadButton8":  ebiten.GamepadButton8,
	"GamepadButton9":  ebiten.GamepadButton9,
	"GamepadButton10": ebiten.GamepadButton10,
	"GamepadButton11": ebiten.GamepadButton11,
	"GamepadButton12": ebiten.GamepadButton12,
	"GamepadButton13": ebiten.GamepadButton13,
	"GamepadButton14": ebiten.GamepadButton14,
	"GamepadButton15": ebiten.GamepadButton15,
	"GamepadButton16": ebiten.GamepadButton16,
	"GamepadButton17": ebiten.GamepadButton17,
	"GamepadButton18": ebiten.GamepadButton18,
	"GamepadButton19": ebiten.GamepadButton19,
	"GamepadButton20": ebiten.GamepadButton20,
	"GamepadButton21": ebiten.GamepadButton21,
	"GamepadButton22": ebiten.GamepadButton22,
	"GamepadButton23": ebiten.GamepadButton23,
	"GamepadButton24": ebiten.GamepadButton24,
	"GamepadButton25": ebiten.GamepadButton25,
	"GamepadButton26": ebiten.GamepadButton26,
	"GamepadButton27": ebiten.GamepadButton27,
	"GamepadButton28": ebiten.GamepadButton28,
	"GamepadButton29": ebiten.GamepadButton29,
	"GamepadButton30": ebiten.GamepadButton30,
	"GamepadButton31": ebiten.GamepadButton31,
}

// MouseButtonMap is a mouse button mapping
var MouseButtonMap = map[string]ebiten.MouseButton{
	"MouseButtonLeft":   ebiten.MouseButtonLeft,
	"MouseButtonRight":  ebiten.MouseButtonRight,
	"MouseButtonMiddle": ebiten.MouseButtonMiddle,
}
