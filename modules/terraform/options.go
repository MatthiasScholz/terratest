package terraform

import "time"

// Options for running Terraform commands
type Options struct {
	TerraformDir             string                 // The path to the folder where the Terraform code is defined.
	Vars                     map[string]interface{} // The vars to pass to Terraform commands using the -var option.
	VarsFiles                []string               // The vars files to pass to Terraform commands using -var-file option.
	EnvVars                  map[string]string      // Environment variables to set when running Terraform
	RetryableTerraformErrors map[string]string      // If Terraform apply fails with one of these (transient) errors, retry. The keys are text to look for in the error and the message is what to display to a user if that error is found.
	MaxRetries               int                    // Maximum number of times to retry errors matching RetryableTerraformErrors
	TimeBetweenRetries       time.Duration          // The amount of time to wait between retries
}
