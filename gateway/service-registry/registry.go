package serviceregistry

// static service registry
var ServiceRegistry = map[string]string{
	"patient": "localhost:50051",
	"doctor":  "localhost:50052",
}
