package main

import (
	"bytes"
	"camunda2flowable/types/camunda"
	"camunda2flowable/types/flowable"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/dlclark/regexp2"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	Version = "0.0.1"
	header  = `<?xml version="1.0" encoding="UTF-8"?>
<definitions xmlns="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:flowable="http://flowable.org/bpmn" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:omgdc="http://www.omg.org/spec/DD/20100524/DC" xmlns:omgdi="http://www.omg.org/spec/DD/20100524/DI" xmlns:design="http://flowable.org/design" typeLanguage="http://www.w3.org/2001/XMLSchema" expressionLanguage="http://www.w3.org/1999/XPath" targetNamespace="http://flowable.org/test" design:palette="flowable-work-process-palette">
`
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
	fSignals := ConvertSignals(cDefinitions.Signal)
	fSignalsInXML, _ := xml.MarshalIndent(fSignals, "  ", "  ")
	fErrors := ConvertErrors(cDefinitions.Error)
	fErrorsInXML, _ := xml.MarshalIndent(fErrors, "  ", "  ")

	fProcess := ConvertProcess(cDefinitions.Process)
	definitionDir := strings.TrimSuffix(*source, filepath.Ext(*source))
	// 查找 BPMN 文件所在目录下是否有同名目录，如果有，将该目录下的 JSON 文件内容赋值给对应的 Documentation
	if _, err := os.Stat(definitionDir); !os.IsNotExist(err) {
		externalJSONDocumentationMap := make(map[string]string)
		files, _ := os.ReadDir(definitionDir)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				objectId := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
				fileContent, _ := os.ReadFile(filepath.Join(definitionDir, f.Name()))
				minifiedContent := &bytes.Buffer{}
				json.Compact(minifiedContent, fileContent)
				externalJSONDocumentationMap[objectId] = minifiedContent.String()
			}
		}
		if len(externalJSONDocumentationMap) != 0 {
			for index, event := range fProcess.StartEvents {
				if externalJSONDocumentationMap[event.Id] != "" {
					if event.Documentation != nil {
						fProcess.StartEvents[index].Documentation.Value = externalJSONDocumentationMap[event.Id]
					} else {
						documentation := flowable.Documentation{
							Value: externalJSONDocumentationMap[event.Id],
						}
						fProcess.StartEvents[index].Documentation = &documentation
					}
				}
			}
			for index, task := range fProcess.UserTasks {
				if externalJSONDocumentationMap[task.Id] != "" {
					if task.Documentation != nil {
						fProcess.UserTasks[index].Documentation.Value = externalJSONDocumentationMap[task.Id]
					} else {
						documentation := flowable.Documentation{
							Value: externalJSONDocumentationMap[task.Id],
						}
						fProcess.UserTasks[index].Documentation = &documentation
					}
				}
			}
			for subprocess_index, subprocess := range fProcess.SubProcesses {
				for task_index, task := range subprocess.UserTasks {
					if externalJSONDocumentationMap[task.Id] != "" {
						if task.Documentation != nil {
							fProcess.SubProcesses[subprocess_index].UserTasks[task_index].Documentation.Value = externalJSONDocumentationMap[task.Id]
						} else {
							documentation := flowable.Documentation{
								Value: externalJSONDocumentationMap[task.Id],
							}
							fProcess.SubProcesses[subprocess_index].UserTasks[task_index].Documentation = &documentation
						}
					}
				}
			}
		}
	}
	fProcessInXML, _ := xml.MarshalIndent(fProcess, "  ", "  ")

	// make tag self-closable
	// 用零宽断言排除 <![CDATA[%s]]>
	content := make([]string, 0)
	if fMessagesInXML != nil {
		content = append(content, string(fMessagesInXML))
	}
	if fSignalsInXML != nil {
		content = append(content, string(fSignalsInXML))
	}
	if fErrorsInXML != nil {
		content = append(content, string(fErrorsInXML))
	}
	regex := regexp2.MustCompile("(?m)(?<=[A-Za-z\"]+)><\\/[A-Za-z:]*>$", regexp2.Multiline)
	content = append(content, string(fProcessInXML))
	selfClosed, _ := regex.Replace(strings.Join(content, "\n"), "/>", -1, -1)

	fDefinitions := header + selfClosed + footer
	fmt.Println(fDefinitions)
}
