package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	// Command-line flags
	monthPtr := flag.String("month", "", "Month (numbers/English words)")
	yearPtr := flag.Int("year", time.Now().Year(), "Year")
	vacationPtr := flag.Int("vacation", 0, "Vacation hours")
	hourlyRatePtr := flag.Float64("hourlyRate", 0.0, "Hourly rate for work")
	moneyPtr := flag.Bool("money", false, "Calculate money based on hourly rate")
	flag.Parse()

	// Validate input
	if *monthPtr == "" || (*vacationPtr < 0 || *vacationPtr > 24) || (*hourlyRatePtr <= 0) {
		log.Fatal("Invalid input parameters")
	}

	// Parse month input
	month := parseMonthInput(*monthPtr)

	// Calculate pre-work hours
	workHours := calculateWorkHours(month, *yearPtr, *vacationPtr)

	// Output results
	fmt.Printf("Total pre-work hours for %s %d (excluding %d vacation hours): %.2f hours\n", month, *yearPtr, *vacationPtr, workHours)

	if *moneyPtr {
		totalMoney := calculateMoney(workHours, *hourlyRatePtr)
		fmt.Printf("Total earnings: $%.2f\n", totalMoney)
	}
}

// Parse month input
func parseMonthInput(monthInput string) string {
	monthInput = strings.ToLower(monthInput)
	switch monthInput {
	case "1", "january":
		return "January"
	case "2", "february":
		return "February"
	case "3", "march":
		return "March"
	case "4", "april":
		return "April"
	case "5", "may":
		return "May"
	case "6", "june":
		return "June"
	case "7", "july":
		return "July"
	case "8", "august":
		return "August"
	case "9", "september":
		return "September"
	case "10", "october":
		return "October"
	case "11", "november":
		return "November"
	case "12", "december":
		return "December"
	default:
		log.Fatal("Invalid month input")
		return ""
	}
}

// Calculate pre-work hours
func calculateWorkHours(month string, year, vacationHours int) float64 {
	monthNew, _ := time.Parse("January", month)
	firstDay := time.Date(year, time.Month(monthNew.Month()), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	workdays := 0

	// Calculate workdays excluding weekends and vacation days
	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			workdays++
		}
	}
	fmt.Printf("Workdays in %s %d: %d\n", month, year, workdays)

	workHours := float64(workdays*8) - float64(vacationHours)
	return workHours
}

// Calculate money based on hourly rate
func calculateMoney(workHours, hourlyRate float64) float64 {
	return workHours * hourlyRate
}
