package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"gopkg.in/ini.v1"
)

// install https://jmeubank.github.io/tdm-gcc/download/ fr g++ and gcc

func main() {
	cfg, err := ini.Load("settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	n := screenshot.NumActiveDisplays()
	fmt.Printf("Number of active displays: %v\n", n)
	display, err := cfg.Section("").Key("display").Int()
	if err != nil {
		fmt.Printf("Fail to read value from ini fo display: %v", err)
		os.Exit(1)
	}
	fmt.Println("Using display #", display, "from settings.ini")

	eventHook := hook.Start()
	var e hook.Event
	var key string
	lastKeyEvent := time.Now().Unix()

	for e = range eventHook {
		now := time.Now().Unix()
		if now-5 > lastKeyEvent {
			if e.Kind == hook.KeyDown {
				// https://www.toptal.com/developers/keycode
				key = string(e.Keychar)
				switch key {
				case "	":
					fmt.Println("pressed tab")
					time.Sleep(100 * time.Millisecond)
					makeScreenshot(display)
					lastKeyEvent = time.Now().Unix()
				default:
					fmt.Printf("pressed %v Event: %v \n", key, e)
				}
			}
			if e.Kind == hook.MouseDown {
				// https://www.toptal.com/developers/keycode
				switch e.Button {
				case 1: //LMB
					fmt.Println("Pressed LMB")
					makeScreenshot(display)
					lastKeyEvent = time.Now().Unix()
				case 2: //RMB
					fmt.Println("Pressed RMB")
				default:
					fmt.Printf("pressed Event: %v \n", e)
				}
			}
		}
	}

	for i := 0; ; i++ {

		// makeScreenshot(display)
		// img := makeScreenshot(display)
		// img := gcv.IMRead("screens\\1_1920x1080_20230324121123.png")
		// map_startscreen_illios_wel := gcv.IMRead("screens\\map_start\\map_startscreen_illios_well.png")
		// minVal, maxVal, minLoc, maxLoc := gcv.FindImg(map_startscreen_illios_wel, img)
		// fmt.Println("find: ", minVal, maxVal, minLoc, maxLoc)
		// time.Sleep(5 * time.Second)
	}
}

func makeScreenshot(display int) image.RGBA {
	bounds := screenshot.GetDisplayBounds(display)
	fmt.Println("Bounds used:", bounds)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("screens\\%d_%dx%d_%v.png", display, bounds.Dx(), bounds.Dy(), time.Now().Format("20060102150405"))
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("#%d : %v \"%s\"\n", display, bounds, fileName)
	return *img
}
