/*
**      SAP CT-LIBRARY
**      Copyright (c) 2013 SAP AG or an SAP affiliate company.  All rights reserved.
**
**      TDSPUBLIC.H
**
**      This is the header file for TDS tokens. It defines all the TDS tokens
**      that are in the TDS Functional Spec, currently adhering to release 5.0
**	subversion 3.8.
**      For details on the syntax and usage of TDS tokens and constants, please
**	see the appropriate reference pages in the TDS Spec.
**
** History:
** Chg#	Date	Description						Resp
** ----	-------	-------------------------------------------------------	----
** 	04Mar92	Created.						amit
** 001	03Nov92	Added definitions for older loginrec versions.		rs
**		Also added TDS_VOID data type token. Also removed
**		ENDPARAM packet type.
** 002	07Jul95	System 11 - Security Services Support			soma
** 003  24Sep97 Added a new capability request TDS_DOL_BULK to		sunilk
**		support the new rowformat (DOL) for EARL server
** 004  29Jul02 Added new date and time date types			steng
** 005  xxJan04 Added new unitext date type				kamphuis
** 006  24Feb04 Added new bigint date type				steng
** 007  11May04 Added new tokens for scrollable cursor			nico
**		TDS version for scrollable cursor is 3.6
** 008	19May04	Added new unsigned integer datetypes			steng
** 009	14Sep04	Added new xml datetype					steng
** 010	21Apr05	Added capability request for CURINFO3 			nico
** 011	11Apr06	OCS/ASE tds tokens integration project  		nico
**		Moved all tokens as previously defined in tds.h into
**		this file, tdspublic.h
** 012	09Sep06 OCS/ASE tds token integration project, phase 2.		nico
**		Sync up with ASE, definitive updates for this project.
** 013	06Mar09	Added new microsecond granularity datetime/time types	pmathy
** 014	05Oct10	Added new TDS cursor flags for "release_locks_on_close"	nico
*/
#ifndef __TDSPUBLIC_H
#define __TDSPUBLIC_H

/*
** For OCS, use typedefs, do not use CS_ types from upper layers.
*/
typedef	unsigned char	TDS_BYTE;
typedef	unsigned short	TDS_SHORT;

/*
** The login record.
*/
#define TDS_PKTLEN	6
#define TDS_PROGNLEN	10
#define TDS_MAXNAME	30
#define TDS_NETBUF	4
#define TDS_VERSIZE	4
#define TDS_OLDSECURE	2
#define TDS_SECURE	2
#define TDS_RPLEN	255
#define TDS_HA		6

/*
** TDS_LOGINREC - This is the 5.0 version of the LOGINREC.
*/
typedef struct tds__loginrec
{
	TDS_BYTE	lhostname[TDS_MAXNAME]; /* name of host or generic */
	TDS_BYTE	lhostnlen;		/* length of lhostname */
	TDS_BYTE	lusername[TDS_MAXNAME];	/* name of user */
	TDS_BYTE	lusernlen;		/* length of lusername */
	TDS_BYTE	lpw[TDS_MAXNAME];	/* password */
	TDS_BYTE	lpwnlen;		/* length of lpw */
	TDS_BYTE	lhostproc[TDS_MAXNAME];	/* host process identification*/
	TDS_BYTE	lhplen;			/* length of host process id */
	TDS_BYTE	lint2;			/* type of int2 on this host */
	TDS_BYTE	lint4;			/* type of int4 on this host */
	TDS_BYTE	lchar;			/* type of char */
	TDS_BYTE	lflt;			/* type of float */
	TDS_BYTE	ldate;			/* type of datetime */
	TDS_BYTE	lusedb;			/* notify on exec of use db cmd */
	TDS_BYTE	ldmpld;			/* disallow use of dump/load
						** and bulk insert
						*/
	TDS_BYTE	linterfacespare;	/* no longer used*/
	TDS_BYTE	ltype;			/* type of network connection */
	TDS_BYTE	lbufsize[TDS_NETBUF];	/* network buffer size */
	TDS_BYTE	spare[3];		/* spare fields */
	TDS_BYTE	lappname[TDS_MAXNAME];	/* application name */
	TDS_BYTE	lappnlen;		/* length of appl name */
	TDS_BYTE	lservname[TDS_MAXNAME];	/* name of server */
	TDS_BYTE	lservnlen;		/* length of lservname */
	TDS_BYTE	lrempw[TDS_RPLEN];	/* passwords for remote servers */
	TDS_BYTE	lrempwlen;		/* length of lrempw */
	TDS_BYTE	ltds[TDS_VERSIZE];	/* tds version */
	TDS_BYTE	lprogname[TDS_PROGNLEN];/* client program name */
	TDS_BYTE	lprognlen;		/* length of client program name */
	TDS_BYTE	lprogvers[TDS_VERSIZE];	/* client program version */
	TDS_BYTE	lnoshort;		/* auto convert of short datatypes */
	TDS_BYTE	lflt4;			/* type of flt4 on this host */
	TDS_BYTE	ldate4;			/* type of flt4 on this host */
	TDS_BYTE	llanguage[TDS_MAXNAME];	/* initial language */
	TDS_BYTE	llanglen;		/* length of language */
	TDS_BYTE	lsetlang;		/* notify on language change */
/*
** The following 13 bytes were used by 1.0 secure servers. Actually 2 bytes in
** the middle are unused. Since we do not support logins to 1.0 secure servers,
** we can re-use these 13 bytes.
** However, non-secure servers, check if the first 2 bytes are non-zero. If they
** are non-zero, they assume that the user want's to login a secure server and
** reject the login.
*/
	TDS_BYTE	loldsecure[TDS_OLDSECURE];/* old secure server field */
	TDS_BYTE	lseclogin;		/* security login options */
	TDS_BYTE	lsecbulk;		/* security bulk copy options */
        TDS_BYTE        lhalogin;               /* HA login options */
        TDS_BYTE        lhasessionid[TDS_HA];   /* HA session id */
	TDS_BYTE	lsecspare[TDS_SECURE];	/* reserved for security */
	TDS_BYTE	lcharset[TDS_MAXNAME]; 	/* character set name */
	TDS_BYTE	lcharsetlen;           	/* length of lcharset */
	TDS_BYTE	lsetcharset;            /* notify character set change */
	TDS_BYTE	lpacketsize[TDS_PKTLEN];/* length of TDS packet desired */
	TDS_BYTE	lpacketsizelen;        	/* length of lpacketsize */
	TDS_BYTE	ldummy[4];		/* pad to longword */
} TDS_LOGINREC;

/*
** 001 - Below are the definitions for older versions of the LOGINREC
** which are still supported.
*/

