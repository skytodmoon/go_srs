/*
The MIT License (MIT)

Copyright (c) 2013-2015 GOSRS(gosrs)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package app

import (
	"go_srs/srs/app/config"
	"go_srs/srs/protocol/rtmp"
	"fmt"
)

type SrsDvrPlan interface {
	OnPublish() error
	OnUnpublish() error
	OnMetaData(metaData *rtmp.SrsRtmpMessage) error
	OnVideo(video *rtmp.SrsRtmpMessage) error
	OnAudio(audio *rtmp.SrsRtmpMessage) error
}

func NewSrsDvrPlan(req *SrsRequest) SrsDvrPlan {
	fmt.Println("xxxxxxxxxxxxx vhost=", req.vhost)
	dvrPlan := config.GetDvrPlan(req.vhost)
	if dvrPlan == "session" {
		return NewSrsSessionDvrPlan(req)
	}
	return nil
}

type SrsSessionDvrPlan struct {
	segment *SrsFlvSegment
}

func NewSrsSessionDvrPlan(req *SrsRequest) *SrsSessionDvrPlan {
	return &SrsSessionDvrPlan{
		segment:NewSrsFlvSegment(req),
	}
}

func (this *SrsSessionDvrPlan) OnPublish() error {
	if err := this.segment.Open(true); err != nil {
		return err
	}
	return nil
}

func (this *SrsSessionDvrPlan) OnUnpublish() error {
	if err := this.segment.Close(); err != nil {
		return err
	}
	return nil
}

func (this *SrsSessionDvrPlan) OnMetaData(metaData *rtmp.SrsRtmpMessage) error {
	fmt.Println("===================================onMetadata")
	return this.segment.WriteMetaData(metaData)
}

func (this *SrsSessionDvrPlan) OnVideo(video *rtmp.SrsRtmpMessage) error {
	return this.segment.WriteVideo(video)
}

func (this *SrsSessionDvrPlan) OnAudio(audio *rtmp.SrsRtmpMessage) error {
	return this.segment.WriteAudio(audio)
}

//
//type SrsDvrPlan struct {
//	segment *SrsFlvSegment
//}
//
//func NewSrsDvrPlan(fname string) *SrsDvrPlan {
//	return &SrsDvrPlan{
//		segment: NewSrsFlvSegment(fname),
//	}
//}
//
//func (this *SrsDvrPlan) Initialize() {
//	this.segment.Initialize()
//}
//
//func (this *SrsDvrPlan) On_video(msg *rtmp.SrsRtmpMessage) error {
//	this.segment.WriteVideo(msg)
//	return nil
//}
//
//func (this *SrsDvrPlan) On_audio(msg *rtmp.SrsRtmpMessage) error {
//	this.segment.WriteAudio(msg)
//	return nil
//}
//
//func (this *SrsDvrPlan) OnMetaData(msg *rtmp.SrsRtmpMessage) error {
//	err := this.segment.WriteMetaData(msg)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (this *SrsDvrPlan) Close() {
//	this.segment.Close()
//}
