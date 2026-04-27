package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid format: expected 'steps,duration', got %d parts", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps: %v", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be greater than 0, got %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %v", err)
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration must be greater than 0, got %d", duration)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil || calories <= 0 {
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories)

	return result
}
