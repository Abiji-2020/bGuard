# Integration in Grafana

## Prometheus

### Prometheus export

bGuard can optionally export metrics for [Prometheus](https://prometheus.io/).

Following metrics will be exported:

| name                                             |   Description                                            |
| ------------------------------------------------ | -------------------------------------------------------- |
| bGuard_denylist_cache / bGuard_allowlist_cache   | Number of entries in denylist/allowlist cache, partitioned by group |
| bGuard_error_total                | Number of total queries that ended in error for any reason |
| bGuard_query_total                | Number of total queries, partitioned by client and DNS request type (A, AAAA, PTR, etc) |
| bGuard_request_duration_ms_bucket | Request duration histogram, partitioned by response type (Blocked, cached, etc)  |
| bGuard_response_total             | Number of responses, partitioned by response type (Blocked, cached, etc), DNS response code, and reason |
| bGuard_blocking_enabled           | 1 if blocking is enabled, 0 otherwise |
| bGuard_cache_entry_count          | Number of entries in cache |
| bGuard_cache_hit_count / bGuard_cache_miss_count | Cache hit/miss counters |
| bGuard_prefetch_count | Amount of prefetched DNS responses |
| bGuard_prefetch_domain_name_cache_count | Amount of domain names being prefetched |
| bGuard_failed_download_count      | Number of failed list downloads |

### Grafana dashboard

Example [Grafana](https://grafana.com/) dashboard
definition [as JSON](bGuard-grafana.json)
or [at grafana.com](https://grafana.com/grafana/dashboards/13768)
![grafana-dashboard](grafana-dashboard.png).

This dashboard shows all relevant statistics and allows enabling and disabling the blocking status.

### Grafana configuration

Please install `grafana-piechart-panel` and
set [disable_sanitize_html](https://grafana.com/docs/grafana/latest/installation/configuration/#disable_sanitize_html)
in config or as env to use control buttons to enable/disable the blocking status.

### Grafana and Prometheus example project

This [repo](https://github.com/Abiji-2020/bGuard-grafana-prometheus-example) contains example docker-compose.yml with
bGuard, prometheus (with configured scraper for bGuard) and grafana with prometheus datasource.

## MySQL / MariaDB

If database query logging is activated (see [Query logging](configuration.md#query-logging)), you can use following
Grafana Dashboard [as JSON](bGuard-query-grafana.json)
or [at grafana.com](https://grafana.com/grafana/dashboards/14980)

![grafana-dashboard](grafana-query-dashboard.png).

Please define the MySQL source in Grafana, which points to the database with bGuard's log entries.

## Postgres

The JSON for a Grafana dashboard equivalent to the MySQL/MariaDB version is located [here](bGuard-query-grafana-postgres.json)

--8<-- "docs/includes/abbreviations.md"
