package systemd

import (
	"context"
	"fmt"
	"github.com/coreos/go-systemd/v22/dbus"
	"time"
)

const (
	jobTimeout = time.Minute
)

type SystemdClient struct {
}

// NewSystemdClient creates a new systemd client
func NewSystemdClient() *SystemdClient {
	return &SystemdClient{}
}

// Reload - systemctl daemon-reload
func (s *SystemdClient) Reload() error {
	conn, err := dbus.New()
	if err != nil {
		return err
	}

	if err = conn.Reload(); err != nil {
		return err
	}

	return nil
}

// Start -  systemctl start <unit>
func (s *SystemdClient) Start(unit string) error {
	conn, err := dbus.New()
	if err != nil {
		return err
	}

	responseChan := make(chan string, 1)
	if _, err = conn.StartUnitContext(context.Background(), unit, "replace", responseChan); err != nil {
		return err
	}

	select {
	case res := <-responseChan:
		switch res {
		case "done":
			return nil
		case "failed":
			return nil
		case "canceled", "timeout", "dependency", "skipped":
			return fmt.Errorf("%v", res)
		default:
			return fmt.Errorf("%v", res)
		}
	case <-time.After(jobTimeout):
		return fmt.Errorf("%v", "job timedout")
	}

	return nil
}

// Stop - systemctl stop <unit>
func (s *SystemdClient) Stop(unit string) error {

	conn, err := dbus.New()
	if err != nil {
		return err
	}

	responseChan := make(chan string, 1)
	if _, err = conn.StopUnitContext(context.Background(), unit, "replace", responseChan); err != nil {
		return err
	}

	select {
	case res := <-responseChan:
		switch res {
		case "done":
			return nil
		case "canceled":
			return nil
		case "timeout", "failed", "dependency", "skipped":
			return fmt.Errorf("%v", res)
		default:
			return fmt.Errorf("%v", res)
		}
	case <-time.After(jobTimeout):
		return fmt.Errorf("%v", "job timeout")
	}

	return nil
}

// Enable - systemctl enable <unit>
func (s *SystemdClient) Enable(unit string) error {

	conn, err := dbus.New()
	if err != nil {
		return err
	}

	if _, _, err = conn.EnableUnitFilesContext(context.Background(), []string{unit}, false, false); err != nil {
		return err
	}

	return nil
}

// Exists - checks if systemd <unit> file exists
func (s *SystemdClient) Exists(unit string) (bool, error) {
	conn, err := dbus.New()
	if err != nil {
		return false, err
	}

	unitStates, err := conn.ListUnitsContext(context.Background())
	if err != nil {
		return false, err
	}

	for _, unitState := range unitStates {
		if unitState.Name == unit {
			return true, nil
		}
	}

	return false, nil
}

// IsActive - checks if <unit> is active
func (s *SystemdClient) IsActive(unit string) (bool, error) {
	conn, err := dbus.New()
	if err != nil {
		return false, err
	}

	unitStates, err := conn.ListUnitsContext(context.Background())
	if err != nil {
		return false, err
	}

	for _, unitState := range unitStates {
		if unitState.Name == unit {
			return unitState.ActiveState == "active", nil
		}
	}

	return false, nil
}
