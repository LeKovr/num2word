
num2word
========

[![GoDoc][gd1]][gd2]
 [![codecov][cc1]][cc2]
 [![Build Status][bs1]][bs2]
 [![GoCard][gc1]][gc2]
 [![GitHub Release][gr1]][gr2]
 [![GitHub license][gl1]][gl2]

[bs1]: https://cloud.drone.io/api/badges/LeKovr/num2word/status.svg
[bs2]: https://cloud.drone.io/LeKovr/num2word
[cc1]: https://codecov.io/gh/LeKovr/num2word/branch/master/graph/badge.svg
[cc2]: https://codecov.io/gh/LeKovr/num2word
[gd1]: https://godoc.org/github.com/LeKovr/num2word?status.svg
[gd2]: https://godoc.org/github.com/LeKovr/num2word
[gc1]: https://goreportcard.com/badge/github.com/LeKovr/num2word
[gc2]: https://goreportcard.com/report/github.com/LeKovr/num2word
[gr1]: https://img.shields.io/github/v/release/LeKovr/num2word.svg
[gr2]: https://github.com/LeKovr/num2word/releases
[gl1]: https://img.shields.io/github/license/LeKovr/num2word.svg
[gl2]: LICENSE

[num2word](https://github.com/LeKovr/num2word) - Числа прописью

golang package для написания сумм прописью.

Features
--------

В текущей версии реализовано:

### RuMoney - деньги прописью на русском

* Оригинальное название: number2word
* Источник на SQL: http://oraclub.trecom.tomsk.su/db/web.page?pid=461
* В конференцию relcom.comp.dbms.oracle поместил "Igor Volkov" (volkov@rdtex.msk.ru)

Usage
-----

```
import (
  "github.com/LeKovr/num2word"
)

...

text := num2word.RusAmount(total, true)
```

License
-------

The MIT License (MIT), see [LICENSE](LICENSE).

Copyright (c) 2016 Alexey Kovrizhkin ak@elfire.ru
