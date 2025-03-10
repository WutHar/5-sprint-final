package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"5-sprint-final/internal/personaldata"
	"5-sprint-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("неверный формат данных")
	}

	stepsStr := strings.TrimSpace(data[0])
	if stepsStr == "" {
		return errors.New("количество шагов не указано")
	}
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	ds.Steps = steps

	durationStr := strings.TrimSpace(data[1])
	if durationStr == "" {
		return errors.New("длительность не указана")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
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

	return fmt.Sprintf("  Количество шагов: %d\n  Дистанция: %.2f км\n  Сожжено калорий: %.2f ккал\n", ds.Steps, distance, calories), nil
}
