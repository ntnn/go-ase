// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

package purego

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/SAP/go-ase/libase/libdsn"
	"github.com/SAP/go-ase/libase/tds"
)

var (
	_ driver.Connector = (*Connector)(nil)
)

type Connector struct {
	DSN            *libdsn.Info
	EnvChangeHooks []tds.EnvChangeHook
	EEDHooks       []tds.EEDHook
}

func NewConnector(dsn *libdsn.Info) (driver.Connector, error) {
	return NewConnectorWithHooks(dsn, nil, nil)
}

func NewConnectorWithHooks(dsn *libdsn.Info, envChangeHooks []tds.EnvChangeHook, eedHooks []tds.EEDHook) (driver.Connector, error) {
	connector := &Connector{
		DSN: dsn,
	}

	conn, err := connector.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error opening test connection: %w", err)
	}
	defer conn.Close()

	pinger, ok := conn.(driver.Pinger)
	if !ok {
		return nil, fmt.Errorf("received conn does not satisfy the pinger interface: %v", conn)
	}

	if err := pinger.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error pinging database server: %w", err)
	}

	// Set the hooks after validating the connection otherwise hooks
	// would get called during the test connection.
	connector.EnvChangeHooks = envChangeHooks
	connector.EEDHooks = eedHooks

	return connector, nil
}

func (c Connector) Driver() driver.Driver {
	return drv
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	return NewConnWithHooks(ctx, c.DSN, c.EnvChangeHooks, c.EEDHooks)
}
