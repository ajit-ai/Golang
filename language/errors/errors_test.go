package main

import (
	"errors"
	"testing"
)

func TestFindUserWrapsSentinel(t *testing.T) {
	name, err := FindUser(1)
	if err != nil || name != "ajit" {
		t.Fatalf("FindUser(1) = %q, %v", name, err)
	}
	_, err = FindUser(99)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("errors.Is(err, ErrNotFound) = false for %v", err)
	}
	if errors.Is(err, ErrForbidden) {
		t.Error("wrong sentinel matched")
	}
}

func TestValidateAgeTypedError(t *testing.T) {
	if err := ValidateAge(30); err != nil {
		t.Fatalf("ValidateAge(30) = %v, want nil", err)
	}
	err := ValidateAge(-5)
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("errors.As failed for %v", err)
	}
	if ve.Field != "age" || ve.Msg != "cannot be negative" {
		t.Errorf("ValidationError = %+v", ve)
	}
	wrapped := ValidateAge(200)
	if !errors.As(wrapped, &ve) {
		t.Error("wrapped ValidationError not detected through wrapping")
	}
}

func TestCheckPermissionWrapsTwice(t *testing.T) {
	if err := CheckPermission("admin"); err != nil {
		t.Fatalf("admin should pass: %v", err)
	}
	err := CheckPermission("guest")
	if !errors.Is(err, ErrForbidden) {
		t.Errorf("errors.Is failed through two wrap levels: %v", err)
	}
}

func TestSafeCallRecoversPanic(t *testing.T) {
	v, err := SafeCall(func() int { return 42 })
	if err != nil || v != 42 {
		t.Fatalf("SafeCall ok path = %d, %v", v, err)
	}
	v, err = SafeCall(func() int { panic("boom") })
	if err == nil || v != 0 {
		t.Fatalf("panic not converted to error: v=%d err=%v", v, err)
	}
}

func TestErrorsMainSmoke(t *testing.T) {
	CustomErrorsMain()
	RecoverMain()
}
