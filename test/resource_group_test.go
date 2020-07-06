package test

import (
	"os"
	"path"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	terraformCore "github.com/hashicorp/terraform/terraform"
)

func readPlan(planPath string) (*terraformCore.Plan, error) {
	f, err := os.Open(planPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	plan, err := terraformCore.ReadPlan(f)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func TestUT_ResourceGroup(t *testing.T) {
	tfOptions := &terraform.Options{
		TerraformDir: "./",
	}
	tfPlanOutput := "terraform.tfplan"
	terraform.Init(t, tfOptions)
	terraform.RunTerraformCommand(t, tfOptions, terraform.FormatArgs(tfOptions, "plan", "-out="+tfPlanOutput)...)

	plan, err := readPlan(path.Join(tfOptions.TerraformDir, tfPlanOutput))
	if err != nil {
		t.Fatal(err)
	}

	expected := "rg-msft-beandrad-test1"
	for _, mod := range plan.Diff.Modules {
		if len(mod.Path) == 2 && mod.Path[0] == "root" && mod.Path[1] == "staticwebpage" {
			actual := mod.Resources["azurerm_resource_group.main"].Attributes["name"].New
			if actual != expected {
				t.Fatalf("Expect %v, but found %v", expected, actual)
			}
		}
	}

}
