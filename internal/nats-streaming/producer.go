package natsstreaming

func (ns *natsStreaming) Publish() error {
	err := ns.sc.Publish(ns.subject, []byte("order"))
	if err != nil {
		return err
	}

	return nil
}