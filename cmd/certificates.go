//Copyright © 2020, SAS Institute Inc., Cary, NC, USA.  All Rights Reserved.
//SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/sassoftware/viya4-orders-cli/lib/assetreqs"

	"github.com/spf13/cobra"
)

// certsCmd represents the certs command
var certificatesCmd = &cobra.Command{
	Use:   "certificates [order number]",
	Short: "Download certificates for the given order number",
	Example: "viya4-orders-cli certs 993456\n"+
		"viya4-orders-cli certs 993456 -p $HOME/sas",
	Aliases: []string{"certs", "cer"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ar := assetreqs.New(token, "certificates", args[0], "","",
			assetFilePath, assetFileName, outFormat)
		ar.GetAsset()
	},
}

func init() {
	rootCmd.AddCommand(certificatesCmd)
}