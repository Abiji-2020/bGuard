# Installation


## Prepare your configuration

bGuard supports single or multiple YAML files as configuration. Create new `config.yml` with your configuration
(see [Configuration](configuration.md) for more details and all configuration options).

Simple configuration file, which enables only basic features:

```yaml
upstream:
  default:
    - 46.182.19.48
    - 80.241.218.68
    - tcp-tls:fdns1.dismail.de:853
    - https://dns.digitale-gesellschaft.ch/dns-query
blocking:
  denylists:
    ads:
      - https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts
  clientGroupsBlock:
    default:
      - ads
ports:
  dns: 53
  http: 4000
```

Add smilar to this in the configration and clone this repo and run ```go run main.go```
