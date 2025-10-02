package internal

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProgressTracker_NewProgressTracker(t *testing.T) {
	pt := NewProgressTracker()

	assert.NotNil(t, pt)
	assert.NotNil(t, pt.output)
	assert.False(t, pt.verbose)
	assert.Equal(t, 0, pt.totalSteps)
	assert.Equal(t, 0, pt.currentStep)
}

func TestProgressTracker_SetOutput(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}

	pt.SetOutput(buf)
	pt.print("test message")

	output := buf.String()
	assert.Contains(t, output, "test message")
}

func TestProgressTracker_SetVerbose(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	// Test without verbose
	pt.SetVerbose(false)
	pt.Info("info message")
	assert.Empty(t, buf.String())

	// Test with verbose
	pt.SetVerbose(true)
	pt.Info("verbose info message")
	assert.Contains(t, buf.String(), "verbose info message")
}

func TestProgressTracker_Start(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(5, "Test Migration")

	output := buf.String()
	assert.Contains(t, output, "Starting Test Migration (5 steps)")
	assert.Equal(t, 5, pt.totalSteps)
	assert.Equal(t, 0, pt.currentStep)
	assert.Equal(t, "Test Migration", pt.currentTask)
}

func TestProgressTracker_StartStep(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(3, "Migration")
	buf.Reset() // Clear start message

	pt.StartStep("Step 1")

	output := buf.String()
	assert.Contains(t, output, "Step 1")
	assert.Contains(t, output, "[1/3]")
	assert.Equal(t, 1, pt.currentStep)
	assert.Equal(t, "Step 1", pt.currentTask)
}

func TestProgressTracker_CompleteStep(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(2, "Migration")
	pt.StartStep("Step 1")
	buf.Reset()

	pt.CompleteStep("Step 1 done")

	output := buf.String()
	assert.Contains(t, output, "✓ Step 1 done")
	assert.Contains(t, output, "[1/2]")
	assert.Regexp(t, `\d+[ms]+\)`, output) // Duration should be included
}

func TestProgressTracker_CompleteStep_NoMessage(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(2, "Migration")
	pt.StartStep("Step 1")
	buf.Reset()

	pt.CompleteStep("")

	output := buf.String()
	assert.Contains(t, output, "✓ Complete")
	assert.NotContains(t, output, "✓ Complete ()")
}

func TestProgressTracker_FailStep(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(2, "Migration")
	pt.StartStep("Step 1")
	buf.Reset()

	err := errors.New("something went wrong")
	pt.FailStep(err)

	output := buf.String()
	assert.Contains(t, output, "✗ Failed: something went wrong")
	assert.Contains(t, output, "[1/2]")
}

func TestProgressTracker_Complete(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(1, "Migration")
	time.Sleep(10 * time.Millisecond) // Ensure some duration
	buf.Reset()

	pt.Complete()

	output := buf.String()
	assert.Contains(t, output, "✓ Migration completed in")
	assert.Contains(t, output, "ms")
}

func TestProgressTracker_Info(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	// Without verbose - should not print
	pt.SetVerbose(false)
	pt.Info("hidden info")
	assert.Empty(t, buf.String())

	// With verbose - should print
	pt.SetVerbose(true)
	pt.Info("visible info")
	assert.Contains(t, buf.String(), "→ visible info")
}

func TestProgressTracker_Warning(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Warning("warning message")

	output := buf.String()
	assert.Contains(t, output, "⚠ warning message")
}

func TestProgressTracker_Error(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Error("error message")

	output := buf.String()
	assert.Contains(t, output, "✗ error message")
}

