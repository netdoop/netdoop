package pm

import (
	"regexp"
	"strconv"
	"time"

	"github.com/netdoop/netdoop/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var durationRegex = regexp.MustCompile(`^P(\d+Y)?(\d+M)?(\d+D)?(T(\d+H)?(\d+M)?(\d+(\.\d+)?S)?)?$`)

func ParseDuration(v string) (time.Duration, error) {
	match := durationRegex.FindStringSubmatch(v)
	utils.GetLogger().Warn("debug", zap.String("v", v), zap.Any("match", match))
	if match == nil {
		return time.Duration(0), errors.New("invalid ISO 8601 duration format")
	}

	years := parseDurationPart(match, 1)
	months := parseDurationPart(match, 2)
	days := parseDurationPart(match, 3)
	hours := parseDurationPart(match, 5)
	minutes := parseDurationPart(match, 6)
	seconds := parseDurationPart(match, 7)

	totalDuration := time.Duration(int(years*365.25*24*60*60+months*30*24*60*60+days*24*60*60+hours*60*60+minutes*60+seconds)) * time.Second

	return totalDuration, nil
}

func parseDurationPart(match []string, index int) float64 {
	if match[index] != "" {
		v, _ := strconv.ParseFloat(match[index][:len(match[index])-1], 64) // excluding the last character which is a letter (Y, M, D, H, M, S)
		return v
	}
	return 0
}
