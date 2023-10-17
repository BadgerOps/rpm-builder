package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Initialize a folder with the base RPM layout.
func initializeRPMFolder(path string) {
	// Create directories expected by rpmbuild.
	for _, dir := range []string{"BUILD", "RPMS", "SOURCES", "SPECS", "SRPMS"} {
		os.MkdirAll(filepath.Join(path, dir), 0755)
	}

	// Write a basic spec file.
	specFilePath := filepath.Join(path, "SPECS", "package.spec")
	specContent := `
Name:           sample-package
Version:        1.0
Release:        1%{?dist}
Summary:        A sample RPM package

License:        MIT
URL:            http://example.com
Source0:        %{name}-%{version}.tar.gz

%description
This is a sample RPM package. You should update this spec file!.

%prep
%setup -q

%build

%install

%files
`

	err := os.WriteFile(specFilePath, []byte(specContent), 0644)
	if err != nil {
		fmt.Println("Error writing the spec file:", err)
		os.Exit(1)
	}
	// Write a README.md file.
	readmeFilePath := filepath.Join(path, "README.md")
	readmeContent := `
	# RPM Package Builder

	This folder structure is designed to help you create RPM packages using ` + "`rpmbuild`" + `.

	## Steps to Build an RPM:

	1. Update the spec file located in the ` + "`SPECS`" + ` directory. A basic template has been provided for you, but you'll need to modify it to fit your package's requirements.
	2. Place any necessary sources in the ` + "`SOURCES`" + ` directory.
	3. Use the ` + "`-b`" + ` flag with this tool to build the RPM from this directory.

	For more detailed information on how to build RPM packages, refer to the [RPM Packaging Guide](https://rpm-packaging-guide.github.io/).
	`

	err = os.WriteFile(readmeFilePath, []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error writing the README file:", err)
		os.Exit(1)
	}
	fmt.Println("Initialized RPM folder at:", path)
	fmt.Println("A basic spec file has been created at", specFilePath)
	fmt.Println("Please refer to README.md for more information.")}


// Build an RPM from the given path.
func buildRPM(path string) {
	// TODO: do some validation here before building, this is bare bones "build an RPM"

	cmd := exec.Command("rpmbuild", "-ba", "--define", "_topdir "+path, filepath.Join(path, "SPECS", "package.spec"))

	// Capture the output.
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error building RPM: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Output of rpmbuild:\n%s\n", output)
}

func main() {
	// Command line flags.
	initializePtr := flag.String("i", "", "initialize a folder with the base RPM layout")
	buildPtr := flag.String("b", "", "build an RPM file from the given path")

	flag.Parse()

	// Check if both flags are not provided.
	if *initializePtr == "" && *buildPtr == "" {
		fmt.Println("Error: Please provide one of the flags -i [initialize] ./directory or -b [build] ./directory.")
		os.Exit(1)
	}

	// If initialize flag is provided.
	if *initializePtr != "" {
		initializeRPMFolder(*initializePtr)
		return
	}

	// If build flag is provided.
	if *buildPtr != "" {
		buildRPM(*buildPtr)
	}
}
