package opencensus

import (
	"log"

	"github.com/Datadog/opencensus-go-exporter-datadog"
	"go.opencensus.io/stats/view"
)

func main() {
	dd, err := datadog.NewExporter(datadog.Options{})
	if err != nil {
		log.Fatalf("Failed to create the Datadog exporter: %v", err)
	}
	// It is imperative to invoke flush before your main function exits
	defer dd.Stop()

	// Register it as a metrics exporter
	view.RegisterExporter(dd)
}
