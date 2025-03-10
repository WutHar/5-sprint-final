package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"5-sprint-final/internal/personaldata"
	"5-sprint-final/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
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
	t.Steps = steps

	trainingType := strings.TrimSpace(data[1])
	if trainingType != "Бег" && trainingType != "Ходьба" {
		return errors.New("неизвестный тип тренировки")
	}
	t.TrainingType = trainingType

	durationStr := strings.TrimSpace(data[2])
	if durationStr == "" {
		return errors.New("длительность не указана")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)
	if t.Duration <= 0 {
		return "", errors.New("продолжительность должна быть больше 0")
	}
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)
	var calories float64
	if t.TrainingType == "Бег" {
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	} else {
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}

func (t Training) Print() {
	fmt.Printf("Тип тренировки: %s\nДлительность: %v\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration, spentenergy.Distance(t.Steps), spentenergy.MeanSpeed(t.Steps, t.Duration), t.Calories())
}

func (t Training) Calories() float64 {
	if t.TrainingType == "Бег" {
		return spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		return spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}
	return 0
}
