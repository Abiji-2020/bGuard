package resolver

import (
	"context"
	"strings"

	"github.com/Abiji-2020/bGuard/config"
	"github.com/Abiji-2020/bGuard/model"
	"github.com/Abiji-2020/bGuard/util"
	"github.com/miekg/dns"
)

type FQDNOnlyResolver struct {
	configurable[*config.FQDNOnly]
	NextResolver
	typed
}

func NewFQDNOnlyResolver(cfg config.FQDNOnly) *FQDNOnlyResolver {
	return &FQDNOnlyResolver{
		configurable: withConfig(&cfg),
		typed:        withType("fqdn_only"),
	}
}

func (r *FQDNOnlyResolver) Resolve(ctx context.Context, request *model.Request) (*model.Response, error) {
	if r.IsEnabled() {
		domainFromQuestion := util.ExtractDomain(request.Req.Question[0])
		if !strings.Contains(domainFromQuestion, ".") {
			response := new(dns.Msg)
			response.Rcode = dns.RcodeNameError

			return &model.Response{Res: response, RType: model.ResponseTypeNOTFQDN, Reason: "NOTFQDN"}, nil
		}
	}

	return r.next.Resolve(ctx, request)
}
