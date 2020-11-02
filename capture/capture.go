package capture

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func Doit() {
	deviceID := 0
	saveFile := "./save.png"

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}

	gocv.PutText(&img, fmt.Sprint("I got you"), image.Pt(30, 70),
		gocv.FontHersheyPlain, 3.2, color.RGBA{255, 99, 71, 0}, 2)

	gocv.IMWrite(saveFile, img)
}
