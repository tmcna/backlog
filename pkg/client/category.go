package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Category is ...
type Category struct {
	space      string
	apiKey     string
	projectKey string
	request    url.Values
}

// NewCategory is ...
func NewCategory(space string, apiKey string, projectKey string) *Category {
	c := new(Category)
	c.space = space
	c.apiKey = apiKey
	c.request = url.Values{}
	c.projectKey = projectKey

	return c
}

// Request is ...
func (c *Category) Request(key string, value string) {
	c.request.Set(key, value)
}

// List function returns list of Categories in the project.
func (c *Category) List() ([]CategoryResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/categories", c.projectKey)
	
	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []CategoryResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Category to the project.
func (c *Category) Add() (*CategoryResponse, error) {

	api := "api/v2/projects/" + c.projectKey + "/categories"
	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Post(api, c.request)
	if err != nil {
		return nil, err
	}
	var r CategoryResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes Category by category ID.
func (c *Category) Delete(id int) (*CategoryResponse, error) {

	api := "api/v2/projects/" + c.projectKey + "/categories/" + strconv.Itoa(id)

	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Delete(api, nil)
	if err != nil {
		return nil, err
	}

	var r CategoryResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from category name.
func (c *Category) GetID(name string) (int, error) {
	api := "api/v2/projects/" + c.projectKey + "/categories"
	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []CategoryResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return -1, err
	}

	for _, n := range r {
		if n.Name == name {
			return n.ID, nil
		}
	}

	return -1, nil
}

// PrintCSV prints list of categories in CSV format.
func (c *Category) PrintCSV(r []CategoryResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s\n", n.ID, n.Name)
	}
}
