package main

import "testing"

var _ DataTransferObject = FlyweightCustomer{}
var _ DataTransferObject = Employee{}
var _ DataTransferObject = Manager{}
var _ DataTransferObject = Address{}

func TestGetDataTransferObject(t *testing.T) {
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	tests := []struct {
		dtoType string
		wantID  string
	}{
		{"customer", "1"},
		{"employee", "2"},
		{"manager", "3"},
		{"address", "4"},
	}
	for _, tt := range tests {
		t.Run(tt.dtoType, func(t *testing.T) {
			dto := factory.getDataTransferObject(tt.dtoType)
			if dto == nil {
				t.Fatalf("getDataTransferObject(%q) = nil, want non-nil", tt.dtoType)
			}
			if got := dto.getId(); got != tt.wantID {
				t.Errorf("getId() = %q, want %q", got, tt.wantID)
			}
		})
	}
}

func TestGetDataTransferObjectReusesInstances(t *testing.T) {
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	first := factory.getDataTransferObject("customer")
	second := factory.getDataTransferObject("customer")
	if len(factory.pool) != 1 {
		t.Errorf("pool size = %d, want 1 (flyweight must reuse instances)", len(factory.pool))
	}
	if first.getId() != second.getId() {
		t.Errorf("ids differ across calls: %q vs %q", first.getId(), second.getId())
	}
}

func TestGetDataTransferObjectUnknownType(t *testing.T) {
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	if dto := factory.getDataTransferObject("unknown"); dto != nil {
		t.Errorf("getDataTransferObject(%q) = %v, want nil", "unknown", dto)
	}
}

func TestFlyweightMain(t *testing.T) {
	FlyweightMain()
}
