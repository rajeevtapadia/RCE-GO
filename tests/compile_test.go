package docker

import (
	"rce-go/docker"
	"rce-go/utils"
	"regexp"
	"strconv"
	"testing"
)

func TestCompileJS(t *testing.T) {
	var data utils.PayLoad
	data.Language = "javascript"
	data.Code = strconv.Quote("console.log(\"js\")")
	got := string(docker.Run(&data))
	got = removeNonPrintableChars(got)
	want := "js"

	if got != want {
		t.Errorf("js test failed\n wanted: %s\n    got: %s", want, got)
	}
}

func TestCompilePython(t *testing.T) {
	var data utils.PayLoad
	data.Language = "python"
	data.Code = strconv.Quote(`print('python test is passing')`)

	got := string(docker.Run(&data))
	got = removeNonPrintableChars(got)
	want := "python test is passing"

	if got != want {
		t.Errorf("python test failed\n wanted: %s\n got: %s", want, got)
	}
}

func TestCompileCpp(t *testing.T) {
	var data utils.PayLoad
	data.Language = "cpp"
	data.Code = strconv.Quote(`#include <iostream>
int main() {
    std::cout << "cpp test is passing" << std::endl;
    return 0;
}`)

	got := string(docker.Run(&data))
	got = removeNonPrintableChars(got)
	want := "cpp test is passing"

	if got != want {
		t.Errorf("cpp test failed\n wanted: %s\n got: %s", want, got)
	}
}

func TestCompileC(t *testing.T) {
	var data utils.PayLoad
	data.Language = "c"
	data.Code = strconv.Quote(`#include <stdio.h>
int main() {
    printf("c test is passing");
    return 0;
}`)

	got := string(docker.Run(&data))
	got = removeNonPrintableChars(got)
	want := "c test is passing"

	if got != want {
		t.Errorf("c test failed\n wanted: %s\n got: %s", want, got)
	}
}

func removeNonPrintableChars(s string) string {
	reg := regexp.MustCompile("[[:^print:]]")
	return reg.ReplaceAllString(s, "")
}
