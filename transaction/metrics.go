package transaction

type DepotMetricsExporter interface {
	DeliveryReceived(td Delivery)
	DeliveryQueued(td Delivery)
	DeliverySent(td Delivery)
}

type depotMetricsExporterNoop struct{}

func (d *depotMetricsExporterNoop) DeliveryReceived(_ Delivery) {}

func (d *depotMetricsExporterNoop) DeliveryQueued(_ Delivery) {}

func (d *depotMetricsExporterNoop) DeliverySent(_ Delivery) {}
