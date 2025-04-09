package main

import (
	"fmt"
	"math"
	"sort"
)

// Essential data structures
type Task struct {
	ID         int
	CPUDemand  float64
	GPUDemand  float64
	Popularity float64
}

type Node struct {
	ID          int
	CPUCapacity float64
	GPUs        []GPU
}

type GPU struct {
	ID       int
	Capacity float64
}

// Result structure
type Assignment struct {
	TaskID    int
	NodeID    int
	GPUIDs    []int
	Objective float64
}

// Main algorithm implementation
func main() {
	// 1. Create test case
	tasks, nodes := generateTestCase()

	// 2. Demonstrate algorithm
	fmt.Println("=== Quantum Annealing-Inspired GPU Scheduler Prototype ===")
	fmt.Println("Demonstration of the algorithm from Section 3.1")
	fmt.Println()

	// 3. Calculate task dominance
	fmt.Println("Step 1: Calculating task dominance relationships...")
	dominanceMatrix := calculateDominance(tasks)
	printDominanceMatrix(dominanceMatrix, tasks)

	// 4. Generate assignment
	fmt.Println("\nStep 2: Solving the resource allocation problem...")
	assignments := solveSchedulingProblem(tasks, nodes, dominanceMatrix)

	// 5. Analyze and visualize results
	fmt.Println("\nStep 3: Analyzing results...")
	printAssignments(assignments, tasks, nodes)
	calculateMetrics(assignments, tasks, nodes, dominanceMatrix)
}

// Calculate dominance relationships between tasks
func calculateDominance(tasks []Task) [][]bool {
	n := len(tasks)
	dominance := make([][]bool, n)
	for i := range dominance {
		dominance[i] = make([]bool, n)
	}

	for i, taskA := range tasks {
		for j, taskB := range tasks {
			if i == j {
				continue
			}

			// Task A dominates Task B if:
			// (A needs more GPU and at least same CPU) OR
			// (A needs at least same GPU and more CPU)
			if (taskA.GPUDemand > taskB.GPUDemand && taskA.CPUDemand >= taskB.CPUDemand) ||
				(taskA.GPUDemand >= taskB.GPUDemand && taskA.CPUDemand > taskB.CPUDemand) {
				dominance[i][j] = true
			}
		}
	}

	return dominance
}

