package mongodb_test

import (
	"os"
	"testing"

	"github.com/PizzaStruct/TenseiAPI/internal/pkg/mongodb"
)

func TestConnect(t *testing.T) {
	os.Setenv("MONGO", "mongodb://localhost:27017/")
	os.Setenv("DB", "TenseiAPI")
	if err := mongodb.Connect(); err != nil {
		t.Fatalf("Connection error: %s", err.Error())
	}
}
