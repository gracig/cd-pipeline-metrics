package lib

import "fmt"

var (
	//ToolName is the name of this tool
	ToolName string
	//Version is the version of this tool
	Version string
	//BuildNumber is the Continuous Integration Build Number
	BuildNumber string
	//BuildTime is the time when this tool was build
	BuildTime string
	//Revision is the Git repository revision of this tool
	Revision string
	//Author is the last developer that commited the code
	Author string
	//Message is the last message for this commit
	Message string
	//CommitDate is the date for this commit
	CommitDate string
)

//PrintVersion prints the version details of this tool
func PrintVersion() {
	fmt.Printf("Tool Name: %v\n", ToolName)
	fmt.Printf("Version: %v\n", Version)
	fmt.Printf("Build Number: %v\n", BuildNumber)
	fmt.Printf("Revision Code %v\n", Revision)
	fmt.Printf("Commit Date: %v\n", CommitDate)
	fmt.Printf("Build Time: %v\n", BuildTime)
	fmt.Printf("Revision Author: %v\n", Author)
	fmt.Printf("Revision Message: %v\n", Message)
}
