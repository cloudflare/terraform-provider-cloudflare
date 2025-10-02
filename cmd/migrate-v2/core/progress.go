package core

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// ProgressTracker tracks the progress of migrations
type ProgressTracker struct {
	mu            sync.Mutex
	totalSteps    int
	currentStep   int
	currentTask   string
	startTime     time.Time
	stepStartTime time.Time
	output        io.Writer
	verbose       bool
}

// NewProgressTracker creates a new progress tracker
func NewProgressTracker() *ProgressTracker {
	return &ProgressTracker{
		output:    os.Stdout,
		startTime: time.Now(),
	}
}

// SetOutput sets the output writer
func (pt *ProgressTracker) SetOutput(w io.Writer) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.output = w
}

// SetVerbose enables verbose output
func (pt *ProgressTracker) SetVerbose(verbose bool) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	pt.verbose = verbose
}

// Start initializes tracking with the total number of steps
func (pt *ProgressTracker) Start(totalSteps int, task string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.totalSteps = totalSteps
	pt.currentStep = 0
	pt.currentTask = task
	pt.startTime = time.Now()
	pt.stepStartTime = time.Now()
	
	pt.print(fmt.Sprintf("Starting %s (%d steps)", task, totalSteps))
}

// StartStep begins tracking a new step
func (pt *ProgressTracker) StartStep(stepName string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.currentStep++
	pt.currentTask = stepName
	pt.stepStartTime = time.Now()
	
	progress := pt.getProgressBar()
	pt.print(fmt.Sprintf("%s %s", progress, stepName))
}

// CompleteStep marks the current step as complete
func (pt *ProgressTracker) CompleteStep(message string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	duration := time.Since(pt.stepStartTime)
	progress := pt.getProgressBar()
	
	if message != "" {
		pt.print(fmt.Sprintf("%s ✓ %s (%s)", progress, message, duration.Round(time.Millisecond)))
	} else {
		pt.print(fmt.Sprintf("%s ✓ Complete (%s)", progress, duration.Round(time.Millisecond)))
	}
}

// FailStep marks the current step as failed
func (pt *ProgressTracker) FailStep(err error) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	progress := pt.getProgressBar()
	errMsg := "unknown error"
	if err != nil {
		errMsg = err.Error()
	}
	pt.print(fmt.Sprintf("%s ✗ Failed: %s", progress, errMsg))
}

// Complete marks the entire migration as complete
func (pt *ProgressTracker) Complete() {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	duration := time.Since(pt.startTime)
	pt.print(fmt.Sprintf("\n✓ Migration completed in %s", duration.Round(time.Millisecond)))
}

// Info logs an informational message
func (pt *ProgressTracker) Info(message string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	if pt.verbose {
		pt.print(fmt.Sprintf("  → %s", message))
	}
}

// Warning logs a warning message
func (pt *ProgressTracker) Warning(message string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.print(fmt.Sprintf("  ⚠ %s", message))
}

// Error logs an error message
func (pt *ProgressTracker) Error(message string) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.print(fmt.Sprintf("  ✗ %s", message))
}

// getProgressBar returns a text progress bar
func (pt *ProgressTracker) getProgressBar() string {
	if pt.totalSteps <= 0 {
		return "[?/?]"
	}
	
	percent := float64(pt.currentStep) / float64(pt.totalSteps) * 100
	barWidth := 20
	filled := int(float64(barWidth) * float64(pt.currentStep) / float64(pt.totalSteps))
	
	bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
	return fmt.Sprintf("[%d/%d] [%s] %.0f%%", pt.currentStep, pt.totalSteps, bar, percent)
}

// print writes output
func (pt *ProgressTracker) print(message string) {
	fmt.Fprintln(pt.output, message)
}

// ProgressReporter provides a simpler interface for progress reporting
type ProgressReporter struct {
	tracker *ProgressTracker
	steps   []string
	current int
}

// NewProgressReporter creates a progress reporter with predefined steps
func NewProgressReporter(steps []string) *ProgressReporter {
	tracker := NewProgressTracker()
	tracker.Start(len(steps), "Migration")
	
	return &ProgressReporter{
		tracker: tracker,
		steps:   steps,
		current: 0,
	}
}

// NextStep advances to the next step
func (pr *ProgressReporter) NextStep() {
	if pr.current < len(pr.steps) {
		pr.tracker.StartStep(pr.steps[pr.current])
		pr.current++
	}
}

// CompleteCurrentStep marks the current step as complete
func (pr *ProgressReporter) CompleteCurrentStep() {
	if pr.current > 0 && pr.current <= len(pr.steps) {
		pr.tracker.CompleteStep(pr.steps[pr.current-1])
	}
}

// Complete marks all steps as complete
func (pr *ProgressReporter) Complete() {
	pr.tracker.Complete()
}

// Info logs an informational message
func (pr *ProgressReporter) Info(message string) {
	pr.tracker.Info(message)
}

// Warning logs a warning
func (pr *ProgressReporter) Warning(message string) {
	pr.tracker.Warning(message)
}

// Error logs an error
func (pr *ProgressReporter) Error(message string) {
	pr.tracker.Error(message)
}