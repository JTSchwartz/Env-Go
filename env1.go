package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	utility := -1
	iFlag := flag.Bool("i", false, "Modifies current environment")
	flag.Parse()

	for i := 0; i < len(flag.Args()); i++ {
		if !strings.Contains(flag.Arg(i), "=") { // Check if argument is an environment variable
			utility = i
			break
		}
	}

	if utility < 0 {
		if len(flag.Args()) == 0 {
			for j := 0; j < len(os.Environ()); j++ {
				fmt.Println(os.Environ()[j])
			}
		} else {
			env := exec.Command(os.Args[0])
			if !*iFlag {
				env.Env = os.Environ()
			}

			for k := 0; k < len(flag.Args()); k++ {
				env.Env = append(env.Env, flag.Arg(k))
				fmt.Println(flag.Arg(k))
			}

			if out, runErr := env.Output(); runErr != nil {
				fmt.Println(runErr)
				os.Exit(3)
			} else {
				fmt.Print(string(out))
			}
		}
	} else {
		cmd := exec.Command(flag.Arg(utility))

		cmd.Args = flag.Args()[utility+1:]

		if !*iFlag {
			cmd.Env = os.Environ()
		}

		for l := 0; l < len(flag.Args()); l++ {
			cmd.Env = append(cmd.Env, flag.Arg(l))
		}

		if out, runErr := cmd.Output(); runErr != nil {
			fmt.Println(runErr)
			os.Exit(2)
		} else {
			fmt.Print(string(out))
		}
	}
}
