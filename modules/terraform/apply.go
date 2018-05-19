package terraform

import (
	"testing"
)

// Run terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApply(t *testing.T, options *Options) string {
	out, err := InitAndApplyE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// Run terraform init and apply with the given options and return stdout/stderr from the apply command. Note that this
// method does NOT call destroy and assumes the caller is responsible for cleaning up any resources created by running
// apply.
func InitAndApplyE(t *testing.T, options *Options) (string, error) {
	if _, err := InitE(t, options); err != nil {
		return "", err
	}

	if _, err := GetE(t, options); err != nil {
		return "", err
	}

	return ApplyE(t, options)
}

// Run terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func Apply(t *testing.T, options *Options) string {
	out, err := ApplyE(t, options)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// Run terraform apply with the given options and return stdout/stderr. Note that this method does NOT call destroy and
// assumes the caller is responsible for cleaning up any resources created by running apply.
func ApplyE(t *testing.T, options *Options) (string, error) {
	return RunTerraformCommandE(t, options, FormatArgs(options.Vars, "apply", "-input=false", "-lock=false", "-auto-approve", FormatArgsFiles(options.VarsFiles))...)
}
