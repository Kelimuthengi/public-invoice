package main

// type Engine struct {
// 	trustedProxies []string
// }

// func (engine *Engine) SetTrustedProxies(trustedProxies []string) error {
// 	engine.trustedProxies = trustedProxies

// 	// return engine.ParseTrustedProxies()
// 	return nil
// }

// func (engine *Engine) ParseTrustedProxies() error {

// 	for _, proxy := range engine.trustedProxies {
// 		if _, _, err := net.ParseCIDR(proxy); err != nil {
// 			if net.ParseIP(proxy) == nil {
// 				return errors.New("invalid proxy address: " + proxy)
// 			}
// 		}
// 	}
// 	return nil
// }