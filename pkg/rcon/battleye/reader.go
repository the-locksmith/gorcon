package battleye

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	be_proto "github.com/playnet-public/battleye/battleye"
	"github.com/playnet-public/gorcon/pkg/rcon"
)

// HandlePacket received from UDP connection
func (c *Connection) HandlePacket(p be_proto.Packet) (err error) {
	defer func() {
		err = errors.Wrap(err, "handling packet")
	}()
	err = c.Protocol.Verify(p)
	if err != nil {
		// TODO: Add logging
		return err
	}
	data, err := c.Protocol.Data(p)
	if err != nil {
		// TODO: Add logging
		return err
	}

	// Handle KeepAlive Pingback
	if len(data) < 1 {
		// TODO: Add logging
		c.AddPingback()
		return nil
	}

	t, err := c.Protocol.Type(p)
	if err != nil {
		// TODO: Add logging
		return err
	}

	switch t {
	case be_proto.Command | be_proto.MultiCommand:
		return c.HandleResponse(p)

	case be_proto.ServerMessage:
		// Handle MultiCommand
		return nil

	}

	return nil
}

// HandleResponse by retrieving the corresponding transmission and updating it
func (c *Connection) HandleResponse(p be_proto.Packet) error {
	s, err := c.Protocol.Sequence(p)
	if err != nil {
		return errors.Wrap(err, "handling response")
	}

	trm := c.GetTransmission(s)
	if trm == nil {
		return errors.New("no transmission for response")
	}

	t, err := c.Protocol.Type(p)
	if err != nil {
		return errors.Wrap(err, "handling response")
	}

	data, err := c.Protocol.Data(p)
	if err != nil {
		return errors.Wrap(err, "handling response")
	}

	last := true
	if t == be_proto.MultiCommand {
		count, index, single := c.Protocol.Multi(p)
		if !single {
			trm.multiBuffer[int(index)] = data
			last = (index+1 >= count)
		}
	} else {
		trm.multiBuffer[0] = data
	}

	if last {
		select {
		case trm.done <- true:
			return nil
		case <-time.After(time.Second):
			// TODO: Add debug log for transmission done timeouts
			return nil
		}
	}

	return nil
}

// HandleServerMessage containing chat and events
func (c *Connection) HandleServerMessage(p be_proto.Packet) error {
	s, err := c.Protocol.Sequence(p)
	if err != nil {
		return errors.Wrap(err, "handling server message")
	}

	var channels = []string{
		"(Group)",
		"(Vehicle)",
		"(Unknown)",
	}
	var t = rcon.TypeEvent
	for _, c := range channels {
		if strings.HasPrefix(string(p), c) {
			t = rcon.TypeChat
		}
	}

	event := &rcon.Event{
		Timestamp: time.Now(),
		Type:      t,
		Message:   string(p),
	}

	_, err = c.UDP.Write(c.Protocol.BuildMsgAckPacket(s))
	if err != nil {
		return errors.Wrap(err, "handling server message")
	}

	c.listenersMutex.RLock()
	defer c.listenersMutex.RUnlock()
	for _, l := range c.listeners {
		go func(l chan *rcon.Event) { l <- event }(l)
	}

	return nil
}