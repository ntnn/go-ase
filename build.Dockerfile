FROM bssdb/ase-sles12 as ase

WORKDIR /builder
COPY . .
RUN /utils/install ./go-1.11.5-1.x86_64.rpm make gcc

ENV GOPROXY=https://athens.mo.sap.corp/

ENV ASE_LIB=/sybase/TST/OCS-16_0/lib
ENV CGO_LDFLAGS="-L${ASE_LIB} -lsybct_r64 -lsybcs_r64"
ENV LD_LIBRARY_PATH=${ASE_LIB}
ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN go mod download; make -C driver build
