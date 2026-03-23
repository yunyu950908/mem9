# CI/CD Workflows

This is a fork of [mem9-ai/mem9](https://github.com/mem9-ai/mem9). Upstream CI workflows (`deploy-dev.yml`, `sync-claude-plugin.yml`) have been removed as they depend on upstream-specific secrets and infrastructure.

## Workflows

| Workflow | Trigger | Description |
|---|---|---|
| `sync-fork.yml` | Daily (UTC 08:00) / Manual | Syncs `main` branch with upstream via GitHub merge-upstream API. Automatically triggers `build-swr` when new commits are synced. |
| `build-swr.yml` | Manual / Auto (via sync-fork) | Builds `mnemo-server` Docker image using `make docker` and pushes to Huawei Cloud SWR with `SHORT_SHA` and `latest` tags. |

## Required Secrets

| Secret | Description |
|---|---|
| `SWR_REGISTRY` | SWR endpoint, e.g. `swr.cn-north-4.myhuaweicloud.com` |
| `SWR_ORGANIZATION` | SWR organization / namespace |
| `SWR_USERNAME` | SWR login username (typically `region@AK`) |
| `SWR_PASSWORD` | SWR login password or temporary token |
