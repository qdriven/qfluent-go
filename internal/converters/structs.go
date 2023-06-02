package converters

type StringFormatConverter interface {
	ToObject(structuredString string, any any)
	ToStructuredString(any any) string
	FileContentToObject(filePath string, any any)
}

type FormatConverter struct {
	Converter StringFormatConverter
}

func (s FormatConverter) UseConverter(converter StringFormatConverter) FormatConverter {
	s.Converter = converter
	return s
}

func NewFormatConverter(converter StringFormatConverter) *FormatConverter {
	return &FormatConverter{Converter: converter}
}
