package main

import "testing"

var _ DecoratorIProcess = &ProcessClass{}
var _ DecoratorIProcess = &ProcessDecorator{}

func TestProcessDecoratorWithoutInstance(t *testing.T) {
	decorator := &ProcessDecorator{}
	decorator.process()
}

func TestProcessDecoratorWithInstance(t *testing.T) {
	decorator := &ProcessDecorator{processInstance: &ProcessClass{}}
	decorator.process()
}

func TestDecoratorMain(t *testing.T) {
	DecoratorMain()
}
