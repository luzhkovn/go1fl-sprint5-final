package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("invalid parameters / неверные параметры")
	}
	ms := MeanSpeed(steps, height, duration)
	return ((duration.Minutes() * ms * weight) / float64(minInH)) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if duration <= 0 || steps <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("invalid parameters / неверные параметры")
	}
	ms := MeanSpeed(steps, height, duration)
	return (duration.Minutes() * ms * weight) / float64(minInH), nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()

}

func Distance(steps int, height float64) float64 {
	stridelength := height * stepLengthCoefficient
	return (float64(steps) * stridelength) / float64(mInKm)
}
