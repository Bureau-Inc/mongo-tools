package mongoexport // ExportDocument writes a line to output with the CSV representation of a document.

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mongodb/mongo-tools/common/bsonutil"
	"github.com/mongodb/mongo-tools/common/json"
	"go.mongodb.org/mongo-driver/bson"
)

// ExportDocumentCustom differs from ExportDocument method of *CSVExportOutput
// while formatting floats. ExportDocument uses fmt.Sprint to format all values,
// whereas ExportDocumentCustom uses strconv.FormatFloat to format float values.
// fmt.Sprint method converts floats to scientific notation while formatting.
// We wanted to avoid this, so we're using strconv.FormatFloat to format floats
// in standard notation.
func (csvExporter *CSVExportOutput) ExportDocumentCustom(document bson.D) error {
	rowOut := make([]string, 0, len(csvExporter.Fields))
	extendedDoc, err := bsonutil.ConvertBSONValueToLegacyExtJSON(document)
	if err != nil {
		return err
	}

	for _, fieldName := range csvExporter.Fields {
		fieldVal := extractFieldByName(fieldName, extendedDoc)
		if fieldVal == nil {
			rowOut = append(rowOut, "")
		} else if reflect.TypeOf(fieldVal) == reflect.TypeOf(bson.M{}) ||
			reflect.TypeOf(fieldVal) == reflect.TypeOf(bson.D{}) ||
			reflect.TypeOf(fieldVal) == marshalDType ||
			reflect.TypeOf(fieldVal) == reflect.TypeOf([]interface{}{}) {
			buf, err := json.Marshal(fieldVal)
			if err != nil {
				rowOut = append(rowOut, "")
			} else {
				rowOut = append(rowOut, string(buf))
			}
		} else if kind := reflect.TypeOf(fieldVal).Kind(); kind == reflect.Float32 || kind == reflect.Float64 {
			f := fieldVal.(json.NumberFloat)
			rowOut = append(rowOut, strconv.FormatFloat(float64(f), 'f', -1, 64))
		} else {
			rowOut = append(rowOut, fmt.Sprintf("%v", fieldVal))
		}
	}
	csvExporter.csvWriter.Write(rowOut)
	csvExporter.NumExported++
	return csvExporter.csvWriter.Error()
}
