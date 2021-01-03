package main

import "github.com/castevet6/snippetbox/pkg/models"

// define wrappertype to hold data for html/templates
type templateData struct {
    Snippet *models.Snippet
}
