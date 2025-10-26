package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	defaults "github.com/mcuadros/go-defaults"
)

type Filter map[string]any

type ListOpts struct {
	Search    string `query:"search" default:""`
	Page      int    `query:"page" default:"1"`
	PerPage   int    `query:"per_page" default:"10"`
	SortBy    string `query:"sort_by" default:"created_at"`
	SortOrder string `query:"sort_order" default:"desc"`
	Select    string `query:"select" default:""`
	Trash     bool   `query:"trash" default:"false"`
	Filters   Filter `query:"filters"`
}
type PaginationInfo struct {
	From  int `json:"from"`
	To    int `json:"to"`
	Total int `json:"total"`
}

type PaginatedResponse[S Schema] struct {
	Data []S            `json:"data"`
	Info PaginationInfo `json:"info"`
}

func NewListOptions(c *fiber.Ctx) (*ListOpts, error) {
	var listOpts ListOpts
	err := c.QueryParser(&listOpts)

	if err != nil {
		return nil, err
	}
	defaults.SetDefaults(&listOpts)
	fmt.Printf("%+v", listOpts)
	return &listOpts, nil
}

func (l *ListOpts) SortOrderInt() int {
	if l.SortOrder == "desc" {
		return -1
	}
	return 1
}