/*
** TDS_40_LOGINREC - This the 4.0 version of the LOGINREC.
*/
typedef struct tds__40_loginrec
{
	TDS_BYTE	lhostname[TDS_MAXNAME];	/* name of host or generic */
	TDS_BYTE	lhostnlen;		/* length of lhostname */
	TDS_BYTE	lusername[TDS_MAXNAME];	/* name of user */
	TDS_BYTE	lusernlen;		/* length of lusername */
	TDS_BYTE	lpw[TDS_MAXNAME];	/* password (plaintext) */
	TDS_BYTE	lpwnlen;		/* length of lpw */
	TDS_BYTE	lhostproc[TDS_MAXNAME];	/* host process identification*/
	TDS_BYTE	lhplen;			/* length of host process id */
	TDS_BYTE	lint2;			/* type of int2 on this host */
	TDS_BYTE	lint4;			/* type of int4 on this host */
	TDS_BYTE	lchar;			/* type of char */
	TDS_BYTE	lflt;			/* type of float */
	TDS_BYTE	ldate;			/* type of datetime */
	TDS_BYTE	lusedb;			/* notify exec of usedb cmd */
	TDS_BYTE	ldmpld;			/* disallow dump/load, bulk */
	TDS_BYTE	linterface;		/* SQL interface type */
	TDS_BYTE	ltype;			/* type of network connection */
	TDS_BYTE	spare[7];		/* spare fields */
	TDS_BYTE	lappname[TDS_MAXNAME];	/* application name */
	TDS_BYTE	lappnlen;		/* length of appl name */
	TDS_BYTE	lservname[TDS_MAXNAME];	/* name of server */
	TDS_BYTE	lservnlen;		/* length of lservname */
	TDS_BYTE	lrempw[0xff];		/* passwords for rmt servers */
	TDS_BYTE	lrempwlen;		/* length of lrempw */
	TDS_BYTE	ltds[4];		/* tds version */
	TDS_BYTE	lprogname[TDS_PROGNLEN];/* client program name */
	TDS_BYTE	lprognlen;		/* length of client prog name */
	TDS_BYTE	lprogvers[4];		/* client program version */
	TDS_BYTE	ldummy[3];		/* pad length to longword */
} TDS_40_LOGINREC;

/*
** TDS_SECLAB - This data structure is used by the 4.0.2 loginrec.
*/
typedef struct tds__seclab
{
        short   	slhier;
        TDS_BYTE    	slcomp[8];
        short   	slspare;
} TDS_SECLAB;

/*
** TDS_402_LOGINREC - This the 4.0.2 version of the LOGINREC.
*/
typedef struct tds__402_loginrec
{
        TDS_BYTE	lhostname[TDS_MAXNAME];	/* name of host or generic */
        TDS_BYTE	lhostnlen;          	/* length of lhostname */
        TDS_BYTE	lusername[TDS_MAXNAME];	/* name of user */
        TDS_BYTE	lusernlen;          	/* length of lusername */
        TDS_BYTE	lpw[TDS_MAXNAME];       /* password (plaintext)	*/
        TDS_BYTE	lpwnlen;            	/* length of lpw */
        TDS_BYTE	lhostproc[TDS_MAXNAME]; /* host process id */
        TDS_BYTE	lhplen;             	/* length of host process id */
        TDS_BYTE	lint2;              	/* type of int2 on this host */
        TDS_BYTE	lint4;              	/* type of int4 on this host */
        TDS_BYTE	lchar;              	/* type of char */
        TDS_BYTE	lflt;               	/* type of float */
        TDS_BYTE	ldate;              	/* type of datetime */
        TDS_BYTE	lusedb;			/* notify on exec of use db cmd */
        TDS_BYTE	ldmpld;    		/* disallow dump/load and bulk */
        TDS_BYTE	linterface;     	/* SQL interface type */
        TDS_BYTE	ltype;          	/* type of network connection */
        TDS_BYTE	spare[7];        	/* spare fields */
        TDS_BYTE	lappname[TDS_MAXNAME]; 	/* application name */
        TDS_BYTE	lappnlen;         	/* length of appl name */
        TDS_BYTE	lservname[TDS_MAXNAME]; /* name of server */
        TDS_BYTE	lservnlen;         	/* length of lservname */
        TDS_BYTE	lrempw[0xff];      	/* passwords for remote servers */
        TDS_BYTE	lrempwlen;         	/* length of lrempw */
	TDS_BYTE	ltds[4];           	/* tds version 	*/
        TDS_BYTE	lprogname[TDS_PROGNLEN];/* client program name */
        TDS_BYTE	lprognlen;         	/* length of client program name*/
        TDS_BYTE	lprogvers[4];      	/* client program version */
        TDS_BYTE	ldummy[2];         	/* pad length to longword */
        TDS_BYTE	spare1;            	/* structure alignment */
        TDS_SECLAB 	lseclab;         	/* login security level	*/
        TDS_BYTE	lrole;             	/* login role(sa = 1,user = 0 */
        TDS_BYTE	spare2[3];         	/* pad length to longword */
} TDS_402_LOGINREC;

/*
** TDS_42_LOGINREC - This the 4.2 version of the LOGINREC.
*/
typedef struct tds__42_loginrec
{
        TDS_BYTE	lhostname[TDS_MAXNAME]; /* name of host or generic */
        TDS_BYTE	lhostnlen;              /* length of lhostname */
        TDS_BYTE	lusername[TDS_MAXNAME]; /* name of user */
        TDS_BYTE	lusernlen;              /* length of lusername */
        TDS_BYTE	lpw[TDS_MAXNAME];       /* password (plaintext)	*/
        TDS_BYTE	lpwnlen;                /* length of lpw */
        TDS_BYTE	lhostproc[TDS_MAXNAME]; /* host process id */
        TDS_BYTE	lhplen;                 /* length of host process id */
        TDS_BYTE	lint2;                  /* type of int2 on this host */
 	TDS_BYTE	lint4;                  /* type of int4 on this host */
        TDS_BYTE	lchar;                  /* type of char */
        TDS_BYTE	lflt;                   /* type of float */
        TDS_BYTE	ldate;                  /* type of datetime */
        TDS_BYTE	lusedb;                 /* notify exec of usedb cmd */
        TDS_BYTE	ldmpld;                 /* disallow dump/load, bulk */
        TDS_BYTE	linterface;             /* SQL interface type */
        TDS_BYTE	ltype;                  /* type of network connection */
	TDS_BYTE	lbufsize[4];		/* network buffer size */
        TDS_BYTE	spare[3];               /* spare fields */
        TDS_BYTE	lappname[TDS_MAXNAME];  /* application name */
        TDS_BYTE	lappnlen;               /* length of appl name */
        TDS_BYTE	lservname[TDS_MAXNAME]; /* name of server */
        TDS_BYTE	lservnlen;              /* length of lservname */
        TDS_BYTE	lrempw[0xff];           /* passwords for rmt servers */
	TDS_BYTE	lrempwlen;		/* length of lrempw */
	TDS_BYTE	ltds[4];                /* tds version */
        TDS_BYTE	lprogname[TDS_PROGNLEN];/* client program name */
        TDS_BYTE	lprognlen;             	/* length of client prog name */
        TDS_BYTE	lprogvers[4];          	/* client program version */
        TDS_BYTE	lnoshort;              	/* auto convert of short datatypes */
	TDS_BYTE	lflt4;                 	/* type of flt4 on this host */
        TDS_BYTE	ldate4;                	/* type of flt4 on this host */
        TDS_BYTE	llanguage[TDS_MAXNAME];	/* initial language */
        TDS_BYTE	llanglen;              	/* length of language */
        TDS_BYTE	lsetlang;              	/* notify on language change */
        TDS_BYTE	slhier[2];             	/* security label hierarchy */
        TDS_BYTE	slcomp[8];             	/* security compartments */
        TDS_BYTE	slspare[2];            	/* security label spare field */
        TDS_BYTE	lrole;                 	/* sec login role (sa=1,user=0) */
        TDS_BYTE	ldummy[3];             	/* extra bytes 	*/
} TDS_42_LOGINREC;

