package writer

import (
	"io"
	"learn/internal/converter/models"
)

type ProtoWriter struct {
	IO io.Writer
}

func (c *ProtoWriter) WriteHead(logs []models.LogRow) error {

}

func (c *ProtoWriter) WriteRow(log models.LogRow) error {

}
