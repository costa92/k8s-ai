#!/usr/bin/env bash

# Run commands requiring root privileges without entering a password.
function k8sai::util::sudo()
{
  echo ${LINUX_PASSWORD} | sudo -S $1
}