/*
** Here are the values that are significant in the login record. They are
** ordered by token value. Please retain ordering when adding a new define.
*/

/*
** lint4 - four-byte integers
*/
#define	TDS_INT4_LSB_HI		0 	/* lsb is hi byte (eg 68000) */
#define	TDS_INT4_LSB_LO 	1	/* lsb is low byte (eg VAX) */

/*
** lint2 - two-byte integers
*/
#define	TDS_INT2_LSB_HI		2	/* lsb is hi byte (eg 68000) */
#define	TDS_INT2_LSB_LO 	3 	/* lsb is low byte (eg VAX ) */

/*
** lflt- Floating point (8-byte) types.
*/
#define	TDS_FLT_IEEE_HI		4	/* IEEE 754 float, lsb in high byte */
#define	TDS_FLT_VAXD 		5 	/* VAX `D' floating point format */

/*
** lchar - Character type.
*/
#define	TDS_CHAR_ASCII		6 	/* ASCII character set */
#define	TDS_CHAR_EBCDIC		7 	/* EBCDIC character set */

/*
** ldate - Datetime type.
*/
#define	TDS_TWO_I4_LSB_HI	8	/* lsb is hi integer */
#define	TDS_TWO_I4_LSB_LO	9 	/* lsb is low integer */

#define	TDS_FLT_IEEE_LO 	10	/* IEEE 754 float, lsb in low byte */
#define	TDS_FLT_ND5000		11 	/* ND5000   float, lsb in hi byte */

/*
** lflt4 - 4-byte float types
*/
#define	TDS_FLT4_IEEE_HI	12	/* IEEE 4-byte float, lsb is hi byte */
#define	TDS_FLT4_IEEE_LO	13	/* IEEE 4-byte float, lsb is lo byte */
#define	TDS_FLT4_VAXF 		14	/* VAX `F' floating point format */
#define	TDS_FLT4_ND50004 	15 	/* ND5000 4-byte float format */

/*
** ldate4 - short datetime type
*/
#define	TDS_TWO_I2_LSB_HI	16	/* lsb is hi integer */
#define	TDS_TWO_I2_LSB_LO	17	/* lsb is lo integer */

/*
** ltype - Connection type
*/
#define	TDS_LSERVER		0x1	/* not a user connecting directly */
#define	TDS_LREMUSER		0x2	/* user login through another server */
#define TDS_LINTERNAL_RPC 	0x4	/* allow an internal RPC to be executed
					** on the connection
					*/

/*
** ltds - TDS version requested
*/

/*
** Values for ltds - 5.1.0.0
*/
#define TDS_5_1_V1		5
#define TDS_5_1_V2		1
#define TDS_5_1_V3		0
#define TDS_5_1_V4		0

/*
** Values for ltds - 5.0.0.0
*/
#define TDS_5_0_V1		5
#define TDS_5_0_V2		0
#define TDS_5_0_V3		0
#define TDS_5_0_V4		0

/*
** Values for ltds - 4.9.5.0
*/
#define TDS_4_9_5_V1		4
#define TDS_4_9_5_V2		9
#define TDS_4_9_5_V3		5
#define TDS_4_9_5_V4		0

/*
** Values for ltds - 4.6.0.0
*/
#define TDS_4_6_V1		4
#define TDS_4_6_V2		6
#define TDS_4_6_V3		0
#define TDS_4_6_V4		0

/*
** Values for ltds - 4.2.0.0
*/
#define TDS_4_2_V1		4
#define TDS_4_2_V2		2
#define TDS_4_2_V3		0
#define TDS_4_2_V4		0

/*
** Values for ltds - 4.0.0.0
*/
#define TDS_4_0_V1		4
#define TDS_4_0_V2		0
#define TDS_4_0_V3		0
#define TDS_4_0_V4		0

/*
** Values for ltds - 4.1.0.0
*/
#define	TDS_4_1_V1		4
#define	TDS_4_1_V2		1
#define	TDS_4_1_V3		0
#define	TDS_4_1_V4		0

/*
** Values for ltds - 3.4.0.0
*/
#define	TDS_3_4_V1		3
#define	TDS_3_4_V2		4
#define	TDS_3_4_V3		0
#define	TDS_3_4_V4		0

/*
** Values for ltds - 2.0.0.0
*/
#define	TDS_2_0_V1		2
#define	TDS_2_0_V2		0
#define	TDS_2_0_V3		0
#define	TDS_2_0_V4		0

/*
** lseclogin - security login options
*/
#define TDS_SEC_LOG_ENCRYPT	0x01
#define TDS_SEC_LOG_CHALLENGE	0x02
#define TDS_SEC_LOG_LABELS	0x04
#define TDS_SEC_LOG_APPDEFINED  0x08
#define TDS_SEC_LOG_SECSESS	0x10
#define TDS_SEC_LOG_ENCRYPT2	0x20
#define TDS_SEC_LOG_ENCRYPT3	0x80

/*
** lsecbulk - bulk copy options
*/
#define TDS_SEC_BULK_LABELED	0x01

/*
** lhalogin - HA login options
*/
#define TDS_HA_LOG_SESSION      0x01
#define TDS_HA_LOG_RESUME       0x02
#define TDS_HA_LOG_FAILOVERSRV  0x04
#define	TDS_HA_LOG_REDIRECT	0x08
#define	TDS_HA_LOG_MIGRATE	0x10

/*
** linterfacespare values, server-to-server negotiation.
*/
#define TDS_LDEFSQL		0	/* server's default SQL will be sent */
#define TDS_LXSQL		1	/* TRANSACT-SQL will be sent */
#define TDS_LSQL		2	/* ANSI SQL, version 1 */
#define TDS_LSQL2_1		3	/* ANSI SQL, version 2, level 1 */
#define TDS_LSQL2_2		4	/* ANSI SQL, version 2, level 2 */

/*
** LOGINACK status values.
*/
#define	TDS_LOG_SUCCEED		5	/* Log in succeeded */
#define TDS_LOG_FAIL		6	/* Log in failed. */
#define	TDS_LOG_NEGOTIATE	7	/* Negotiate further. */

