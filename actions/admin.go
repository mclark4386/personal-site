package actions

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gofrs/uuid"
	"github.com/mclark4386/personal-site/models"
)

// AdminPanel default implementation.
func AdminPanel(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	config := models.SiteConfig{}
	tx.Where("key = ?", "AllowNewUsers").First(&config)
	if config.ID != uuid.Nil && strings.ToLower(config.Value) != "true" {
		c.Set("AllowNewUsers", config.Value)
	} else {
		c.Set("AllowNewUsers", "true")
		spew.Printf("config:%+v\n", config)
	}
	return c.Render(200, r.HTML("admin/panel.html"))
}

//SetSiteConfigs is behind Authorize and will set the AllowNewUsers key to whatever we are given.
func SetSiteConfigs(c buffalo.Context) error {
	allow := c.Param("AllowNewUsers")
	tx := c.Value("tx").(*pop.Connection)
	if allow != "" {
		config := models.SiteConfig{}

		err := tx.Where("key = ?", "AllowNewUsers").First(&config)
		if err != nil {
			config.Key = "AllowNewUsers"
			config.Value = allow
			err = nil
		} else {
			config.Value = allow
		}
		err = tx.Save(&config)
		if err != nil {
			return c.Error(402, err)
		}
		print("good!\n")
	}
	return c.Redirect(302, "/")
}
