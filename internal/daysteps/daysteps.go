package daysteps

import (
	"5-sprint-final/internal/personaldata"
	"5-sprint-final/internal/spentenergy"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("неверный формат данных")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return err
	}
	ds.Duration = duration

	return nil
}
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше 0")
	}
	distance := float64(ds.Steps) * StepLength / 1000
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
