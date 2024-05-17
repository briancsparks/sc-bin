package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(name string, args ...string) int {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("    !!! Command: %v failed\n", name)
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		// Return a non-zero exit code for other errors
		return cmd.ProcessState.ExitCode()
	}

	return 0
}

func main() {

	// Example: Get the current working directory (project directory)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}
	// fmt.Println("Current working directory:", cwd)
	_ = cwd

	if _, err := os.Stat("scripts/workstation-make"); os.IsNotExist(err) {
		fmt.Println("Error: scripts/workstation-make script not found or not executable")
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		exitCode := runCommand("scripts/workstation-make")
		os.Exit(exitCode)
		//		cmd := exec.Command("scripts/workstation-make")
		//		cmd.Stdout = os.Stdout
		//		cmd.Stderr = os.Stderr
		//		if err := cmd.Run(); err != nil {
		//			fmt.Println("Error: scripts/workstation-make script failed")
		//			os.Exit(1)
		//		}

	case os.Args[1] == "r":
		exitCode := runCommand("scripts/workstation-make", os.Args[2:]...)
		os.Exit(exitCode)

		//cmd := exec.Command("scripts/workstation-make")
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//if err := cmd.Run(); err != nil {
		//	fmt.Println("Error: scripts/workstation-make script failed")
		//	os.Exit(1)
		//}

		//checkAurieOut()

		cmd := exec.Command("./build-workstation/aurie.out")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("FAILURE!")
			fmt.Printf("aurie.out failed with exit code: %v\n", cmd.ProcessState.ExitCode())
			fmt.Println("Please check the logs for more details")
			os.Exit(cmd.ProcessState.ExitCode())
		}
		fmt.Println("Success: aurie.out executed successfully")

	case os.Args[1] == "s":
		exitCode := runCommand("scripts/workstation-make", os.Args[2:]...)
		if exitCode != 0 {
			os.Exit(exitCode)
			return
		}

		//checkAurieOut()

		exitCode = runCommand("./build-workstation/aurie.out", "--smoke-test")
		if exitCode != 0 {
			os.Exit(exitCode)
			return
		}
		fmt.Println("Success: Smoke test passed")

	case os.Args[1] == "t":
		cmd := exec.Command("scripts/workstation-make")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: scripts/workstation-make script failed")
			os.Exit(1)
		}

		if _, err := os.Stat("./build-workstation/tests/unit_tests"); os.IsNotExist(err) {
			fmt.Println("Error: unit_tests not found or not executable")
			fmt.Println("Please ensure that the build completed successfully")
			os.Exit(1)
		}

		cmd = exec.Command("./build-workstation/tests/unit_tests")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("FAILURE!")
			fmt.Printf("Unit tests failed with exit code: %v\n", cmd.ProcessState.ExitCode())
			fmt.Println("Please check the logs for more details")
			os.Exit(cmd.ProcessState.ExitCode())
		}
		fmt.Println("Success: Unit tests passed")

	case os.Args[1] == "clean":
		cmd := exec.Command("scripts/workstation-make", "clean")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: scripts/workstation-make clean failed")
			os.Exit(1)
		}

	}
}