/*
** LOGINACK status bits. Note that this bit can be set and one of
** the above status values may be returned in the same byte. i.e.
** 0x05, 0x06, 0x07, 0x85, 0x86, and 0x87 are the possible values
** for the status.
*/
#define TDS_LOG_SECSESS_ACK	0x80

/*
** MSG datastream status and types. 0 through 32,767 are reserved
** for SAP CSI types.
*/
#define TDS_MSG_HASARGS		0x01

#define TDS_MSG_SEC_ENCRYPT	1
#define TDS_MSG_SEC_LOGPWD	2
#define TDS_MSG_SEC_REMPWD	3
#define TDS_MSG_SEC_CHALLENGE	4
#define TDS_MSG_SEC_RESPONSE	5
#define TDS_MSG_SEC_GETLABEL	6
#define TDS_MSG_SEC_LABEL	7
#define	TDS_MSG_SQL_TBLNAME	8
#define	TDS_MSG_GW_RESERVED	9
#define	TDS_MSG_OMNI_CAPABILITIES 10
#define TDS_MSG_SEC_OPAQUE	11
#define TDS_MSG_HAFAILOVER	12
#define	TDS_MSG_EMPTY		13
#define	TDS_MSG_SEC_ENCRYPT2	14
#define	TDS_MSG_SEC_LOGPWD2	15
#define	TDS_MSG_SEC_SUP_CIPHER2	16
#define	TDS_MSG_MIG_REQ		17
#define	TDS_MSG_MIG_SYNC	18
#define	TDS_MSG_MIG_CONT	19
#define	TDS_MSG_MIG_IGN		20
#define	TDS_MSG_MIG_FAIL	21
#define TDS_MSG_SEC_REMPWD2	22
#define	TDS_MSG_MIG_RESUME	23

#define	TDS_MSG_SEC_ENCRYPT3	30
#define	TDS_MSG_SEC_LOGPWD3	31
#define TDS_MSG_SEC_REMPWD3	32
#define TDS_MSG_DR_MAP		33
#define TDS_MSG_SEC_SYMKEY	34
#define TDS_MSG_SEC_ENCRYPT4	35

/*
** TDS_MSG_SEC_OPAQUE message types
*/
#define	TDS_SEC_SECSESS		1
#define TDS_SEC_FORWARD		2
#define TDS_SEC_SIGN		3
#define TDS_SEC_OTHER		4

/*
** Keyword values for OFFSET datastream.
*/
#define TDS_OFF_SELECT		0x016D
#define TDS_OFF_FROM		0x014F
#define TDS_OFF_ORDER		0x0165
#define TDS_OFF_COMPUTE		0x0139
#define TDS_OFF_TABLE		0x0173
#define TDS_OFF_PROC		0x016A
#define TDS_OFF_STMT		0x01CB
#define TDS_OFF_PARAM		0x01C4

/*
** ENVCHANGE types
*/
#define TDS_ENV_DB		1
#define TDS_ENV_LANG		2
#define TDS_ENV_CHARSET		3
#define TDS_ENV_PACKSIZE	4

/*
** Special TDS EED message numbers and values.
** These values are used in server to client library control
** messages sent via TDS_EED tokens.
*/
#define	TDS_REDIRECT		1

/*
** State values for a redirect message.
*/
#define	TDS_EED_IMMEDIATE_REDIRECT	0x01
#define	TDS_EED_SET_REDIRECT		0x02

/*
** Message buffer types.
*/
#define TDS_BUF_LANG		1
#define TDS_BUF_LOGIN		2
#define TDS_BUF_RPC		3
#define TDS_BUF_RESPONSE	4
#define TDS_BUF_UNFMT		5
#define TDS_BUF_ATTN		6
#define TDS_BUF_BULK		7
#define TDS_BUF_SETUP		8
#define TDS_BUF_CLOSE		9
#define TDS_BUF_ERROR		10
#define TDS_BUF_PROTACK		11
#define TDS_BUF_ECHO		12
#define TDS_BUF_LOGOUT		13
#define TDS_BUF_ENDPARAM	14
#define TDS_BUF_NORMAL		15
#define TDS_BUF_URGENT		16
#define TDS_BUF_MIGRATE		17
#define	TDS_BUF_CMDSEC_NORMAL	24
#define	TDS_BUF_CMDSEC_LOGIN	25
#define	TDS_BUF_CMDSEC_LIVENESS	26
#define	TDS_BUF_CMDSEC_RESERVED1 27
#define	TDS_BUF_CMDSEC_RESERVED2 28

/*
** Message buffer status values.
*/
#define TDS_BUFSTAT_EOM		0x01
#define TDS_BUFSTAT_ATTNACK	0x02
#define TDS_BUFSTAT_ATTN	0x04
#define TDS_BUFSTAT_EVENT	0x08
#define TDS_BUFSTAT_SEAL	0x10
#define TDS_BUFSTAT_ENCRYPT	0x20
#define TDS_BUFSTAT_SYMENCRYPT	0x40

/*
** Defined for physical layout of header structure that is compiler
** independent. This allow memcpys to extract the data from the network
** stream.
*/
#define	TDS_HDR_MSGTYPE_LEN	1
#define	TDS_HDR_STATUS_LEN	1
#define	TDS_HDR_LENGTH_LEN	2
#define	TDS_HDR_CHANNEL_LEN	2
#define	TDS_HDR_PACKETNO_LEN	1
#define	TDS_HDR_WINDOW_LEN	1
#define	TDS_HDR_MSGTYPE		0
#define	TDS_HDR_STATUS		(TDS_HDR_MSGTYPE + TDS_HDR_MSGTYPE_LEN)
#define	TDS_HDR_LENGTH		(TDS_HDR_STATUS + TDS_HDR_STATUS_LEN)
#define	TDS_HDR_CHANNEL		(TDS_HDR_LENGTH + TDS_HDR_LENGTH_LEN)
#define	TDS_HDR_PACKETNO	(TDS_HDR_CHANNEL + TDS_HDR_CHANNEL_LEN)
#define	TDS_HDR_WINDOW		(TDS_HDR_PACKETNO + TDS_HDR_PACKETNO_LEN)

