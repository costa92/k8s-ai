#!/usr/bin/env bash

# Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/k8s-ai.

cmd=$1
shell=$2

function add() 
{
  cat << EOF >> ${HOME}/.bashrc

# ${cmd} shell completion 
if [ -f \${HOME}/.${cmd}-completion.bash ]; then 
    source \${HOME}/.${cmd}-completion.bash
fi 
EOF
}

${cmd} completion ${shell} > ${HOME}/.${cmd}-completion.bash

if ! grep -q "# ${cmd} shell completion" ${HOME}/.bashrc;then
  add
fi