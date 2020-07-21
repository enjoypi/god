package service

type supervisor struct {
	Actor
	children map[uint32]*Actor
}