// For prototype: Simplified solver that implements key algorithm concepts
func solveSchedulingProblem(tasks []Task, nodes []Node, dominance [][]bool) []Assignment {
	// For demonstration, implement a simplified version of the algorithm
	// that captures the essence of the approach without requiring a full QP solver

	// Sort tasks by "importance" (determined by dominance relationships)
	taskImportance := calculateTaskImportance(tasks, dominance)

	// Create sorted task indices by importance
	taskOrder := make([]int, len(tasks))
	for i := range taskOrder {
		taskOrder[i] = i
	}
	sort.Slice(taskOrder, func(i, j int) bool {
		return taskImportance[taskOrder[i]] > taskImportance[taskOrder[j]]
	})

	// Track node resource availability
	nodeCPURemaining := make([]float64, len(nodes))
	nodeGPURemaining := make([][]float64, len(nodes))
	for i, node := range nodes {
		nodeCPURemaining[i] = node.CPUCapacity
		nodeGPURemaining[i] = make([]float64, len(node.GPUs))
		for j, gpu := range node.GPUs {
			nodeGPURemaining[i][j] = gpu.Capacity
		}
	}

	// Assign tasks in order of importance
	assignments := []Assignment{}
	for _, taskIdx := range taskOrder {
		task := tasks[taskIdx]
		bestNode := -1
		bestGPUs := []int{}
		bestObjective := math.Inf(-1)

		// Try each node
		for nodeIdx, node := range nodes {
			// Check CPU capacity
			if nodeCPURemaining[nodeIdx] < task.CPUDemand {
				continue
			}

			// Check GPU capacity and find best allocation
			selectedGPUs := []int{}
			remainingGPUDemand := task.GPUDemand

			// Sort GPUs by available capacity (descending)
			type GPUWithCapacity struct {
				GPUID    int
				Capacity float64
			}
			gpus := make([]GPUWithCapacity, len(node.GPUs))
			for i, gpu := range node.GPUs {
				gpus[i] = GPUWithCapacity{
					GPUID:    gpu.ID,
					Capacity: nodeGPURemaining[nodeIdx][i],
				}
			}
			sort.Slice(gpus, func(i, j int) bool {
				return gpus[i].Capacity > gpus[j].Capacity
			})

			// Allocate GPUs greedily
			for _, gpu := range gpus {
				if remainingGPUDemand <= 0 {
					break
				}
				gpuIdx := -1
				for i, nodeGPU := range node.GPUs {
					if nodeGPU.ID == gpu.GPUID {
						gpuIdx = i
						break
					}
				}

				allocation := math.Min(gpu.Capacity, remainingGPUDemand)
				if allocation > 0 {
					selectedGPUs = append(selectedGPUs, gpu.GPUID)
					remainingGPUDemand -= allocation
				}
			}

			// If we couldn't allocate all needed GPU resources, skip this node
			if remainingGPUDemand > 1e-6 {
				continue
			}

			// Calculate objective value for this assignment
			// In a full implementation, this would evaluate the quadratic objective
			// For the prototype, use a simplified objective:
			// 1. Maximize tasks scheduled (constant +1)
			// 2. Minimize out-of-order penalties
			objective := 1.0 // Base value for scheduling a task

			// Calculate penalties for out-of-order assignments
			outOfOrderPenalty := 0.0
			for otherTaskIdx, otherTask := range tasks {
				if otherTaskIdx == taskIdx {
					continue
				}

				// Check if this task dominates another task that's already assigned
				if dominance[taskIdx][otherTaskIdx] {
					// Search if the dominated task is already assigned
					for _, assign := range assignments {
						if assign.TaskID == otherTask.ID {
							// This is an out-of-order assignment, penalize
							outOfOrderPenalty += 0.5
							break
						}
					}
				}
			}

			// Final objective for this assignment
			assignmentObjective := objective - outOfOrderPenalty

			// Update best assignment if this is better
			if assignmentObjective > bestObjective {
				bestNode = nodeIdx
				bestGPUs = selectedGPUs
				bestObjective = assignmentObjective
			}
		}

		// If we found a valid assignment
		if bestNode >= 0 {
			// Update resource tracking
			nodeCPURemaining[bestNode] -= task.CPUDemand

			// Update GPU capacity
			remainingGPUDemand := task.GPUDemand
			for _, gpuID := range bestGPUs {
				gpuIdx := -1
				for i, gpu := range nodes[bestNode].GPUs {
					if gpu.ID == gpuID {
						gpuIdx = i
						break
					}
				}

				allocation := math.Min(nodeGPURemaining[bestNode][gpuIdx], remainingGPUDemand)
				nodeGPURemaining[bestNode][gpuIdx] -= allocation
				remainingGPUDemand -= allocation
			}

			// Record assignment
			assignments = append(assignments, Assignment{
				TaskID:    task.ID,
				NodeID:    nodes[bestNode].ID,
				GPUIDs:    bestGPUs,
				Objective: bestObjective,
			})
		}
	}

	return assignments
}

// Calculate task importance based on dominance relationships
func calculateTaskImportance(tasks []Task, dominance [][]bool) []float64 {
	importance := make([]float64, len(tasks))

	for i := range tasks {
		// Count how many tasks this one dominates
		dominatesCount := 0
		for j := range tasks {
			if dominance[i][j] {
				dominatesCount++
			}
		}

		// Count how many tasks dominate this one
		dominatedByCount := 0
		for j := range tasks {
			if dominance[j][i] {
				dominatedByCount++
			}
		}

		// Importance is based on net dominance plus resource demands
		importance[i] = float64(dominatesCount-dominatedByCount) +
			tasks[i].GPUDemand + tasks[i].CPUDemand/10.0
	}

	return importance
}

// Print the dominance matrix between tasks
func printDominanceMatrix(dominance [][]bool, tasks []Task) {
	fmt.Println("Dominance Matrix (Task i dominates Task j):")
	fmt.Print("     ")
	for j := range tasks {
		fmt.Printf("T%-3d ", tasks[j].ID)
	}
	fmt.Println()

	for i, row := range dominance {
		fmt.Printf("T%-3d ", tasks[i].ID)
		for _, dominates := range row {
			if dominates {
				fmt.Print("1    ")
			} else {
				fmt.Print("0    ")
			}
		}
		fmt.Println()
	}
}

