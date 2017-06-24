package goosrelease

import (
	"fmt"
	"strconv"
	"testing"
)

func TestReadNoFile(t *testing.T) {
	_, err := ReadFile("test/nosuchfile")
	if err == nil {
		t.Errorf("Read() on non-existant file returned nil, should not be nil")
	}
}

func TestReadFile(t *testing.T) {
	expectedResults := map[int]map[string]string{
		1: {
			"NAME":        "void",
			"ID":          "void",
			"DISTRIB_ID":  "void",
			"PRETTY_NAME": "void",
		},
		2: {},
		3: {
			"NAME":           "CentOS Linux",
			"VERSION":        "7 (Core)",
			"ID":             "centos",
			"ID_LIKE":        "rhel fedora",
			"VERSION_ID":     "7",
			"PRETTY_NAME":    "CentOS Linux 7 (Core)",
			"ANSI_COLOR":     "0;31",
			"CPE_NAME":       "cpe:/o:centos:centos:7",
			"HOME_URL":       "https://www.centos.org/",
			"BUG_REPORT_URL": "https://bugs.centos.org/",
		},
		4: {
			"NAME":           "CoreOS",
			"ID":             "coreos",
			"VERSION":        "1185.3.0",
			"VERSION_ID":     "1185.3.0",
			"BUILD_ID":       "2016-11-01-0605",
			"PRETTY_NAME":    "CoreOS 1185.3.0 (MoreOS)",
			"ANSI_COLOR":     "1;32",
			"HOME_URL":       "https://coreos.com/",
			"BUG_REPORT_URL": "https://github.com/coreos/bugs/issues",
		},
		5: {
			"NAME":           "Container Linux by CoreOS",
			"ID":             "coreos",
			"VERSION":        "1235.6.0",
			"VERSION_ID":     "1235.6.0",
			"BUILD_ID":       "2017-01-10-0545",
			"PRETTY_NAME":    "Container Linux by CoreOS 1235.6.0 (Ladybug)",
			"ANSI_COLOR":     "38;5;75",
			"HOME_URL":       "https://coreos.com/",
			"BUG_REPORT_URL": "https://github.com/coreos/bugs/issues",
		},
		6: {
			"NAME":        "dummy",
			"DISTRIB_ID":  "\"foobar",
			"VERSION":     "17\"",
			"ID":          ",;: ",
			"PRETTY_NAME": "$ ` \\ \"",
			"ANSI_COLOR":  "",
		},
	}

	for test := 1; test <= len(expectedResults); test++ {
		filename := "test/os-release." + strconv.Itoa(test)
		osrelease, err := ReadFile(filename)
		if err != nil {
			t.Fatalf("Error reading test file '%v': %v", filename, err)
		} else {
			for key, value := range expectedResults[test] {
				if osrelease[key] != value {
					t.Errorf("In file 'test/os-release.%v', Read() returned '%v' = '%v', should be '%v'", test, key, osrelease[key], value)
				}
			}
		}
	}
}

func ExampleRead() {
	osrelease, err := Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("PRETTY_NAME = %v\n", osrelease["PRETTY_NAME"])
}

func ExampleReadFile() {
	osrelease, err := ReadFile("/etc/os-release")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("PRETTY_NAME = %v\n", osrelease["PRETTY_NAME"])
}
