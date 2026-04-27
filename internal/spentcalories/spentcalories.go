package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid format: expected 'steps,activity,duration', got %d parts", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps: %v", err)
	}

	activity := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid duration: %v", err)
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient

	distanceMeters := float64(steps) * stepLength

	distanceKm := distanceMeters / float64(mInKm)

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)

	durationHours := duration.Hours()
	if durationHours == 0 {
		return 0
	}

	speed := dist / durationHours

	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if weight <= 0 {
		return "", fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return "", fmt.Errorf("height must be positive, got %.2f", height)
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	switch activity {
	case "Бег":
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км.\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f\n",
			activity, durationHours, dist, speed, calories,
		), nil

	case "Ходьба":
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
				"Длительность: %.2f ч.\n"+
				"Дистанция: %.2f км.\n"+
				"Скорость: %.2f км/ч\n"+
				"Сожгли калорий: %.2f\n",
			activity, durationHours, dist, speed, calories,
		), nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	meanSpeedValue := meanSpeed(steps, height, duration)

	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeedValue * durationMinutes) / float64(minInH)
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive, got %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be positive, got %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	meanSpeedValue := meanSpeed(steps, height, duration)

	durationMinutes := duration.Minutes()
	calories := (weight * meanSpeedValue * durationMinutes) / float64(minInH)
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