/*
** The basic TDS tokens. Each of these tokens has its own datastream.
** Note: these have been put in numerical order by tokenvalue, please
** retain this when adding a new token definition.
*/
# define TDS_CURDECLARE3  	0x10
# define TDS_PARAMFMT2    	0x20
# define TDS_LANGUAGE     	0x21
# define TDS_ORDERBY2     	0x22
# define TDS_CURDECLARE2  	0x23
# define TDS_COLFMTOLD    	0x2A
# define TDS_DEBUGCMD    	0x60
# define TDS_ROWFMT2      	0x61
# define TDS_DYNAMIC2     	0x62
# define TDS_MSG          	0x65
# define TDS_LOGOUT       	0x71
# define TDS_OFFSET       	0x78
# define TDS_RETURNSTATUS 	0x79
# define TDS_PROCID       	0x7C
# define TDS_CURCLOSE     	0x80
# define TDS_CURDELETE    	0x81
# define TDS_CURFETCH     	0x82
# define TDS_CURINFO      	0x83
# define TDS_CUROPEN      	0x84
# define TDS_CURUPDATE    	0x85
# define TDS_CURDECLARE   	0x86
# define TDS_CURINFO2     	0x87
# define TDS_CURINFO3     	0x88
# define TDS_COLNAME      	0xA0
# define TDS_COLFMT       	0xA1
# define TDS_EVENTNOTICE  	0xA2
# define TDS_TABNAME      	0xA4
# define TDS_COLINFO      	0xA5
# define TDS_OPTIONCMD    	0xA6
# define TDS_ALTNAME      	0xA7
# define TDS_ALTFMT       	0xA8
# define TDS_ORDERBY      	0xA9
# define TDS_ERROR        	0xAA
# define TDS_INFO         	0xAB
# define TDS_RETURNVALUE  	0xAC
# define TDS_LOGINACK     	0xAD
# define TDS_CONTROL      	0xAE
# define TDS_ALTCONTROL   	0xAF
# define TDS_KEY          	0xCA
# define TDS_ROW          	0xD1
# define TDS_ALTROW       	0xD3
# define TDS_PARAMS       	0xD7
# define TDS_RPC          	0xE0
# define TDS_CAPABILITY   	0xE2
# define TDS_ENVCHANGE    	0xE3
# define TDS_EED          	0xE5
# define TDS_DBRPC        	0xE6
# define TDS_DYNAMIC      	0xE7
# define TDS_DBRPC2       	0xE8
# define TDS_PARAMFMT     	0xEC
# define TDS_ROWFMT       	0xEE
# define TDS_DONE         	0xFD
# define TDS_DONEPROC     	0xFE
# define TDS_DONEINPROC   	0xFF

/*
** The TDS datatypes, ordered by tokenvalue.
*/
#define TDS_VOID		0x1f
#define TDS_IMAGE		0x22
#define TDS_TEXT		0x23
#define	TDS_BLOB		0x24
#define TDS_VARBINARY		0x25
#define TDS_INTN		0x26
#define TDS_VARCHAR		0x27
#define TDS_BINARY		0x2D
#define TDS_CHAR		0x2F
#define TDS_INT1		0x30
#define TDS_DATE		0x31
#define TDS_BIT			0x32
#define TDS_TIME		0x33
#define TDS_INT2		0x34
#define TDS_INT4 		0x38
#define TDS_SHORTDATE		0x3A
#define TDS_FLT4		0x3B
#define TDS_MONEY		0x3C
#define TDS_DATETIME		0x3D
#define TDS_FLT8		0x3E
#define TDS_UINT2		0x41
#define TDS_UINT4		0x42
#define TDS_UINT8		0x43
#define TDS_UINTN		0x44
#define TDS_SENSITIVITY         0x67
#define TDS_BOUNDARY         	0x68
#define TDS_DECN		0x6A
#define TDS_NUMN		0x6C
#define TDS_FLTN		0x6D
#define TDS_MONEYN	 	0x6E
#define TDS_DATETIMN		0x6F
#define TDS_SHORTMONEY		0x7A
#define TDS_DATEN		0x7B
#define TDS_TIMEN		0x93
#define TDS_XML			0xA3
#define TDS_UNITEXT		0xAE
#define TDS_LONGCHAR		0xAF
#define TDS_BIGDATETIMEN	0xBB
#define TDS_BIGTIMEN		0xBC
#define TDS_INT8 		0xBF
#define TDS_LONGBINARY		0xE1

/*
** TDS usertypes.
*/
#define TDS_USER_TEXT		19
#define TDS_USER_IMAGE		20
#define TDS_USER_UNITEXT	36

/*
** TDS_BLOB BlobTypes.
*/
#define TDS_BLOB_FULLCLASSNAME	0x01
#define TDS_BLOB_DBID_CLASSDEF	0x02
#define TDS_BLOB_CHAR		0x03
#define TDS_BLOB_BINARY		0x04
#define TDS_BLOB_UNICHAR	0x05
#define TDS_LOBLOC_CHAR		0x06
#define TDS_LOBLOC_BINARY	0x07
#define TDS_LOBLOC_UNICHAR	0x08

/*
** Capability types
*/
#define TDS_CAP_REQUEST		1
#define TDS_CAP_RESPONSE	2

/*
** Capability request values
*/
#define TDS_REQ_LANG			1
#define TDS_REQ_RPC			2
#define TDS_REQ_EVT			3
#define TDS_REQ_MSTMT			4
#define TDS_REQ_BCP			5
#define TDS_REQ_CURSOR			6
#define TDS_REQ_DYNF			7
#define TDS_REQ_MSG			8
#define TDS_REQ_PARAM			9
#define TDS_DATA_INT1			10
#define TDS_DATA_INT2			11
#define TDS_DATA_INT4			12
#define TDS_DATA_BIT			13
#define TDS_DATA_CHAR			14
#define TDS_DATA_VCHAR			15
#define TDS_DATA_BIN			16
#define TDS_DATA_VBIN			17
#define TDS_DATA_MNY8			18
#define TDS_DATA_MNY4			19
#define TDS_DATA_DATE8			20
#define TDS_DATA_DATE4			21
#define TDS_DATA_FLT4			22
#define TDS_DATA_FLT8			23
#define TDS_DATA_NUM			24
#define TDS_DATA_TEXT			25
#define TDS_DATA_IMAGE			26
#define TDS_DATA_DEC			27
#define TDS_DATA_LCHAR			28
#define TDS_DATA_LBIN			29
#define TDS_DATA_INTN			30
#define TDS_DATA_DATETIMEN		31
#define TDS_DATA_MONEYN			32
#define TDS_CSR_PREV			33
#define TDS_CSR_FIRST			34
#define TDS_CSR_LAST			35
#define TDS_CSR_ABS			36
#define TDS_CSR_REL			37
#define TDS_CSR_MULTI			38
#define TDS_CON_OOB			39
#define TDS_CON_INBAND			40
#define TDS_CON_LOGICAL			41
#define TDS_PROTO_TEXT			42
#define TDS_PROTO_BULK			43
#define TDS_REQ_URGEVT			44
#define TDS_DATA_SENSITIVITY		45
#define TDS_DATA_BOUNDARY		46
#define TDS_PROTO_DYNAMIC		47
#define TDS_PROTO_DYNPROC		48
#define TDS_DATA_FLTN			49
#define TDS_DATA_BITN			50
#define TDS_DATA_INT8			51
#define TDS_DATA_VOID			52
#define TDS_DOL_BULK			53
#define TDS_OBJECT_JAVA1		54
#define TDS_OBJECT_CHAR			55
#define TDS_REQ_RESERVED1		56
#define TDS_OBJECT_BINARY		57
#define TDS_DATA_COLUMNSTATUS		58
#define TDS_WIDETABLES			59
#define TDS_REQ_RESERVED2		60
#define TDS_DATA_UINT2			61
#define TDS_DATA_UINT4			62
#define TDS_DATA_UINT8			63
#define TDS_DATA_UINTN			64
#define TDS_CUR_IMPLICIT		65
#define TDS_DATA_NLBIN			66
#define TDS_IMAGE_NCHAR			67
#define TDS_BLOB_NCHAR_16		68
#define TDS_BLOB_NCHAR_8		69
#define TDS_BLOB_NCHAR_SCSU		70
#define TDS_DATA_DATE			71
#define TDS_DATA_TIME			72
#define TDS_DATA_INTERVAL		73
#define TDS_CSR_SCROLL			74
#define TDS_CSR_SENSITIVE		75
#define TDS_CSR_INSENSITIVE		76
#define TDS_CSR_SEMISENSITIVE		77
#define TDS_CSR_KEYSETDRIVEN		78
#define TDS_REQ_SRVPKTSIZE		79
#define TDS_DATA_UNITEXT		80
#define TDS_CAP_CLUSTERFAILOVER		81
#define TDS_DATA_SINT1			82
#define TDS_REQ_LARGEIDENT		83
#define TDS_REQ_BLOB_NCHAR_16		84
#define TDS_DATA_XML			85
#define	TDS_REQ_CURINFO3		86
#define TDS_REQ_DBRPC2			87
#define TDS_UNUSED_REQ			88
#define TDS_REQ_MIGRATE			89
#define TDS_MULTI_REQUESTS		90
#define TDS_REQ_RESERVED_91		91
#define TDS_REQ_RESERVED_92		92
#define TDS_DATA_BIGDATETIME		93
#define TDS_DATA_USECS			94
#define TDS_RPCPARAM_LOB		95
#define TDS_REQ_INSTID			96
#define TDS_REQ_GRID			97
#define TDS_REQ_DYN_BATCH		98
#define TDS_REQ_LANG_BATCH		99
#define TDS_REQ_RPC_BATCH		100
#define TDS_DATA_LOBLOCATOR		101
#define TDS_REQ_ROWCOUNT_FOR_SELECT	102
#define TDS_REQ_LOGPARAMS		103
#define TDS_REQ_DYNAMIC_SUPPRESS_PARAMFMT 104
#define	TDS_REQ_READONLY		105
#define	TDS_REQ_COMMAND_ENCRYPTION	106

