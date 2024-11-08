# Syntax

- casing does not matter.
- search modifiers have to respect the case sensitivity.
- search values does not need to respect the case sensitivity.
- dates are stored as int64 (unix timestamps)
- dates are in ISO format: 2023-09-12T14:30:00

## Example:

- type=pe and tag=upx // and is optional, by default, we always AND 2 search sub-expressions
- type=pe or tag=upx
- ( type=pe or tag=upx ) and avast!=locky
- size >= 1000000 // by default, it's bytes
- size < 1000KB
- size > 1MB
- fs >= 2009
- fs <= 2020-12
- fs <= 2020-01-30
- fs < 2012-08-21T1
- fs < 2012-08-21T16:59
- fs < 2012-08-21T16:59:20 // UTC
- fs < 2012-08-21T16:59:20Z // UTC explicit
- fs < 2012-08-21T16:59:20+02:00 // 2 hours ahead of UTC
- fs < 3d
- ls has the same syntax as `fs`.

## Search Modifiers

- size (int)
- type (string) enum { pe, elf, macho, txt}
- extension (string) enum { dll, exe, ps1 }
- name: (string)
- trid: (array)
- packer: (array)
- magic: (string)
- tag: (string)
- fs (first seen): (string date)
- ls (last scanned): (string date)
- positives: int (max = 14)
- crc32: (string)
- engines (string), search in all AntiVirus
- avast: (string)
- avira: (string)
- bitdefender: (string)
- clamav: (string)
- comodo: (string)
- drweb: (string)
- eset: (string)
- fsecure: (string)
- kaspersky: (string)
- mcafee: (string)
- sophos: (string)
- symantec: (string)
- trendmicro: (string)
- windefender: (string)
