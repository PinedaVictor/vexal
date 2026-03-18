/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"vx/config"
	"vx/internal/keyrings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:     "context",
	Aliases: []string{"ctx"},
	Short:   "Manage Vexal contexts for external service configuration",
	Long: `Manage Vexal contexts used to configure external services.

Contexts allow you to define different environments such as personal
and work setups. Each context can store service configuration like
GitHub credentials and other integrations used by Vexal commands.

When running commands such as "vx pr", Vexal will attempt to resolve
the correct context automatically based on the repository's git remote.
Repository configuration (.vx.yaml) always takes precedence over
global contexts stored in ~/.vx/context.json.
		
Examples:

  vx context list
      List all configured contexts.

  vx context add work
      Add or update a context.

  vx context use work
      Set the active context.`,
}

var contextListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured contexts",
	Run: func(cmd *cobra.Command, args []string) {
		ctxCfg, err := config.LoadContextConfig()
		if err != nil || len(ctxCfg.Contexts) == 0 {
			fmt.Println("No contexts configured. Run 'vx context add <name>' to get started.")
			return
		}
		for name, ctx := range ctxCfg.Contexts {
			active := ""
			if name == ctxCfg.Active {
				active = " (active)"
			}
			fmt.Printf("%s%s\n", name, active)
			fmt.Printf("  github_user: %s\n", ctx.GithubUser)
			if _, err := keyrings.GetSecret(name, "github_key"); err == nil {
				fmt.Println("  github_key:  configured")
			} else {
				fmt.Println("  github_key:  not set")
			}
			if _, err := keyrings.GetSecret(name, "openai_key"); err == nil {
				fmt.Println("  openai_key:  configured")
			} else {
				fmt.Println("  openai_key:  not set")
			}
		}
	},
}

func readMaskedInput() (string, error) {
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	var input []byte
	buf := make([]byte, 1)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			return "", err
		}
		switch buf[0] {
		case '\r', '\n':
			fmt.Print("\r\n")
			return strings.TrimSpace(string(input)), nil
		case 127, '\b':
			if len(input) > 0 {
				input = input[:len(input)-1]
				fmt.Print("\b \b")
			}
		default:
			input = append(input, buf[0])
			fmt.Print("*")
		}
	}
}

var contextAddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Add or update a context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		githubUser, _ := cmd.Flags().GetString("github-user")
		if githubUser == "" {
			fmt.Println("GitHub username is required. Use --github-user or -u.")
			os.Exit(1)
		}

		fmt.Print("GitHub token: ")
		githubToken, err := readMaskedInput()
		if err != nil {
			fmt.Println("Error reading GitHub token:", err)
			os.Exit(1)
		}
		if githubToken == "" {
			fmt.Println("GitHub token cannot be empty.")
			os.Exit(1)
		}

		fmt.Print("OpenAI key: ")
		openaiKey, err := readMaskedInput()
		if err != nil {
			fmt.Println("Error reading OpenAI key:", err)
			os.Exit(1)
		}
		if openaiKey == "" {
			fmt.Println("OpenAI key cannot be empty.")
			os.Exit(1)
		}

		if err := keyrings.SetSecret(name, "github_key", githubToken); err != nil {
			fmt.Println("Error saving GitHub token:", err)
			os.Exit(1)
		}
		if err := keyrings.SetSecret(name, "openai_key", openaiKey); err != nil {
			fmt.Println("Error saving OpenAI key:", err)
			os.Exit(1)
		}

		ctxCfg, _ := config.LoadContextConfig()
		ctxCfg.Contexts[name] = config.Context{GithubUser: githubUser}
		if ctxCfg.Active == "" {
			ctxCfg.Active = name
		}
		if err := config.SaveContextConfig(ctxCfg); err != nil {
			fmt.Println("Error saving context:", err)
			os.Exit(1)
		}
		fmt.Printf("Context '%s' saved.\n", name)
	},
}

var contextUseCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Set the active context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctxCfg, err := config.LoadContextConfig()
		if err != nil || len(ctxCfg.Contexts) == 0 {
			fmt.Println("No contexts found. Run 'vx context add <name>' first.")
			os.Exit(1)
		}
		if _, ok := ctxCfg.Contexts[name]; !ok {
			fmt.Printf("Context '%s' not found. Run 'vx context list' to see available contexts.\n", name)
			os.Exit(1)
		}
		ctxCfg.Active = name
		if err := config.SaveContextConfig(ctxCfg); err != nil {
			fmt.Println("Error saving context:", err)
			os.Exit(1)
		}
		fmt.Printf("Active context set to '%s'.\n", name)
	},
}

var contextRemoveCmd = &cobra.Command{
	Use:   "remove [name]",
	Short: "Remove a context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		ctxCfg, err := config.LoadContextConfig()
		if err != nil || len(ctxCfg.Contexts) == 0 {
			fmt.Println("No contexts found. Run 'vx context list' to see available contexts.")
			os.Exit(1)
		}
		if _, ok := ctxCfg.Contexts[name]; !ok {
			fmt.Printf("Context '%s' not found. Run 'vx context list' to see available contexts.\n", name)
			os.Exit(1)
		}

		keyrings.DeleteSecret(name, "github_key")
		keyrings.DeleteSecret(name, "openai_key")

		delete(ctxCfg.Contexts, name)
		if ctxCfg.Active == name {
			ctxCfg.Active = ""
		}
		if err := config.SaveContextConfig(ctxCfg); err != nil {
			fmt.Println("Error saving context:", err)
			os.Exit(1)
		}
		fmt.Printf("Context '%s' removed.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
	contextCmd.AddCommand(contextListCmd)
	contextCmd.AddCommand(contextAddCmd)
	contextCmd.AddCommand(contextUseCmd)
	contextCmd.AddCommand(contextRemoveCmd)

	contextAddCmd.Flags().StringP("github-user", "u", "", "GitHub username for this context")
}
