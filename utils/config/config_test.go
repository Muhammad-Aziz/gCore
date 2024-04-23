package config

import (
	"context"
	"fmt"
	"testing"
)

/*
// MockFileHandler simulates the FileHandler interface for testing.
type MockFileHandler struct {
	Exists map[string]bool
}

// Mock FilePathExists
func (m MockFileHandler) FilePathExists(filePath string) bool {
	fmt.Println(filePath)
	fmt.Println(m.Exists[filePath])
	exists, found := m.Exists[filePath]
	return found && exists
}
*/
// setupInitConfigContext creates a context with a mock FileHandler injected.
func setupInitConfigContext() context.Context {
	ctx := context.Background()
	//mockHandler := MockFileHandler{Exists: map[string]bool{
	//	"exists.txt": true, // Simulates "exists.txt" as an existing file
	//}}
	//return file_handler.AddHandlerToContext(ctx, mockHandler)
	return ctx
}

// TestInitConfigReturn tests the initialization of ConfigPackage.
func TestInitConfigReturn(t *testing.T) {
	configPkg := InitConfig(&struct{}{}, "valid/path")
	if configPkg == nil {
		t.Fatal("Expected non-nil ConfigPackage")
	}
}

// TestInitConfigAppConfigPointer tests the AppConfigPointer initialization.
func TestInitConfigAppConfigPointer(t *testing.T) {
	appConfigPointer := &struct{}{}
	configPkg := InitConfig(appConfigPointer, "valid/path")

	if got := configPkg.GetAppConfigPointer(); got != appConfigPointer {
		t.Error("AppConfigPointer not set correctly")
	}
}

// TestSetAppConfigPointer tests setting the app config pointer.
func TestSetAppConfigPointer(t *testing.T) {
	tests := []struct {
		name    string
		ptr     interface{}
		wantErr bool
	}{
		{"ValidStructPtr", &struct{}{}, false},
		{"NilPointer", nil, true},
		{"NonStructPtr", new(int), true},
		{"NonPtr", struct{}{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPkg := &ConfigPackage{}
			err := configPkg.SetAppConfigObjectPointer(tt.ptr)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetAppConfigPointer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestInitConfigAppConfigFiles tests the AppConfigFiles initialization.
func TestInitConfigAppConfigFiles(t *testing.T) {
	configPkg := InitConfig(&struct{}{}, "valid/path")

	if configPkg.GetAppConfigFile() != "valid/path" {
		t.Error("AppConfigFile not set correctly")
	}
}

// TestSetAppConfigFiles tests setting app config files
func TestSetAppConfigFiles(t *testing.T) {

	tests := []struct {
		name          string
		appConfigFile string
		wantErr       bool
	}{
		{"NoFile", "", true},
		{"FileExist", "exists.txt", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ConfigPackage{}
			err := cp.SetAppConfigFile(tt.appConfigFile)
			fmt.Println(err != nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s: SetAppConfigFile error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
