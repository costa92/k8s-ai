# Makefile 命令参考文档

本文档详细说明了项目中所有可用的 Makefile 命令及其用途。

## 基础命令

| 命令 | 说明 | 示例 |
|------|------|------|
| **build** | 构建项目源代码 | `make build` |
| **tidy** | 整理并更新 Go 依赖 | `make tidy` |
| **help** | 显示帮助信息 | `make help` |

## 工具安装命令

| 命令 | 说明 | 示例 |
|------|------|------|
| **install-tools** | 安装所有开发工具 | `make install-tools` |
| **tools.install.%** | 安装指定工具 | `make tools.install.golangci-lint` |
| **tools.verify.%** | 验证工具安装 | `make tools.verify.uplift` |

## 发布相关命令

| 命令 | 说明 | 示例 |
|------|------|------|
| **release** | 执行完整发布流程 | `make release` |
| **release.run** | 生成 CHANGELOG 并打标签 | `make release.run` |

## 开发工具集

| 工具 | 用途 | 安装命令 |
|------|------|----------|
| **golangci-lint** | Go 代码质量检查 | `make _install.golangci-lint` |
| **gotests** | 自动生成单元测试 | `make _install.gotests` |
| **wire** | 依赖注入工具 | `make _install.wire` |
| **mockgen** | Mock 代码生成器 | `make _install.mockgen` |
| **goimports** | Go 导入语句格式化 | `make _install.goimports` |
| **git-chglog** | CHANGELOG 生成工具 | `make _install.git-chglog` |
| **uplift** | 版本发布管理工具 | `make _install.uplift` |
| **license** | LICENSE 文件生成工具 | `make _install.license` |
| **gsemver** | 语义版本管理工具 | `make _install.gsemver` |

## CI/CD 工具

| 命令 | 说明 |
|------|------|
| **_install.ci** | 安装 CI/CD 必需工具 |
| **_install.code-generator** | 安装代码生成工具 |
| **tools.verify.code-generator** | 验证代码生成器 |

### 示例用法
