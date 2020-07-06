provider "azurerm" {
 version = "=2.16.0"
 features { }  
}

resource "azurerm_resource_group" "rg" {
 name     = var.resource_group_name
 location = var.location
}
