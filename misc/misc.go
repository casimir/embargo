package misc

import (
	"time"
)

func formatTime(layout string) string {
	return time.Now().Format(layout)
}
