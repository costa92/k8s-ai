#!/usr/bin/env bash

# Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/k8s-ai.


# Run commands requiring root privileges without entering a password.
function k8sai::util::sudo()
{
  echo ${LINUX_PASSWORD} | sudo -S $1
}
