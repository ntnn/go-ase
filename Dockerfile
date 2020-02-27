FROM bssdb/dbcat:latest

RUN /utils/install tcsh libaio1 vim

# Required for tput etc.pp. to work correctly
ENV TERM=linux

# Defaults for ASE installation
ENV SID=TST
ENV ASE_VERSION=tactical
ENV ASE_CPU=3
ENV ASE_RAM=16
ENV ASE_APP=empty

ADD docker/base.cnf.local.sh /opt/dbcat/etc/
ADD docker/ase.cnf.local.sh /opt/dbcat/etc/
ADD docker/entrypoint.bash /entrypoint.bash

RUN /entrypoint.bash install stop

ENTRYPOINT /entrypoint.bash
CMD default
