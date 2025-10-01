package main

import (
	"os"
)

func main() {
	if os.Args[0] != "lemur" {
		panic("Invalid executable name, its should be lemur.")
	}
	
	switch len(os.Args) {
		case 1:
			VersionPanel();
			
		case 2:
			switch os.Args[1] {
				case "-v": VersionPanel();
				case "-h": HelpPanel();
				case "stdout": ReadPipe(); 
				default: Fail(nil);
			}
			
		case 3: 
			switch os.Args[1] {
				case "-f": ReadFile(os.Args[2]);
				case "-p": ReadPath(os.Args[2]);
				default: Fail(nil);
			}
		
		default:
			Fail(nil)
	}
}
