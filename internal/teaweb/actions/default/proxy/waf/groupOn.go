package waf

import (
	"github.com/TeaWeb/build/internal/teaconfigs"
	"github.com/TeaWeb/build/internal/teaweb/actions/default/proxy/proxyutils"
	"github.com/TeaWeb/build/internal/teaweb/actions/default/proxy/waf/wafutils"
	"github.com/iwind/TeaGo/actions"
)

type GroupOnAction actions.Action

//启用分组
func (this *GroupOnAction) RunPost(params struct {
	WafId   string
	GroupId string
}) {
	wafList := teaconfigs.SharedWAFList()
	waf := wafList.FindWAF(params.WafId)
	if waf == nil {
		this.Fail("找不到WAF")
	}

	group := waf.FindRuleGroup(params.GroupId)
	if group == nil {
		this.Fail("找不到分组")
	}

	group.On = true
	err := wafList.SaveWAF(waf)
	if err != nil {
		this.Fail("保存失败：" + err.Error())
	}

	// 通知刷新
	if wafutils.IsPolicyUsed(waf.Id) {
		proxyutils.NotifyChange()
	}

	this.Success()
}