/*
** Capability response values
*/
#define TDS_RES_NOMSG			1
#define TDS_RES_NOEED			2
#define TDS_RES_NOPARAM			3
#define TDS_DATA_NOINT1			4
#define TDS_DATA_NOINT2			5
#define TDS_DATA_NOINT4			6
#define TDS_DATA_NOBIT			7
#define TDS_DATA_NOCHAR			8
#define TDS_DATA_NOVCHAR		9
#define TDS_DATA_NOBIN			10
#define TDS_DATA_NOVBIN			11
#define TDS_DATA_NOMNY8			12
#define TDS_DATA_NOMNY4			13
#define TDS_DATA_NODATE8		14
#define TDS_DATA_NODATE4		15
#define TDS_DATA_NOFLT4			16
#define TDS_DATA_NOFLT8			17
#define TDS_DATA_NONUM			18
#define TDS_DATA_NOTEXT			19
#define TDS_DATA_NOIMAGE		20
#define TDS_DATA_NODEC			21
#define TDS_DATA_NOLCHAR		22
#define TDS_DATA_NOLBIN			23
#define TDS_DATA_NOINTN			24
#define TDS_DATA_NODATETIMEN		25
#define TDS_DATA_NOMONEYN		26
#define TDS_CON_NOOOB			27
#define TDS_CON_NOINBAND		28
#define TDS_PROTO_NOTEXT		29
#define TDS_PROTO_NOBULK		30
#define TDS_DATA_NOSENSITIVITY		31
#define TDS_DATA_NOBOUNDARY		32
#define TDS_RES_NOTDSDEBUG		33
#define TDS_RES_NOSTRIPBLANKS		34
#define TDS_DATA_NOINT8			35
#define TDS_OBJECT_NOJAVA1		36
#define TDS_OBJECT_NOCHAR		37
#define TDS_DATA_NOCOLUMNSTATUS		38
#define TDS_OBJECT_NOBINARY		39
#define TDS_RES_RESERVED		40
#define TDS_DATA_NOUINT2		41
#define TDS_DATA_NOUINT4		42
#define TDS_DATA_NOUINT8		43
#define TDS_DATA_NOUINTN		44
#define TDS_NOWIDETABLES		45
#define TDS_DATA_NONLBIN		46
#define TDS_IMAGE_NONCHAR		47
#define TDS_BLOB_NONCHAR_16		48
#define TDS_BLOB_NONCHAR_8		49
#define TDS_BLOB_NONCHAR_SCSU		50
#define TDS_DATA_NODATE			51
#define TDS_DATA_NOTIME			52
#define TDS_DATA_NOINTERVAL		53
#define TDS_DATA_NOUNITEXT		54
#define TDS_DATA_NOSINT1		55
#define TDS_NO_LARGEIDENT		56
#define TDS_NO_BLOB_NCHAR_16		57
#define TDS_NO_SRVPKTSIZE		58
#define TDS_DATA_NOXML			59
#define TDS_NONINT_RETURN_VALUE		60
#define TDS_RES_NOXNLMETADATA		61
#define TDS_RES_SUPPRESS_FMT		62
#define TDS_RES_SUPPRESS_DONEINPROC	63
#define TDS_UNUSED_RES			64
#define TDS_DATA_NOBIGDATETIME		65
#define TDS_DATA_NOUSECS		66
#define TDS_RES_NO_TDSCONTROL		67
#define TDS_RPCPARAM_NOLOB		68
#define TDS_DATA_NOLOBLOCATOR		69
#define TDS_RES_NOROWCOUNT_FOR_SELECT	70
#define	TDS_RES_LIST_DR_MAP		72
#define	TDS_RES_DR_NOKILL		73

/*
** Client/Server option commands and types. Types 1-200 are reserved for
** SAP CSI.
*/
#define TDS_OPT_SET		1
#define TDS_OPT_DEFAULT		2
#define TDS_OPT_LIST		3
#define TDS_OPT_INFO		4

