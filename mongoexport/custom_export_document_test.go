package mongoexport // ExportDocument writes a line to output with the CSV representation of a document.

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExportDocumentCustom(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	csvOut := NewCSVExportOutput([]string{"f1", "f2", "f3", "f4"}, true, output)
	doc := primitive.D{
		primitive.E{Key: "f1", Value: float64(1.234)},
		primitive.E{Key: "f2", Value: float64(1)},
		primitive.E{Key: "f3", Value: float32(-12334.2433)},
		primitive.E{Key: "f4", Value: float64(844737823984723.4430948230)},
	}
	err := csvOut.ExportDocumentCustom(doc)
	assert.NoError(t, err)
	csvOut.Flush()

	assert.Equal(t, "1.234,1,-12334.2431640625,844737823984723.5\n", output.String())
}
