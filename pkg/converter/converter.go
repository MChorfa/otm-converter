package converter

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"goa.design/model/mdl"
)

type Converter struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Nonce     string
	design    *mdl.Design
}

func NewConverter(design *mdl.Design) *Converter {
	var c Converter
	c.ID = uuid.New()
	c.CreatedAt = time.Now()
	c.design = design
	return &c
}

func (c *Converter) Convert(ouputPath string) error {

	for _, system := range c.design.Model.Systems {
		fmt.Printf("%v", system)
	}

	otmFile := path.Join(ouputPath, "otm.yaml")
	f, err := os.Create(otmFile)
	defer func() { _ = f.Close() }()
	if err != nil {
		msg := fmt.Sprintf("Saving failed, can't write to %s: %s!\n", otmFile, err)
		fmt.Println(msg)
	}

	return nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