/*
** All options take parameter
*/
#define TDS_OPT_UNUSED				0 /* Unused option */
#define TDS_OPT_DATEFIRST			1 /* Set first day of week */
#define TDS_OPT_TEXTSIZE			2 /* Text size */
#define	TDS_OPT_STAT_TIME			3 /* Server time statistics */
#define	TDS_OPT_STAT_IO				4 /* Server I/O statistics */
#define TDS_OPT_ROWCOUNT			5 /* Maximum row count */
#define TDS_OPT_NATLANG				6 /* National Language */
#define TDS_OPT_DATEFORMAT			7 /* Date format */
#define TDS_OPT_ISOLATION			8 /* Transaction isolation level */
#define TDS_OPT_AUTHON				9 /* Set authority level */
#define TDS_OPT_CHARSET				10 /* Character set */
#define TDS_OPT_PLAN				11 /* Define plan to use */
#define TDS_OPT_ERRLVL				12 /* Error level */
#define TDS_OPT_SHOWPLAN 			13 /* show execution plan	  */
#define TDS_OPT_NOEXEC 				14 /* don't execute query */
#define TDS_OPT_ARITHIGNORE			15 /* turn on arithmetic exceptions */
#define TDS_OPT_TRUNCIGNORE			16 /* turn on arithmetic exceptions */
#define TDS_OPT_ARITHABORT			17 /* turn on arithmetic abort */
#define TDS_OPT_PARSEONLY			18 /* parse only, return error msgs */
#define TDS_OPT_ESTIMATE			19 /* estimate of query time */
#define TDS_OPT_GETDATA				20 /* return trigger data */
#define TDS_OPT_NOCOUNT				21 /* don't print done count */
#define TDS_OPT_FORCEPLAN			23 /* force variable substitute order */
#define TDS_OPT_FORMATONLY			24 /* send format w/o row */
#define TDS_OPT_CHAINXACTS			25 /* chained transaction mode */
#define TDS_OPT_CURCLOSEONXACT			26 /* close cursor on end trans */
#define TDS_OPT_FIPSFLAG			27 /* FIPS flag */
#define TDS_OPT_RESTREES			28 /* return resolution trees */
#define	TDS_OPT_IDENTITYON			29 /* turn on explicit identity */
#define	TDS_OPT_CURREAD				30 /* Set session label @@curread */
#define	TDS_OPT_CURWRITE			31 /* Set session label @@curwrite */
#define	TDS_OPT_IDENTITYOFF			32 /* turn off explicit identity */
#define TDS_OPT_AUTHOFF				33 /* Set authority level off */
#define TDS_OPT_ANSINULL			34 /* ANSI NULLS behaviour */
#define TDS_OPT_QUOTED_IDENT			35 /* Quoted Identifiers */
#define TDS_OPT_ANSIPERM			36 /* ANSI permission checking */
#define TDS_OPT_STR_RTRUNC			37 /* ANSI right truncation */
#define TDS_OPT_SORTMERGE			38 /* Sort merge behaviour */
#define	TDS_OPT_JTC				39 /* Set JTC for session */
#define TDS_OPT_CLIENTREALNAME			40 /* Set Client Real Name */
#define TDS_OPT_CLIENTHOSTNAME			41 /* Set Client Host Name */
#define TDS_OPT_CLIENTAPPLNAME			42 /* Set Client Appl Name */
#define	TDS_OPT_IDENTITYUPD_ON  		43 /* turn on identity update */
#define	TDS_OPT_IDENTITYUPD_OFF 		44 /* turn off identity update */
#define TDS_OPT_NODATA				45 /* turn on/off "nodata" option */
#define TDS_OPT_CIPHERTEXT			46
#define TDS_OPT_SHOW_FI				47 /* Expose Functional Indexes */
#define TDS_OPT_HIDE_VCC			48 /* Hide Virtual Computed Columns */
#define TDS_OPT_LOBLOCATOR			49 /* Turn LOB Locators on */
#define TDS_OPT_LOBLOCATORFETCHSIZE		50 /* Set initial text fetchsize */
#define TDS_OPT_ISOLATION_MODE			52 /* Transaction isolation mode */

/*
** The supported options are summarized below
** with their defined values for `ArgLength' and `OptionArg'.
**
** Option			ArgLength	OptionArg
** ---------------		---------	---------
** TDS_OPT_DATEFIRST (1)	1 byte		Defines below
** TDS_OPT_TEXTSIZE (2)		4 bytes		Size in bytes
** TDS_OPT_ROWCOUNT (5) 	4 bytes		Number of rows
** TDS_OPT_NATLANG (6)		OptionArg Len	National Lang (string)
** TDS_OPT_DATEFORMAT (7)	1 byte		Defines below
** TDS_OPT_ISOLATION (8)	1 byte		Defines below
** TDS_OPT_CHARSET (10)		OptionArg Len	Character set (string)
** TDS_OPT_IDENTITYON (29)	OptionArg Len	Table Name (string)
** TDS_OPT_CURREAD (30)		OptionArg Len	Read Label(string)
** TDS_OPT_CURWRITE (31)	OptionArg Len	Write Label(string)
** TDS_OPT_IDENTITYOFF (32)	OptionArg Len	Table Name (string)
** TDS_OPT_AUTHON (9)		OptionArg Len	Table Name (string)
** TDS_OPT_AUTHOFF (33)		OptionArg Len	Table Name (string)
** TDS_OPT_IDENTITYUPD_ON (43)	OptionArg Len	Table Name (string)
** TDS_OPT_IDENTITYUPD_OFF (44)	OptionArg Len	Table Name (string)
** TDS_OPT_ISOLATION_MODE (51)	1 byte		Defines below
** (All remaining options)	1 byte		Boolean value defines below
**
** All string values must be sent in 7 bit ASCII.
**
*/

/*
** Boolean Value
*/
#define	TDS_OPT_FALSE			0
#define	TDS_OPT_TRUE			1

/*
** TDS_OPT_DATEFIRST
*/
#define TDS_OPT_MONDAY			1
#define TDS_OPT_TUESDAY			2
#define TDS_OPT_WEDNESDAY		3
#define TDS_OPT_THURSDAY		4
#define TDS_OPT_FRIDAY			5
#define TDS_OPT_SATURDAY		6
#define TDS_OPT_SUNDAY			7

/*
** TDS_OPT_DATEFORMAT
*/
#define TDS_OPT_FMTMDY			1
#define TDS_OPT_FMTDMY			2
#define TDS_OPT_FMTYMD			3
#define TDS_OPT_FMTYDM			4
#define TDS_OPT_FMTMYD			5
#define TDS_OPT_FMTDYM			6

/*
** TDS_OPT_ISOLATION
*/
#define TDS_OPT_LEVEL0				0
#define TDS_OPT_LEVEL1				1
#define TDS_OPT_LEVEL2				2
#define TDS_OPT_LEVEL3				3
#define TDS_OPT_STATEMENT_SNAPSHOT		11
#define TDS_OPT_TRANSACTION_SNAPSHOT		12
#define TDS_OPT_READONLY_STATEMENT_SNAPSHOT	13

/*
** TDS_OPT_ISOLATION_MODE
*/
#define TDS_OPT_MODE_DEFAULT		0
#define TDS_OPT_MODE_SNAPSHOT		1
#define TDS_OPT_MODE_READONLY_SNAPSHOT	2

/*
** Compute operators used in ALTFMT.
*/
#define	TDS_ALT_AVG			0x4F	/* The average value */
#define	TDS_ALT_COUNT			0x4B	/* The summary count value */
#define	TDS_ALT_MAX			0x52	/* The maximum value */
#define	TDS_ALT_MIN			0x51	/* The minimum value */
#define	TDS_ALT_SUM			0x4d	/* The sum value */

