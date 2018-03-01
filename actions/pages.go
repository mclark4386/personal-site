package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/mclark4386/personal-site/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Page)
// DB Table: Plural (pages)
// Resource: Plural (Pages)
// Path: Plural (/pages)
// View Template Folder: Plural (/templates/pages/)

// PagesResource is the resource for the Page model
type PagesResource struct {
	buffalo.Resource
}

// List gets all Pages. This function is mapped to the path
// GET /pages
func (v PagesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	pages := &models.Pages{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Pages from the DB
	if err := q.All(pages); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, pages))
}

// Show gets the data for one Page. This function is mapped to
// the path GET /pages/{page_id}
func (v PagesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Page
	page := &models.Page{}

	// To find the Page the parameter page_id is used.
	if err2 := tx.Where("slug = ?", c.Param("page_id")).First(page); err2 != nil {
		if err := tx.Find(page, c.Param("page_id")); err != nil {
			return c.Error(404, err)
		}
	}

	return c.Render(200, r.Auto(c, page))
}

// New renders the form for creating a new Page.
// This function is mapped to the path GET /pages/new
func (v PagesResource) New(c buffalo.Context) error {
	// Make page available inside the html template
	c.Set("page", &models.Page{})

	return c.Render(200, r.HTML("pages/new.html"))
}

// Create adds a Page to the DB. This function is mapped to the
// path POST /pages
func (v PagesResource) Create(c buffalo.Context) error {
	// Allocate an empty Page
	page := &models.Page{}

	// Bind page to the html form elements
	if err := c.Bind(page); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(page)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make page available inside the html template
		c.Set("page", page)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("pages/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Page was created successfully")

	// and redirect to the pages index page
	return c.Redirect(302, "/pages/%s", page.ID)
}

// Edit renders a edit form for a Page. This function is
// mapped to the path GET /pages/{page_id}/edit
func (v PagesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Page
	page := &models.Page{}

	if err := tx.Find(page, c.Param("page_id")); err != nil {
		return c.Error(404, err)
	}

	// Make page available inside the html template
	c.Set("page", page)
	return c.Render(200, r.HTML("pages/edit.html"))
}

// Update changes a Page in the DB. This function is mapped to
// the path PUT /pages/{page_id}
func (v PagesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Page
	page := &models.Page{}

	if err := tx.Find(page, c.Param("page_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Page to the html form elements
	if err := c.Bind(page); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(page)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make page available inside the html template
		c.Set("page", page)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("pages/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Page was updated successfully")

	// and redirect to the pages index page
	return c.Redirect(302, "/pages/%s", page.ID)
}

// Destroy deletes a Page from the DB. This function is mapped
// to the path DELETE /pages/{page_id}
func (v PagesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Page
	page := &models.Page{}

	// To find the Page the parameter page_id is used.
	if err := tx.Find(page, c.Param("page_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(page); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Page was destroyed successfully")

	// Redirect to the pages index page
	return c.Redirect(302, "/pages")
}
