package clientevent

type BasePlayerEvent struct {
	SenderId byte
	Raw      []byte
}

func (p *BasePlayerEvent) Read(data []byte) {
	p.Raw = data
	p.SenderId = data[1]
}
