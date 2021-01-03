package new

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/ChrisRx/splits/pkg/prompt"
	"github.com/ChrisRx/splits/pkg/srapi"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "new [filename]",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inprompt := &survey.Input{
				Message: "Search for game",
			}
			var answer string
			if err := survey.AskOne(inprompt, &answer); err != nil {
				return err
			}
			games, err := srapi.GetGameByName(answer)
			if err != nil {
				return err
			}
			if len(games) == 0 {
				return errors.Errorf("No games found for query: %q\n", answer)
			}
			opts := make([]string, 0)
			for _, game := range games {
				opts = append(opts, game.Names.International)
			}
			selprompt := &survey.Select{
				Message: fmt.Sprintf("Found %d games:", len(games)),
				Options: opts,
				VimMode: true,
			}
			if err := survey.AskOne(selprompt, &answer); err != nil {
				return err
			}
			var game *srapi.Game
			for _, g := range games {
				if g.Names.International == answer {
					game = g
				}
			}

			categories, err := game.Categories()
			if err != nil {
				return err
			}
			opts = opts[:0]
			for _, cat := range categories {
				opts = append(opts, cat.Name)
			}
			if len(games) == 0 {
				return errors.Errorf("No categories found for game: %q\n", answer)
			}
			selprompt = &survey.Select{
				Message: fmt.Sprintf("Found %d categories:", len(categories)),
				Options: opts,
				VimMode: true,
			}
			if err := survey.AskOne(selprompt, &answer); err != nil {
				return err
			}

			var cat *srapi.Category
			for _, c := range categories {
				if c.Name == answer {
					cat = c
				}
			}
			if prompt.Confirmf("write this file to path: %s", args[0]) {
				data, err := json.MarshalIndent(struct {
					ID       string
					Name     string
					Category string
				}{
					ID:       game.ID,
					Name:     game.Names.International,
					Category: cat.Name,
				}, "", "    ")
				if err != nil {
					return err
				}
				data = append(data, '\n')
				return ioutil.WriteFile(args[0], data, 0644)
			}
			return nil
		},
	}
	return cmd
}
