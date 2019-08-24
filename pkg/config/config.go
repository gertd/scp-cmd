package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GetConfig -- impl
func GetConfig(cmd *cobra.Command, args []string) error {
	fmt.Printf("called GetConfig\n")
	return nil
}

// SetConfig -- impl
func SetConfig(cmd *cobra.Command, args []string) error {
	fmt.Printf("called SetConfig\n")
	return nil
}

// RemoveConfig -- impl
func RemoveConfig(cmd *cobra.Command, args []string) error {
	fmt.Printf("called RemoveConfig\n")
	return nil
}

// ListConfig -- impl
func ListConfig(cmd *cobra.Command, args []string) error {
	fmt.Printf("called ListConfig\n")
	return nil
}
