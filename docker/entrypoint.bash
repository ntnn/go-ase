#!/usr/bin/env bash

. /utils/common.sh

_set_dns() {
    ip="$(ip addr show dev eth0 | awk '/inet/ { print $2 }' |cut -d/ -f1)"

    # SWPM sapinst needs to resolve $hostname.$hostname if no domain is set.
    # Replacing /etc/hosts is not possible since it is mounted into the
    # container by docker. So only in place modification is possible.
    # Real inplace modification is only supported by vim and echo...
    vim /etc/hosts +g/^$ip/d +wq &>/dev/null
    echo "$ip	$(hostname) $(hostname).$(hostname)" >> /etc/hosts

    vim /etc/hosts +g/^127.0.0.1/d +wq &>/dev/null
    echo "127.0.0.1 localhost localhost.localhost" >> /etc/hosts
}

install() {
    _set_dns

    /opt/dbcat/bin/ase/aseFactory.sh \
        -a install -i $SID -r $ASE_VERSION -F -o CPU=$ASE_CPU,RAM=$ASE_RAM,app=$ASE_APP,hn=localhost \
        -A -M
}

stop() {
    /opt/dbcat/bin/ase/aseFactory.sh -A -i $SID -a stop
}

start() {
    /opt/dbcat/bin/ase/aseFactory.sh -A -i $SID -a start
}

logs() {
    tail -f /sybase/$SID/ASE-16_0/install/$SID.log
}

isql() {
    su --login syb$(echo $SID | tr '[A-Z]' '[a-z]') \
        ./OCS-16_0/bin/isql -S $SID -U sapsa \
        -P "$(awk -F= '/^cnfAseDefaultPass/ { print $2 }' < /opt/dbcat/etc/ase.cnf.sh)"
}

default() {
    start
    if [[ $# -ne 0 ]]; then
        echo "Failed to start ASE" >&2
        exit 1
    fi

    trap stop EXIT
    logs
}

_execute() {
    if ! declare -F | awk '{ print $3 }' | grep -q $1; then
        die "No action '$1' available"
    fi
    "$@"
}

if [[ $# -eq 0 ]]; then
    log "No arguments passed"
else
    _execute "$@"
fi
