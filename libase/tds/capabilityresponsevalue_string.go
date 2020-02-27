// Code generated by "stringer -type=CapabilityResponseValue"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_RES_NOMSG-1]
	_ = x[TDS_RES_NOEED-2]
	_ = x[TDS_RES_NOPARAM-3]
	_ = x[TDS_DATA_NOINT1-4]
	_ = x[TDS_DATA_NOINT2-5]
	_ = x[TDS_DATA_NOINT4-6]
	_ = x[TDS_DATA_NOBIT-7]
	_ = x[TDS_DATA_NOCHAR-8]
	_ = x[TDS_DATA_NOVCHAR-9]
	_ = x[TDS_DATA_NOBIN-10]
	_ = x[TDS_DATA_NOVBIN-11]
	_ = x[TDS_DATA_NOMNY8-12]
	_ = x[TDS_DATA_NOMNY4-13]
	_ = x[TDS_DATA_NODATE8-14]
	_ = x[TDS_DATA_NODATE4-15]
	_ = x[TDS_DATA_NOFLT4-16]
	_ = x[TDS_DATA_NOFLT8-17]
	_ = x[TDS_DATA_NONUM-18]
	_ = x[TDS_DATA_NOTEXT-19]
	_ = x[TDS_DATA_NOIMAGE-20]
	_ = x[TDS_DATA_NODEC-21]
	_ = x[TDS_DATA_NOLCHAR-22]
	_ = x[TDS_DATA_NOLBIN-23]
	_ = x[TDS_DATA_NOINTN-24]
	_ = x[TDS_DATA_NODATETIMEN-25]
	_ = x[TDS_DATA_NOMONEYN-26]
	_ = x[TDS_CON_NOOOB-27]
	_ = x[TDS_CON_NOINBAND-28]
	_ = x[TDS_PROTO_NOTEXT-29]
	_ = x[TDS_PROTO_NOBULK-30]
	_ = x[TDS_DATA_NOSENSITIVITY-31]
	_ = x[TDS_DATA_NOBOUNDARY-32]
	_ = x[TDS_RES_NOTDSDEBUG-33]
	_ = x[TDS_RES_NOSTRIPBLANKS-34]
	_ = x[TDS_DATA_NOINT8-35]
	_ = x[TDS_OBJECT_NOJAVA1-36]
	_ = x[TDS_OBJECT_NOCHAR-37]
	_ = x[TDS_DATA_NOCOLUMNSTATUS-38]
	_ = x[TDS_OBJECT_NOBINARY-39]
	_ = x[TDS_RES_RESERVED-40]
	_ = x[TDS_DATA_NOUINT2-41]
	_ = x[TDS_DATA_NOUINT4-42]
	_ = x[TDS_DATA_NOUINT8-43]
	_ = x[TDS_DATA_NOUINTN-44]
	_ = x[TDS_NOWIDETABLES-45]
	_ = x[TDS_DATA_NONLBIN-46]
	_ = x[TDS_IMAGE_NONCHAR-47]
	_ = x[TDS_BLOB_NONCHAR_16-48]
	_ = x[TDS_BLOB_NONCHAR_8-49]
	_ = x[TDS_BLOB_NONCHAR_SCSU-50]
	_ = x[TDS_DATA_NODATE-51]
	_ = x[TDS_DATA_NOTIME-52]
	_ = x[TDS_DATA_NOINTERVAL-53]
	_ = x[TDS_DATA_NOUNITEXT-54]
	_ = x[TDS_DATA_NOSINT1-55]
	_ = x[TDS_NO_LARGEIDENT-56]
	_ = x[TDS_NO_BLOB_NCHAR_16-57]
	_ = x[TDS_NO_SRVPKTSIZE-58]
	_ = x[TDS_DATA_NOXML-59]
	_ = x[TDS_NONINT_RETURN_VALUE-60]
	_ = x[TDS_RES_NOXNLMETADATA-61]
	_ = x[TDS_RES_SUPPRESS_FMT-62]
	_ = x[TDS_RES_SUPPRESS_DONEINPROC-63]
	_ = x[TDS_UNUSED_RES-64]
	_ = x[TDS_DATA_NOBIGDATETIME-65]
	_ = x[TDS_DATA_NOUSECS-66]
	_ = x[TDS_RES_NO_TDSCONTROL-67]
	_ = x[TDS_RPCPARAM_NOLOB-68]
	_ = x[TDS_DATA_NOLOBLOCATOR-69]
	_ = x[TDS_RES_NOROWCOUNT_FOR_SELECT-70]
	_ = x[TDS_RES_LIST_DR_MAP-71]
	_ = x[TDS_RES_DR_NOKILL-72]
}

