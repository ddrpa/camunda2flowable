package main

import (
	"camunda2flowable/types/camunda"
	"encoding/xml"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/dlclark/regexp2"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	Version = "0.0.1"

	header = `<?xml version="1.0" encoding="UTF-8"?>
<definitions xmlns="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:flowable="http://flowable.org/bpmn" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:omgdc="http://www.omg.org/spec/DD/20100524/DC" xmlns:omgdi="http://www.omg.org/spec/DD/20100524/DI" xmlns:design="http://flowable.org/design" typeLanguage="http://www.w3.org/2001/XMLSchema" expressionLanguage="http://www.w3.org/1999/XPath" targetNamespace="http://flowable.org/test" design:palette="flowable-work-process-palette">`
	footer = `
</definitions>`
)

func main() {
	parser := argparse.NewParser("camunda2flowable", "camunda2flowable, version "+Version)
	source := parser.String("s", "source", &argparse.Options{Required: true, Help: "source of camunda BPMN definition XML file, like: example.bpmn"})
	//diagram := parser.Flag("d", "diagram", &argparse.Options{Required: false, Help: "convert diagram information"})
	//checkMode := parser.Flag("c", "check", &argparse.Options{Required: false, Help: "check mode"})
	//target := parser.String("t", "target", &argparse.Options{Required: false, Help: "target of flowable BPMN definition XML file, like: process_1_flowable.bpmn"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		panic(err)
	}

	file, err := os.Open(*source)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	byteValue, _ := io.ReadAll(file)

	var cDefinitions camunda.Definitions
	_ = xml.Unmarshal(byteValue, &cDefinitions)

	fMessages := ConvertMessages(cDefinitions.Message)
	fMessagesInXML, _ := xml.MarshalIndent(fMessages, "  ", "  ")
	fProcess := ConvertProcess(cDefinitions.Process)

	definitionDir := strings.TrimSuffix(*source, filepath.Ext(*source))
	//fileName := strings.TrimSuffix(filepath.Base(*source), filepath.Ext(*source))
	// if same name directory exists
	if _, err := os.Stat(definitionDir); !os.IsNotExist(err) {
		activitieyIds := make([]string, 0)
		files, _ := os.ReadDir(definitionDir)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				// if f is JSON file, add to array
				activityId := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				activitieyIds = append(activitieyIds, activityId)
			}
		}
		if len(activitieyIds) != 0 {
			for _, event := range fProcess.StartEvents {
				if slices.Contains(activitieyIds, event.Id) {
					documentation, _ := os.ReadFile(filepath.Join(definitionDir, event.Id+".json"))
					event.Documentation.Value = string(documentation)
				}
			}
			for _, task := range fProcess.UserTasks {
				if slices.Contains(activitieyIds, task.Id) {
					documentation, _ := os.ReadFile(filepath.Join(definitionDir, task.Id+".json"))
					task.Documentation.Value = string(documentation)
				}
			}
		}
	}

	fProcessInXML, _ := xml.MarshalIndent(fProcess, "  ", "  ")

	// make tag self-closable
	// 用零宽断言排除 <![CDATA[%s]]>
	regex := regexp2.MustCompile("(?m)(?<=[A-Za-z\"]+)><\\/[A-Za-z:]*>$", regexp2.Multiline)
	selfClosed, _ := regex.Replace(string(fMessagesInXML)+"\n"+string(fProcessInXML), "/>", -1, -1)

	fDefinitions := header + selfClosed + footer
	fmt.Println(fDefinitions)
}
