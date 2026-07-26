// There are k worker and a list of n workloads, divide workloads among worker such that worker will get minimized maximum workload
// example, like split the workload among all workers such that each will get <= x workload.
package main

import "fmt"

func MaximumNSum(workloads []int) (int, int) {
	max := workloads[0]
	sum := workloads[0]
	for i := 1; i < len(workloads); i++ {
		if max < workloads[i] {
			max = workloads[i]
		}
		sum += workloads[i]
	}

	return max, sum
}

func checkPossibility(workloads []int, worker int, target int) bool {

	currentWorker := 1
	currentWorkerLoad := 0

	for i := 0; i < len(workloads); i++ {

		if (currentWorkerLoad + workloads[i]) <= target {
			currentWorkerLoad = currentWorkerLoad + workloads[i]
		} else {

			// no worker is left to assign remaining task
			if currentWorker == worker {
				return false
			}
			// have one worker but single workload is greater than target
			if workloads[i] > target {
				return false
			}
			currentWorkerLoad = workloads[i]
			currentWorker++
		}
	}

	return true
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

func main() {
	workloads := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	k := 3 //workers

	minimumWorkload, maximumWorkload := MaximumNSum(workloads)
	fmt.Printf("Min %d, Max %d", minimumWorkload, maximumWorkload)
	ans := 10000

	for minimumWorkload <= maximumWorkload {

		mid := minimumWorkload + (maximumWorkload-minimumWorkload)/2

		res := checkPossibility(workloads, k, mid)

		if res {
			ans = min(ans, mid)
			maximumWorkload = mid - 1
		} else {
			minimumWorkload = mid + 1
		}
	}

	fmt.Println("Possible : %d", ans)

}
