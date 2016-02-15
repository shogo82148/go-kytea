#!/bin/bash

base_dir=`pwd`

wget http://www.phontron.com/kytea/download/kytea-0.4.7.tar.gz
tar zxfv kytea-0.4.7.tar.gz
cd kytea-0.4.7
./configure
make
sudo make install
sudo ldconfig

cd $base_dir
rm -rf kytea-0.4.7
