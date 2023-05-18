package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"gotodo/pkg/infrastructure/client"
)

func NewLoginCommand(login client.Login) *cobra.Command {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login and save token",
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")

			claims, err := login.Execute(client.Credentials{
				Username: username,
				Password: password,
			})

			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			fmt.Println("Login successful")
			fmt.Printf("Welcome, %s\n", claims.Name)
		},
	}

	loginCmd.Flags().String("username", "", "Username")
	loginCmd.Flags().String("password", "", "Password")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")

	return loginCmd
}
