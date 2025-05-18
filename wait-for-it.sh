#!/usr/bin/env bash
# wait-for-it.sh: Wait until a TCP host:port becomes available

WAITFORIT_cmdname=${0##*/}

echoerr() { [[ $WAITFORIT_QUIET -ne 1 ]] && echo "$@" 1>&2; }

usage() {
    cat << USAGE >&2
Usage:
    $WAITFORIT_cmdname host:port [-s] [-t timeout] [-q] [-- command args]
    -h HOST | --host=HOST       Host or IP under test
    -p PORT | --port=PORT       TCP port under test
                                 Alternatively use host:port as a single argument
    -s | --strict               Only execute subcommand if the test succeeds
    -q | --quiet                Don't output any status messages
    -t TIMEOUT | --timeout=TIMEOUT
                                 Timeout in seconds, zero for no timeout
    -- COMMAND ARGS             Command to execute after the test succeeds
USAGE
    exit 1
}

wait_for() {
    [[ $WAITFORIT_TIMEOUT -gt 0 ]] \
        && echoerr "$WAITFORIT_cmdname: waiting $WAITFORIT_TIMEOUT seconds for $WAITFORIT_HOST:$WAITFORIT_PORT" \
        || echoerr "$WAITFORIT_cmdname: waiting for $WAITFORIT_HOST:$WAITFORIT_PORT without a timeout"

    local start_ts=$(date +%s)
    while :; do
        if [[ $WAITFORIT_ISBUSY -eq 1 ]]; then
            nc -z "$WAITFORIT_HOST" "$WAITFORIT_PORT" >/dev/null 2>&1
        else
            (echo > /dev/tcp/"$WAITFORIT_HOST"/"$WAITFORIT_PORT") >/dev/null 2>&1
        fi
        local result=$?
        if [[ $result -eq 0 ]]; then
            local end_ts=$(date +%s)
            echoerr "$WAITFORIT_cmdname: $WAITFORIT_HOST:$WAITFORIT_PORT is available after $((end_ts - start_ts)) seconds"
            return 0
        fi
        sleep 1
    done
}

wait_for_wrapper() {
    if command -v gtimeout >/dev/null 2>&1; then
        local timeout_cmd=gtimeout
    elif command -v timeout >/dev/null 2>&1; then
        local timeout_cmd=timeout
    else
        echoerr "$WAITFORIT_cmdname: timeout command not found"
        exit 1
    fi

    $timeout_cmd $WAITFORIT_TIMEOUT "$0" --child --host="$WAITFORIT_HOST" --port="$WAITFORIT_PORT" ${WAITFORIT_QUIET:+--quiet} >/dev/null &
    local pid=$!
    trap "kill -INT -$pid" INT
    wait $pid
    return $?
}

# --- Parse args ---
WAITFORIT_TIMEOUT=15
WAITFORIT_STRICT=0
WAITFORIT_QUIET=0
WAITFORIT_CHILD=0

while [[ $# -gt 0 ]]; do
    case "$1" in
        *:*)
            IFS=':' read -r WAITFORIT_HOST WAITFORIT_PORT <<< "$1"
            shift ;;
        --child)
            WAITFORIT_CHILD=1
            shift ;;
        -q|--quiet)
            WAITFORIT_QUIET=1
            shift ;;
        -s|--strict)
            WAITFORIT_STRICT=1
            shift ;;
        -h|--host)
            WAITFORIT_HOST="${2:-}"
            shift 2 ;;
        --host=*)
            WAITFORIT_HOST="${1#*=}"
            shift ;;
        -p|--port)
            WAITFORIT_PORT="${2:-}"
            shift 2 ;;
        --port=*)
            WAITFORIT_PORT="${1#*=}"
            shift ;;
        -t|--timeout)
            WAITFORIT_TIMEOUT="${2:-}"
            shift 2 ;;
        --timeout=*)
            WAITFORIT_TIMEOUT="${1#*=}"
            shift ;;
        --)
            shift
            WAITFORIT_CLI=("$@")
            break ;;
        --help)
            usage ;;
        *)
            echoerr "Unknown argument: $1"
            usage ;;
    esac
done

[[ -z "$WAITFORIT_HOST" || -z "$WAITFORIT_PORT" ]] && {
    echoerr "Error: you must specify a host and port to test."
    usage
}

# Detect whether we need to use nc or /dev/tcp
if nc -z 127.0.0.1 1 >/dev/null 2>&1; then
    WAITFORIT_ISBUSY=1
else
    WAITFORIT_ISBUSY=0
fi

# --- Main ---
if [[ $WAITFORIT_CHILD -eq 1 ]]; then
    wait_for
    exit $?
fi

if [[ $WAITFORIT_TIMEOUT -gt 0 ]]; then
    wait_for_wrapper
    WAITFORIT_RESULT=$?
else
    wait_for
    WAITFORIT_RESULT=$?
fi

if [[ -n "${WAITFORIT_CLI[*]}" ]]; then
    if [[ $WAITFORIT_RESULT -ne 0 && $WAITFORIT_STRICT -eq 1 ]]; then
        echoerr "$WAITFORIT_cmdname: strict mode, refusing to execute subprocess"
        exit $WAITFORIT_RESULT
    fi
    exec "${WAITFORIT_CLI[@]}"
else
    exit $WAITFORIT_RESULT
fi