const _CapabilityResponseValue_name = "TDS_RES_NOMSGTDS_RES_NOEEDTDS_RES_NOPARAMTDS_DATA_NOINT1TDS_DATA_NOINT2TDS_DATA_NOINT4TDS_DATA_NOBITTDS_DATA_NOCHARTDS_DATA_NOVCHARTDS_DATA_NOBINTDS_DATA_NOVBINTDS_DATA_NOMNY8TDS_DATA_NOMNY4TDS_DATA_NODATE8TDS_DATA_NODATE4TDS_DATA_NOFLT4TDS_DATA_NOFLT8TDS_DATA_NONUMTDS_DATA_NOTEXTTDS_DATA_NOIMAGETDS_DATA_NODECTDS_DATA_NOLCHARTDS_DATA_NOLBINTDS_DATA_NOINTNTDS_DATA_NODATETIMENTDS_DATA_NOMONEYNTDS_CON_NOOOBTDS_CON_NOINBANDTDS_PROTO_NOTEXTTDS_PROTO_NOBULKTDS_DATA_NOSENSITIVITYTDS_DATA_NOBOUNDARYTDS_RES_NOTDSDEBUGTDS_RES_NOSTRIPBLANKSTDS_DATA_NOINT8TDS_OBJECT_NOJAVA1TDS_OBJECT_NOCHARTDS_DATA_NOCOLUMNSTATUSTDS_OBJECT_NOBINARYTDS_RES_RESERVEDTDS_DATA_NOUINT2TDS_DATA_NOUINT4TDS_DATA_NOUINT8TDS_DATA_NOUINTNTDS_NOWIDETABLESTDS_DATA_NONLBINTDS_IMAGE_NONCHARTDS_BLOB_NONCHAR_16TDS_BLOB_NONCHAR_8TDS_BLOB_NONCHAR_SCSUTDS_DATA_NODATETDS_DATA_NOTIMETDS_DATA_NOINTERVALTDS_DATA_NOUNITEXTTDS_DATA_NOSINT1TDS_NO_LARGEIDENTTDS_NO_BLOB_NCHAR_16TDS_NO_SRVPKTSIZETDS_DATA_NOXMLTDS_NONINT_RETURN_VALUETDS_RES_NOXNLMETADATATDS_RES_SUPPRESS_FMTTDS_RES_SUPPRESS_DONEINPROCTDS_UNUSED_RESTDS_DATA_NOBIGDATETIMETDS_DATA_NOUSECSTDS_RES_NO_TDSCONTROLTDS_RPCPARAM_NOLOBTDS_DATA_NOLOBLOCATORTDS_RES_NOROWCOUNT_FOR_SELECTTDS_RES_LIST_DR_MAPTDS_RES_DR_NOKILL"

var _CapabilityResponseValue_index = [...]uint16{0, 13, 26, 41, 56, 71, 86, 100, 115, 131, 145, 160, 175, 190, 206, 222, 237, 252, 266, 281, 297, 311, 327, 342, 357, 377, 394, 407, 423, 439, 455, 477, 496, 514, 535, 550, 568, 585, 608, 627, 643, 659, 675, 691, 707, 723, 739, 756, 775, 793, 814, 829, 844, 863, 881, 897, 914, 934, 951, 965, 988, 1009, 1029, 1056, 1070, 1092, 1108, 1129, 1147, 1168, 1197, 1216, 1233}

func (i CapabilityResponseValue) String() string {
	i -= 1
	if i < 0 || i >= CapabilityResponseValue(len(_CapabilityResponseValue_index)-1) {
		return "CapabilityResponseValue(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CapabilityResponseValue_name[_CapabilityResponseValue_index[i]:_CapabilityResponseValue_index[i+1]]
}
