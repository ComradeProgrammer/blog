package main

import (
	"fmt"
	"os"

	"github.com/ComradeProgrammer/blog/internal/blog/model"
	"github.com/ComradeProgrammer/blog/internal/pkg/prettyprint"
)

/*
get category
create category

	input description

create blog

	title
	markdownfile
	categoryid
*/
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("sqlite file is required")
		return
	}

	if len(args) < 3 || args[2] != "get" && args[2] != "create" {
		fmt.Println("operation should be get / create")
		return
	}

	if len(args) < 4 || args[3] != "category" && args[3] != "blog" {
		fmt.Println("operation should be category / blog")
		return
	}
	database := args[1]
	operation := args[2]
	resource := args[3]
	model.ConnectDatabase(database)
	if operation == "get" && resource == "category" {
		categories, err := model.GetCategories()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		prettyprint.PrintJsonln(categories)
		return
	}

	if operation == "get" && resource == "blog" {
		blogs, err := model.GetBlogs()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// for _,b:=range blogs{
		// 	b.Content=""
		// }
		prettyprint.PrintJsonln(blogs)
		return
	}

	if operation == "create" && resource == "category" {
		var name, description string
		fmt.Println("enter name:")
		n, err := fmt.Scanf("%s", &name)
		if err != nil || n == 0 || name == "" {
			fmt.Println("Error: invalid name")
			return
		}

		fmt.Println("enter description:")
		n, err = fmt.Scanf("%s", &description)
		if err != nil || n == 0 || description == "" {
			fmt.Println("Error: invalid description")
			return
		}

		err = model.CreateCategory(
			&model.Category{
				Name:        name,
				Description: description,
			},
		)
		if err != nil {
			fmt.Println("Failed to create category" + err.Error())
			return
		}
		fmt.Println("ok")
		return
	}
	if operation == "create" && resource == "blog" {
		var title, contentFile string
		var categoryID int
		fmt.Println("enter title:")
		n, err := fmt.Scanf("%s", &title)
		if err != nil || n == 0 || title == "" {
			fmt.Println("Error: invalid title")
			return
		}

		fmt.Println("enter name for markdown file:")
		n, err = fmt.Scanf("%s", &contentFile)
		if err != nil || n == 0 || contentFile == "" {
			fmt.Println("Error: invalid description")
			return
		}
		data, err := os.ReadFile(contentFile)
		if err != nil {
			fmt.Println("failed to read file: " + err.Error())
			return
		}
		content := string(data)

		fmt.Println("enter category ID:")
		n, err = fmt.Scanf("%d", &categoryID)
		if err != nil || n == 0 {
			fmt.Println("Error: invalid title")
			return
		}

		err = model.CreateBlog(
			&model.Blog{
				Title:      title,
				Content:    content,
				CategoryID: categoryID,
			},
		)
		if err != nil {
			fmt.Println("Failed to create blog" + err.Error())
			return
		}
		fmt.Println("ok")
		return

	}
}
