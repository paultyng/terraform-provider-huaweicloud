config {
  disabled_by_default = false
}

rule "terraform_deprecated_index" {
  enabled = true
}
rule "terraform_unused_declarations" {
  enabled = true
}
rule "terraform_comment_syntax" {
  enabled = true
}
rule "terraform_documented_outputs" {
  enabled = true
}
rule "terraform_documented_variables" {
  enabled = false
}
rule "terraform_typed_variables" {
  enabled = false
}
# https://github.com/terraform-linters/tflint/blob/master/docs/rules/terraform_required_version.md
rule "terraform_required_version" {
  enabled = false
}
