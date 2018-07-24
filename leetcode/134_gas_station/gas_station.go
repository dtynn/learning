package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

STATION_LOOP:
	for i := 0; i < n; i++ {
		amount := 0
		for step := 0; step < n; step++ {
			stop := (i + step) % n
			amount += gas[stop]
			amount -= cost[stop]
			if amount < 0 {
				continue STATION_LOOP
			}
		}

		return i
	}

	return -1
}
