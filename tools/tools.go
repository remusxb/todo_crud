//go:build tools

package tools

import (
	_ "github.com/gofiber/fiber/v2"
)

//go:generate go get -u github.com/gofiber/fiber/v2@latest
