# NFC-e Compact

This program move and compact all `-nfe` files to send your accounting

#### how to use

* `nfcecompact -path="/tmp/nfetest"` or `nfcecompact -path="/tmp/nfetest" -year=2018 -month=10`

After run this command a file will be generate on path defined with year and month folders and a `.zip` file will be generated too.

ex:
```
/tmp/nfetest
    |_/2018
    |    |_/9
    |_/2018-9.zip
```
