package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

// "qwertyuiopasdfghjklçzxcvbnm123456789"
// "qwertyuiopa;.,/*)(&)"
// "qqqwwweee   222333555666999*+/   hhhh      ;;;;.......     "
const (
	ASCIILIST           = "....... ...   .......eerrrtttyyyuuuiiiooopppaaa          "
	LENASCII            = uint16(len(ASCIILIST) - 1)
	MAXGRAYCOLOR uint16 = 65535
	QUALITY = 100
)

func nearestNeighborScaling(img image.Image, w, h int) image.Image {
	widthimg := img.Bounds().Max.X
	heightimg := img.Bounds().Max.Y
	MinPoint := image.Point{0, 0}
	MaxPoint := image.Point{w, h}
	scalonateImg := image.NewRGBA(image.Rectangle{MinPoint, MaxPoint})
	for i := 0; i < widthimg; i++ {
		for j := 0; j < heightimg; j++ {
			srcX := float32(i) / float32(w) * float32(widthimg)
			srcY := float32(j) / float32(h) * float32(heightimg)
			srcX = min(srcX, float32(widthimg))
			srcY = min(srcY, float32(heightimg))

			r, g, b, a := img.At(int(srcX), int(srcY)).RGBA()
			fmt.Printf("valor anterior: %d\n",r)
			newPix := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
			fmt.Printf("valor final: %d\n", newPix.R)
			scalonateImg.SetRGBA(i, j, newPix)
		}
	}
	return scalonateImg
}

func RGB2GrayColor(c color.Color) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r + g + b) / 3)
}

func graySacaleImage(img image.Image) image.Gray16 {
	newImage := image.NewGray16(image.Rectangle{img.Bounds().Min, img.Bounds().Max})
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			newImage.SetGray16(i, j, color.Gray16{RGB2GrayColor(img.At(i, j))})
		}
	}
	return *newImage
}

func accisImage(img image.Gray16) []string {
	ascciImg := make([]string, img.Bounds().Max.Y)
	for i := 0; i < img.Bounds().Max.X; i++ {
		for j := 0; j < img.Bounds().Max.Y; j++ {
			//fmt.Println(img.Gray16At(i,j).Y)
			convertValue := float32(LENASCII) * (float32(img.Gray16At(i, j).Y) / float32(MAXGRAYCOLOR))
			//fmt.Println(convertValue)
			ascciImg[j] += " " + string(ASCIILIST[uint16(convertValue)])
		}
	}
	return ascciImg
}

func main() {
	// Read image from file that already exists

	fmt.Println("Start Code")
	existingImageFile, err := os.Open("images.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)

	loadedImage, err := jpeg.Decode(existingImageFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	newScalonate := nearestNeighborScaling(loadedImage, 500, 500)

	fileImageScalonate, err := os.Create("scalonate.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	jpeg.Encode(fileImageScalonate, newScalonate, &jpeg.Options{QUALITY})

	newGray := graySacaleImage(newScalonate)
	fileImageGray, err := os.Create("grayImage.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	jpeg.Encode(fileImageGray, image.Image(&newGray), &jpeg.Options{QUALITY})
	
	imgascii := accisImage(newGray)
	fileTxt, err := os.Create("texto.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileTxt.Close()

	fileTxt.Seek(0, 0)

	for _, v := range imgascii {
		fmt.Println(v)
		fileTxt.WriteString(v + "\n")
	}
}
