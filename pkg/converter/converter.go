package converter

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/MChorfa/otm-converter/pkg/converter/types"
	"github.com/google/uuid"
	"goa.design/model/mdl"
	"gopkg.in/yaml.v2"
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

	project := &types.Project{
		Name:         c.design.Name,
		ID:           fmt.Sprintf("%v", c.ID),
		Description:  c.design.Description,
		Owner:        "",
		OwnerContact: "",
		Attributes:   types.Attributes{},
	}

	otm := &types.Otm{
		OtmVersion: c.design.Version,
		Project:    *project,
	}

	for _, system := range c.design.Model.Systems {
		fmt.Printf("%v", system)
		trustZone := &types.TrustZone{
			Name:            system.Name,
			ID:              system.ID,
			Type:            system.Technology,
			Description:     system.Description,
			Risk:            types.TrustZoneRisk{},
			Representations: []types.TrustZoneRepresentation{},
			Attributes:      system.Properties,
			Parent:          &types.Parent{},
		}
		otm.TrustZones = append(otm.TrustZones, *trustZone)
	}

	otmData, err := yaml.Marshal(&otm)
	handleError(err)

	otmFile := path.Join(ouputPath, "otm.yaml")

	f, err := os.Create(otmFile)
	handleError(err)

	f.Write(otmData)

	defer func() { _ = f.Close() }()
	handleError(err)

	return nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
