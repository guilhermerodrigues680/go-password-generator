package cmd

import (
	"fmt"
	"passwordgenerator/passwordgenerator"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	// enableLowecase  bool
	// enableUppercase bool
	// enableNumbers   bool
	// enableSymbols   bool

	rootCmd = &cobra.Command{
		Use:   "passwordgenerator",
		Short: "Password Generator é um gerador de senha de linha de comando",
		Long:  "Password Generator é um gerador de senha de linha de comando",

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requer um comprimento")
			}

			_, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("comprimento inválido especificado: %q", args[0])
			}

			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			enableLowecase, err := cmd.Flags().GetBool("lowecase")
			cobra.CheckErr(err)
			enableUppercase, err := cmd.Flags().GetBool("uppercase")
			cobra.CheckErr(err)
			enableNumbers, err := cmd.Flags().GetBool("numbers")
			cobra.CheckErr(err)
			enableSymbols, err := cmd.Flags().GetBool("symbols")
			cobra.CheckErr(err)

			if verbose {
				fmt.Println("Tudo certo", enableLowecase, args)
			}

			pwd, err := passwordgenerator.Generate(8, enableLowecase, enableUppercase, enableNumbers, enableSymbols)
			cobra.CheckErr(err)

			fmt.Println(pwd)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verboso")
	rootCmd.PersistentFlags().BoolP("lowecase", "l", false, "usar caracteres minusculos")
	rootCmd.PersistentFlags().BoolP("uppercase", "u", false, "usar caracteres maiusculos")
	rootCmd.PersistentFlags().BoolP("numbers", "n", false, "usar numeros")
	rootCmd.PersistentFlags().BoolP("symbols", "s", false, "usar simbolos")
	// rootCmd.PersistentFlags().BoolVarP(&enableLowecase, "lowecase", "l", false, "usar caracteres minusculos")
	// rootCmd.PersistentFlags().BoolVarP(&enableUppercase, "uppercase", "u", false, "usar caracteres maiusculos")
	// rootCmd.PersistentFlags().BoolVarP(&enableNumbers, "numbers", "n", false, "usar numeros")
	// rootCmd.PersistentFlags().BoolVarP(&enableSymbols, "symbols", "s", false, "usar simbolos")
}

func Execute() {
	rootCmd.Execute()
}
