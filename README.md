
num2word
========

[![GoDoc][1]][2]
[![GoCard][3]][4]

[1]: https://godoc.org/github.com/LeKovr/num2word?status.svg
[2]: https://godoc.org/github.com/LeKovr/num2word
[3]: https://goreportcard.com/badge/LeKovr/dbcc
[4]: https://goreportcard.com/report/github.com/LeKovr/dbcc

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
