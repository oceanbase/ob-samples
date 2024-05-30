package cmd

import (
	"encoding/hex"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Convert hex string to byte",
	RunE:  writeE,
}

var in string
var out string

func init() {
	hexCmd.Flags().StringVarP(&in, "in", "i", "", "hex string")
	hexCmd.Flags().StringVarP(&out, "out", "o", "", "out file name")

	hexCmd.MarkFlagRequired("in")
	hexCmd.MarkFlagRequired("out")

	rootCmd.AddCommand(hexCmd)
}

func writeE(cmd *cobra.Command, args []string) error {
	log.Printf(">>> %s", in)

	decodeString, err := hex.DecodeString(in)
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(decodeString)
	if err != nil {
		return err
	}

	return nil
}