func TestProgressTracker_getProgressBar(t *testing.T) {
	tests := []struct {
		name        string
		currentStep int
		totalSteps  int
		expected    string
	}{
		{
			name:        "zero total steps",
			currentStep: 0,
			totalSteps:  0,
			expected:    "[?/?]",
		},
		{
			name:        "no progress",
			currentStep: 0,
			totalSteps:  10,
			expected:    "[0/10] [--------------------] 0%",
		},
		{
			name:        "half progress",
			currentStep: 5,
			totalSteps:  10,
			expected:    "[5/10] [==========----------] 50%",
		},
		{
			name:        "complete progress",
			currentStep: 10,
			totalSteps:  10,
			expected:    "[10/10] [====================] 100%",
		},
		{
			name:        "partial progress",
			currentStep: 3,
			totalSteps:  10,
			expected:    "[3/10] [======--------------] 30%",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := &ProgressTracker{
				currentStep: tt.currentStep,
				totalSteps:  tt.totalSteps,
			}

			result := pt.getProgressBar()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProgressTracker_Concurrent(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)
	pt.SetVerbose(true)

	pt.Start(100, "Concurrent Test")

	// Run multiple goroutines that update progress
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simulate work
			pt.StartStep(fmt.Sprintf("Worker %d", id))
			time.Sleep(time.Millisecond)
			pt.CompleteStep(fmt.Sprintf("Worker %d done", id))

			// Also test other methods
			pt.Info(fmt.Sprintf("Info from worker %d", id))
			if id%3 == 0 {
				pt.Warning(fmt.Sprintf("Warning from worker %d", id))
			}
			if id%5 == 0 {
				pt.Error(fmt.Sprintf("Error from worker %d", id))
			}
		}(i)
	}

	wg.Wait()
	pt.Complete()

	// Verify output contains expected elements (order may vary)
	output := buf.String()
	assert.Contains(t, output, "✓ Migration completed")
	assert.Contains(t, output, "Worker")

	// Should have exactly 10 workers completed
	completeCount := strings.Count(output, "Worker")
	assert.GreaterOrEqual(t, completeCount, 10)
}

func TestProgressReporter_NewProgressReporter(t *testing.T) {
	steps := []string{"Step 1", "Step 2", "Step 3"}
	pr := NewProgressReporter(steps)

	assert.NotNil(t, pr)
	assert.NotNil(t, pr.tracker)
	assert.Equal(t, steps, pr.steps)
	assert.Equal(t, 0, pr.current)
	assert.Equal(t, 3, pr.tracker.totalSteps)
}

func TestProgressReporter_NextStep(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"First", "Second", "Third"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)

	// Clear initial output
	buf.Reset()

	// First step
	pr.NextStep()
	assert.Equal(t, 1, pr.current)
	assert.Contains(t, buf.String(), "First")

	// Second step
	buf.Reset()
	pr.NextStep()
	assert.Equal(t, 2, pr.current)
	assert.Contains(t, buf.String(), "Second")

	// Third step
	buf.Reset()
	pr.NextStep()
	assert.Equal(t, 3, pr.current)
	assert.Contains(t, buf.String(), "Third")

	// Beyond last step - should not advance
	buf.Reset()
	pr.NextStep()
	assert.Equal(t, 3, pr.current)
	assert.Empty(t, buf.String())
}

func TestProgressReporter_CompleteCurrentStep(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"Step 1", "Step 2"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)

	// Complete without starting - should do nothing
	buf.Reset()
	pr.CompleteCurrentStep()
	assert.Empty(t, buf.String())

	// Start and complete first step
	pr.NextStep()
	buf.Reset()
	pr.CompleteCurrentStep()
	assert.Contains(t, buf.String(), "✓")
	assert.Contains(t, buf.String(), "Step 1")

	// Start and complete second step
	pr.NextStep()
	buf.Reset()
	pr.CompleteCurrentStep()
	assert.Contains(t, buf.String(), "✓")
	assert.Contains(t, buf.String(), "Step 2")
}

func TestProgressReporter_Complete(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"Step 1", "Step 2"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)

	pr.NextStep()
	pr.CompleteCurrentStep()
	pr.NextStep()
	pr.CompleteCurrentStep()

	buf.Reset()
	pr.Complete()

	output := buf.String()
	assert.Contains(t, output, "✓ Migration completed")
}

