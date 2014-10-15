package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// PkgLine is a structure for holding the parsed contents of a single line  of 
// output from the reprepro tool
type PkgLine struct {
	Distro  string `json:"distro"`
	Version string `json:"version"`
}

// DistroLine is a structure for holding the parsed contents of a single line  of 
// output from the reprepro tool
type DistroLine struct {
	Package string `json:"package"`
	Version string `json:"version"`
}

// PackageMessage is a structure for holding the main JSON contents of the
// marshalled output of package info
type PackageMessage struct {
	Package string      `json:"package"`
	Info    interface{} `json:"info"`
}

// DistroMessage is a structure for holding the main JSON contents of the
// marshalled output of distro info
type DistroMessage struct {
	Distro string      `json:"distro"`
	Info   interface{} `json:"info"`
}

// shell accepts a single string parameter used to execute a shell command.
// It returns a byte array of the command output
func shell(cmd string) []byte {
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Println(err)
	}
	return out
}

// marshallPackage returns the output of the command "reprepro ls pkg" as
// marshalled JSON
func marshallPackage(pkg string) []byte {
	out := shell(setCommand("ls") + pkg)
	buf := bytes.NewBuffer(out)
	lines := make([]PkgLine, 1, 24)
	for i := 0; i < cap(lines); i++ {
		l, err := buf.ReadString(0x0A)
		if err == io.EOF {
			break
		}
		if i+1 > len(lines) {
			lines = lines[0 : i+1]
		}
		if err != nil {
			log.Fatal("buf.ReadString: ", err)
		}
		j := strings.Fields(l)
		lines[i] = PkgLine{j[4], j[2]}
	}
	return doMarshall(PackageMessage{pkg, lines})
}

// marshallDistro returns the output of the command "reprepro lsit distro" as
// marshalled JSON
func marshallDistro(dist string) []byte {
	out := shell(setCommand("list") + dist)
	buf := bytes.NewBuffer(out)
	lines := make([]DistroLine, 1, 100)
	for i := 0; i < cap(lines); i++ {
		l, err := buf.ReadString(0x0A)
		if i+1 > len(lines) {
			lines = lines[0 : i+1]
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("buf.ReadString: ", err)
		}
		j := strings.Fields(l)
		lines[i] = DistroLine{j[1], j[2]}
	}
	return doMarshall(DistroMessage{dist, lines})
}

// marshalCreate returns the output of the attempt to create a new distro as
// marshalled JSON
func marshalCreate(dist string) []byte {
	file := base + "/conf/distributions"

	f, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		log.Println("os.OpenFile: ", err)
		return doMarshall(DistroMessage{dist, "os.OpenFile error"})
	}

	defer f.Close()

	text := "\n" +
		"Origin: Skybox Imaging\n" +
		"Label: Skybox\n" +
		"Suite: " + dist + "\n" +
		"Codename: " + dist + "\n" +
		"Version: 0\n" +
		"Architectures: amd64\n" +
		"Components: main\n" +
		"Description: " + dist + "\n" +
		"Log: " + dist + "\n" +
		"Update: " + dist + "\n" +
		"SignWith: 7C752D58\n"

	if _, err = f.WriteString(text); err != nil {
		log.Println("f.WriteString: ", err)
		return doMarshall(DistroMessage{dist, "f.WriteString error"})
	}

	return doMarshall(DistroMessage{dist, "Created"})
}
