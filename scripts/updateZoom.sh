#!/bin/sh
################################################################
# 'updateZoom.sh'                                              #
#                                                              #
# by David Skinner                                             #
# on May 9, 2020                                               #
# for Davsk hugosite                                           #
#                                                              #
# Copyright © 2020. Davsk Ltd Co                               #
# All Rights Reserved. Licensed under the MIT License.         #
################################################################

################################################################################
# Permission is hereby granted, free of charge, to any person obtaining a copy #
# of this software and associated documentation files (the "Software"),        #
# to deal in the Software without restriction, including without limitation    #
# the rights to use, copy, modify, merge, publish, distribute, sublicense,     #
# and/or sell copies of the Software, and to permit persons to whom the        #
# Software is furnished to do so, subject to the following conditions:         #
#                                                                              #
# The above copyright notice and this permission notice shall be included in   #
# all copies or substantial portions of the Software.                          #
#                                                                              #
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS      #
# OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,  #
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE  #
#A UTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER       #
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING      #
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER          #
# DEALINGS IN THE SOFTWARE.                                                    #
################################################################################

################################################################
# Start Install/Update
echo 'updateZoom.sh'
set ZOOM_DIR = /opt/zoom

# Deps
sudo eopkg upgrade
#sudo eopkg install

# Download
cd ~/Downloads || exit
mkdir Zoom && cd Zoom
curl -LO https://zoom.us/client/latest/zoom_x86_64.tar.xz
if [ ${#} -eq 0 ]; then
  if [ -d "$ZOOM_DIR" ]; then
    rm -Rf $ZOOM_DIR
  fi

  sudo tar -xvf zoom*.tar.xz -C $ZOOM_DIR
  cd $ZOOM_DIR
  ./ZoomLaunchertrue
else
  echo "Download failed!"
  exit # curl failed
fi
################################################################
