package helper

import (
	"fmt"
	"github.com/Jacob953/behavior3gen/internal/generate"
	"strings"
)

// GenerateJSON 生成 NodeMeta 结构体数组的 JSON 字符串，支持缩进和前缀
func GenerateJSON(nodes []*generate.NodeMeta, indent string, prefix string) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, node := range nodes {
		if i > 0 {
			sb.WriteString(",") // 添加逗号分隔
		}
		sb.WriteString(fmt.Sprintf("\n%s%s{", prefix, indent))
		pre := fmt.Sprintf("\n%s%s%s", prefix, indent, indent)
		sb.WriteString(fmt.Sprintf("%s\"version\": \"%s\",", pre, node.Version))
		sb.WriteString(fmt.Sprintf("%s\"scope\": \"%s\",", pre, node.Scope))
		sb.WriteString(fmt.Sprintf("%s\"name\": \"%s\",", pre, node.Name))
		sb.WriteString(fmt.Sprintf("%s\"category\": \"%s\",", pre, node.Category))
		sb.WriteString(fmt.Sprintf("%s\"title\": \"%s\",", pre, node.Title))
		sb.WriteString(fmt.Sprintf("%s\"description\": \"%s\",", pre, node.Description))

		// 处理 Properties 字段
		sb.WriteString(fmt.Sprintf(`%s"properties": {`, pre))
		propCount := 0
		secondPre := fmt.Sprintf("%s%s", pre, indent)
		for key, value := range node.Properties {
			if propCount > 0 {
				sb.WriteString(",") // 添加逗号分隔
			}
			sb.WriteString(fmt.Sprintf("%s\"%s\": \"%v\"", secondPre, key, value))
			propCount++
		}
		sb.WriteString(fmt.Sprintf("%s}", pre))

		sb.WriteString(fmt.Sprintf("\n%s%s}", prefix, indent))
	}

	sb.WriteString("\n]")
	return sb.String()
}
