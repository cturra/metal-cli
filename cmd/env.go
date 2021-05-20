/*
Copyright © 2020 Equinix Metal Developers <support@equinixmetal.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// envCmd represents a command that, when run, generates a
// set of environment variables, for use in shell environments
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Generate environment variables",
	Long: fmt.Sprintf(`Currently emitted variables:
	- %s

	To load environment variables:

	Bash, Zsh:

	$ source <(packet env)

	Bash (3.2.x):

	$ eval "$(packet env)"

	Fish:

	$ packet env | source
	`, apiTokenEnvVar),
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s=%s\n", apiTokenEnvVar, packetToken)
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
