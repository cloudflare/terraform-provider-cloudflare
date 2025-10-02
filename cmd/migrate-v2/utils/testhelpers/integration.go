package testhelpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RunIntegrationTests provides a dead-simple interface for migration testing.
// Just provide source and target directories with matching files, and this handles everything else.
//
// Usage:
//   func TestMyResourceMigration(t *testing.T) {
//       testhelpers.RunIntegrationTests(t, testhelpers.IntegrationTestOptions{
//           FromVersion: "v4",
//           ToVersion:   "v5",
//       })
//   }
func RunIntegrationTests(t *testing.T, opts ...IntegrationTestOptions) {
	// Use defaults if no options provided
	options := IntegrationTestOptions{
		TestDataDir: "testdata",
		FromVersion: "v4",
		ToVersion:   "v5",
	}
	if len(opts) > 0 {
		// Merge provided options with defaults
		provided := opts[0]
		if provided.TestDataDir != "" {
			options.TestDataDir = provided.TestDataDir
		}
		if provided.FromVersion != "" {
			options.FromVersion = provided.FromVersion
		}
		if provided.ToVersion != "" {
			options.ToVersion = provided.ToVersion
		}
		if provided.MigrateBinary != "" {
			options.MigrateBinary = provided.MigrateBinary
		}
		options.SkipStateTests = provided.SkipStateTests
	}
	
	test := &integrationTest{
		options: options,
	}
	test.run(t)
}

// IntegrationTestOptions configures the integration test runner
type IntegrationTestOptions struct {
	// FromVersion is the source version (e.g., "v4")
	// Defaults to "v4"
	FromVersion string
	
	// ToVersion is the target version (e.g., "v5")
	// Defaults to "v5"
	ToVersion string
	
	// TestDataDir is the path to the directory containing version subdirectories
	// Defaults to "testdata"
	TestDataDir string
	
	// MigrateBinary is the path to a pre-built migrate-v2 binary (optional)
	// If not provided, the binary will be built automatically
	MigrateBinary string
	
	// SkipStateTests skips state file migration tests
	SkipStateTests bool
}

// integrationTest handles the test execution
type integrationTest struct {
	options       IntegrationTestOptions
	migrateBinary string
}

func (it *integrationTest) run(t *testing.T) {
	// Build or locate the migrate binary
	if it.options.MigrateBinary != "" {
		it.migrateBinary = it.options.MigrateBinary
	} else {
		it.migrateBinary = it.buildMigrateBinary(t)
	}
	
	// Discover and run all test cases
	testCases := it.discoverTestCases(t)
	
	if len(testCases) == 0 {
		t.Fatal("No test cases found. Ensure testdata/v4 and testdata/v5 contain matching .tf files")
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it.runTestCase(t, tc)
		})
	}
}

type testCase struct {
	name      string
	v4Config  string
	v4State   string
	v5Config  string
	v5State   string
}

func (it *integrationTest) discoverTestCases(t *testing.T) []testCase {
	fromDir := filepath.Join(it.options.TestDataDir, it.options.FromVersion)
	toDir := filepath.Join(it.options.TestDataDir, it.options.ToVersion)
	
	// Verify directories exist
	require.DirExists(t, fromDir, "%s test data directory not found", it.options.FromVersion)
	require.DirExists(t, toDir, "%s test data directory not found", it.options.ToVersion)
	
	var testCases []testCase
	
	// Find all .tf files in source directory
	sourceFiles, err := filepath.Glob(filepath.Join(fromDir, "*.tf"))
	require.NoError(t, err)
	
	for _, sourceFile := range sourceFiles {
		baseName := filepath.Base(sourceFile)
		testName := strings.TrimSuffix(baseName, ".tf")
		
		// Check for corresponding target file
		targetFile := filepath.Join(toDir, baseName)
		if _, err := os.Stat(targetFile); os.IsNotExist(err) {
			t.Logf("Warning: No %s file for %s, skipping", it.options.ToVersion, baseName)
			continue
		}
		
		tc := testCase{
			name:     testName,
			v4Config: sourceFile,
			v5Config: targetFile,
		}
		
		// Check for state files (optional)
		if !it.options.SkipStateTests {
			sourceState := filepath.Join(fromDir, testName+".tfstate")
			targetState := filepath.Join(toDir, testName+".tfstate")
			if _, err := os.Stat(sourceState); err == nil {
				tc.v4State = sourceState
				if _, err := os.Stat(targetState); err == nil {
					tc.v5State = targetState
				}
			}
		}
		
		testCases = append(testCases, tc)
	}
	
	return testCases
}

