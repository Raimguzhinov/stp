package main

import "fmt"

func main() {
	etaValues := []int{300, 400, 512} // Значения η*₂
	programmers := 5                  // Число программистов
	v := 20                           // Число команд ассемблера, отлаживаемых в день
	tauFactor := 0.5                  // τ = 0.5 * Tk

	results := processEtaValues(etaValues, programmers, v, tauFactor)

	fmt.Printf("%-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-10s\n",
		"η*₂", "k", "i", "K", "N", "V", "P", "Tk", "B", "tn")
	for _, r := range results {
		fmt.Printf("%-10d %-10d %-10d %-10d %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f %-10.2f\n",
			r.ModulesLowerLevel*8, r.ModulesLowerLevel, r.HierarchyLevels, r.TotalModules,
			r.ProgramLength, r.ProgramVolume, r.AssemblerCommands,
			r.ProgrammingTime, r.InitialErrors, r.Reliability)
	}
}
