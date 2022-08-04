# go-utils

Various Open Source utility functions

## Subdirectory: cadence

Common interface code for setting Cadence client and Cadence workers

## Subdirectory: s3client

Simple client for uploading Athena compatible JSON records to S3
in with a path formated as: `bucket/folder/YYYY/MM/DD/HH/PREFIX-TIMESTAMP-UUID`

## Subdirectory: types

sqlTypes package - types to interface to SQL types

- TagSet - maintain unique tags in SQL array

dict package - model an object/list from generic JSON decoding

- T - the generic obj/list type
