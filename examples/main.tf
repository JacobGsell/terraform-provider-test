terraform {
  required_providers {
    shell = {
      source  = "local/shell"
      version = "0.1.0"
    }
  }
}

provider "shell" {}

resource "shell_command" "example" {
  command = "echo test %PATH%"
}