func (it *integrationTest) runTestCase(t *testing.T, tc testCase) {
	// Create temp directory and copy files
	tempDir := t.TempDir()
	
	tempConfig := filepath.Join(tempDir, filepath.Base(tc.v4Config))
	require.NoError(t, copyFile(tc.v4Config, tempConfig))
	
	var tempState string
	if tc.v4State != "" {
		tempState = filepath.Join(tempDir, filepath.Base(tc.v4State))
		require.NoError(t, copyFile(tc.v4State, tempState))
	}
	
	// Run migration
	args := []string{
		"-config", tempDir,
		"-from", it.options.FromVersion,
		"-to", it.options.ToVersion,
		"-backup=false",
	}
	if tempState != "" {
		args = append(args, "-state", tempState)
	}
	
	cmd := exec.Command(it.migrateBinary, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Migration failed: %v\nOutput: %s", err, output)
	}
	
	// Compare configs
	expected, err := os.ReadFile(tc.v5Config)
	require.NoError(t, err)
	
	actual, err := os.ReadFile(tempConfig)
	require.NoError(t, err)
	
	expectedNorm := NormalizeHCL(string(expected))
	actualNorm := NormalizeHCL(string(actual))
	
	assert.Equal(t, expectedNorm, actualNorm, 
		"Config mismatch for %s", tc.name)
	
	// Compare states if applicable
	if tc.v5State != "" && tempState != "" {
		it.compareStates(t, tc.v5State, tempState)
	}
}

func (it *integrationTest) compareStates(t *testing.T, expectedPath, actualPath string) {
	var expected, actual map[string]interface{}
	
	expectedBytes, err := os.ReadFile(expectedPath)
	require.NoError(t, err)
	err = json.Unmarshal(expectedBytes, &expected)
	require.NoError(t, err)
	
	actualBytes, err := os.ReadFile(actualPath)
	require.NoError(t, err)
	err = json.Unmarshal(actualBytes, &actual)
	require.NoError(t, err)
	
	compareJSON(t, expected, actual, "")
}

func compareJSON(t *testing.T, expected, actual interface{}, path string) {
	switch e := expected.(type) {
	case map[string]interface{}:
		a, ok := actual.(map[string]interface{})
		require.True(t, ok, "Type mismatch at %s", path)
		
		for key, eval := range e {
			aval, exists := a[key]
			require.True(t, exists, "Missing key at %s.%s", path, key)
			compareJSON(t, eval, aval, path+"."+key)
		}
		
		for key := range a {
			_, exists := e[key]
			require.True(t, exists, "Unexpected key at %s.%s", path, key)
		}
		
	case []interface{}:
		a, ok := actual.([]interface{})
		require.True(t, ok, "Type mismatch at %s", path)
		require.Equal(t, len(e), len(a), "Length mismatch at %s", path)
		
		for i := range e {
			compareJSON(t, e[i], a[i], fmt.Sprintf("%s[%d]", path, i))
		}
		
	default:
		assert.Equal(t, expected, actual, "Value mismatch at %s", path)
	}
}

func (it *integrationTest) buildMigrateBinary(t *testing.T) string {
	tempBinary := filepath.Join(t.TempDir(), "migrate-v2")
	
	// Find migrate-v2 main.go
	mainPath := findMigrateMain(t)
	
	cmd := exec.Command("go", "build", "-o", tempBinary, ".")
	cmd.Dir = filepath.Dir(mainPath)
	
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build: %v\nStderr: %s", err, stderr.String())
	}
	
	return tempBinary
}

func findMigrateMain(t *testing.T) string {
	// Start from current directory and search upward
	dir, _ := os.Getwd()
	
	for i := 0; i < 10; i++ {
		// Check for cmd/migrate-v2/main.go
		mainPath := filepath.Join(dir, "cmd", "migrate-v2", "main.go")
		if _, err := os.Stat(mainPath); err == nil {
			return mainPath
		}
		
		// Check if we're already in migrate-v2
		mainPath = filepath.Join(dir, "main.go")
		if _, err := os.Stat(mainPath); err == nil {
			content, _ := os.ReadFile(mainPath)
			if strings.Contains(string(content), "migrate-v2") {
				return mainPath
			}
		}
		
		// Move up
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	
	t.Fatal("Cannot find migrate-v2 main.go")
	return ""
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	
	dest, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dest.Close()
	
	_, err = io.Copy(dest, source)
	return err
}

