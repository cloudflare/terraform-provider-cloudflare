package cloudflare

func expandInterfaceToStringList(list interface{}) []string {
	ifaceList := list.([]interface{})
	vs := make([]string, 0, len(ifaceList))
	for _, v := range ifaceList {
		vs = append(vs, v.(string))
	}
	return vs
}

func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}
