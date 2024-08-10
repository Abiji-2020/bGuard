# Interfaces

## REST API


??? abstract "OpenAPI specification"

    ```yaml
    --8<-- "docs/api/openapi.yaml"
    ```

If http listener is enabled, bGuard provides REST API. You can download the [OpenAPI YAML](api/openapi.yaml) interface specification. 

You can also browse the interactive API documentation (RapiDoc) documentation [online](rapidoc.html).

## CLI

bGuard provides a CLI interface to control. This interface uses internally the REST API.

To run the CLI, please ensure, that bGuard DNS server is running, then execute `bGuard help` for help or

- `./bGuard blocking enable` to enable blocking
- `./bGuard blocking disable` to disable blocking
- `./bGuard blocking disable --duration [duration]` to disable blocking for a certain amount of time (30s, 5m, 10m30s,
  ...)
- `./bGuard blocking disable --groups ads,othergroup` to disable blocking only for special groups
- `./bGuard blocking status` to print current status of blocking
- `./bGuard query <domain>` execute DNS query (A) (simple replacement for dig, useful for debug purposes)
- `./bGuard query <domain> --type <queryType>` execute DNS query with passed query type (A, AAAA, MX, ...)
- `./bGuard lists refresh` reloads all allow/denylists
- `./bGuard validate [--config /path/to/config.yaml]` validates configuration file

!!! tip 

    To run this inside docker run `docker exec bGuard ./bGuard blocking status`

--8<-- "docs/includes/abbreviations.md"