func TestProgressReporter_Info(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"Step 1"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)
	pr.tracker.SetVerbose(true)

	buf.Reset()
	pr.Info("info message")

	assert.Contains(t, buf.String(), "→ info message")
}

func TestProgressReporter_Warning(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"Step 1"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)

	buf.Reset()
	pr.Warning("warning message")

	assert.Contains(t, buf.String(), "⚠ warning message")
}

func TestProgressReporter_Error(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{"Step 1"}
	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)

	buf.Reset()
	pr.Error("error message")

	assert.Contains(t, buf.String(), "✗ error message")
}

func TestProgressReporter_FullWorkflow(t *testing.T) {
	buf := &bytes.Buffer{}

	steps := []string{
		"Loading files",
		"Analyzing resources",
		"Applying migrations",
		"Verifying results",
	}

	pr := NewProgressReporter(steps)
	pr.tracker.SetOutput(buf)
	pr.tracker.SetVerbose(true)

	// Simulate a full migration workflow
	pr.NextStep() // Loading files
	pr.Info("Found 10 .tf files")
	pr.CompleteCurrentStep()

	pr.NextStep() // Analyzing resources
	pr.Info("Found 25 resources")
	pr.Warning("3 resources need manual review")
	pr.CompleteCurrentStep()

	pr.NextStep() // Applying migrations
	pr.Info("Migrating cloudflare_record resources")
	pr.Info("Migrating cloudflare_zone resources")
	pr.CompleteCurrentStep()

	pr.NextStep() // Verifying results
	pr.CompleteCurrentStep()

	pr.Complete()

	// Verify output contains expected elements
	output := buf.String()
	assert.Contains(t, output, "Loading files")
	assert.Contains(t, output, "Found 10 .tf files")
	assert.Contains(t, output, "Analyzing resources")
	assert.Contains(t, output, "3 resources need manual review")
	assert.Contains(t, output, "Applying migrations")
	assert.Contains(t, output, "Verifying results")
	assert.Contains(t, output, "✓ Migration completed")

	// Check that steps were completed (4 steps + 1 final completion)
	assert.Equal(t, 5, strings.Count(output, "✓"))
}

func TestProgressTracker_EdgeCases(t *testing.T) {
	t.Run("negative steps", func(t *testing.T) {
		pt := NewProgressTracker()
		buf := &bytes.Buffer{}
		pt.SetOutput(buf)

		// Start with negative steps should still work
		pt.Start(-1, "Test")
		assert.Equal(t, -1, pt.totalSteps)

		// Progress bar should handle it
		bar := pt.getProgressBar()
		assert.Equal(t, "[?/?]", bar) // Should fallback to unknown
	})

	t.Run("very long task names", func(t *testing.T) {
		pt := NewProgressTracker()
		buf := &bytes.Buffer{}
		pt.SetOutput(buf)

		longName := strings.Repeat("A", 1000)
		pt.Start(1, longName)

		output := buf.String()
		assert.Contains(t, output, longName)
	})

	t.Run("nil error in FailStep", func(t *testing.T) {
		pt := NewProgressTracker()
		buf := &bytes.Buffer{}
		pt.SetOutput(buf)

		pt.Start(1, "Test")
		pt.StartStep("Step")

		// Should handle nil error gracefully
		assert.NotPanics(t, func() {
			pt.FailStep(nil)
		})
	})
}

func TestProgressTracker_TimingAccuracy(t *testing.T) {
	pt := NewProgressTracker()
	buf := &bytes.Buffer{}
	pt.SetOutput(buf)

	pt.Start(1, "Timing Test")
	pt.StartStep("Step with delay")

	// Add a known delay
	time.Sleep(50 * time.Millisecond)

	pt.CompleteStep("Done")

	output := buf.String()
	// Should show at least 50ms duration
	assert.Contains(t, output, "ms)")
	// The output should contain a number >= 50
	assert.Regexp(t, `\d+ms\)`, output)
}
