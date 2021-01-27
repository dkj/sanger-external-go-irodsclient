package common

import "fmt"

// ErrorCode is an error code type
type ErrorCode int

// error codes
const (
	SYS_SOCK_OPEN_ERR                        ErrorCode = -1000
	SYS_SOCK_BIND_ERR                        ErrorCode = -2000
	SYS_SOCK_ACCEPT_ERR                      ErrorCode = -3000
	SYS_HEADER_READ_LEN_ERR                  ErrorCode = -4000
	SYS_HEADER_WRITE_LEN_ERR                 ErrorCode = -5000
	SYS_HEADER_TPYE_LEN_ERR                  ErrorCode = -6000
	SYS_CAUGHT_SIGNAL                        ErrorCode = -7000
	SYS_GETSTARTUP_PACK_ERR                  ErrorCode = -8000
	SYS_EXCEED_CONNECT_CNT                   ErrorCode = -9000
	SYS_USER_NOT_ALLOWED_TO_CONN             ErrorCode = -10000
	SYS_READ_MSG_BODY_INPUT_ERR              ErrorCode = -11000
	SYS_UNMATCHED_API_NUM                    ErrorCode = -12000
	SYS_NO_API_PRIV                          ErrorCode = -13000
	SYS_API_INPUT_ERR                        ErrorCode = -14000
	SYS_PACK_INSTRUCT_FORMAT_ERR             ErrorCode = -15000
	SYS_MALLOC_ERR                           ErrorCode = -16000
	SYS_GET_HOSTNAME_ERR                     ErrorCode = -17000
	SYS_OUT_OF_FILE_DESC                     ErrorCode = -18000
	SYS_FILE_DESC_OUT_OF_RANGE               ErrorCode = -19000
	SYS_UNRECOGNIZED_REMOTE_FLAG             ErrorCode = -20000
	SYS_INVALID_SERVER_HOST                  ErrorCode = -21000
	SYS_SVR_TO_SVR_CONNECT_FAILED            ErrorCode = -22000
	SYS_BAD_FILE_DESCRIPTOR                  ErrorCode = -23000
	SYS_INTERNAL_NULL_INPUT_ERR              ErrorCode = -24000
	SYS_CONFIG_FILE_ERR                      ErrorCode = -25000
	SYS_INVALID_ZONE_NAME                    ErrorCode = -26000
	SYS_COPY_LEN_ERR                         ErrorCode = -27000
	SYS_PORT_COOKIE_ERR                      ErrorCode = -28000
	SYS_KEY_VAL_TABLE_ERR                    ErrorCode = -29000
	SYS_INVALID_RESC_TYPE                    ErrorCode = -30000
	SYS_INVALID_FILE_PATH                    ErrorCode = -31000
	SYS_INVALID_RESC_INPUT                   ErrorCode = -32000
	SYS_INVALID_PORTAL_OPR                   ErrorCode = -33000
	SYS_PARA_OPR_NO_SUPPORT                  ErrorCode = -34000
	SYS_INVALID_OPR_TYPE                     ErrorCode = -35000
	SYS_NO_PATH_PERMISSION                   ErrorCode = -36000
	SYS_NO_ICAT_SERVER_ERR                   ErrorCode = -37000
	SYS_AGENT_INIT_ERR                       ErrorCode = -38000
	SYS_PROXYUSER_NO_PRIV                    ErrorCode = -39000
	SYS_NO_DATA_OBJ_PERMISSION               ErrorCode = -40000
	SYS_DELETE_DISALLOWED                    ErrorCode = -41000
	SYS_OPEN_REI_FILE_ERR                    ErrorCode = -42000
	SYS_NO_RCAT_SERVER_ERR                   ErrorCode = -43000
	SYS_UNMATCH_PACK_INSTRUCTI_NAME          ErrorCode = -44000
	SYS_SVR_TO_CLI_MSI_NO_EXIST              ErrorCode = -45000
	SYS_COPY_ALREADY_IN_RESC                 ErrorCode = -46000
	SYS_RECONN_OPR_MISMATCH                  ErrorCode = -47000
	SYS_INPUT_PERM_OUT_OF_RANGE              ErrorCode = -48000
	SYS_FORK_ERROR                           ErrorCode = -49000
	SYS_PIPE_ERROR                           ErrorCode = -50000
	SYS_EXEC_CMD_STATUS_SZ_ERROR             ErrorCode = -51000
	SYS_PATH_IS_NOT_A_FILE                   ErrorCode = -52000
	SYS_UNMATCHED_SPEC_COLL_TYPE             ErrorCode = -53000
	SYS_TOO_MANY_QUERY_RESULT                ErrorCode = -54000
	SYS_SPEC_COLL_NOT_IN_CACHE               ErrorCode = -55000
	SYS_SPEC_COLL_OBJ_NOT_EXIST              ErrorCode = -56000
	SYS_REG_OBJ_IN_SPEC_COLL                 ErrorCode = -57000
	SYS_DEST_SPEC_COLL_SUB_EXIST             ErrorCode = -58000
	SYS_SRC_DEST_SPEC_COLL_CONFLICT          ErrorCode = -59000
	SYS_UNKNOWN_SPEC_COLL_CLASS              ErrorCode = -60000
	SYS_DUPLICATE_XMSG_TICKET                ErrorCode = -61000
	SYS_UNMATCHED_XMSG_TICKET                ErrorCode = -62000
	SYS_NO_XMSG_FOR_MSG_NUMBER               ErrorCode = -63000
	SYS_COLLINFO_2_FORMAT_ERR                ErrorCode = -64000
	SYS_CACHE_STRUCT_FILE_RESC_ERR           ErrorCode = -65000
	SYS_NOT_SUPPORTED                        ErrorCode = -66000
	SYS_TAR_STRUCT_FILE_EXTRACT_ERR          ErrorCode = -67000
	SYS_STRUCT_FILE_DESC_ERR                 ErrorCode = -68000
	SYS_TAR_OPEN_ERR                         ErrorCode = -69000
	SYS_TAR_EXTRACT_ALL_ERR                  ErrorCode = -70000
	SYS_TAR_CLOSE_ERR                        ErrorCode = -71000
	SYS_STRUCT_FILE_PATH_ERR                 ErrorCode = -72000
	SYS_MOUNT_MOUNTED_COLL_ERR               ErrorCode = -73000
	SYS_COLL_NOT_MOUNTED_ERR                 ErrorCode = -74000
	SYS_STRUCT_FILE_BUSY_ERR                 ErrorCode = -75000
	SYS_STRUCT_FILE_INMOUNTED_COLL           ErrorCode = -76000
	SYS_COPY_NOT_EXIST_IN_RESC               ErrorCode = -77000
	SYS_RESC_DOES_NOT_EXIST                  ErrorCode = -78000
	SYS_COLLECTION_NOT_EMPTY                 ErrorCode = -79000
	SYS_OBJ_TYPE_NOT_STRUCT_FILE             ErrorCode = -80000
	SYS_WRONG_RESC_POLICY_FOR_BUN_OPR        ErrorCode = -81000
	SYS_DIR_IN_VAULT_NOT_EMPTY               ErrorCode = -82000
	SYS_OPR_FLAG_NOT_SUPPORT                 ErrorCode = -83000
	SYS_TAR_APPEND_ERR                       ErrorCode = -84000
	SYS_INVALID_PROTOCOL_TYPE                ErrorCode = -85000
	SYS_UDP_CONNECT_ERR                      ErrorCode = -86000
	SYS_UDP_TRANSFER_ERR                     ErrorCode = -89000
	SYS_UDP_NO_SUPPORT_ERR                   ErrorCode = -90000
	SYS_READ_MSG_BODY_LEN_ERR                ErrorCode = -91000
	CROSS_ZONE_SOCK_CONNECT_ERR              ErrorCode = -92000
	SYS_NO_FREE_RE_THREAD                    ErrorCode = -93000
	SYS_BAD_RE_THREAD_INX                    ErrorCode = -94000
	SYS_CANT_DIRECTLY_ACC_COMPOUND_RESC      ErrorCode = -95000
	SYS_SRC_DEST_RESC_COMPOUND_TYPE          ErrorCode = -96000
	SYS_CACHE_RESC_NOT_ON_SAME_HOST          ErrorCode = -97000
	SYS_NO_CACHE_RESC_IN_GRP                 ErrorCode = -98000
	SYS_UNMATCHED_RESC_IN_RESC_GRP           ErrorCode = -99000
	SYS_CANT_MV_BUNDLE_DATA_TO_TRASH         ErrorCode = -100000
	SYS_CANT_MV_BUNDLE_DATA_BY_COPY          ErrorCode = -101000
	SYS_EXEC_TAR_ERR                         ErrorCode = -102000
	SYS_CANT_CHKSUM_COMP_RESC_DATA           ErrorCode = -103000
	SYS_CANT_CHKSUM_BUNDLED_DATA             ErrorCode = -104000
	SYS_RESC_IS_DOWN                         ErrorCode = -105000
	SYS_UPDATE_REPL_INFO_ERR                 ErrorCode = -106000
	SYS_COLL_LINK_PATH_ERR                   ErrorCode = -107000
	SYS_LINK_CNT_EXCEEDED_ERR                ErrorCode = -108000
	SYS_CROSS_ZONE_MV_NOT_SUPPORTED          ErrorCode = -109000
	SYS_RESC_QUOTA_EXCEEDED                  ErrorCode = -110000
	USER_AUTH_SCHEME_ERR                     ErrorCode = -300000
	USER_AUTH_STRING_EMPTY                   ErrorCode = -301000
	USER_RODS_HOST_EMPTY                     ErrorCode = -302000
	USER_RODS_HOSTNAME_ERR                   ErrorCode = -303000
	USER_SOCK_OPEN_ERR                       ErrorCode = -304000
	USER_SOCK_CONNECT_ERR                    ErrorCode = -305000
	USER_STRLEN_TOOLONG                      ErrorCode = -306000
	USER_API_INPUT_ERR                       ErrorCode = -307000
	USER_PACKSTRUCT_INPUT_ERR                ErrorCode = -308000
	USER_NO_SUPPORT_ERR                      ErrorCode = -309000
	USER_FILE_DOES_NOT_EXIST                 ErrorCode = -310000
	USER_FILE_TOO_LARGE                      ErrorCode = -311000
	OVERWRITE_WITHOUT_FORCE_FLAG             ErrorCode = -312000
	UNMATCHED_KEY_OR_INDEX                   ErrorCode = -313000
	USER_CHKSUM_MISMATCH                     ErrorCode = -314000
	USER_BAD_KEYWORD_ERR                     ErrorCode = -315000
	USER__NULL_INPUT_ERR                     ErrorCode = -316000
	USER_INPUT_PATH_ERR                      ErrorCode = -317000
	USER_INPUT_OPTION_ERR                    ErrorCode = -318000
	USER_INVALID_USERNAME_FORMAT             ErrorCode = -319000
	USER_DIRECT_RESC_INPUT_ERR               ErrorCode = -320000
	USER_NO_RESC_INPUT_ERR                   ErrorCode = -321000
	USER_PARAM_LABEL_ERR                     ErrorCode = -322000
	USER_PARAM_TYPE_ERR                      ErrorCode = -323000
	BASE64_BUFFER_OVERFLOW                   ErrorCode = -324000
	BASE64_INVALID_PACKET                    ErrorCode = -325000
	USER_MSG_TYPE_NO_SUPPORT                 ErrorCode = -326000
	USER_RSYNC_NO_MODE_INPUT_ERR             ErrorCode = -337000
	USER_OPTION_INPUT_ERR                    ErrorCode = -338000
	SAME_SRC_DEST_PATHS_ERR                  ErrorCode = -339000
	USER_RESTART_FILE_INPUT_ERR              ErrorCode = -340000
	RESTART_OPR_FAILED                       ErrorCode = -341000
	BAD_EXEC_CMD_PATH                        ErrorCode = -342000
	EXEC_CMD_OUTPUT_TOO_LARGE                ErrorCode = -343000
	EXEC_CMD_ERROR                           ErrorCode = -344000
	BAD_INPUT_DESC_INDEX                     ErrorCode = -345000
	USER_PATH_EXCEEDS_MAX                    ErrorCode = -346000
	USER_SOCK_CONNECT_TIMEDOUT               ErrorCode = -347000
	USER_API_VERSION_MISMATCH                ErrorCode = -348000
	USER_INPUT_FORMAT_ERR                    ErrorCode = -349000
	USER_ACCESS_DENIED                       ErrorCode = -350000
	CANT_RM_MV_BUNDLE_TYPE                   ErrorCode = -351000
	NO_MORE_RESULT                           ErrorCode = -352000
	NO_KEY_WD_IN_MS_INP_STR                  ErrorCode = -353000
	CANT_RM_NON_EMPTY_HOME_COLL              ErrorCode = -354000
	CANT_UNREG_IN_VAULT_FILE                 ErrorCode = -355000
	NO_LOCAL_FILE_RSYNC_IN_MSI               ErrorCode = -356000
	FILE_INDEX_LOOKUP_ERR                    ErrorCode = -500000
	UNIX_FILE_OPEN_ERR                       ErrorCode = -510000
	UNIX_FILE_CREATE_ERR                     ErrorCode = -511000
	UNIX_FILE_READ_ERR                       ErrorCode = -512000
	UNIX_FILE_WRITE_ERR                      ErrorCode = -513000
	UNIX_FILE_CLOSE_ERR                      ErrorCode = -514000
	UNIX_FILE_UNLINK_ERR                     ErrorCode = -515000
	UNIX_FILE_STAT_ERR                       ErrorCode = -516000
	UNIX_FILE_FSTAT_ERR                      ErrorCode = -517000
	UNIX_FILE_LSEEK_ERR                      ErrorCode = -518000
	UNIX_FILE_FSYNC_ERR                      ErrorCode = -519000
	UNIX_FILE_MKDIR_ERR                      ErrorCode = -520000
	UNIX_FILE_RMDIR_ERR                      ErrorCode = -521000
	UNIX_FILE_OPENDIR_ERR                    ErrorCode = -522000
	UNIX_FILE_CLOSEDIR_ERR                   ErrorCode = -523000
	UNIX_FILE_READDIR_ERR                    ErrorCode = -524000
	UNIX_FILE_STAGE_ERR                      ErrorCode = -525000
	UNIX_FILE_GET_FS_FREESPACE_ERR           ErrorCode = -526000
	UNIX_FILE_CHMOD_ERR                      ErrorCode = -527000
	UNIX_FILE_RENAME_ERR                     ErrorCode = -528000
	UNIX_FILE_TRUNCATE_ERR                   ErrorCode = -529000
	UNIX_FILE_LINK_ERR                       ErrorCode = -530000
	UNIV_MSS_SYNCTOARCH_ERR                  ErrorCode = -550000
	UNIV_MSS_STAGETOCACHE_ERR                ErrorCode = -551000
	UNIV_MSS_UNLINK_ERR                      ErrorCode = -552000
	UNIV_MSS_MKDIR_ERR                       ErrorCode = -553000
	UNIV_MSS_CHMOD_ERR                       ErrorCode = -554000
	UNIV_MSS_STAT_ERR                        ErrorCode = -555000
	HPSS_AUTH_NOT_SUPPORTED                  ErrorCode = -600000
	HPSS_FILE_OPEN_ERR                       ErrorCode = -610000
	HPSS_FILE_CREATE_ERR                     ErrorCode = -611000
	HPSS_FILE_READ_ERR                       ErrorCode = -612000
	HPSS_FILE_WRITE_ERR                      ErrorCode = -613000
	HPSS_FILE_CLOSE_ERR                      ErrorCode = -614000
	HPSS_FILE_UNLINK_ERR                     ErrorCode = -615000
	HPSS_FILE_STAT_ERR                       ErrorCode = -616000
	HPSS_FILE_FSTAT_ERR                      ErrorCode = -617000
	HPSS_FILE_LSEEK_ERR                      ErrorCode = -618000
	HPSS_FILE_FSYNC_ERR                      ErrorCode = -619000
	HPSS_FILE_MKDIR_ERR                      ErrorCode = -620000
	HPSS_FILE_RMDIR_ERR                      ErrorCode = -621000
	HPSS_FILE_OPENDIR_ERR                    ErrorCode = -622000
	HPSS_FILE_CLOSEDIR_ERR                   ErrorCode = -623000
	HPSS_FILE_READDIR_ERR                    ErrorCode = -624000
	HPSS_FILE_STAGE_ERR                      ErrorCode = -625000
	HPSS_FILE_GET_FS_FREESPACE_ERR           ErrorCode = -626000
	HPSS_FILE_CHMOD_ERR                      ErrorCode = -627000
	HPSS_FILE_RENAME_ERR                     ErrorCode = -628000
	HPSS_FILE_TRUNCATE_ERR                   ErrorCode = -629000
	HPSS_FILE_LINK_ERR                       ErrorCode = -630000
	HPSS_AUTH_ERR                            ErrorCode = -631000
	HPSS_WRITE_LIST_ERR                      ErrorCode = -632000
	HPSS_READ_LIST_ERR                       ErrorCode = -633000
	HPSS_TRANSFER_ERR                        ErrorCode = -634000
	HPSS_MOVER_PROT_ERR                      ErrorCode = -635000
	S3_INIT_ERROR                            ErrorCode = -701000
	S3_PUT_ERROR                             ErrorCode = -702000
	S3_GET_ERROR                             ErrorCode = -703000
	S3_FILE_UNLINK_ERR                       ErrorCode = -715000
	S3_FILE_STAT_ERR                         ErrorCode = -716000
	S3_FILE_COPY_ERR                         ErrorCode = -717000
	CATALOG_NOT_CONNECTED                    ErrorCode = -801000
	CAT_ENV_ERR                              ErrorCode = -802000
	CAT_CONNECT_ERR                          ErrorCode = -803000
	CAT_DISCONNECT_ERR                       ErrorCode = -804000
	CAT_CLOSE_ENV_ERR                        ErrorCode = -805000
	CAT_SQL_ERR                              ErrorCode = -806000
	CAT_GET_ROW_ERR                          ErrorCode = -807000
	CAT_NO_ROWS_FOUND                        ErrorCode = -808000
	CATALOG_ALREADY_HAS_ITEM_BY_THAT_NAME    ErrorCode = -809000
	CAT_INVALID_RESOURCE_TYPE                ErrorCode = -810000
	CAT_INVALID_RESOURCE_CLASS               ErrorCode = -811000
	CAT_INVALID_RESOURCE_NET_ADDR            ErrorCode = -812000
	CAT_INVALID_RESOURCE_VAULT_PATH          ErrorCode = -813000
	CAT_UNKNOWN_COLLECTION                   ErrorCode = -814000
	CAT_INVALID_DATA_TYPE                    ErrorCode = -815000
	CAT_INVALID_ARGUMENT                     ErrorCode = -816000
	CAT_UNKNOWN_FILE                         ErrorCode = -817000
	CAT_NO_ACCESS_PERMISSION                 ErrorCode = -818000
	CAT_SUCCESS_BUT_WITH_NO_INFO             ErrorCode = -819000
	CAT_INVALID_USER_TYPE                    ErrorCode = -820000
	CAT_COLLECTION_NOT_EMPTY                 ErrorCode = -821000
	CAT_TOO_MANY_TABLES                      ErrorCode = -822000
	CAT_UNKNOWN_TABLE                        ErrorCode = -823000
	CAT_NOT_OPEN                             ErrorCode = -824000
	CAT_FAILED_TO_LINK_TABLES                ErrorCode = -825000
	CAT_INVALID_AUTHENTICATION               ErrorCode = -826000
	CAT_INVALID_USER                         ErrorCode = -827000
	CAT_INVALID_ZONE                         ErrorCode = -828000
	CAT_INVALID_GROUP                        ErrorCode = -829000
	CAT_INSUFFICIENT_PRIVILEGE_LEVEL         ErrorCode = -830000
	CAT_INVALID_RESOURCE                     ErrorCode = -831000
	CAT_INVALID_CLIENT_USER                  ErrorCode = -832000
	CAT_NAME_EXISTS_AS_COLLECTION            ErrorCode = -833000
	CAT_NAME_EXISTS_AS_DATAOBJ               ErrorCode = -834000
	CAT_RESOURCE_NOT_EMPTY                   ErrorCode = -835000
	CAT_NOT_A_DATAOBJ_AND_NOT_A_COLLECTION   ErrorCode = -836000
	CAT_RECURSIVE_MOVE                       ErrorCode = -837000
	CAT_LAST_REPLICA                         ErrorCode = -838000
	CAT_OCI_ERROR                            ErrorCode = -839000
	CAT_PASSWORD_EXPIRED                     ErrorCode = -840000
	CAT_PASSWORD_ENCODING_ERROR              ErrorCode = -850000
	CAT_TABLE_ACCESS_DENIED                  ErrorCode = -851000
	CAT_UNKNOWN_SPECIFIC_QUERY               ErrorCode = -853000
	CAT_STATEMENT_TABLE_FULL                 ErrorCode = -860000
	RDA_NOT_COMPILED_IN                      ErrorCode = -880000
	RDA_NOT_CONNECTED                        ErrorCode = -881000
	RDA_ENV_ERR                              ErrorCode = -882000
	RDA_CONNECT_ERR                          ErrorCode = -883000
	RDA_DISCONNECT_ERR                       ErrorCode = -884000
	RDA_CLOSE_ENV_ERR                        ErrorCode = -885000
	RDA_SQL_ERR                              ErrorCode = -886000
	RDA_CONFIG_FILE_ERR                      ErrorCode = -887000
	RDA_ACCESS_PROHIBITED                    ErrorCode = -888000
	RDA_NAME_NOT_FOUND                       ErrorCode = -889000
	FILE_OPEN_ERR                            ErrorCode = -900000
	FILE_READ_ERR                            ErrorCode = -901000
	FILE_WRITE_ERR                           ErrorCode = -902000
	PASSWORD_EXCEEDS_MAX_SIZE                ErrorCode = -903000
	ENVIRONMENT_VAR_HOME_NOT_DEFINED         ErrorCode = -904000
	UNABLE_TO_STAT_FILE                      ErrorCode = -905000
	AUTH_FILE_NOT_ENCRYPTED                  ErrorCode = -906000
	AUTH_FILE_DOES_NOT_EXIST                 ErrorCode = -907000
	UNLINK_FAILED                            ErrorCode = -908000
	NO_PASSWORD_ENTERED                      ErrorCode = -909000
	REMOTE_SERVER_AUTHENTICATION_FAILURE     ErrorCode = -910000
	REMOTE_SERVER_AUTH_NOT_PROVIDED          ErrorCode = -911000
	REMOTE_SERVER_AUTH_EMPTY                 ErrorCode = -912000
	REMOTE_SERVER_SID_NOT_DEFINED            ErrorCode = -913000
	GSI_NOT_COMPILED_IN                      ErrorCode = -921000
	GSI_NOT_BUILT_INTO_CLIENT                ErrorCode = -922000
	GSI_NOT_BUILT_INTO_SERVER                ErrorCode = -923000
	GSI_ERROR_IMPORT_NAME                    ErrorCode = -924000
	GSI_ERROR_INIT_SECURITY_CONTEXT          ErrorCode = -925000
	GSI_ERROR_SENDING_TOKEN_LENGTH           ErrorCode = -926000
	GSI_ERROR_READING_TOKEN_LENGTH           ErrorCode = -927000
	GSI_ERROR_TOKEN_TOO_LARGE                ErrorCode = -928000
	GSI_ERROR_BAD_TOKEN_RCVED                ErrorCode = -929000
	GSI_SOCKET_READ_ERROR                    ErrorCode = -930000
	GSI_PARTIAL_TOKEN_READ                   ErrorCode = -931000
	GSI_SOCKET_WRITE_ERROR                   ErrorCode = -932000
	GSI_ERROR_FROM_GSI_LIBRARY               ErrorCode = -933000
	GSI_ERROR_IMPORTING_NAME                 ErrorCode = -934000
	GSI_ERROR_ACQUIRING_CREDS                ErrorCode = -935000
	GSI_ACCEPT_SEC_CONTEXT_ERROR             ErrorCode = -936000
	GSI_ERROR_DISPLAYING_NAME                ErrorCode = -937000
	GSI_ERROR_RELEASING_NAME                 ErrorCode = -938000
	GSI_DN_DOES_NOT_MATCH_USER               ErrorCode = -939000
	GSI_QUERY_INTERNAL_ERROR                 ErrorCode = -940000
	GSI_NO_MATCHING_DN_FOUND                 ErrorCode = -941000
	GSI_MULTIPLE_MATCHING_DN_FOUND           ErrorCode = -942000
	KRB_NOT_COMPILED_IN                      ErrorCode = -951000
	KRB_NOT_BUILT_INTO_CLIENT                ErrorCode = -952000
	KRB_NOT_BUILT_INTO_SERVER                ErrorCode = -953000
	KRB_ERROR_IMPORT_NAME                    ErrorCode = -954000
	KRB_ERROR_INIT_SECURITY_CONTEXT          ErrorCode = -955000
	KRB_ERROR_SENDING_TOKEN_LENGTH           ErrorCode = -956000
	KRB_ERROR_READING_TOKEN_LENGTH           ErrorCode = -957000
	KRB_ERROR_TOKEN_TOO_LARGE                ErrorCode = -958000
	KRB_ERROR_BAD_TOKEN_RCVED                ErrorCode = -959000
	KRB_SOCKET_READ_ERROR                    ErrorCode = -960000
	KRB_PARTIAL_TOKEN_READ                   ErrorCode = -961000
	KRB_SOCKET_WRITE_ERROR                   ErrorCode = -962000
	KRB_ERROR_FROM_KRB_LIBRARY               ErrorCode = -963000
	KRB_ERROR_IMPORTING_NAME                 ErrorCode = -964000
	KRB_ERROR_ACQUIRING_CREDS                ErrorCode = -965000
	KRB_ACCEPT_SEC_CONTEXT_ERROR             ErrorCode = -966000
	KRB_ERROR_DISPLAYING_NAME                ErrorCode = -967000
	KRB_ERROR_RELEASING_NAME                 ErrorCode = -968000
	KRB_USER_DN_NOT_FOUND                    ErrorCode = -969000
	KRB_NAME_MATCHES_MULTIPLE_USERS          ErrorCode = -970000
	KRB_QUERY_INTERNAL_ERROR                 ErrorCode = -971000
	OBJPATH_EMPTY_IN_STRUCT_ERR              ErrorCode = -1000000
	RESCNAME_EMPTY_IN_STRUCT_ERR             ErrorCode = -1001000
	DATATYPE_EMPTY_IN_STRUCT_ERR             ErrorCode = -1002000
	DATASIZE_EMPTY_IN_STRUCT_ERR             ErrorCode = -1003000
	CHKSUM_EMPTY_IN_STRUCT_ERR               ErrorCode = -1004000
	VERSION_EMPTY_IN_STRUCT_ERR              ErrorCode = -1005000
	FILEPATH_EMPTY_IN_STRUCT_ERR             ErrorCode = -1006000
	REPLNUM_EMPTY_IN_STRUCT_ERR              ErrorCode = -1007000
	REPLSTATUS_EMPTY_IN_STRUCT_ERR           ErrorCode = -1008000
	DATAOWNER_EMPTY_IN_STRUCT_ERR            ErrorCode = -1009000
	DATAOWNERZONE_EMPTY_IN_STRUCT_ERR        ErrorCode = -1010000
	DATAEXPIRY_EMPTY_IN_STRUCT_ERR           ErrorCode = -1011000
	DATACOMMENTS_EMPTY_IN_STRUCT_ERR         ErrorCode = -1012000
	DATACREATE_EMPTY_IN_STRUCT_ERR           ErrorCode = -1013000
	DATAMODIFY_EMPTY_IN_STRUCT_ERR           ErrorCode = -1014000
	DATAACCESS_EMPTY_IN_STRUCT_ERR           ErrorCode = -1015000
	DATAACCESSINX_EMPTY_IN_STRUCT_ERR        ErrorCode = -1016000
	NO_RULE_FOUND_ERR                        ErrorCode = -1017000
	NO_MORE_RULES_ERR                        ErrorCode = -1018000
	UNMATCHED_ACTION_ERR                     ErrorCode = -1019000
	RULES_FILE_READ_ERROR                    ErrorCode = -1020000
	ACTION_ARG_COUNT_MISMATCH                ErrorCode = -1021000
	MAX_NUM_OF_ARGS_IN_ACTION_EXCEEDED       ErrorCode = -1022000
	UNKNOWN_PARAM_IN_RULE_ERR                ErrorCode = -1023000
	DESTRESCNAME_EMPTY_IN_STRUCT_ERR         ErrorCode = -1024000
	BACKUPRESCNAME_EMPTY_IN_STRUCT_ERR       ErrorCode = -1025000
	DATAID_EMPTY_IN_STRUCT_ERR               ErrorCode = -1026000
	COLLID_EMPTY_IN_STRUCT_ERR               ErrorCode = -1027000
	RESCGROUPNAME_EMPTY_IN_STRUCT_ERR        ErrorCode = -1028000
	STATUSSTRING_EMPTY_IN_STRUCT_ERR         ErrorCode = -1029000
	DATAMAPID_EMPTY_IN_STRUCT_ERR            ErrorCode = -1030000
	USERNAMECLIENT_EMPTY_IN_STRUCT_ERR       ErrorCode = -1031000
	RODSZONECLIENT_EMPTY_IN_STRUCT_ERR       ErrorCode = -1032000
	USERTYPECLIENT_EMPTY_IN_STRUCT_ERR       ErrorCode = -1033000
	HOSTCLIENT_EMPTY_IN_STRUCT_ERR           ErrorCode = -1034000
	AUTHSTRCLIENT_EMPTY_IN_STRUCT_ERR        ErrorCode = -1035000
	USERAUTHSCHEMECLIENT_EMPTY_IN_STRUCT_ERR ErrorCode = -1036000
	USERINFOCLIENT_EMPTY_IN_STRUCT_ERR       ErrorCode = -1037000
	USERCOMMENTCLIENT_EMPTY_IN_STRUCT_ERR    ErrorCode = -1038000
	USERCREATECLIENT_EMPTY_IN_STRUCT_ERR     ErrorCode = -1039000
	USERMODIFYCLIENT_EMPTY_IN_STRUCT_ERR     ErrorCode = -1040000
	USERNAMEPROXY_EMPTY_IN_STRUCT_ERR        ErrorCode = -1041000
	RODSZONEPROXY_EMPTY_IN_STRUCT_ERR        ErrorCode = -1042000
	USERTYPEPROXY_EMPTY_IN_STRUCT_ERR        ErrorCode = -1043000
	HOSTPROXY_EMPTY_IN_STRUCT_ERR            ErrorCode = -1044000
	AUTHSTRPROXY_EMPTY_IN_STRUCT_ERR         ErrorCode = -1045000
	USERAUTHSCHEMEPROXY_EMPTY_IN_STRUCT_ERR  ErrorCode = -1046000
	USERINFOPROXY_EMPTY_IN_STRUCT_ERR        ErrorCode = -1047000
	USERCOMMENTPROXY_EMPTY_IN_STRUCT_ERR     ErrorCode = -1048000
	USERCREATEPROXY_EMPTY_IN_STRUCT_ERR      ErrorCode = -1049000
	USERMODIFYPROXY_EMPTY_IN_STRUCT_ERR      ErrorCode = -1050000
	COLLNAME_EMPTY_IN_STRUCT_ERR             ErrorCode = -1051000
	COLLPARENTNAME_EMPTY_IN_STRUCT_ERR       ErrorCode = -1052000
	COLLOWNERNAME_EMPTY_IN_STRUCT_ERR        ErrorCode = -1053000
	COLLOWNERZONE_EMPTY_IN_STRUCT_ERR        ErrorCode = -1054000
	COLLEXPIRY_EMPTY_IN_STRUCT_ERR           ErrorCode = -1055000
	COLLCOMMENTS_EMPTY_IN_STRUCT_ERR         ErrorCode = -1056000
	COLLCREATE_EMPTY_IN_STRUCT_ERR           ErrorCode = -1057000
	COLLMODIFY_EMPTY_IN_STRUCT_ERR           ErrorCode = -1058000
	COLLACCESS_EMPTY_IN_STRUCT_ERR           ErrorCode = -1059000
	COLLACCESSINX_EMPTY_IN_STRUCT_ERR        ErrorCode = -1060000
	COLLMAPID_EMPTY_IN_STRUCT_ERR            ErrorCode = -1062000
	COLLINHERITANCE_EMPTY_IN_STRUCT_ERR      ErrorCode = -1063000
	RESCZONE_EMPTY_IN_STRUCT_ERR             ErrorCode = -1065000
	RESCLOC_EMPTY_IN_STRUCT_ERR              ErrorCode = -1066000
	RESCTYPE_EMPTY_IN_STRUCT_ERR             ErrorCode = -1067000
	RESCTYPEINX_EMPTY_IN_STRUCT_ERR          ErrorCode = -1068000
	RESCCLASS_EMPTY_IN_STRUCT_ERR            ErrorCode = -1069000
	RESCCLASSINX_EMPTY_IN_STRUCT_ERR         ErrorCode = -1070000
	RESCVAULTPATH_EMPTY_IN_STRUCT_ERR        ErrorCode = -1071000
	NUMOPEN_ORTS_EMPTY_IN_STRUCT_ERR         ErrorCode = -1072000
	PARAOPR_EMPTY_IN_STRUCT_ERR              ErrorCode = -1073000
	RESCID_EMPTY_IN_STRUCT_ERR               ErrorCode = -1074000
	GATEWAYADDR_EMPTY_IN_STRUCT_ERR          ErrorCode = -1075000
	RESCMAX_BJSIZE_EMPTY_IN_STRUCT_ERR       ErrorCode = -1076000
	FREESPACE_EMPTY_IN_STRUCT_ERR            ErrorCode = -1077000
	FREESPACETIME_EMPTY_IN_STRUCT_ERR        ErrorCode = -1078000
	FREESPACETIMESTAMP_EMPTY_IN_STRUCT_ERR   ErrorCode = -1079000
	RESCINFO_EMPTY_IN_STRUCT_ERR             ErrorCode = -1080000
	RESCCOMMENTS_EMPTY_IN_STRUCT_ERR         ErrorCode = -1081000
	RESCCREATE_EMPTY_IN_STRUCT_ERR           ErrorCode = -1082000
	RESCMODIFY_EMPTY_IN_STRUCT_ERR           ErrorCode = -1083000
	INPUT_ARG_NOT_WELL_FORMED_ERR            ErrorCode = -1084000
	INPUT_ARG_OUT_OF_ARGC_RANGE_ERR          ErrorCode = -1085000
	INSUFFICIENT_INPUT_ARG_ERR               ErrorCode = -1086000
	INPUT_ARG_DOES_NOT_MATCH_ERR             ErrorCode = -1087000
	RETRY_WITHOUT_RECOVERY_ERR               ErrorCode = -1088000
	CUT_ACTION_PROCESSED_ERR                 ErrorCode = -1089000
	ACTION_FAILED_ERR                        ErrorCode = -1090000
	FAIL_ACTION_ENCOUNTERED_ERR              ErrorCode = -1091000
	VARIABLE_NAME_TOO_LONG_ERR               ErrorCode = -1092000
	UNKNOWN_VARIABLE_MAP_ERR                 ErrorCode = -1093000
	UNDEFINED_VARIABLE_MAP_ERR               ErrorCode = -1094000
	NULL_VALUE_ERR                           ErrorCode = -1095000
	DVARMAP_FILE_READ_ERROR                  ErrorCode = -1096000
	NO_RULE_OR_MSI_FUNCTION_FOUND_ERR        ErrorCode = -1097000
	FILE_CREATE_ERROR                        ErrorCode = -1098000
	FMAP_FILE_READ_ERROR                     ErrorCode = -1099000
	DATE_FORMAT_ERR                          ErrorCode = -1100000
	RULE_FAILED_ERR                          ErrorCode = -1101000
	NO_MICROSERVICE_FOUND_ERR                ErrorCode = -1102000
	INVALID_REGEXP                           ErrorCode = -1103000
	INVALID_OBJECT_NAME                      ErrorCode = -1104000
	INVALID_OBJECT_TYPE                      ErrorCode = -1105000
	NO_VALUES_FOUND                          ErrorCode = -1106000
	NO_COLUMN_NAME_FOUND                     ErrorCode = -1107000
	BREAK_ACTION_ENCOUNTERED_ERR             ErrorCode = -1108000
	CUT_ACTION_ON_SUCCESS_PROCESSED_ERR      ErrorCode = -1109000
	MSI_OPERATION_NOT_ALLOWED                ErrorCode = -1110000
	PHP_EXEC_SCRIPT_ERR                      ErrorCode = -1600000
	PHP_REQUEST_STARTUP_ERR                  ErrorCode = -1601000
	PHP_OPEN_SCRIPT_FILE_ERR                 ErrorCode = -1602000
	PAM_AUTH_NOT_BUILT_INTO_CLIENT           ErrorCode = -991000
	PAM_AUTH_NOT_BUILT_INTO_SERVER           ErrorCode = -992000
	PAM_AUTH_PASSWORD_FAILED                 ErrorCode = -993000
	PAM_AUTH_PASSWORD_INVALID_TTL            ErrorCode = -994000
)

// MakeIRODSError creates an error from error code
func MakeIRODSError(code ErrorCode) error {
	return fmt.Errorf("iRODS Error - %v", code)
}

// MakeIRODSErrorFromString creates an error from error code and message
func MakeIRODSErrorFromString(code ErrorCode, message string) error {
	return fmt.Errorf("iRODS Error - %v: %s", code, message)
}
