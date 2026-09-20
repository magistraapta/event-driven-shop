package jobs

type ProcessOrder struct {
}

// Signature The name and signature of the job.
func (receiver *ProcessOrder) Signature() string {
	return "process_order"
}

// Handle Execute the job.
func (receiver *ProcessOrder) Handle(args ...any) error {
	return nil
}
