package screenController

import (
	"fmt"
	"github.com/solanafish/millibelle/internal/text"
	"time"
)

func Draw(x int, color[3]uint8, pixels[32][8][3]uint8) [32][8][3]uint8 {
	position := x
	hours, minutes, seconds := time.Now().Local().Clock()

	hourString := fmt.Sprintf("%02d", hours)
	step1 := text.DrawString(hourString, position, color, pixels)
	position += 2 * (text.FontWidth + 1) - 1

	var sign byte
	if seconds % 2 == 0 {
		sign = ' '
	} else {
		sign = ':'
	}
	step2 := text.DrawSign(sign, position, color, step1)
	position += 3

	minuteString := fmt.Sprintf("%02d", minutes)
	step3 := text.DrawString(minuteString, position, color, step2)

	return step3
}
