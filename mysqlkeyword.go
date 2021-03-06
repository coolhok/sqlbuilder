package sqlbuilder

var mysqlkeyword = map[string]string{
	"ACCESSIBLE":                    "ACCESSIBLE",
	"ADD":                           "ADD",
	"AGGREGATE":                     "AGGREGATE",
	"ALTER":                         "ALTER",
	"ANALYZE":                       "ANALYZE",
	"AS":                            "AS",
	"ASENSITIVE":                    "ASENSITIVE",
	"AUTO_INCREMENT":                "AUTO_INCREMENT",
	"BACKUP":                        "BACKUP",
	"BETWEEN":                       "BETWEEN",
	"BINLOG":                        "BINLOG",
	"BLOCK":                         "BLOCK",
	"BOTH":                          "BOTH",
	"BYTE":                          "BYTE",
	"CASCADE":                       "CASCADE",
	"CATALOG_NAME":                  "CATALOG_NAME",
	"CHANGED":                       "CHANGED",
	"CHARACTER":                     "CHARACTER",
	"CHECKSUM":                      "CHECKSUM",
	"CLIENT":                        "CLIENT",
	"CODE":                          "CODE",
	"COLUMN":                        "COLUMN",
	"COLUMN_NAME":                   "COLUMN_NAME",
	"COMMITTED":                     "COMMITTED",
	"COMPRESSED":                    "COMPRESSED",
	"CONDITION":                     "CONDITION",
	"CONSTRAINT":                    "CONSTRAINT",
	"CONSTRAINT_SCHEMA":             "CONSTRAINT_SCHEMA",
	"CONTINUE":                      "CONTINUE",
	"CREATE":                        "CREATE",
	"CURRENT":                       "CURRENT",
	"CURRENT_TIMESTAMP":             "CURRENT_TIMESTAMP",
	"CURSOR_NAME":                   "CURSOR_NAME",
	"DATABASES":                     "DATABASES",
	"DATETIME":                      "DATETIME",
	"DAY_MICROSECOND":               "DAY_MICROSECOND",
	"DEALLOCATE":                    "DEALLOCATE",
	"DECLARE":                       "DECLARE",
	"DEFINER":                       "DEFINER",
	"DELETE":                        "DELETE",
	"DES_KEY_FILE":                  "DES_KEY_FILE",
	"DIRECTORY":                     "DIRECTORY",
	"DISK":                          "DISK",
	"DIV":                           "DIV",
	"DROP":                          "DROP",
	"DUPLICATE":                     "DUPLICATE",
	"ELSE":                          "ELSE",
	"ENCLOSED":                      "ENCLOSED",
	"ENDS":                          "ENDS",
	"ENUM":                          "ENUM",
	"ESCAPE":                        "ESCAPE",
	"EVENTS":                        "EVENTS",
	"EXECUTE":                       "EXECUTE",
	"EXPANSION":                     "EXPANSION",
	"EXPORT":                        "EXPORT",
	"FALSE":                         "FALSE",
	"FETCH":                         "FETCH",
	"FILE_BLOCK_SIZE":               "FILE_BLOCK_SIZE",
	"FIXED":                         "FIXED",
	"FLOAT8":                        "FLOAT8",
	"FOR":                           "FOR",
	"FORMAT":                        "FORMAT",
	"FULL":                          "FULL",
	"GENERAL":                       "GENERAL",
	"GEOMETRYCOLLECTION":            "GEOMETRYCOLLECTION",
	"GLOBAL":                        "GLOBAL",
	"GROUP":                         "GROUP",
	"HASH":                          "HASH",
	"HIGH_PRIORITY":                 "HIGH_PRIORITY",
	"HOUR":                          "HOUR",
	"HOUR_SECOND":                   "HOUR_SECOND",
	"IGNORE":                        "IGNORE",
	"IN":                            "IN",
	"INFILE":                        "INFILE",
	"INOUT":                         "INOUT",
	"INSERT_METHOD":                 "INSERT_METHOD",
	"INT":                           "INT",
	"INT3":                          "INT3",
	"INTEGER":                       "INTEGER",
	"INVOKER":                       "INVOKER",
	"IO_BEFORE_GTIDS":               "IO_BEFORE_GTIDS",
	"IS":                            "IS",
	"ITERATE":                       "ITERATE",
	"KEY":                           "KEY",
	"KILL":                          "KILL",
	"LEADING":                       "LEADING",
	"LEFT":                          "LEFT",
	"LIKE":                          "LIKE",
	"LINES":                         "LINES",
	"LOAD":                          "LOAD",
	"LOCALTIMESTAMP":                "LOCALTIMESTAMP",
	"LOGFILE":                       "LOGFILE",
	"LONGBLOB":                      "LONGBLOB",
	"LOW_PRIORITY":                  "LOW_PRIORITY",
	"MASTER_BIND":                   "MASTER_BIND",
	"MASTER_HEARTBEAT_PERIOD":       "MASTER_HEARTBEAT_PERIOD",
	"MASTER_LOG_POS":                "MASTER_LOG_POS",
	"MASTER_RETRY_COUNT":            "MASTER_RETRY_COUNT",
	"MASTER_SSL_CA":                 "MASTER_SSL_CA",
	"MASTER_SSL_CIPHER":             "MASTER_SSL_CIPHER",
	"MASTER_SSL_KEY":                "MASTER_SSL_KEY",
	"MASTER_USER":                   "MASTER_USER",
	"MAX_CONNECTIONS_PER_HOUR":      "MAX_CONNECTIONS_PER_HOUR",
	"MAX_SIZE":                      "MAX_SIZE",
	"MAX_USER_CONNECTIONS":          "MAX_USER_CONNECTIONS",
	"MEDIUMINT":                     "MEDIUMINT",
	"MERGE":                         "MERGE",
	"MIDDLEINT":                     "MIDDLEINT",
	"MINUTE_MICROSECOND":            "MINUTE_MICROSECOND",
	"MOD":                           "MOD",
	"MODIFY":                        "MODIFY",
	"MULTIPOINT":                    "MULTIPOINT",
	"MYSQL_ERRNO":                   "MYSQL_ERRNO",
	"NATIONAL":                      "NATIONAL",
	"NDB":                           "NDB",
	"NEW":                           "NEW",
	"NODEGROUP":                     "NODEGROUP",
	"NOT":                           "NOT",
	"NULL":                          "NULL",
	"NVARCHAR":                      "NVARCHAR",
	"ON":                            "ON",
	"OPEN":                          "OPEN",
	"OPTION":                        "OPTION",
	"OR":                            "OR",
	"OUTER":                         "OUTER",
	"PACK_KEYS":                     "PACK_KEYS",
	"PARSE_GCOL_EXPR":               "PARSE_GCOL_EXPR",
	"PARTITIONING":                  "PARTITIONING",
	"PHASE":                         "PHASE",
	"PLUGIN_DIR":                    "PLUGIN_DIR",
	"PORT":                          "PORT",
	"PREPARE":                       "PREPARE",
	"PRIMARY":                       "PRIMARY",
	"PROCESSLIST":                   "PROCESSLIST",
	"PROXY":                         "PROXY",
	"QUERY":                         "QUERY",
	"READ":                          "READ",
	"READ_WRITE":                    "READ_WRITE",
	"RECOVER":                       "RECOVER",
	"REDUNDANT":                     "REDUNDANT",
	"RELAY":                         "RELAY",
	"RELAY_LOG_POS":                 "RELAY_LOG_POS",
	"RELOAD":                        "RELOAD",
	"REORGANIZE":                    "REORGANIZE",
	"REPEATABLE":                    "REPEATABLE",
	"REPLICATE_DO_TABLE":            "REPLICATE_DO_TABLE",
	"REPLICATE_REWRITE_DB":          "REPLICATE_REWRITE_DB",
	"REPLICATION":                   "REPLICATION",
	"RESIGNAL":                      "RESIGNAL",
	"RESUME":                        "RESUME",
	"RETURNS":                       "RETURNS",
	"RIGHT":                         "RIGHT",
	"ROLLUP":                        "ROLLUP",
	"ROW":                           "ROW",
	"ROW_FORMAT":                    "ROW_FORMAT",
	"SCHEDULE":                      "SCHEDULE",
	"SCHEMA_NAME":                   "SCHEMA_NAME",
	"SECURITY":                      "SECURITY",
	"SEPARATOR":                     "SEPARATOR",
	"SERVER":                        "SERVER",
	"SHARE":                         "SHARE",
	"SIGNAL":                        "SIGNAL",
	"SLAVE":                         "SLAVE",
	"SNAPSHOT":                      "SNAPSHOT",
	"SONAME":                        "SONAME",
	"SPATIAL":                       "SPATIAL",
	"SQLEXCEPTION":                  "SQLEXCEPTION",
	"SQL_AFTER_GTIDS":               "SQL_AFTER_GTIDS",
	"SQL_BIG_RESULT":                "SQL_BIG_RESULT",
	"SQL_CALC_FOUND_ROWS":           "SQL_CALC_FOUND_ROWS",
	"SQL_THREAD":                    "SQL_THREAD",
	"SQL_TSI_MINUTE":                "SQL_TSI_MINUTE",
	"SQL_TSI_SECOND":                "SQL_TSI_SECOND",
	"SSL":                           "SSL",
	"STARTING":                      "STARTING",
	"STATS_PERSISTENT":              "STATS_PERSISTENT",
	"STOP":                          "STOP",
	"STRAIGHT_JOIN":                 "STRAIGHT_JOIN",
	"SUBJECT":                       "SUBJECT",
	"SUPER":                         "SUPER",
	"SWITCHES":                      "SWITCHES",
	"TABLESPACE":                    "TABLESPACE",
	"TEMPORARY":                     "TEMPORARY",
	"TEXT":                          "TEXT",
	"TIME":                          "TIME",
	"TIMESTAMPDIFF":                 "TIMESTAMPDIFF",
	"TINYTEXT":                      "TINYTEXT",
	"TRANSACTION":                   "TRANSACTION",
	"TRUE":                          "TRUE",
	"TYPES":                         "TYPES",
	"UNDO":                          "UNDO",
	"UNICODE":                       "UNICODE",
	"UNIQUE":                        "UNIQUE",
	"UNSIGNED":                      "UNSIGNED",
	"UPGRADE":                       "UPGRADE",
	"USER":                          "USER",
	"USING":                         "USING",
	"UTC_TIMESTAMP":                 "UTC_TIMESTAMP",
	"VALUES":                        "VALUES",
	"VARCHARACTER":                  "VARCHARACTER",
	"VIEW":                          "VIEW",
	"WARNINGS":                      "WARNINGS",
	"WHEN":                          "WHEN",
	"WITH":                          "WITH",
	"WRAPPER":                       "WRAPPER",
	"XA":                            "XA",
	"XOR":                           "XOR",
	"ZEROFILL":                      "ZEROFILL",
	"ACCOUNT":                       "ACCOUNT",
	"AFTER":                         "AFTER",
	"ALGORITHM":                     "ALGORITHM",
	"ALWAYS":                        "ALWAYS",
	"AND":                           "AND",
	"ASC":                           "ASC",
	"AT":                            "AT",
	"AVG":                           "AVG",
	"BEFORE":                        "BEFORE",
	"BIGINT":                        "BIGINT",
	"BIT":                           "BIT",
	"BOOL":                          "BOOL",
	"BTREE":                         "BTREE",
	"CACHE":                         "CACHE",
	"CASCADED":                      "CASCADED",
	"CHAIN":                         "CHAIN",
	"CHANNEL":                       "CHANNEL",
	"CHARSET":                       "CHARSET",
	"CIPHER":                        "CIPHER",
	"CLOSE":                         "CLOSE",
	"COLLATE":                       "COLLATE",
	"COLUMNS":                       "COLUMNS",
	"COMMENT":                       "COMMENT",
	"COMPACT":                       "COMPACT",
	"COMPRESSION":                   "COMPRESSION",
	"CONNECTION":                    "CONNECTION",
	"CONSTRAINT_CATALOG":            "CONSTRAINT_CATALOG",
	"CONTAINS":                      "CONTAINS",
	"CONVERT":                       "CONVERT",
	"CROSS":                         "CROSS",
	"CURRENT_DATE":                  "CURRENT_DATE",
	"CURRENT_USER":                  "CURRENT_USER",
	"DATA":                          "DATA",
	"DATAFILE":                      "DATAFILE",
	"DAY":                           "DAY",
	"DAY_MINUTE":                    "DAY_MINUTE",
	"DEC":                           "DEC",
	"DEFAULT":                       "DEFAULT",
	"DELAYED":                       "DELAYED",
	"DESC":                          "DESC",
	"DETERMINISTIC":                 "DETERMINISTIC",
	"DISABLE":                       "DISABLE",
	"DISTINCT":                      "DISTINCT",
	"DO":                            "DO",
	"DUAL":                          "DUAL",
	"DYNAMIC":                       "DYNAMIC",
	"ELSEIF":                        "ELSEIF",
	"ENCRYPTION":                    "ENCRYPTION",
	"ENGINE":                        "ENGINE",
	"ERROR":                         "ERROR",
	"ESCAPED":                       "ESCAPED",
	"EVERY":                         "EVERY",
	"EXISTS":                        "EXISTS",
	"EXPIRE":                        "EXPIRE",
	"EXTENDED":                      "EXTENDED",
	"FAST":                          "FAST",
	"FIELDS":                        "FIELDS",
	"FILTER":                        "FILTER",
	"FLOAT":                         "FLOAT",
	"FLUSH":                         "FLUSH",
	"FORCE":                         "FORCE",
	"FOUND":                         "FOUND",
	"FULLTEXT":                      "FULLTEXT",
	"GENERATED":                     "GENERATED",
	"GET":                           "GET",
	"GRANT":                         "GRANT",
	"GROUP_REPLICATION":             "GROUP_REPLICATION",
	"HAVING":                        "HAVING",
	"HOST":                          "HOST",
	"HOUR_MICROSECOND":              "HOUR_MICROSECOND",
	"IDENTIFIED":                    "IDENTIFIED",
	"IGNORE_SERVER_IDS":             "IGNORE_SERVER_IDS",
	"INDEX":                         "INDEX",
	"INITIAL_SIZE":                  "INITIAL_SIZE",
	"INSENSITIVE":                   "INSENSITIVE",
	"INSTALL":                       "INSTALL",
	"INT1":                          "INT1",
	"INT4":                          "INT4",
	"INTERVAL":                      "INTERVAL",
	"IO":                            "IO",
	"IO_THREAD":                     "IO_THREAD",
	"ISOLATION":                     "ISOLATION",
	"JOIN":                          "JOIN",
	"KEYS":                          "KEYS",
	"LANGUAGE":                      "LANGUAGE",
	"LEAVE":                         "LEAVE",
	"LESS":                          "LESS",
	"LIMIT":                         "LIMIT",
	"LINESTRING":                    "LINESTRING",
	"LOCAL":                         "LOCAL",
	"LOCK":                          "LOCK",
	"LOGS":                          "LOGS",
	"LONGTEXT":                      "LONGTEXT",
	"MASTER":                        "MASTER",
	"MASTER_CONNECT_RETRY":          "MASTER_CONNECT_RETRY",
	"MASTER_HOST":                   "MASTER_HOST",
	"MASTER_PASSWORD":               "MASTER_PASSWORD",
	"MASTER_SERVER_ID":              "MASTER_SERVER_ID",
	"MASTER_SSL_CAPATH":             "MASTER_SSL_CAPATH",
	"MASTER_SSL_CRL":                "MASTER_SSL_CRL",
	"MASTER_SSL_VERIFY_SERVER_CERT": "MASTER_SSL_VERIFY_SERVER_CERT",
	"MATCH":                       "MATCH",
	"MAX_QUERIES_PER_HOUR":        "MAX_QUERIES_PER_HOUR",
	"MAX_STATEMENT_TIME":          "MAX_STATEMENT_TIME",
	"MEDIUM":                      "MEDIUM",
	"MEDIUMTEXT":                  "MEDIUMTEXT",
	"MESSAGE_TEXT":                "MESSAGE_TEXT",
	"MIGRATE":                     "MIGRATE",
	"MINUTE_SECOND":               "MINUTE_SECOND",
	"MODE":                        "MODE",
	"MONTH":                       "MONTH",
	"MULTIPOLYGON":                "MULTIPOLYGON",
	"NAME":                        "NAME",
	"NATURAL":                     "NATURAL",
	"NDBCLUSTER":                  "NDBCLUSTER",
	"NEXT":                        "NEXT",
	"NONBLOCKING":                 "NONBLOCKING",
	"NO_WAIT":                     "NO_WAIT",
	"NUMBER":                      "NUMBER",
	"OFFSET":                      "OFFSET",
	"ONE":                         "ONE",
	"OPTIMIZE":                    "OPTIMIZE",
	"OPTIONALLY":                  "OPTIONALLY",
	"ORDER":                       "ORDER",
	"OUTFILE":                     "OUTFILE",
	"PAGE":                        "PAGE",
	"PARTIAL":                     "PARTIAL",
	"PARTITIONS":                  "PARTITIONS",
	"PLUGIN":                      "PLUGIN",
	"POINT":                       "POINT",
	"PRECEDES":                    "PRECEDES",
	"PRESERVE":                    "PRESERVE",
	"PRIVILEGES":                  "PRIVILEGES",
	"PROFILE":                     "PROFILE",
	"PURGE":                       "PURGE",
	"QUICK":                       "QUICK",
	"READS":                       "READS",
	"REAL":                        "REAL",
	"REDOFILE":                    "REDOFILE",
	"REFERENCES":                  "REFERENCES",
	"RELAYLOG":                    "RELAYLOG",
	"RELAY_THREAD":                "RELAY_THREAD",
	"REMOVE":                      "REMOVE",
	"REPAIR":                      "REPAIR",
	"REPLACE":                     "REPLACE",
	"REPLICATE_IGNORE_DB":         "REPLICATE_IGNORE_DB",
	"REPLICATE_WILD_DO_TABLE":     "REPLICATE_WILD_DO_TABLE",
	"REQUIRE":                     "REQUIRE",
	"RESTORE":                     "RESTORE",
	"RETURN":                      "RETURN",
	"REVERSE":                     "REVERSE",
	"RLIKE":                       "RLIKE",
	"ROTATE":                      "ROTATE",
	"ROWS":                        "ROWS",
	"RTREE":                       "RTREE",
	"SCHEMA":                      "SCHEMA",
	"SECOND":                      "SECOND",
	"SELECT":                      "SELECT",
	"SERIAL":                      "SERIAL",
	"SESSION":                     "SESSION",
	"SHOW":                        "SHOW",
	"SIGNED":                      "SIGNED",
	"SLOW":                        "SLOW",
	"SOCKET":                      "SOCKET",
	"SOUNDS":                      "SOUNDS",
	"SPECIFIC":                    "SPECIFIC",
	"SQLSTATE":                    "SQLSTATE",
	"SQL_AFTER_MTS_GAPS":          "SQL_AFTER_MTS_GAPS",
	"SQL_BUFFER_RESULT":           "SQL_BUFFER_RESULT",
	"SQL_NO_CACHE":                "SQL_NO_CACHE",
	"SQL_TSI_DAY":                 "SQL_TSI_DAY",
	"SQL_TSI_MONTH":               "SQL_TSI_MONTH",
	"SQL_TSI_WEEK":                "SQL_TSI_WEEK",
	"STACKED":                     "STACKED",
	"STARTS":                      "STARTS",
	"STATS_SAMPLE_PAGES":          "STATS_SAMPLE_PAGES",
	"STORAGE":                     "STORAGE",
	"STRING":                      "STRING",
	"SUBPARTITION":                "SUBPARTITION",
	"SUSPEND":                     "SUSPEND",
	"TABLE":                       "TABLE",
	"TABLE_CHECKSUM":              "TABLE_CHECKSUM",
	"TEMPTABLE":                   "TEMPTABLE",
	"THAN":                        "THAN",
	"TIMESTAMP":                   "TIMESTAMP",
	"TINYBLOB":                    "TINYBLOB",
	"TO":                          "TO",
	"TRIGGER":                     "TRIGGER",
	"TRUNCATE":                    "TRUNCATE",
	"UNCOMMITTED":                 "UNCOMMITTED",
	"UNDOFILE":                    "UNDOFILE",
	"UNINSTALL":                   "UNINSTALL",
	"UNKNOWN":                     "UNKNOWN",
	"UNTIL":                       "UNTIL",
	"USAGE":                       "USAGE",
	"USER_RESOURCES":              "USER_RESOURCES",
	"UTC_DATE":                    "UTC_DATE",
	"VALIDATION":                  "VALIDATION",
	"VARBINARY":                   "VARBINARY",
	"VARIABLES":                   "VARIABLES",
	"VIRTUAL":                     "VIRTUAL",
	"WEEK":                        "WEEK",
	"WHERE":                       "WHERE",
	"WITHOUT":                     "WITHOUT",
	"WRITE":                       "WRITE",
	"XID":                         "XID",
	"YEAR":                        "YEAR",
	"ACTION":                      "ACTION",
	"AGAINST":                     "AGAINST",
	"ALL":                         "ALL",
	"ANALYSE":                     "ANALYSE",
	"ANY":                         "ANY",
	"ASCII":                       "ASCII",
	"AUTOEXTEND_SIZE":             "AUTOEXTEND_SIZE",
	"AVG_ROW_LENGTH":              "AVG_ROW_LENGTH",
	"BEGIN":                       "BEGIN",
	"BINARY":                      "BINARY",
	"BLOB":                        "BLOB",
	"BOOLEAN":                     "BOOLEAN",
	"BY":                          "BY",
	"CALL":                        "CALL",
	"CASE":                        "CASE",
	"CHANGE":                      "CHANGE",
	"CHAR":                        "CHAR",
	"CHECK":                       "CHECK",
	"CLASS_ORIGIN":                "CLASS_ORIGIN",
	"COALESCE":                    "COALESCE",
	"COLLATION":                   "COLLATION",
	"COLUMN_FORMAT":               "COLUMN_FORMAT",
	"COMMIT":                      "COMMIT",
	"COMPLETION":                  "COMPLETION",
	"CONCURRENT":                  "CONCURRENT",
	"CONSISTENT":                  "CONSISTENT",
	"CONSTRAINT_NAME":             "CONSTRAINT_NAME",
	"CONTEXT":                     "CONTEXT",
	"CPU":                         "CPU",
	"CUBE":                        "CUBE",
	"CURRENT_TIME":                "CURRENT_TIME",
	"CURSOR":                      "CURSOR",
	"DATABASE":                    "DATABASE",
	"DATE":                        "DATE",
	"DAY_HOUR":                    "DAY_HOUR",
	"DAY_SECOND":                  "DAY_SECOND",
	"DECIMAL":                     "DECIMAL",
	"DEFAULT_AUTH":                "DEFAULT_AUTH",
	"DELAY_KEY_WRITE":             "DELAY_KEY_WRITE",
	"DESCRIBE":                    "DESCRIBE",
	"DIAGNOSTICS":                 "DIAGNOSTICS",
	"DISCARD":                     "DISCARD",
	"DISTINCTROW":                 "DISTINCTROW",
	"DOUBLE":                      "DOUBLE",
	"DUMPFILE":                    "DUMPFILE",
	"EACH":                        "EACH",
	"ENABLE":                      "ENABLE",
	"END":                         "END",
	"ENGINES":                     "ENGINES",
	"ERRORS":                      "ERRORS",
	"EVENT":                       "EVENT",
	"EXCHANGE":                    "EXCHANGE",
	"EXIT":                        "EXIT",
	"EXPLAIN":                     "EXPLAIN",
	"EXTENT_SIZE":                 "EXTENT_SIZE",
	"FAULTS":                      "FAULTS",
	"FILE":                        "FILE",
	"FIRST":                       "FIRST",
	"FLOAT4":                      "FLOAT4",
	"FOLLOWS":                     "FOLLOWS",
	"FOREIGN":                     "FOREIGN",
	"FROM":                        "FROM",
	"FUNCTION":                    "FUNCTION",
	"GEOMETRY":                    "GEOMETRY",
	"GET_FORMAT":                  "GET_FORMAT",
	"GRANTS":                      "GRANTS",
	"HANDLER":                     "HANDLER",
	"HELP":                        "HELP",
	"HOSTS":                       "HOSTS",
	"HOUR_MINUTE":                 "HOUR_MINUTE",
	"IF":                          "IF",
	"IMPORT":                      "IMPORT",
	"INDEXES":                     "INDEXES",
	"INNER":                       "INNER",
	"INSERT":                      "INSERT",
	"INSTANCE":                    "INSTANCE",
	"INT2":                        "INT2",
	"INT8":                        "INT8",
	"INTO":                        "INTO",
	"IO_AFTER_GTIDS":              "IO_AFTER_GTIDS",
	"IPC":                         "IPC",
	"ISSUER":                      "ISSUER",
	"JSON":                        "JSON",
	"KEY_BLOCK_SIZE":              "KEY_BLOCK_SIZE",
	"LAST":                        "LAST",
	"LEAVES":                      "LEAVES",
	"LEVEL":                       "LEVEL",
	"LINEAR":                      "LINEAR",
	"LIST":                        "LIST",
	"LOCALTIME":                   "LOCALTIME",
	"LOCKS":                       "LOCKS",
	"LONG":                        "LONG",
	"LOOP":                        "LOOP",
	"MASTER_AUTO_POSITION":        "MASTER_AUTO_POSITION",
	"MASTER_DELAY":                "MASTER_DELAY",
	"MASTER_LOG_FILE":             "MASTER_LOG_FILE",
	"MASTER_PORT":                 "MASTER_PORT",
	"MASTER_SSL":                  "MASTER_SSL",
	"MASTER_SSL_CERT":             "MASTER_SSL_CERT",
	"MASTER_SSL_CRLPATH":          "MASTER_SSL_CRLPATH",
	"MASTER_TLS_VERSION":          "MASTER_TLS_VERSION",
	"MAXVALUE":                    "MAXVALUE",
	"MAX_ROWS":                    "MAX_ROWS",
	"MAX_UPDATES_PER_HOUR":        "MAX_UPDATES_PER_HOUR",
	"MEDIUMBLOB":                  "MEDIUMBLOB",
	"MEMORY":                      "MEMORY",
	"MICROSECOND":                 "MICROSECOND",
	"MINUTE":                      "MINUTE",
	"MIN_ROWS":                    "MIN_ROWS",
	"MODIFIES":                    "MODIFIES",
	"MULTILINESTRING":             "MULTILINESTRING",
	"MUTEX":                       "MUTEX",
	"NAMES":                       "NAMES",
	"NCHAR":                       "NCHAR",
	"NEVER":                       "NEVER",
	"NO":                          "NO",
	"NONE":                        "NONE",
	"NO_WRITE_TO_BINLOG":          "NO_WRITE_TO_BINLOG",
	"NUMERIC":                     "NUMERIC",
	"OLD_PASSWORD":                "OLD_PASSWORD",
	"ONLY":                        "ONLY",
	"OPTIMIZER_COSTS":             "OPTIMIZER_COSTS",
	"OPTIONS":                     "OPTIONS",
	"OUT":                         "OUT",
	"OWNER":                       "OWNER",
	"PARSER":                      "PARSER",
	"PARTITION":                   "PARTITION",
	"PASSWORD":                    "PASSWORD",
	"PLUGINS":                     "PLUGINS",
	"POLYGON":                     "POLYGON",
	"PRECISION":                   "PRECISION",
	"PREV":                        "PREV",
	"PROCEDURE":                   "PROCEDURE",
	"PROFILES":                    "PROFILES",
	"QUARTER":                     "QUARTER",
	"RANGE":                       "RANGE",
	"READ_ONLY":                   "READ_ONLY",
	"REBUILD":                     "REBUILD",
	"REDO_BUFFER_SIZE":            "REDO_BUFFER_SIZE",
	"REGEXP":                      "REGEXP",
	"RELAY_LOG_FILE":              "RELAY_LOG_FILE",
	"RELEASE":                     "RELEASE",
	"RENAME":                      "RENAME",
	"REPEAT":                      "REPEAT",
	"REPLICATE_DO_DB":             "REPLICATE_DO_DB",
	"REPLICATE_IGNORE_TABLE":      "REPLICATE_IGNORE_TABLE",
	"REPLICATE_WILD_IGNORE_TABLE": "REPLICATE_WILD_IGNORE_TABLE",
	"RESET":              "RESET",
	"RESTRICT":           "RESTRICT",
	"RETURNED_SQLSTATE":  "RETURNED_SQLSTATE",
	"REVOKE":             "REVOKE",
	"ROLLBACK":           "ROLLBACK",
	"ROUTINE":            "ROUTINE",
	"ROW_COUNT":          "ROW_COUNT",
	"SAVEPOINT":          "SAVEPOINT",
	"SCHEMAS":            "SCHEMAS",
	"SECOND_MICROSECOND": "SECOND_MICROSECOND",
	"SENSITIVE":          "SENSITIVE",
	"SERIALIZABLE":       "SERIALIZABLE",
	"SET":                "SET",
	"SHUTDOWN":           "SHUTDOWN",
	"SIMPLE":             "SIMPLE",
	"SMALLINT":           "SMALLINT",
	"SOME":               "SOME",
	"SOURCE":             "SOURCE",
	"SQL":                "SQL",
	"SQLWARNING":         "SQLWARNING",
	"SQL_BEFORE_GTIDS":   "SQL_BEFORE_GTIDS",
	"SQL_CACHE":          "SQL_CACHE",
	"SQL_SMALL_RESULT":   "SQL_SMALL_RESULT",
	"SQL_TSI_HOUR":       "SQL_TSI_HOUR",
	"SQL_TSI_QUARTER":    "SQL_TSI_QUARTER",
	"SQL_TSI_YEAR":       "SQL_TSI_YEAR",
	"START":              "START",
	"STATS_AUTO_RECALC":  "STATS_AUTO_RECALC",
	"STATUS":             "STATUS",
	"STORED":             "STORED",
	"SUBCLASS_ORIGIN":    "SUBCLASS_ORIGIN",
	"SUBPARTITIONS":      "SUBPARTITIONS",
	"SWAPS":              "SWAPS",
	"TABLES":             "TABLES",
	"TABLE_NAME":         "TABLE_NAME",
	"TERMINATED":         "TERMINATED",
	"THEN":               "THEN",
	"TIMESTAMPADD":       "TIMESTAMPADD",
	"TINYINT":            "TINYINT",
	"TRAILING":           "TRAILING",
	"TRIGGERS":           "TRIGGERS",
	"TYPE":               "TYPE",
	"UNDEFINED":          "UNDEFINED",
	"UNDO_BUFFER_SIZE":   "UNDO_BUFFER_SIZE",
	"UNION":              "UNION",
	"UNLOCK":             "UNLOCK",
	"UPDATE":             "UPDATE",
	"USE":                "USE",
	"USE_FRM":            "USE_FRM",
	"UTC_TIME":           "UTC_TIME",
	"VALUE":              "VALUE",
	"VARCHAR":            "VARCHAR",
	"VARYING":            "VARYING",
	"WAIT":               "WAIT",
	"WEIGHT_STRING":      "WEIGHT_STRING",
	"WHILE":              "WHILE",
	"WORK":               "WORK",
	"X509":               "X509",
	"XML":                "XML",
	"YEAR_MONTH":         "YEAR_MONTH",
}