/*
** PARAMFMT status.
*/
#define TDS_PARAM_RETURN		0x01
#define TDS_PARAM_NULLALLOWED		0x20
#define TDS_PARAM_RETURN_NULLALLOWED    0x21

/*
** RPC options.
*/
#define	TDS_RPC_UNUSED			0x0000
#define TDS_RPC_RECOMPILE		0x0001

/*
** DBRPC/DBRPC2 options, in addition to the
** TDS_RPC_UNUSED and TDS_RPC_RECOMPILE values from above.
*/
#define TDS_RPC_PARAMS			0x0002
#define TDS_RPC_BATCH_PARAMS		0x0004

/*
** RPC status values.
*/
#define	TDS_RPC_STATUS_UNUSED		0x00
#define TDS_RPC_OUTPUT			0x01
#define TDS_RPC_NODEF			0x02

/*
** ROWFMT options.
*/
#define TDS_ROW_HIDDEN			0x01
#define TDS_ROW_KEY			0x02
#define TDS_ROW_VERSION			0x04
#define TDS_ROW_NODATA			0x08
#define TDS_ROW_UPDATABLE		0x10
#define TDS_ROW_NULLALLOWED		0x20
#define TDS_ROW_IDENTITY		0x40
#define TDS_ROW_PADCHAR			0x80
#define TDS_ROW_RETURN          	0xc0
#define TDS_ROW_RETURN_CANBENULL 	0xe0

/*
** COLINFO status bits for columns & parameters.
*/
#define	TDS_STAT_EXPR			0x04	/* COLINFO */
#define	TDS_STAT_KEY			0x08	/* COLINFO */
#define	TDS_STAT_HIDDEN			0x10	/* COLINFO */
#define	TDS_STAT_RENAME			0x20	/* COLINFO */

/*
** Cursor info status bit settings.
*/
#define	TDS_CUR_ISTAT_UNUSED		0x00000000
#define	TDS_CUR_ISTAT_DECLARED		0x00000001
#define	TDS_CUR_ISTAT_OPEN		0x00000002
#define	TDS_CUR_ISTAT_CLOSED		0x00000004
#define	TDS_CUR_ISTAT_RDONLY		0x00000008
#define	TDS_CUR_ISTAT_UPDATABLE		0x00000010
#define	TDS_CUR_ISTAT_ROWCNT		0x00000020
#define	TDS_CUR_ISTAT_DEALLOC		0x00000040
#define	TDS_CUR_ISTAT_SCROLLABLE	0x00000080
#define	TDS_CUR_ISTAT_IMPLICIT		0x00000100
#define	TDS_CUR_ISTAT_SENSITIVE		0x00000200
#define	TDS_CUR_ISTAT_INSENSITIVE	0x00000400
#define	TDS_CUR_ISTAT_SEMISENSITIVE	0x00000800
#define	TDS_CUR_ISTAT_KEYSETDRIVEN	0x00001000
#define	TDS_CUR_ISTAT_RELLOCKSONCLOSE	0x00002000

/*
** Cursor info enumerated commands.
*/
#define	TDS_CUR_CMD_SETCURROWS		1
#define	TDS_CUR_CMD_INQUIRE		2
#define	TDS_CUR_CMD_INFORM		3
#define	TDS_CUR_CMD_LISTALL		4

/*
** Cursor close options.
*/
#define	TDS_CUR_COPT_UNUSED		0x00
#define	TDS_CUR_COPT_DEALLOC		0x01

/*
** Cursor declare status bits.
*/
#define	TDS_CUR_DSTAT_UNUSED		0x00
#define	TDS_CUR_DSTAT_HASARGS		0x01

/*
** Cursor declare options.
*/
#define TDS_CUR_DOPT_UNUSED		0x00000000
#define TDS_CUR_DOPT_RDONLY		0x00000001
#define TDS_CUR_DOPT_UPDATABLE		0x00000002
#define TDS_CUR_DOPT_SENSITIVE		0x00000004 /* no support */
#define TDS_CUR_DOPT_DYNAMIC		0x00000008
#define TDS_CUR_DOPT_IMPLICIT		0x00000010
#define TDS_CUR_DOPT_INSENSITIVE	0x00000020
#define TDS_CUR_DOPT_SEMISENSITIVE	0x00000040
#define TDS_CUR_DOPT_KEYSETDRIVEN	0x00000080 /* no support */
#define TDS_CUR_DOPT_SCROLLABLE		0x00000100
#define TDS_CUR_DOPT_RELLOCKSONCLOSE	0x00000200

/*
** Cursor open status bits.
*/
#define	TDS_CUR_OSTAT_UNUSED		0x00
#define	TDS_CUR_OSTAT_HASARGS		0x01

/*
** Cursor fetch enumerated types.
*/
#define	TDS_CUR_NEXT			1
#define	TDS_CUR_PREV			2
#define	TDS_CUR_FIRST			3
#define	TDS_CUR_LAST			4
#define	TDS_CUR_ABS			5
#define	TDS_CUR_REL			6

/*
** Status bits for DONE.
** NOTE: use == to test TDS_DONE_FINAL.
*/
#define TDS_DONE_FINAL			0x0000
#define TDS_DONE_MORE			0x0001
#define TDS_DONE_ERROR			0x0002
#define TDS_DONE_INXACT			0x0004
#define TDS_DONE_PROC			0x0008
#define TDS_DONE_COUNT			0x0010
#define TDS_DONE_ATTN			0x0020
#define TDS_DONE_EVENT			0x0040

/*
** Command types for DYNAMIC
*/
#define TDS_DYN_PREPARE			0x01
#define TDS_DYN_EXEC			0x02
#define TDS_DYN_DEALLOC			0x04
#define TDS_DYN_EXEC_IMMED		0x08
#define TDS_DYN_PROCNAME		0x10
#define TDS_DYN_ACK			0x20
#define TDS_DYN_DESCIN			0x40
#define TDS_DYN_DESCOUT			0x80

/*
** Status values for DYNAMIC/DYNAMIC2
*/
#define TDS_DYNAMIC_UNUSED		0x00
#define TDS_DYNAMIC_HASARGS		0x01
#define TDS_DYNAMIC_SUPPRESS_FMT	0x02
#define TDS_DYNAMIC_BATCH_PARAMS	0x04
#define TDS_DYNAMIC_SUPPRESS_PARAMFMT	0x08

/*
** LANGUAGE status values.
*/
#define TDS_LANGUAGE_UNUSED             0x00
#define TDS_LANGUAGE_HASARGS            0x01
#define TDS_LANG_BATCH_PARAMS		0x04

/*
** Extended error status values
*/
#define TDS_NO_EED			0x00
#define TDS_EED_FOLLOWS			0x01
#define	TDS_EED_INFO			0x02

#endif /* __TDSPUBLIC_H */
