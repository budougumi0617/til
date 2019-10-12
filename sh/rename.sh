#!/bin/bash -x
for f in *
do
mv $f `echo $f | tr ‘[A-Z]’ ‘[a-z]’`
done
