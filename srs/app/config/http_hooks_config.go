/*
The MIT License (MIT)

Copyright (c) 2019 GOSRS(gosrs)

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

package config

type HttpHooksConf struct {
	Enabled     string `json:"enabled"`
	OnConnect   string `json:"on_connect"`
	OnClose     string `json:"on_close"`
	OnPublish   string `json:"on_publish"`
	OnUnpublish string `json:"on_unpublish"`
	OnPlay      string `json:"on_play"`
	OnStop      string `json:"on_stop"`
	OnDvr       string `json:"on_dvr"`
	OnHls       string `json:"on_hls"`
	OnHlsNotify string `json:"on_hls_notify"`
}

func (this *HttpHooksConf) initDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}
}