// Print the assignments in a clear format
func printAssignments(assignments []Assignment, tasks []Task, nodes []Node) {
	fmt.Println("Task Assignments:")
	fmt.Println("----------------")
	fmt.Printf("Scheduled %d out of %d tasks\n", len(assignments), len(tasks))
	fmt.Println()

	fmt.Printf("%-6s %-6s %-10s %-10s %-20s\n",
		"Task", "Node", "CPU", "GPU", "Assigned GPUs")
	fmt.Printf("%-6s %-6s %-10s %-10s %-20s\n",
		"----", "----", "---", "---", "------------")

	for _, assign := range assignments {
		var task Task
		for _, t := range tasks {
			if t.ID == assign.TaskID {
				task = t
				break
			}
		}

		fmt.Printf("%-6d %-6d %-10.2f %-10.2f %v\n",
			task.ID, assign.NodeID, task.CPUDemand, task.GPUDemand, assign.GPUIDs)
	}
}

// Calculate metrics to evaluate the quality of the solution
func calculateMetrics(assignments []Assignment, tasks []Task, nodes []Node, dominance [][]bool) {
	// Track which tasks are scheduled
	scheduled := make(map[int]bool)
	for _, assign := range assignments {
		scheduled[assign.TaskID] = true
	}

	// Count out-of-order pairs
	outOfOrderPairs := 0
	for i, taskA := range tasks {
		for j, taskB := range tasks {
			if i == j {
				continue
			}

			// If A dominates B, but B is scheduled and A is not
			if dominance[i][j] && scheduled[taskB.ID] && !scheduled[taskA.ID] {
				outOfOrderPairs++
			}
		}
	}

	// Calculate resource utilization
	totalCPUCapacity := 0.0
	totalGPUCapacity := 0.0
	usedCPU := 0.0
	usedGPU := 0.0

	for _, node := range nodes {
		totalCPUCapacity += node.CPUCapacity
		for _, gpu := range node.GPUs {
			totalGPUCapacity += gpu.Capacity
		}
	}

	for _, assign := range assignments {
		var task Task
		for _, t := range tasks {
			if t.ID == assign.TaskID {
				task = t
				break
			}
		}

		usedCPU += task.CPUDemand
		usedGPU += task.GPUDemand
	}

	// Print metrics
	fmt.Println("\nPerformance Metrics:")
	fmt.Println("-----------------")
	fmt.Printf("Task utilization:     %.1f%% (%d/%d tasks scheduled)\n",
		float64(len(assignments))/float64(len(tasks))*100, len(assignments), len(tasks))
	fmt.Printf("CPU utilization:      %.1f%% (%.2f/%.2f units)\n",
		usedCPU/totalCPUCapacity*100, usedCPU, totalCPUCapacity)
	fmt.Printf("GPU utilization:      %.1f%% (%.2f/%.2f units)\n",
		usedGPU/totalGPUCapacity*100, usedGPU, totalGPUCapacity)
	fmt.Printf("Out-of-order pairs:   %d\n", outOfOrderPairs)
}

// Generate a test case for demonstration
func generateTestCase() ([]Task, []Node) {
	tasks := []Task{
		{ID: 1, CPUDemand: 2.0, GPUDemand: 1.0, Popularity: 0.2},
		{ID: 2, CPUDemand: 3.0, GPUDemand: 0.5, Popularity: 0.1},
		{ID: 3, CPUDemand: 1.0, GPUDemand: 0.7, Popularity: 0.3},
		{ID: 4, CPUDemand: 4.0, GPUDemand: 2.0, Popularity: 0.1},
		{ID: 5, CPUDemand: 2.5, GPUDemand: 1.5, Popularity: 0.2},
		{ID: 6, CPUDemand: 1.5, GPUDemand: 0.3, Popularity: 0.1},
	}

	nodes := []Node{
		{
			ID:          1,
			CPUCapacity: 4.0,
			GPUs: []GPU{
				{ID: 1, Capacity: 1.0},
				{ID: 2, Capacity: 1.0},
			},
		},
		{
			ID:          2,
			CPUCapacity: 6.0,
			GPUs: []GPU{
				{ID: 3, Capacity: 1.0},
				{ID: 4, Capacity: 1.0},
				{ID: 5, Capacity: 1.0},
			},
		},
	}

	return tasks, nodes
}
