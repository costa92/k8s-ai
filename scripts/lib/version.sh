#!/usr/bin/env bash


# Grovels through git to set a set of env variables.
#
# If PROJ_GIT_VERSION_FILE, this function will load from that file instead of
# querying git.
proj::version::get_version_vars() {
  if [[ -n ${PROJ_GIT_VERSION_FILE-} ]]; then
    proj::version::load_version_vars "${PROJ_GIT_VERSION_FILE}"
    return
  fi

  # If the proj source was exported through git archive, then
  # we likely don't have a git tree, but these magic values may be filled in.
  # shellcheck disable=SC2016,SC2050
  # Disabled as we're not expanding these at runtime, but rather expecting
  # that another tool may have expanded these and rewritten the source (!)
  if [[ '$Format:%%$' == "%" ]]; then
    PROJ_GIT_COMMIT='$Format:%H$'
    PROJ_GIT_TREE_STATE="archive"
    # When a 'git archive' is exported, the '$Format:%D$' below will look
    # something like 'HEAD -> release-1.8, tag: v1.8.3' where then 'tag: '
    # can be extracted from it.
    if [[ '$Format:%D$' =~ tag:\ (v[^ ,]+) ]]; then
     PROJ_GIT_VERSION="${BASH_REMATCH[1]}"
    fi
  fi

  local git=(git --work-tree "${PROJ_ROOT_DIR}")

  if [[ -n ${PROJ_GIT_COMMIT-} ]] || PROJ_GIT_COMMIT=$("${git[@]}" rev-parse "HEAD^{commit}" 2>/dev/null); then
    if [[ -z ${PROJ_GIT_TREE_STATE-} ]]; then
      # Check if the tree is dirty.  default to dirty
      if git_status=$("${git[@]}" status --porcelain 2>/dev/null) && [[ -z ${git_status} ]]; then
        PROJ_GIT_TREE_STATE="clean"
      else
        PROJ_GIT_TREE_STATE="dirty"
      fi
    fi

    # Use git describe to find the version based on tags.
    if [[ -n ${PROJ_GIT_VERSION-} ]] || PROJ_GIT_VERSION=$("${git[@]}" describe --tags --match='v*' --abbrev=14 "${PROJ_GIT_COMMIT}^{commit}" 2>/dev/null); then
      # This translates the "git describe" to an actual semver.org
      # compatible semantic version that looks something like this:
      #   v1.1.0-alpha.0.6+84c76d1142ea4d
      #
      # TODO: We continue calling this "git version" because so many
      # downstream consumers are expecting it there.
      #
      # These regexes are painful enough in sed...
      # We don't want to do them in pure shell, so disable SC2001
      # shellcheck disable=SC2001
      DASHES_IN_VERSION=$(echo "${PROJ_GIT_VERSION}" | sed "s/[^-]//g")
      if [[ "${DASHES_IN_VERSION}" == "---" ]] ; then
        # shellcheck disable=SC2001
        # We have distance to subversion (v1.1.0-subversion-1-gCommitHash)
        PROJ_GIT_VERSION=$(echo "${PROJ_GIT_VERSION}" | sed "s/-\([0-9]\{1,\}\)-g\([0-9a-f]\{14\}\)$/.\1\+\2/")
      elif [[ "${DASHES_IN_VERSION}" == "--" ]] ; then
        # shellcheck disable=SC2001
        # We have distance to base tag (v1.1.0-1-gCommitHash)
        PROJ_GIT_VERSION=$(echo "${PROJ_GIT_VERSION}" | sed "s/-g\([0-9a-f]\{14\}\)$/+\1/")
      fi
      # We don't want to be inadvertently added a -dirty suffix
      # TODO: fix bug mentioned in comments
      #if [[ "${PROJ_GIT_TREE_STATE}" == "dirty" ]]; then
        # git describe --dirty only considers changes to existing files, but
        # that is problematic since new untracked .go files affect the build,
        # so use our idea of "dirty" from git status instead.
        #PROJ_GIT_VERSION+="-dirty"
      #fi


      # Try to match the "git describe" output to a regex to try to extract
      # the "major" and "minor" versions and whether this is the exact tagged
      # version or whether the tree is between two tagged versions.
      if [[ "${PROJ_GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?([-].*)?([+].*)?$ ]]; then
        PROJ_GIT_MAJOR=${BASH_REMATCH[1]}
        PROJ_GIT_MINOR=${BASH_REMATCH[2]}
        if [[ -n "${BASH_REMATCH[4]}" ]]; then
          PROJ_GIT_MINOR+="+"
        fi
      fi

      # If PROJ_GIT_VERSION is not a valid Semantic Version, then refuse to build.
      if ! [[ "${PROJ_GIT_VERSION}" =~ ^v([0-9]+)\.([0-9]+)(\.[0-9]+)?(-[0-9A-Za-z.-]+)?(\+[0-9A-Za-z.-]+)?$ ]]; then
          proj::log::error "PROJ_GIT_VERSION should be a valid Semantic Version. Current value: ${PROJ_GIT_VERSION}"
          proj::log::error "Please see more details here: https://semver.org"
          exit 1
      fi
    fi
  fi
}

# Saves the environment flags to $1
proj::version::save_version_vars() {
  local version_file=${1-}
  [[ -n ${version_file} ]] || {
    echo "!!! Internal error.  No file specified in proj::version::save_version_vars"
    return 1
  }

  cat <<EOF >"${version_file}"
PROJ_GIT_COMMIT='${PROJ_GIT_COMMIT-}'
PROJ_GIT_TREE_STATE='${PROJ_GIT_TREE_STATE-}'
PROJ_GIT_VERSION='${PROJ_GIT_VERSION-}'
PROJ_GIT_MAJOR='${PROJ_GIT_MAJOR-}'
PROJ_GIT_MINOR='${PROJ_GIT_MINOR-}'
EOF
}

# Loads up the version variables from file $1
proj::version::load_version_vars() {
  local version_file=${1-}
  [[ -n ${version_file} ]] || {
    echo "!!! Internal error.  No file specified in proj::version::load_version_vars"
    return 1
  }

  source "${version_file}"
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the Kubernetes based on the git tree status.
# IMPORTANT: if you update any of these, also update the lists in
# hack/print-workspace-status.sh.
proj::version::ldflags() {
  proj::version::get_version_vars

  local -a ldflags
  function add_ldflag() {
    local key=${1}
    local val=${2}
    ldflags+=(
      "-X '${PROJ_GO_PACKAGE}/pkg/version.${key}=${val}'"
    )
  }

  proj::util::ensure-gnu-date

  # 兼容 mac 和 linux 的 date 命令
  if [[ "$(uname)" == "Darwin" ]]; then
    add_ldflag "buildDate" "$(date -u "+%Y-%m-%dT%H:%M:%SZ")"
  else
    add_ldflag "buildDate" "$(${DATE} ${SOURCE_DATE_EPOCH:+"--date=@${SOURCE_DATE_EPOCH}"} -u +'%Y-%m-%dT%H:%M:%SZ')"
  fi

  if [[ -n ${PROJ_GIT_COMMIT-} ]]; then
    add_ldflag "gitCommit" "${PROJ_GIT_COMMIT}"
    add_ldflag "gitTreeState" "${PROJ_GIT_TREE_STATE}"
  fi

  if [[ -n ${PROJ_GIT_VERSION-} ]]; then
    add_ldflag "gitVersion" "${PROJ_GIT_VERSION}"
  fi

  if [[ -n ${PROJ_GIT_MAJOR-} && -n ${PROJ_GIT_MINOR-} ]]; then
    add_ldflag "gitMajor" "${PROJ_GIT_MAJOR}"
    add_ldflag "gitMinor" "${PROJ_GIT_MINOR}"
  fi

  # The -ldflags parameter takes a single string, so join the output.
  echo "${ldflags[*]-}"
}
