package credit

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gitlab.com/velo-labs/cen/cmd/gvel/layers/logic"
	"gitlab.com/velo-labs/cen/cmd/gvel/utils/config"
	"gitlab.com/velo-labs/cen/cmd/gvel/utils/console"
)

type CommandHandler struct {
	Logic     logic.Logic
	Prompt    console.Prompt
	AppConfig config.Configuration
}

func NewCommandHandler(logic logic.Logic, prompt console.Prompt, config config.Configuration) *CommandHandler {
	return &CommandHandler{
		Logic:     logic,
		Prompt:    prompt,
		AppConfig: config,
	}
}

func (creditCommand *CommandHandler) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "credit <arg>",
		Short: "Use credit command for managing the stable credit on Velo",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !creditCommand.AppConfig.Exists() {
				console.ExitWithError(console.ExitError, errors.New("config file not found, please run `gvel init`"))
			}
		},
	}

	command.AddCommand(
		creditCommand.GetSetupCommand(),
	)

	return command
}

func (creditCommand *CommandHandler) GetSetupCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "setup",
		Short: "Setup a stable credit on Velo",
		Run:   creditCommand.Setup,
	}

	return command
}