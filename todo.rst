## Todo
libase/api/tds -> Sybase TDS Protokoll
   Sybase TDS 5.0 implementation

libase/api -> Sybase API via TDS

go/ -> go driver, wrapping libase/api

#####
- money/money4 with cs_convert

- info on null values, clientlibrary seemingly doesn't have a way to
  signal a null value in ASE to the user (aside from using default
  values, which would be valid values in turn)

- info on empty values for text, unichar, univarchar, ...
  Empty values for these columns are being passed to ASE as NULL instead
  of empty

- Decimal/Numeric:
   1. Write shim that only transports information
   2. Use a complete implementation (e.g. github.com/shopspring/decimal MIT)

- Todo: columnconverter on statement
