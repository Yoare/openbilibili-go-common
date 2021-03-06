// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathGiftRoomGiftList = "/live.xgift.v1.Gift/room_gift_list"
var PathGiftGiftConfig = "/live.xgift.v1.Gift/gift_config"
var PathGiftDiscountGiftList = "/live.xgift.v1.Gift/discount_gift_list"
var PathGiftDailyBag = "/live.xgift.v1.Gift/daily_bag"

// ==============
// Gift Interface
// ==============

type GiftBMServer interface {
	RoomGiftList(ctx context.Context, req *RoomGiftListReq) (resp *RoomGiftListResp, err error)

	GiftConfig(ctx context.Context, req *GiftConfigReq) (resp *GiftConfigResp, err error)

	DiscountGiftList(ctx context.Context, req *DiscountGiftListReq) (resp *DiscountGiftListResp, err error)

	DailyBag(ctx context.Context, req *DailyBagReq) (resp *DailyBagResp, err error)
}

var v1GiftSvc GiftBMServer

func giftRoomGiftList(c *bm.Context) {
	p := new(RoomGiftListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1GiftSvc.RoomGiftList(c, p)
	c.JSON(resp, err)
}

func giftGiftConfig(c *bm.Context) {
	p := new(GiftConfigReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1GiftSvc.GiftConfig(c, p)
	c.JSON(resp, err)
}

func giftDiscountGiftList(c *bm.Context) {
	p := new(DiscountGiftListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1GiftSvc.DiscountGiftList(c, p)
	c.JSON(resp, err)
}

func giftDailyBag(c *bm.Context) {
	p := new(DailyBagReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1GiftSvc.DailyBag(c, p)
	c.JSON(resp, err)
}

// RegisterGiftBMServer Register the blademaster route
func RegisterGiftBMServer(e *bm.Engine, server GiftBMServer) {
	e.GET("/live.xgift.v1.Gift/room_gift_list", giftRoomGiftList)
	e.GET("/live.xgift.v1.Gift/gift_config", giftGiftConfig)
	e.GET("/live.xgift.v1.Gift/discount_gift_list", giftDiscountGiftList)
	e.GET("/live.xgift.v1.Gift/daily_bag", giftDailyBag)
}
