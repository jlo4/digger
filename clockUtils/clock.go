package clockUtils

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func UpdateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("Time: 03:04:05")
	clock.SetText(formattedTime)
}
