package user

import (
	"fmt"
	"time"
	"regexp"
	"github.com/spf13/cobra"
	"../../model"
)

func AddCmd() *cobra.Command {
	return &cobra.Command {
		Use: "add",
		Short: "add user",
		Run: func(cmd *cobra.Command, args []string) {

			if args[1] == "" {
				fmt.Println("Please give user's Nickname")
				return
			}
			match, err := regexp.MatchString("^d{3}-d{3,4}-d{4}$/", args[2])
			if err != nil {
				fmt.Errorf("error in user add by user's phone match")
				return
			} else if !match {
				fmt.Println("Please give proper phone number")
				return
			}

			newUser := new(model.User)
			newUser.Nickname = args[1]
			newUser.Phone = args[2]
			newUser.PhoneDevice = args[3]
			newUser.CreatedAt = time.Now()


		},
	}
}