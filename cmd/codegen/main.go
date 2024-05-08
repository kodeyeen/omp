package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const srcTpl = `#include "include/%s"

#ifdef __cplusplus
extern "C" {
#endif

%s

#ifdef __cplusplus
}
#endif
`

const funcTpl = `%s%s %s(%s) {
%s%sreturn call<%s>(%s);
%s}`

// ^\s*([\w\s\*]+)\s+(\w+)\s*\(([^)]*)\)\s*;
var funcPrototypeRe = regexp.MustCompile(`(?m)` +
	`^` +
	`\s*` +
	`([\w\s\*]+)` + // type
	`\s+` +
	`(\w+)` + // name
	`\s*` +
	`\(` +
	`([^)]*)` + // arguments
	`\)` +
	`\s*` +
	`;`,
)

func main() {
	genSrcFiles("include")
}

func genSrcFiles(headersPath string) error {
	headers, err := os.ReadDir(headersPath)
	if err != nil {
		return err
	}

	for _, header := range headers {
		if header.Name() == "omp.h" {
			continue
		}

		headerPath := filepath.Join(headersPath, header.Name())
		genSrcFile(headerPath)
	}

	return nil
}

func genSrcFile(headerPath string) error {
	contents, err := os.ReadFile(headerPath)
	if err != nil {
		return err
	}

	matches := funcPrototypeRe.FindAllStringSubmatch(string(contents), -1)

	funcs := make([]string, 0, len(matches))

	for _, match := range matches {
		retType := match[1]
		name := match[2]
		args := match[3]

		funcs = append(funcs, newFunc(retType, name, args, "    "))
	}

	header := filepath.Base(headerPath)
	headerName := strings.TrimSuffix(header, filepath.Ext(header))

	src := fmt.Sprintf(srcTpl, header, strings.Join(funcs, "\n\n"))

	// srcFile, _ := os.Create("test.cpp")
	srcFile, _ := os.Create(fmt.Sprintf("%s.cpp", headerName))
	srcFile.WriteString(src)

	return nil
}

func newFunc(retType, name, args, indent string) string {
	callArgs := []string{
		fmt.Sprintf(`"%s"`, name),
	}

	if args != "" {
		callArgs = append(callArgs, removeArgTypes(args)...)
	}

	return fmt.Sprintf(
		funcTpl,
		indent,
		retType,
		name,
		args,
		indent, indent,
		retType,
		strings.Join(callArgs, ", "),
		indent,
	)
}

func removeArgTypes(rawArgs string) []string {
	args := strings.Split(rawArgs, ",")
	result := make([]string, 0, len(args))

	for _, arg := range args {
		argWords := strings.Fields(strings.TrimSpace(arg))
		argName := argWords[len(argWords)-1]

		result = append(result, argName)
	}

	return result
}
